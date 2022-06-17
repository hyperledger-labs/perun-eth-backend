// Copyright 2021 - See NOTICE file for copyright holders.
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

package bindings

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	"github.com/perun-network/perun-eth-backend/bindings/assetholder"
	"github.com/perun-network/perun-eth-backend/bindings/assetholdererc20"
	"github.com/perun-network/perun-eth-backend/bindings/assetholdereth"
	"github.com/perun-network/perun-eth-backend/bindings/peruntoken"
	"github.com/perun-network/perun-eth-backend/bindings/trivialapp"
)

// ABI contains all the parsed ABI definitions of our contracts.
// Use it together with `bind.NewBoundContract` to create a bound contract.
var ABI = struct {
	// PerunToken is the parsed ABI definition of contract PerunToken.
	PerunToken abi.ABI
	// Adjudicator is the parsed ABI definition of contract Adjudicator.
	Adjudicator abi.ABI
	// AssetHolder is the parsed ABI definition of contract AssetHolder.
	AssetHolder abi.ABI
	// ETHAssetHolder is the parsed ABI definition of contract ETHAssetHolder.
	ETHAssetHolder abi.ABI
	// ERC20AssetHolder is the parsed ABI definition of contract ERC20AssetHolder.
	ERC20AssetHolder abi.ABI
	// TrivialApp is the parsed ABI definition of contract TrivialApp.
	TrivialApp abi.ABI
}{}

// Events contains the event names for specific events.
var Events = struct {
	// AdjChannelUpdate is the ChannelUpdate event of the Adjudicator contract.
	AdjChannelUpdate string
	// AhDeposited is the Deposited event of the Assetholder contract.
	AhDeposited string
	// AhWithdrawn is the Withdrawn event of the Assetholder contract.
	AhWithdrawn string
	// PerunTokenApproval is the Approval event of the PerunToken contract.
	PerunTokenApproval string
}{}

func init() {
	parseABIs()
	extractEvents()
}

func parseABIs() {
	parse := func(raw string) abi.ABI {
		abi, err := abi.JSON(strings.NewReader(raw))
		if err != nil {
			panic(err)
		}
		return abi
	}

	ABI.PerunToken = parse(peruntoken.PeruntokenMetaData.ABI)
	ABI.Adjudicator = parse(adjudicator.AdjudicatorMetaData.ABI)
	ABI.AssetHolder = parse(assetholder.AssetholderMetaData.ABI)
	ABI.ETHAssetHolder = parse(assetholdereth.AssetholderethMetaData.ABI)
	ABI.ERC20AssetHolder = parse(assetholdererc20.Assetholdererc20MetaData.ABI)
	ABI.TrivialApp = parse(trivialapp.TrivialappMetaData.ABI)
}

// extractEvents sets the event names and panics if any event does not exist.
func extractEvents() {
	extract := func(abi abi.ABI, eName string) string {
		e, ok := abi.Events[eName]
		if !ok {
			panic(fmt.Sprintf("Event '%s' not found.", eName))
		}
		return e.Name
	}

	Events.AdjChannelUpdate = extract(ABI.Adjudicator, "ChannelUpdate")
	Events.AhDeposited = extract(ABI.AssetHolder, "Deposited")
	Events.AhWithdrawn = extract(ABI.AssetHolder, "Withdrawn")
	Events.PerunTokenApproval = extract(ABI.PerunToken, "Approval")
}
