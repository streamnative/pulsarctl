// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package token

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var algorithmList = []string{
	"HS256",
	"HS384",
	"HS512",
	"RS256",
	"RS384",
	"RS512",
	"ES256",
	"ES384",
	"ES512",
	"PS256",
	"PS384",
	"PS512",
}

var keyFiles = []string{
	"../../../test/key/pulsar-admin-hs256-secret.key",
	"../../../test/key/pulsar-admin-hs256-base64-secret.key",
	"../../../test/key/pulsarctl-hs256-secret.key",
	"../../../test/key/pulsarctl-hs256-base64-secret.key",

	"../../../test/key/pulsar-admin-hs384-secret.key",
	"../../../test/key/pulsar-admin-hs384-base64-secret.key",
	"../../../test/key/pulsarctl-hs384-secret.key",
	"../../../test/key/pulsarctl-hs384-base64-secret.key",

	"../../../test/key/pulsar-admin-hs512-secret.key",
	"../../../test/key/pulsar-admin-hs512-base64-secret.key",
	"../../../test/key/pulsarctl-hs512-secret.key",
	"../../../test/key/pulsarctl-hs512-base64-secret.key",

	"../../../test/key/pulsar-admin-rs256-private.key",
	"../../../test/key/pulsar-admin-rs384-private.key",
	"../../../test/key/pulsar-admin-rs512-private.key",

	"../../../test/key/pulsarctl-rs256-private.key",
	"../../../test/key/pulsarctl-rs384-private.key",
	"../../../test/key/pulsarctl-rs512-private.key",

	"../../../test/key/pulsar-admin-es256-private.key",
	"../../../test/key/pulsar-admin-es384-private.key",
	"../../../test/key/pulsar-admin-es512-private.key",

	"../../../test/key/pulsarctl-es256-private.key",
	"../../../test/key/pulsarctl-es384-private.key",
	"../../../test/key/pulsarctl-es512-private.key",
}

func TestCreateTokenWithSecretKeyFileCmd(t *testing.T) {
	for _, sa := range algorithmList {
		for _, keyFile := range keyFiles {
			t.Logf("Signature algorithm: %s, key file: %s", sa, keyFile)
			if strings.Contains(keyFile, "secret") {
				doTestCreateTokenWithSecretKey(t, sa, keyFile)
			} else {
				doTestCreateTokenWithPrivateKey(t, sa, keyFile)
			}
		}
	}
}

func doTestCreateTokenWithSecretKey(t *testing.T, signatureAlgorithm, secretKeyFile string) {
	args := []string{
		"create",
		"--signature-algorithm", signatureAlgorithm,
		"--secret-key-file", secretKeyFile,
		"--subject", "test-create-with-secret-key-file",
		"--expire", "1s",
	}

	if strings.Contains(secretKeyFile, "base") {
		args = append(args, "--base64")
	}

	out, execErr, err := testTokenCommands(create, args)
	assert.Nil(t, err)
	switch {
	case strings.HasPrefix(signatureAlgorithm, "HS"):
		assert.Nil(t, execErr)
		t.Log(out.String())
	case strings.HasPrefix(signatureAlgorithm, "RS"):
		assert.NotNil(t, execErr)
		assert.Equal(t, "key is invalid", execErr.Error())
	case strings.HasPrefix(signatureAlgorithm, "ES"):
		fallthrough
	case strings.HasPrefix(signatureAlgorithm, "PS"):
		assert.NotNil(t, execErr)
		assert.Equal(t, "key is of invalid type", execErr.Error())
	default:
		t.Fatal("invalid case for testing create token with secret key")
	}
}

func doTestCreateTokenWithPrivateKey(t *testing.T, signatureAlgorithm, privateKeyFile string) {
	args := []string{
		"create",
		"--signature-algorithm", signatureAlgorithm,
		"--private-key-file", privateKeyFile,
		"--subject", "test-create-with-private-key-file",
		"--expire", "1s",
	}

	out, execErr, err := testTokenCommands(create, args)
	assert.Nil(t, err)
	switch {
	case strings.HasPrefix(signatureAlgorithm, "HS"):
		assert.NotNil(t, execErr)
		assert.Equal(t, "invalid type of the signature algorithm", execErr.Error())
	case strings.HasPrefix(signatureAlgorithm, "RS") && strings.Contains(privateKeyFile, "rs"):
		fallthrough
	case strings.HasPrefix(signatureAlgorithm, "PS") && strings.Contains(privateKeyFile, "ps"):
		assert.Nil(t, execErr)
		t.Log(out.String())
	case strings.HasPrefix(signatureAlgorithm, "ES256") && strings.Contains(privateKeyFile, "es256"):
		fallthrough
	case strings.HasPrefix(signatureAlgorithm, "ES384") && strings.Contains(privateKeyFile, "es384"):
		fallthrough
	case strings.HasPrefix(signatureAlgorithm, "ES512") && strings.Contains(privateKeyFile, "es512"):
		assert.Nil(t, execErr)
		t.Logf(out.String())
	default:
		t.Logf("invalid case for testing create token with private key")
	}
}

func TestNoKeySpecifiedErr(t *testing.T) {
	args := []string{"create", "--subject", "subject"}
	_, execErr, err := testTokenCommands(create, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errNoKeySpecified.Error(), execErr.Error())
}

func TestKeySpecifiedMoreThanOneErr(t *testing.T) {
	args := []string{"create", "--secret-key-string", "secret-key", "--private-key-file", "private-key",
		"--subject", "subject"}
	_, execErr, err := testTokenCommands(create, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errKeySpecifiedMoreThanOne.Error(), execErr.Error())
}

func TestTrimSpaceForCreadCmdArgs(t *testing.T) {
	args := []string{"create", "--secret-key-string", "   ", "--subject", "   "}
	_, execErr, err := testTokenCommands(create, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errNoKeySpecified.Error(), execErr.Error())
}
