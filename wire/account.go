// Copyright 2024 - See NOTICE file for copyright holders.
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

package wire

import (
	"crypto/ecdsa"
	"log"
	"math/rand"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/perun-network/perun-eth-backend/wallet"
	"github.com/pkg/errors"
	"perun.network/go-perun/wire"
)

// SigLen length of a signature in byte.
// ref https://godoc.org/github.com/ethereum/go-ethereum/crypto/secp256k1#Sign
// ref https://github.com/ethereum/go-ethereum/blob/54b271a86dd748f3b0bcebeaf678dc34e0d6177a/crypto/signature_cgo.go#L66
const SigLen = 65

// sigVSubtract value that is subtracted from the last byte of a signature if
// the last bytes exceeds it.
const sigVSubtract = 27

// Account is a wire account.
type Account struct {
	addr *Address
	key  *ecdsa.PrivateKey
}

// Sign signs the given message with the account's private key.
func (acc *Account) Sign(data []byte) ([]byte, error) {
	hash := PrefixedHash(data)
	sig, err := crypto.Sign(hash, acc.key)
	if err != nil {
		return nil, errors.Wrap(err, "SignHash")
	}
	sig[64] += 27
	return sig, nil
}

// Address returns the account's address.
func (acc *Account) Address() wire.Address {
	return acc.addr
}

// NewRandomAccount generates a new random account.
func NewRandomAccount(rng *rand.Rand) *Account {
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rng)
	if err != nil {
		panic(err)
	}
	log.Print("Generated new account with address ", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())

	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &Account{
		addr: &Address{wallet.AsWalletAddr(addr)},
		key:  privateKey,
	}
}
