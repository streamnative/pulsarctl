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
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	InvalidAlgorithm   bool
	encode             bool
	SignatureAlgorithm string
	outputFile         string
}{
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS256", encode: false, outputFile: ""},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS256", encode: true, outputFile: ""},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS256", encode: false, outputFile: "test-HS256-secret.key"},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS256", encode: true, outputFile: "test-HS256-secret.key"},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS384", encode: false, outputFile: ""},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS384", encode: true, outputFile: ""},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS384", encode: false, outputFile: "test-HS384-secret.key"},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS384", encode: true, outputFile: "test-HS384-secret.key"},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS512", encode: false, outputFile: ""},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS512", encode: true, outputFile: ""},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS512", encode: false, outputFile: "test-HS512-secret.key"},
	{InvalidAlgorithm: false, SignatureAlgorithm: "HS512", encode: true, outputFile: "test-HS512-secret.key"},
	{InvalidAlgorithm: true, SignatureAlgorithm: "INVALID", encode: false, outputFile: ""},
}

func TestCreateSecretKeyCommand(t *testing.T) {
	for _, data := range testData {
		if data.InvalidAlgorithm {
			testInvalidError(t, data.SignatureAlgorithm)
			continue
		}
		log.Printf("test case: %+v", data)
		testNormalCase(t, data.SignatureAlgorithm, data.outputFile, data.encode)
	}
}

func testNormalCase(t *testing.T, signatureAlgorithm, outputFile string, encode bool) {
	args := []string{"create-secret-key", "--signature-algorithm", signatureAlgorithm}
	if encode {
		args = append(args, "--base64")
	}
	if outputFile != "" {
		args = append(args, "--output-file", outputFile)
		defer os.Remove(outputFile)
	}

	out, execErr, _ := testTokenCommands(createSecretKey, args)
	assert.Nil(t, execErr)

	if outputFile != "" {
		assert.Equal(t,
			fmt.Sprintf("Write the secret key to the file %s successfully.\n", outputFile),
			out.String())
		return
	}

	var output []byte
	if encode {
		output, _ = base64.StdEncoding.DecodeString(out.String())
	} else {
		output = out.Bytes()[:len(out.Bytes())-1]
	}

	switch {
	case signatureAlgorithm == "HS256":
		assert.Equal(t, 32, len(output))
	case signatureAlgorithm == "HS384":
		assert.Equal(t, 48, len(output))
	case signatureAlgorithm == "HS512":
		assert.Equal(t, 64, len(output))
	default:
		log.Fatal("Test error")
	}
}

func testInvalidError(t *testing.T, signatureAlgorithm string) {
	args := []string{"create-secret-key", "--signature-algorithm", signatureAlgorithm}
	_, execErr, _ := testTokenCommands(createSecretKey, args)
	assert.NotNil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("the signature algorithm '%s' is invalid. Valid options are: 'HS256', "+
			"'HS384', 'HS512'\n", signatureAlgorithm),
		execErr.Error())
}
