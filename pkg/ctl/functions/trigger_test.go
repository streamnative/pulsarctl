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
//go:build function
// +build function

package functions

import (
	"encoding/json"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestTriggerFunctions(t *testing.T) {
	fName := "trigger-f" + test.RandomSuffix()
	jarName := path.Join(ResourceDir(), "api-examples.jar")

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.WordCountFunction",
		"--jar", jarName,
	}

	out, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Equal(t, out.String(), fmt.Sprintf("Created %s successfully\n", fName))

	statsArgs := []string{"stats",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
	}
	outStats, execErr, err := TestFunctionsCommands(statsFunctionsCmd, statsArgs)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	var stats utils.FunctionStats
	err = json.Unmarshal(outStats.Bytes(), &stats)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.ReceivedTotal)
	assert.Equal(t, int64(0), stats.ProcessedSuccessfullyTotal)

	// send trigger cmd to broker
	triggerArgs := []string{"trigger",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--topic", "test-input-topic",
		"--trigger-value", "hello pulsar",
	}

	for i := 0; i < 2; i++ {
		_, execErr, err = TestFunctionsCommands(triggerFunctionsCmd, triggerArgs)
		assert.Nil(t, err)
		if execErr != nil {
			t.Error(execErr.Error())
		}
	}
}

func TestTriggerFunctionsFailure(t *testing.T) {
	fName := "trigger-f" + test.RandomSuffix()
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
		"--name", fName,
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
