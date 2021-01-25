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

package functions

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadWithoutRequiredArguments(t *testing.T)  {
	args := []string{
		"upload",
	}
	_, _, err := TestFunctionsCommands(uploadFunctionsCmd, args)
	assert.NotNil(t, err)
}

func TestUploadWithEmptyArguments(t *testing.T)  {
	args := []string{
		"upload",
		"--path", "    ",
		"--source-file", "    ",
	}
	_, execErr, err := TestFunctionsCommands(uploadFunctionsCmd, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.True(t, strings.Contains(execErr.Error(), "empty"))
}

func TestUploadAndDownloadCommands(t *testing.T)  {
	testFile := "upload.go"
	fileHash, err := getFileSha256(testFile)
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}

	pulsarPath := "public/default/test"
	args := []string{
		"upload",
		"--path", pulsarPath,
		"--source-file", testFile,
	}
	out, execErr, err := TestFunctionsCommands(uploadFunctionsCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "successfully"))

	downloadFilePath := "download-upload-file"
	args = []string{
		"download",
		"--destination-file", pulsarPath,
		"--path", downloadFilePath,
	}
	out, execErr, err = TestFunctionsCommands(downloadFunctionsCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "successfully"))

	downloadFileSha, err := getFileSha256(downloadFilePath)
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	assert.Equal(t, fileHash, downloadFileSha)
}

func getFileSha256(filename string) (string, error) {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	hasher.Write(s)
	fileHash := hex.EncodeToString(hasher.Sum(nil))
	return fileHash, err
}