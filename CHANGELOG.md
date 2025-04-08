# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.6.0] Bromo - 2025-04-08 [:boom:]
Support for EVM to Non-EVM payment channels and other fixes.

### Added
- Egoistic Funding Test [#52]
- Support for EVM to Non-EVM payment channels including go-perun version update and new contract binaries [#54] [:boom:]

### Changed
- Update action cache [#55]

### Fixed
- Fixed AppID handling [#53]

[#52]: https://github.com/hyperledger-labs/perun-eth-backend/pull/52
[#53]: https://github.com/hyperledger-labs/perun-eth-backend/pull/53
[#54]: https://github.com/hyperledger-labs/perun-eth-backend/pull/54
[#55]: https://github.com/hyperledger-labs/perun-eth-backend/pull/55

## [0.5.0] Athos - 2023-02-22 [:boom:]
Improved Contract Compatibility, Swap Parallelization and other fixes.

### Added [:boom:]
- AppID Type to generalize App identifiers instead of using Ethereum addresses: [#44] [:boom:]
- Add a shared nonce map to ensure that the nonce is incremented for all transactions, including those sent in parallel: [#43]
- Add ability to set the gas limits instead of using fixed values: [#46]
- Add pointer receiver to Asset.UnmarshallBinary to return the unmarshalled value instead of the default: [#36]

### Fixed
- Fix the order from Asset.MarshalBinary: [#36]
- Restore compatibility between go-perun post-0.10.6 and eth-backend: [#44]

### Changed
- Use git diff instead of go test to check if the bindings are generated correctly: [#33]
- Remove IncreaseAllowance and use Approve instead to ensure compatibility with more contracts in the depositing process: [#43]
- Include secondary settling to reduce contract calls: [#47]

[#33]: https://github.com/hyperledger-labs/perun-eth-backend/pull/33
[#36]: https://github.com/hyperledger-labs/perun-eth-backend/pull/36
[#43]: https://github.com/hyperledger-labs/perun-eth-backend/pull/43
[#44]: https://github.com/hyperledger-labs/perun-eth-backend/pull/44
[#46]: https://github.com/hyperledger-labs/perun-eth-backend/pull/46
[#47]: https://github.com/hyperledger-labs/perun-eth-backend/pull/47

## [0.1.0] - 2022-08.26
### Added
- Compatibility with go-perun version 0.10.7
- Add Virtual channel test: [#32]
- Add test for fund recovery: [#29]
- Add tests for Subchannels: [#21]
- Add support for multi-backends: [#19]

### Fixed
- Address consistency.
- Fix multi ledger event subscription for conclude and deposit: [#30]

### Changed
- Get nonce from auth.Form for Transactors: [#31]

[#3]: https://github.com/hyperledger-labs/perun-eth-backend/pull/3
[#18]: https://github.com/hyperledger-labs/perun-eth-backend/pull/18
[#19]: https://github.com/hyperledger-labs/perun-eth-backend/pull/19
[#21]: https://github.com/hyperledger-labs/perun-eth-backend/pull/21
[#26]: https://github.com/hyperledger-labs/perun-eth-backend/pull/26
[#29]: https://github.com/hyperledger-labs/perun-eth-backend/pull/29
[#30]: https://github.com/hyperledger-labs/perun-eth-backend/pull/30
[#31]: https://github.com/hyperledger-labs/perun-eth-backend/pull/31
[#32]: https://github.com/hyperledger-labs/perun-eth-backend/pull/32

## Legend
- <span id="warning">:warning:</span> This is a pre-release and not intended for usage with real funds.
- <span id="breaking">:boom:</span> This is a breaking change, e.g., it changes the external API.

[:warning:]: #warning
[:boom:]: #breaking