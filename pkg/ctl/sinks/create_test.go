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

package sinks

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSinkFailedByEmptyFile(t *testing.T) {
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	failedArgs := []string{
		"create",
		"--name", "failed-create",
		"--inputs", "failed-inputs",
		"--archive", file.Name(),
	}
	_, execErr, err := TestSinksCommands(createSinksCmd, failedArgs)
	failImmediatelyIfErrorNotNil(t, err)
	assert.NotNil(t, execErr)
}

func TestCreateSinkFailedByNotExistSinkConfigFile(t *testing.T) {
	failedArgs := []string{
		"create",
		"--name", "failed-create",
		"--inputs", "failed-inputs",
		"--sink-config-file", "/tmp/failed-config.yaml",
	}
	_, execErr, err := TestSinksCommands(createSinksCmd, failedArgs)
	failImmediatelyIfErrorNotNil(t, err)
	assert.NotNil(t, execErr)
}
