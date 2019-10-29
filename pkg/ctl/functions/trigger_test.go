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

// TODO re-enable the test: https://github.com/streamnative/pulsarctl/issues/60
// +build functions

package functions

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestTriggerFunctions(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-trigger",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.WordCountFunction",
		"--jar", basePath + "/test/functions/api-examples.jar",
	}

	out, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)
	if execErr != nil {
		t.Errorf("create fucntions error value: %s", execErr.Error())
	}
	assert.Equal(t, out.String(), "Created test-functions-trigger successfully\n")

	statsArgs := []string{"stats",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-trigger",
	}
	outStats := new(bytes.Buffer)

	outStats, _, _ = TestFunctionsCommands(statsFunctionsCmd, statsArgs)
	var stats pulsar.FunctionStats
	err = json.Unmarshal(outStats.Bytes(), &stats)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.ReceivedTotal)
	assert.Equal(t, int64(0), stats.ProcessedSuccessfullyTotal)

	// send trigger cmd to broker
	triggerArgs := []string{"trigger",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-trigger",
		"--topic", "test-input-topic",
		"--trigger-value", "hello pulsar",
	}

	triggerOut := new(bytes.Buffer)
	for i := 0; i < 2; i++ {
		triggerOut, execErr, err = TestFunctionsCommands(triggerFunctionsCmd, triggerArgs)
		assert.Nil(t, err)
		if execErr != nil {
			t.Error(execErr.Error())
		}
		t.Log(triggerOut.String())
	}
}

func TestTriggerFunctionsFailure(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-trigger-failure",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", basePath + "/test/functions/api-examples.jar",
	}

	out, _, err := TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, out.String(), "Created test-functions-trigger-failure successfully\n")
	// wait the function create successfully
	time.Sleep(time.Second * 3)

	triggerArgs := []string{"trigger",
		"--name", "not-exist",
		"--topic", "test-input-topic",
		"--trigger-value", "hello pulsar",
	}

	_, errMsg, _ := TestFunctionsCommands(triggerFunctionsCmd, triggerArgs)
	errorMessage := "code: 404 reason: Function not-exist doesn't exist"
	assert.Equal(t, errorMessage, errMsg.Error())

	triggerArgsNoTopic := []string{"trigger",
		"--name", "test-functions-trigger-failure",
		"--topic", "test-input-topic-failure",
		"--trigger-value", "hello pulsar",
	}
	_, errMsg, _ = TestFunctionsCommands(triggerFunctionsCmd, triggerArgsNoTopic)
	noTopicErr := "code: 400 reason: Function in trigger function has unidentified topic"
	assert.Equal(t, noTopicErr, errMsg.Error())

	triggerArgsNoValueOrFile := []string{"trigger",
		"--name", "test-functions-trigger-failure",
		"--topic", "test-input-topic",
	}
	_, errMsg, _ = TestFunctionsCommands(triggerFunctionsCmd, triggerArgsNoValueOrFile)
	noValueOrFile := "either a trigger value or a trigger filepath needs to be specified"
	assert.Equal(t, noValueOrFile, errMsg.Error())
}
