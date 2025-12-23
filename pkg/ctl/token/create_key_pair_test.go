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
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	signatureAlgorithm    string
	outputPrivateFilePath string
	outputPublicFilePath  string
}{
	{"RS256", "rs256-private.key", "rs256-public.key"},
	{"RS384", "rs384-private.key", "rs384-public.key"},
	{"RS512", "rs512-private.key", "rs512-public.key"},
	{"ES256", "es256-private.key", "es256-public.key"},
	{"ES384", "es384-private.key", "es384-public.key"},
	{"ES512", "es512-private.key", "es512-public.key"},
	{"INVALID", "private.key", "public.key"},
}

func TestCreateKeyPair(t *testing.T) {
	for _, data := range testData {
		args := []string{"create-key-pair", "--signature-algorithm", data.signatureAlgorithm,
			"--output-private-key", data.outputPrivateFilePath, "--output-public-key", data.outputPublicFilePath}
		out, execErr, err := testTokenCommands(createKeyPair, args)
		assert.Nil(t, err)
		if data.signatureAlgorithm == "INVALID" {
			assert.NotNil(t, execErr)
			assert.Equal(t,
				fmt.Sprintf("the signature algorithm '%s' is invalid. Valid options are: "+
					"'HS256', 'HS384', 'HS512', 'RS256', 'RS384', 'RS512', 'ES256', 'ES384', 'ES512'\n",
					data.signatureAlgorithm),
				execErr.Error())
			continue
		}

		assert.Nil(t, execErr)
		assert.Equal(t,
			fmt.Sprintf("The private key and public key are generated to %s and %s successfully.\n",
				data.outputPrivateFilePath, data.outputPublicFilePath),
			out.String())
		_ = os.Remove(data.outputPrivateFilePath)
		_ = os.Remove(data.outputPublicFilePath)
	}
}

func TestSpaceInOutputFileParams(t *testing.T) {
	args := []string{"create-key-pair", "--signature-algorithm", "RS256",
		"--output-private-key", " ", "--output-public-key", " "}
	_, execErr, err := testTokenCommands(createKeyPair, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the private key file path and the public key file path can not be empty", execErr.Error())
}
