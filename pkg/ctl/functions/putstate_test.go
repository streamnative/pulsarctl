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
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestStateFunctions(t *testing.T) {
	fName := "f" + test.RandomSuffix()
	jarName := path.Join(ResourceDir(), "api-examples.jar")

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}
	out, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Equal(t, out.String(), fmt.Sprintf("Created %s successfully\n", fName))

	putstateArgs := []string{"putstate",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"pulsar", "-", "hello",
	}

	task := func(args []string, obj interface{}) bool {
		out, execErr, _ := TestFunctionsCommands(putstateFunctionsCmd, args)
		if execErr != nil {
			return false
		}

		return strings.Contains(out.String(), "successfully\n")
	}

	err = cmdutils.RunFuncWithTimeout(task, true, 1*time.Minute, putstateArgs, nil)
	if err != nil {
		t.Fatal(err)
	}

	// test failure case for put state
	failureStateArgs := []string{"putstate",
		"--name", "not-exist",
		"pulsar", "-", "hello",
	}

	stateArgsErrInFormat := []string{"putstate",
		"--name", fName,
		"pulsar", "hello",
	}

	_, execErrMsg, _ := TestFunctionsCommands(putstateFunctionsCmd, failureStateArgs)
	assert.NotNil(t, execErrMsg)
	exceptMsg := "'not-exist' is not found"
	assert.True(t, strings.Contains(execErrMsg.Error(), exceptMsg))

	_, errMsg, _ := TestFunctionsCommands(putstateFunctionsCmd, stateArgsErrInFormat)
	assert.NotNil(t, errMsg)
	exceptErrMsg := "error input format"
	assert.True(t, strings.Contains(errMsg.Error(), exceptErrMsg))

	// query state
	queryStateArgs := []string{"querystate",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--key", "pulsar",
	}

	outQueryState, _, err := TestFunctionsCommands(querystateFunctionsCmd, queryStateArgs)
	if err != nil {
		t.Fatal(err)
	}

	var state utils.FunctionState
	err = json.Unmarshal(outQueryState.Bytes(), &state)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "pulsar", state.Key)
	assert.Equal(t, "hello", state.StringValue)
	assert.Equal(t, int64(0), state.Version)
	// put state again
	outPutStateAgain, _, err := TestFunctionsCommands(putstateFunctionsCmd, putstateArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, strings.Contains(outPutStateAgain.String(), "successfully\n"))

	// query state again
	outQueryStateAgain, _, err := TestFunctionsCommands(querystateFunctionsCmd, queryStateArgs)
	if err != nil {
		t.Fatal(err)
	}

	var stateAgain utils.FunctionState
	err = json.Unmarshal(outQueryStateAgain.Bytes(), &stateAgain)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(1), stateAgain.Version)
}

func TestByteValue(t *testing.T) {
	fName := "f" + test.RandomSuffix()
	jarName := path.Join(ResourceDir(), "api-examples.jar")

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	out, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Equal(t, out.String(), fmt.Sprintf("Created %s successfully\n", fName))

	buf := "hello pulsar!"
	file, err := os.CreateTemp("", "byte-value-functions")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte(buf)); err != nil {
		t.Fatal(err)
	}

	putstateArgs := []string{"putstate",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"pulsar", "=", file.Name(),
	}

	task := func(args []string, obj interface{}) bool {
		out, execErr, _ := TestFunctionsCommands(putstateFunctionsCmd, args)
		if execErr != nil {
			return false
		}
		return strings.Contains(out.String(), "successfully\n")
	}

	err = cmdutils.RunFuncWithTimeout(task, true, 1*time.Minute, putstateArgs, nil)
	if err != nil {
		t.Fatal(err)
	}

	// query state
	queryStateArgs := []string{"querystate",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--key", "pulsar",
	}

	outQueryState, execErr, err := TestFunctionsCommands(querystateFunctionsCmd, queryStateArgs)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

	var state utils.FunctionState
	err = json.Unmarshal(outQueryState.Bytes(), &state)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "pulsar", state.Key)
	assert.Equal(t, "hello pulsar!", state.StringValue)
}
