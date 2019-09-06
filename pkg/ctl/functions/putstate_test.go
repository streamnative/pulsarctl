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
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestStateFunctions(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-putstate",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", basePath + "/test/functions/api-examples.jar",
	}

	out, _, err := TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, out.String(), "Created test-functions-putstate successfully")

	// wait the function create successfully
	time.Sleep(time.Second * 5)

	putstateArgs := []string{"putstate",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-putstate",
		"--state", "{\"key\":\"pulsar\", \"stringValue\":\"hello\"}",
	}

	outPutState, _, err := TestFunctionsCommands(putstateFunctionsCmd, putstateArgs)
	assert.Nil(t, err)
	t.Logf("outPutState:%s", outPutState.String())
	assert.Equal(t, outPutState.String(), "PutState successfully")

	// test failure case for put state
	failureStateArgs := []string{"putstate",
		"--name", "not-exist",
		"--state", "{\"key\":\"pulsar\", \"stringValue\":\"hello\"}",
	}

	_, execErrMsg, _ := TestFunctionsCommands(putstateFunctionsCmd, failureStateArgs)
	assert.NotNil(t, execErrMsg)
	exceptMsg := "'not-exist' is not found"
	assert.True(t, strings.Contains(execErrMsg.Error(), exceptMsg))

	// query state
	queryStateArgs := []string{"querystate",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-putstate",
		"--key", "pulsar",
	}

	outQueryState, _, err := TestFunctionsCommands(querystateFunctionsCmd, queryStateArgs)
	assert.Nil(t, err)
	t.Logf("outQueryState:%s", outQueryState.String())

	var state pulsar.FunctionState
	err = json.Unmarshal(outQueryState.Bytes(), &state)
	assert.Nil(t, err)

	assert.Equal(t, "pulsar", state.Key)
	assert.Equal(t, "hello", state.StringValue)
	assert.Equal(t, int64(0), state.Version)
	// put state again
	outPutStateAgain, _, err := TestFunctionsCommands(putstateFunctionsCmd, putstateArgs)
	assert.Nil(t, err)
	assert.Equal(t, outPutStateAgain.String(), "PutState successfully")

	// query state again
	outQueryStateAgain, _, err := TestFunctionsCommands(querystateFunctionsCmd, queryStateArgs)
	assert.Nil(t, err)

	var stateAgain pulsar.FunctionState
	err = json.Unmarshal(outQueryStateAgain.Bytes(), &stateAgain)
	assert.Nil(t, err)

	assert.Equal(t, int64(1), stateAgain.Version)
}
