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
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadWithoutRequiredArguments(t *testing.T) {
	args := []string{
		"upload",
	}
	_, _, err := TestFunctionsCommands(uploadFunctionsCmd, args)
	assert.NotNil(t, err)
}

func TestUploadWithEmptyArguments(t *testing.T) {
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

const fileContent = `
#!/usr/bin/env bash
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

BINDIR=$(dirname "$0")
export PULSAR_HOME=$(cd -P $BINDIR/..;pwd)
. "$PULSAR_HOME/bin/pulsar-admin-common.sh"

#Change to PULSAR_HOME to support relative paths
cd "$PULSAR_HOME"
exec $JAVA $OPTS org.apache.pulsar.admin.cli.PulsarAdminTool $PULSAR_CLIENT_CONF "$@"
`

func TestUploadAndDownloadCommands(t *testing.T) {
	f, err := os.CreateTemp(".", "test")
	if err != nil {
		log.Fatal(err)
		t.Fail()
		return
	}
	defer os.RemoveAll(f.Name())
	err = os.WriteFile(f.Name(), []byte(fileContent), os.ModePerm)
	if err != nil {
		log.Panic(err)
		t.Fail()
		return
	}

	testFile := f.Name()
	fileHash, err := getFileSha256(testFile)
	if err != nil {
		log.Panic(err)
		t.Fail()
	}

	pulsarPath := "public/default/test"
	args := []string{
		"upload",
		"--path", pulsarPath,
		"--source-file", testFile,
	}
	out, execErr, err := TestFunctionsCommands(uploadFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	assert.True(t, strings.Contains(out.String(), "successfully"))

	downloadFilePath := "download-upload-file"
	args = []string{
		"download",
		"--destination-file", downloadFilePath,
		"--path", pulsarPath,
	}
	out, execErr, err = TestFunctionsCommands(downloadFunctionsCmd, args)
	defer os.RemoveAll(downloadFilePath)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	assert.True(t, strings.Contains(out.String(), "successfully"))

	downloadFileSha, err := getFileSha256(downloadFilePath)
	if err != nil {
		log.Panic(err)
		t.Fail()
	}
	assert.Equal(t, fileHash, downloadFileSha)
}

func getFileSha256(filename string) (string, error) {
	hasher := sha256.New()
	s, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	hasher.Write(s)
	fileHash := hex.EncodeToString(hasher.Sum(nil))
	return fileHash, err
}
