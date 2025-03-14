// Copyright 2025 - See NOTICE file for copyright holders.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package channel

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/assetholder"
	"github.com/perun-network/perun-eth-backend/subscription"
	"github.com/perun-network/perun-eth-backend/wallet"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
	perunwallet "perun.network/go-perun/wallet"
	pcontext "polycry.pt/poly-go/context"
	perror "polycry.pt/poly-go/errors"
)

type assetHolder struct {
	*assetholder.Assetholder
	*common.Address
	contract   *bind.BoundContract
	assetIndex channel.Index
}

// Funder implements the channel.Funder interface for Ethereum.
//
// In addition to the `Fund` method required by the `Funder` interface, it also
// provides additional functions for convenience.
//
// All the exported methods are thread-safe and can be invoked concurrently.
type Funder struct {
	mtx sync.RWMutex

	// Egoistic Part discloses if a participant should fund last.
	EgoisticPart []bool

	ContractBackend
	// accounts associates an Account to every AssetIndex.
	accounts map[AssetMapKey]accounts.Account
	// depositors associates a Depositor to every AssetIndex.
	depositors map[AssetMapKey]Depositor
	log        log.Logger // structured logger
}

const funderEventBufSize = 10

// compile time check that we implement the perun funder interface.
var _ channel.Funder = (*Funder)(nil)

// NewFunder creates a new ethereum funder.
func NewFunder(backend ContractBackend) *Funder {
	return &Funder{
		ContractBackend: backend,
		accounts:        make(map[AssetMapKey]accounts.Account),
		depositors:      make(map[AssetMapKey]Depositor),
		log:             log.Default(),
	}
}

// SetEgoisticPart sets the egoistic part of the funder.
func (f *Funder) SetEgoisticPart(idx channel.Index, numParts int) {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	f.EgoisticPart = make([]bool, numParts)
	f.EgoisticPart[idx] = true
}

// RegisterAsset registers the depositor and account for the specified asset in
// the funder.
//
// Deposits for this asset will be sent using the depositors from the
// specified account when funding. Hence, it is the responsibility of the
// caller to ensure, the account has sufficient balance in the asset.
//
// It returns true if the asset was successfully registered, false if it was already
// present.
func (f *Funder) RegisterAsset(asset Asset, d Depositor, acc accounts.Account) bool {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	// Both the maps (f.accounts & f.assets) are always modified together such
	// that they will have the same set of keys. Hence, it is okay to check one
	// of the two.
	if _, ok := f.accounts[asset.MapKey()]; ok {
		return false
	}
	f.accounts[asset.MapKey()] = acc
	f.depositors[asset.MapKey()] = d
	return true
}

// IsAssetRegistered returns if the specified asset is registered in the funder or not.
// If is registered, then the corresponding depositor and account will also be
// returned.
func (f *Funder) IsAssetRegistered(asset Asset) (Depositor, accounts.Account, bool) {
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	// Both the maps (f.accounts & f.assets) are always modified togethe such
	// that they will have the same set of keys. Hence, it is okay to check one
	// of the two.
	if acc, ok := f.accounts[asset.MapKey()]; ok {
		return f.depositors[asset.MapKey()], acc, true
	}
	return nil, accounts.Account{}, false
}

// Fund implements the channel.Funder interface. It funds all assets in
// parallel. If not all participants successfully fund within a timeframe of
// ChallengeDuration seconds, Fund returns a FundingTimeoutError.
//
// If funding on a real blockchain, make sure that the passed context doesn't
// cancel before the funding period of length ChallengeDuration elapses, or
// funding will be canceled prematurely.
//
//nolint:funlen
func (f *Funder) Fund(ctx context.Context, request channel.FundingReq) error {
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	channelID := request.Params.ID()
	f.log.WithField("channel", channelID).Debug("Funding Channel.")

	// We wait for the funding timeout in a go routine and cancel the funding
	// context if the timeout elapses.
	ctx, cancel, err := f.fundingTimeoutContext(ctx, request)
	if err != nil {
		return err
	}
	defer func() {
		cancel() // Cancel the context if we return before the block timeout.
	}()

	// Extract only ethereum Assets to fund
	var ethAssets []*Asset
	for _, asset := range request.State.Assets {
		ethAsset, ok := asset.(*Asset)
		if ok {
			ethAssets = append(ethAssets, ethAsset)
		}
	}

	// Fund each asset, saving the TX in `txs` and the errors in `errg`.
	txs, errg := f.fundAssets(ctx, ethAssets, channelID, request)

	// Wait for the TXs to be mined.
	for a, asset := range ethAssets {
		for i, tx := range txs[a] {
			acc := f.accounts[asset.MapKey()]
			if _, err := f.ConfirmTransaction(ctx, tx, acc); err != nil {
				if errors.Is(err, errTxTimedOut) {
					err = client.NewTxTimedoutError(Fund.String(), tx.Hash().Hex(), err.Error())
				}
				return errors.WithMessagef(err, "sending %dth funding TX for asset %d", i, a)
			}
			f.log.Debugf("Mined TX: %v", tx.Hash().Hex())
		}
	}

	// Wait for the funding events or timeout.
	var fundingErrs []*channel.AssetFundingError
	nonFundingErrg := perror.NewGatherer()
	for _, err := range perror.Causes(errg.Wait()) {
		if channel.IsAssetFundingError(err) && err != nil {
			fundingErr, ok := err.(*channel.AssetFundingError)
			if !ok {
				return fmt.Errorf("wrong type: expected %T, got %T", &channel.AssetFundingError{}, err)
			}
			fundingErrs = append(fundingErrs, fundingErr)
		} else if err != nil {
			nonFundingErrg.Add(err)
		}
	}
	// Prioritize funding errors over other errors.
	if len(fundingErrs) != 0 {
		return channel.NewFundingTimeoutError(fundingErrs)
	}
	return nonFundingErrg.Err()
}

func (f *Funder) fundingTimeoutContext(ctx context.Context, req channel.FundingReq) (context.Context, context.CancelFunc, error) {
	timeout, err := NewBlockTimeoutDuration(ctx, f.ContractInterface, req.Params.ChallengeDuration)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "creating block timeout")
	}
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		if err := timeout.Wait(ctx); err != nil && !pcontext.IsContextError(err) {
			f.log.Warn("Fund: BlockTimeout.Wait runtime error: ", err)
		}
		cancel() // cancel funding context on funding timeout
	}()
	return ctx, cancel, nil
}

// fundAssets funds each asset of the funding agreement in the `req`.
// Sends the transactions and returns them. Wait on the returned gatherer
// to ensure that all `funding` events were received.
func (f *Funder) fundAssets(ctx context.Context, assets []*Asset, channelID channel.ID, req channel.FundingReq) ([]types.Transactions, *perror.Gatherer) {
	txs := make([]types.Transactions, len(assets))
	errg := perror.NewGatherer()
	fundingIDs := FundingIDs(channelID, req.Params.Parts...)

	egoistic := false
	if len(f.EgoisticPart) != 0 {
		egoistic = f.EgoisticPart[req.Idx]
		err := checkEgoisticPart(f.EgoisticPart)
		if err != nil {
			errg.Add(errors.WithMessage(err, "checking egoistic part"))
		}
	}

	contracts := make([]assetHolder, len(assets))
	for i, asset := range assets {
		// Bind contract.
		assetIdx, ok := assetIdx(req.State.Assets, asset)
		if !ok {
			errg.Add(errors.New("asset not found in funding request"))
			continue
		}
		contracts[i] = bindAssetHolder(f.ContractBackend, asset, assetIdx)
	}

	if egoistic {
		err := f.WaitForOthersFundingConfirmation(ctx, req, contracts, fundingIDs)
		if err != nil {
			errg.Add(errors.WithMessage(err, "funding asset"))
		}
	}

	for i, asset := range assets {
		// Bind contract.
		assetIdx, ok := assetIdx(req.State.Assets, asset)
		if !ok {
			errg.Add(errors.New("asset not found in funding request"))
			continue
		}
		contract := bindAssetHolder(f.ContractBackend, asset, assetIdx)

		// Wait for the funding event in a goroutine.
		errg.Go(func() error {
			return f.waitForFundingConfirmation(ctx, req, contract, fundingIDs)
		})

		// Send the funding TX.
		tx, err := f.sendFundingTx(ctx, asset, req, contract, fundingIDs[req.Idx])
		if err != nil {
			f.log.WithField("asset", asset).WithError(err).Errorf("Could not fund asset %v", req.Params.Parts[req.Idx])
			errg.Add(errors.WithMessage(err, "funding asset"))
			continue
		}
		txs[i] = tx
	}
	return txs, errg
}

// sendFundingTx sends and returns the TXs that are needed to fulfill the
// funding request. It is idempotent.
func (f *Funder) sendFundingTx(
	ctx context.Context,
	asset channel.Asset,
	request channel.FundingReq,
	contract assetHolder,
	fundingID [32]byte,
) (txs []*types.Transaction, fatal error) {
	bal := request.Agreement[contract.assetIndex][request.Idx]
	if bal == nil || bal.Sign() <= 0 {
		f.log.WithFields(log.Fields{"channel": request.Params.ID(), "idx": request.Idx}).Debug("Skipped zero funding.")
		return nil, nil
	}

	alreadyFunded, err := f.checkFunded(ctx, bal, contract, fundingID)
	if err != nil {
		return nil, errors.WithMessage(err, "checking funded")
	} else if alreadyFunded {
		f.log.WithFields(log.Fields{"channel": request.Params.ID(), "idx": request.Idx}).Debug("Skipped second funding.")
		return nil, nil
	}

	assetTyped, ok := asset.(*Asset)
	if !ok {
		return nil, fmt.Errorf("wrong type: expected %T, got %T", &Asset{}, asset)
	}
	return f.deposit(ctx, bal, *assetTyped, fundingID)
}

// deposit deposits funds for one funding-ID by calling the associated Depositor.
// Returns an error if no matching Depositor or Account could be found.
func (f *Funder) deposit(ctx context.Context, bal *big.Int, asset Asset, fundingID [32]byte) (types.Transactions, error) {
	depositor, ok := f.depositors[asset.MapKey()]
	if !ok {
		return nil, errors.Errorf("could not find Depositor for asset #%d", asset)
	}
	acc, ok := f.accounts[asset.MapKey()]
	if !ok {
		return nil, errors.Errorf("could not find account for asset #%d", asset)
	}
	return depositor.Deposit(ctx, *NewDepositReq(bal, f.ContractBackend, asset, acc, fundingID))
}

// checkFunded returns whether `fundingID` holds at least `amount` funds.
func (f *Funder) checkFunded(ctx context.Context, amount *big.Int, asset assetHolder, fundingID [32]byte) (bool, error) {
	deposited := make(chan *subscription.Event, funderEventBufSize)
	subErr := make(chan error, 1)
	// Subscribe to events.
	sub, err := f.depositedSub(ctx, asset.contract, fundingID)
	if err != nil {
		return false, errors.WithMessage(err, "subscribing to deposited event")
	}
	defer sub.Close()
	// Read from the sub.
	go func() {
		defer close(deposited)
		subErr <- sub.ReadPast(ctx, deposited)
	}()

	left := new(big.Int).Set(amount)
	for _event := range deposited {
		event, ok := _event.Data.(*assetholder.AssetholderDeposited)
		if !ok {
			log.Panic("wrong event type")
		}
		left.Sub(left, event.Amount)
	}
	return left.Sign() != 1, errors.WithMessagef(<-subErr, "filtering old Funding events for asset %d", asset.assetIndex)
}

func (f *Funder) depositedSub(ctx context.Context, contract *bind.BoundContract, fundingIDs ...[32]byte) (*subscription.ResistantEventSub, error) {
	filter := make([]interface{}, len(fundingIDs))
	for i, fundingID := range fundingIDs {
		filter[i] = fundingID
	}
	event := func() *subscription.Event {
		return &subscription.Event{
			Name:   bindings.Events.AhDeposited,
			Data:   new(assetholder.AssetholderDeposited),
			Filter: [][]interface{}{filter},
		}
	}
	sub, err := subscription.Subscribe(ctx, f, contract, event, startBlockOffset, f.txFinalityDepth)
	return sub, errors.WithMessage(err, "subscribing to deposited event")
}

func (f *Funder) subscribeDeposited(ctx context.Context, contract *bind.BoundContract, fundingIDs ...[32]byte) (chan *subscription.Event, *subscription.ResistantEventSub, chan error, error) {
	deposited := make(chan *subscription.Event)
	subErr := make(chan error, 1)
	// Subscribe to events.
	sub, err := f.depositedSub(ctx, contract, fundingIDs...)
	if err != nil {
		return nil, nil, nil, errors.WithMessage(err, "subscribing to deposited event")
	}
	// Read from the sub.
	go func() {
		subErr <- sub.Read(ctx, deposited)
	}()
	return deposited, sub, subErr, nil
}

// waitForFundingConfirmation waits for the confirmation events on the blockchain that
// both we and all peers successfully funded the channel for the specified asset
// according to the funding agreement.
func (f *Funder) waitForFundingConfirmation(ctx context.Context, request channel.FundingReq, asset assetHolder, fundingIDs [][32]byte) error {
	// If asset on different ledger, return.
	a := request.State.Assets[asset.assetIndex]
	ethAsset, ok := a.(*Asset)
	if !ok {
		return fmt.Errorf("wrong type: expected *Asset, got %T", a)
	}
	if ethAsset.LedgerID().MapKey() != f.chainID.MapKey() {
		return nil
	}

	// Subscribe to events.
	deposited, sub, subErr, err := f.subscribeDeposited(ctx, asset.contract, fundingIDs...)
	if err != nil {
		return errors.WithMessage(err, "subscribing to deposited event")
	}
	defer sub.Close()

	// Wait until funding complete.
	remaining := request.Agreement.Clone()[asset.assetIndex]
	remainingTotal := channel.Balances([][]*big.Int{remaining}).Sum()[0]
	if remainingTotal.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}
loop:
	for {
		select {
		case rawEvent := <-deposited:
			event, ok := rawEvent.Data.(*assetholder.AssetholderDeposited)
			if !ok {
				log.Panic("wrong event type")
			}
			log := f.log.WithField("fundingID", event.FundingID)

			// Subtract amount.
			idx := partIdx(event.FundingID, fundingIDs)
			remainingForPart := remaining[idx]
			remainingForPart.Sub(remainingForPart, event.Amount)
			log.Debugf("peer[%d]: got: %v, remaining for [%d, %d] = %v", request.Idx, event.Amount, asset.assetIndex, idx, remainingForPart)

			// Exit loop if fully funded.
			remainingTotal := channel.Balances([][]*big.Int{remaining}).Sum()[0]
			if remainingTotal.Cmp(big.NewInt(0)) <= 0 {
				break loop
			}
		case <-ctx.Done():
			return fundingTimeoutError(remaining, asset)
		case err := <-subErr:
			// Resolve race between ctx and subErr, as ctx fires both events.
			select {
			case <-ctx.Done():
				return fundingTimeoutError(remaining, asset)
			default:
			}
			return err
		}
	}
	return nil
}

// WaitForOthersFundingConfirmation waits for the confirmation events on the blockchain that
// all peers except oneself (request.Idx) successfully funded the channel for all assets
// according to the funding agreement.
func (f *Funder) WaitForOthersFundingConfirmation(ctx context.Context, request channel.FundingReq, assets []assetHolder, fundingIDs [][32]byte) error {
	totalBalanceForOther := calculateTotalBalances(request)
	for _, asset := range assets {
		// If asset on different ledger, return.
		a := request.State.Assets[asset.assetIndex]
		ethAsset, ok := a.(*Asset)
		if !ok {
			return fmt.Errorf("wrong type: expected *Asset, got %T", a)
		}
		if ethAsset.LedgerID().MapKey() != f.chainID.MapKey() {
			return nil
		}

		// Subscribe to events.
		deposited, sub, subErr, err := f.subscribeDeposited(ctx, asset.contract, fundingIDs...)
		if err != nil {
			return errors.WithMessage(err, "subscribing to deposited event")
		}
		defer sub.Close()

		remainingTotal, remainingOthers := compareBalances(request, asset.assetIndex)
		if remainingTotal.Cmp(big.NewInt(0)) <= 0 {
			continue
		}

		if totalBalanceForOther.Cmp(big.NewInt(0)) <= 0 {
			return nil
		}

		newBalance, err := f.waitForFundingEvents(ctx, deposited, subErr, remainingOthers, totalBalanceForOther, fundingIDs, request, asset)
		if err != nil {
			return err
		}
		totalBalanceForOther = newBalance
	}
	return nil
}

// waitForFundingEvents waits for the confirmation events and returns the updated balance.
func (f *Funder) waitForFundingEvents(ctx context.Context, deposited <-chan *subscription.Event, subErr <-chan error, remainingOthers []*big.Int, totalBalanceForOther *big.Int, fundingIDs [][32]byte, request channel.FundingReq, asset assetHolder) (*big.Int, error) {
loop:
	for {
		select {
		case rawEvent := <-deposited:
			event, ok := rawEvent.Data.(*assetholder.AssetholderDeposited)
			if !ok {
				log.Panic("wrong event type")
			}

			idx := partIdx(event.FundingID, fundingIDs)
			// Ignore if the current participant should have deposited.
			if channel.Index(idx) != request.Idx {
				totalBalanceForOther.Sub(totalBalanceForOther, event.Amount)
			}
			if totalBalanceForOther.Cmp(big.NewInt(0)) <= 0 {
				break loop
			}
		case <-ctx.Done():
			return totalBalanceForOther, fundingTimeoutError(remainingOthers, asset)
		case err := <-subErr:
			select {
			case <-ctx.Done():
				return totalBalanceForOther, fundingTimeoutError(remainingOthers, asset)
			default:
			}
			return totalBalanceForOther, err
		}
	}
	return totalBalanceForOther, nil
}

func fundingTimeoutError(remaining []channel.Bal, asset assetHolder) error {
	var indices []channel.Index
	for k, bals := range remaining {
		if bals.Sign() == 1 {
			indices = append(indices, channel.Index(k))
		}
	}
	if len(indices) != 0 {
		return &channel.AssetFundingError{Asset: asset.assetIndex, TimedOutPeers: indices}
	}
	return nil
}

func partIdx(partID [32]byte, fundingIDs [][32]byte) int {
	for i, id := range fundingIDs {
		if id == partID {
			return i
		}
	}
	return -1
}

// FundingIDs returns a slice the same size as the number of passed participants
// where each entry contains the hash Keccak256(channel id || participant address).
func FundingIDs(channelID channel.ID, participants ...map[perunwallet.BackendID]perunwallet.Address) [][32]byte {
	ids := make([][32]byte, len(participants))
	for idx, pID := range participants {
		address, ok := pID[wallet.BackendID].(*wallet.Address)
		if !ok {
			log.Panic("wrong address type")
		}
		ids[idx] = FundingID(channelID, address)
	}
	return ids
}

// FundingID returns the funding identifier for a participant, i.e.,
// Keccak256(channel id || participant address).
func FundingID(channelID channel.ID, participant perunwallet.Address) [32]byte {
	args := abi.Arguments{{Type: abiBytes32}, {Type: abiAddress}}
	address, ok := participant.(*wallet.Address)
	if !ok {
		log.Panic("wrong address type")
	}
	bytes, err := args.Pack(channelID, common.Address(*address))
	if err != nil {
		log.Panicf("error packing values: %v", err)
	}
	return crypto.Keccak256Hash(bytes)
}

// NumTX returns how many Transactions are needed for the funding request.
func (f *Funder) NumTX(req channel.FundingReq) (sum uint32, err error) {
	f.mtx.RLock()
	defer f.mtx.RUnlock()

	for _, a := range req.State.Assets {
		depositor, ok := f.depositors[a.(*Asset).MapKey()]
		if !ok {
			return 0, errors.Errorf("could not find Depositor for asset #%d", a)
		}
		sum += depositor.NumTX()
	}
	return
}

// checkEgoisticPart checks if more than one entries are set to true. If so it returns an error.
func checkEgoisticPart(egoisticPart []bool) error {
	trueCount := 0
	for _, v := range egoisticPart {
		if v {
			trueCount++
			if trueCount > 1 {
				return errors.New("more than one entry is true")
			}
		}
	}
	return nil
}

// calculateTotalBalances calculates the total balance for other participants.
func calculateTotalBalances(request channel.FundingReq) *big.Int {
	totalBalanceForOther := big.NewInt(0)
	// Iterate over each asset to sum up the total balance for other participants.
	for _, asset := range request.Agreement {
		for i, bal := range asset {
			if channel.Index(i) != request.Idx {
				totalBalanceForOther.Add(totalBalanceForOther, bal)
			}
		}
	}
	return totalBalanceForOther
}

// compareBalances creates a slice from the balances without the current participant and returns it and the total sum over it.
func compareBalances(request channel.FundingReq, assetIndex channel.Index) (*big.Int, []*big.Int) {
	remainingAll := request.Agreement.Clone()[assetIndex]
	// Create a new remainingOthers slice excluding the balance of the current participant.
	remainingOthers := make([]*big.Int, 0, len(remainingAll)-1)
	for i, bal := range remainingAll {
		if channel.Index(i) != request.Idx {
			remainingOthers = append(remainingOthers, bal)
		}
	}
	remainingTotal := channel.Balances([][]*big.Int{remainingOthers}).Sum()[0]
	return remainingTotal, remainingOthers
}
