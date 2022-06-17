# Perun's Ethereum Backend

<p>
  <a href="https://www.apache.org/licenses/LICENSE-2.0.txt"><img src="https://img.shields.io/badge/license-Apache%202-blue" alt="License: Apache 2.0"></a>
  <a href="https://github.com/hyperledger-labs/perun-eth-backend/actions/workflows/ci.yml"><img src="https://github.com/hyperledger-labs/perun-eth-backend/actions/workflows/ci.yml/badge.svg?branch=main" alt="CI status"></a>
  <a href="https://pkg.go.dev/github.com/perun-network/perun-eth-backend"><img src="https://pkg.go.dev/badge/github.com/perun-network/perun-eth-backend.svg" alt="Go Reference"></a>
</p>

This repository provides an [Ethereum] blockchain module for the [go-perun] state channel library.
It thereby enables Perun channels for EVM-compatible networks.

## Project structure
* `bindings/`: Contract bindings.
* `channel/`: Channel interface implementations.
* `client/`: Client tests.
* `wallet/`: Wallet interface implementations.

## Development

1. Clone the repository.
```sh
git clone https://github.com/perun-network/perun-eth-backend
cd perun-eth-backend
```

2. Run the tests. This step needs a working [Go distribution](https://golang.org), see [go.mod](go.mod) for the required version.

```sh
go test ./...
```

## Demo

The [perun-eth-demo] project demonstrates Perun payment channels on Ethereum.

## Security Disclaimer

The authors take no responsibility for any loss of digital assets or other damage caused by the use of this software.

## Copyright

Copyright 2022 PolyCrypt GmbH.  
Use of the source code is governed by the Apache 2.0 license that can be found in the [LICENSE file](LICENSE).

<!--- Links -->

[Ethereum]: https://ethereum.org/
[go-perun]: https://github.com/hyperledger-labs/go-perun
[perun-eth-contracts]: https://github.com/hyperledger-labs/perun-eth-contracts
[perun-eth-demo]: https://github.com/perun-network/perun-eth-demo
