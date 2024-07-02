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

package wire_test

import (
	"math/rand"
	"testing"

	"github.com/perun-network/perun-eth-backend/wire"
	"github.com/stretchr/testify/assert"
	perunwire "perun.network/go-perun/wire"
	"perun.network/go-perun/wire/test"
	pkgtest "polycry.pt/poly-go/test"
)

var dataToSign = []byte("SomeLongDataThatShouldBeSignedPlease")

func TestAddress(t *testing.T) {
	test.TestAddressImplementation(t, func() perunwire.Address {
		return wire.NewAddress()
	}, func(rng *rand.Rand) perunwire.Address {
		return wire.NewRandomAddress(rng)
	})
}

func TestSignatures_Success(t *testing.T) {
	acc := wire.NewRandomAccount(pkgtest.Prng(t))
	sig, err := acc.Sign(dataToSign)
	assert.NoError(t, err, "Sign with new account should succeed")
	assert.NotNil(t, sig)
	assert.Equal(t, len(sig), wire.SigLen, "Ethereum signature has wrong length")
	err = acc.Address().Verify(dataToSign, sig)
	assert.NoError(t, err, "Verification should succeed")
}

func TestSignatures_ModifyData_Failure(t *testing.T) {
	acc := wire.NewRandomAccount(pkgtest.Prng(t))
	sig, err := acc.Sign(dataToSign)
	assert.NoError(t, err, "Sign with new account should succeed")
	assert.NotNil(t, sig)

	// Modify a single byte of the signed data
	modifiedData := make([]byte, len(dataToSign))
	copy(modifiedData, dataToSign)
	modifiedData[0] ^= 0x01

	err = acc.Address().Verify(modifiedData, sig)
	assert.Error(t, err, "Verification should fail with modified data")
}

func TestSignatures_ModifySignature_Failure(t *testing.T) {
	acc := wire.NewRandomAccount(pkgtest.Prng(t))
	sig, err := acc.Sign(dataToSign)
	assert.NoError(t, err, "Sign with new account should succeed")
	assert.NotNil(t, sig)

	// Modify a single byte of the signature (first 64 bytes)
	modifiedSig := make([]byte, len(sig))
	copy(modifiedSig, sig)
	modifiedSig[0] ^= 0x01

	err = acc.Address().Verify(dataToSign, modifiedSig)
	assert.Error(t, err, "Verification should fail with modified signature")
}

func TestSignatures_ModifyLastByteOfSignature_Failure(t *testing.T) {
	acc := wire.NewRandomAccount(pkgtest.Prng(t))
	sig, err := acc.Sign(dataToSign)
	assert.NoError(t, err, "Sign with new account should succeed")
	assert.NotNil(t, sig)

	// Modify the last byte of the signature
	modifiedSig := make([]byte, len(sig))
	copy(modifiedSig, sig)
	modifiedSig[len(sig)-1] ^= 0x01

	err = acc.Address().Verify(dataToSign, modifiedSig)
	assert.Error(t, err, "Verification should fail with modified signature")
}

func TestSignatures_WrongAccount_Failure(t *testing.T) {
	acc := wire.NewRandomAccount(pkgtest.Prng(t))
	sig, err := acc.Sign(dataToSign)
	assert.NoError(t, err, "Sign with new account should succeed")
	assert.NotNil(t, sig)

	// Verify with a wrong account
	wrongAcc := wire.NewRandomAccount(pkgtest.Prng(t))
	err = wrongAcc.Address().Verify(dataToSign, sig)
	assert.Error(t, err, "Verification should fail with wrong account")
}
