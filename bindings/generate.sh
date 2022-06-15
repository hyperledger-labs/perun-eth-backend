#!/usr/bin/env sh

# Copyright 2021 - See NOTICE file for copyright holders.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

# Define ABIGEN and SOLC default values.
ABIGEN="${ABIGEN-abigen}"
SOLC="${SOLC-solc}"

echo 'Please ensure that solc v0.7.6+ and abigen v1.10.18+ are installed.'

if ! $ABIGEN --version
then
    echo "'abigen' not found. Please add to PATH or set ABIGEN='path_to_abigen'."
    exit 1
fi

if ! $SOLC --version
then
    echo "'solc' not found. Please add to PATH or set SOLC='path_to_solc'."
    exit 1
fi

echo "Please ensure that the repository was cloned with submodules: 'git submodule update --init --recursive'."

# Generates optimized golang bindings and runtime binaries for sol contracts.
# $1  solidity file path, relative to ../contracts/contracts/.
# $2  golang package name.
generate() {
    FILE=$1; PKG=$2; CONTRACT=$FILE
    echo "Generating $PKG bindings..."

    rm -r $PKG
    mkdir $PKG

    # Compile and generate binary runtime.
    $SOLC --abi --bin --bin-runtime --optimize --allow-paths contracts/vendor, contracts/contracts/$FILE.sol -o $PKG/
    BIN_RUNTIME=$(cat ${PKG}/${CONTRACT}.bin-runtime)
    OUT_FILE="$PKG/${CONTRACT}BinRuntime.go"
    echo "package $PKG // import \"github.com/perun-network/perun-eth-backend/bindings/$PKG\"" > $OUT_FILE
    echo >> $OUT_FILE
    echo "// ${CONTRACT}BinRuntime is the runtime part of the compiled bytecode used for deploying new contracts." >> $OUT_FILE
    echo "var ${CONTRACT}BinRuntime = \"$BIN_RUNTIME\"" >> $OUT_FILE

    # Generate bindings.
    $ABIGEN --pkg $PKG --abi $PKG/$FILE.abi --bin $PKG/$FILE.bin --out $PKG/$FILE.go
}

# Adjudicator
generate "Adjudicator" "adjudicator"

# Asset holders
generate "AssetHolder" "assetholder"
generate "AssetHolderETH" "assetholdereth"
generate "AssetHolderERC20" "assetholdererc20"

# Tokens
generate "PerunToken" "peruntoken"

# Applications
generate "TrivialApp" "trivialapp"

echo "Bindings generated successfully."
