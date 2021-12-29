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

package topic

import (
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/stretchr/testify/assert"
)

func TestMaxUnackMessagesPerConsumer(t *testing.T) {
	topicName := "persistent://public/default/test-max-unacked-messages-per-consumer-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	setArgs := []string{"set-max-unacked-messages-per-consumer", topicName, "-m", "20"}
	setOut, execErr, _, _ := TestTopicCommands(SetMaxUnackMessagesPerConsumerCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set max unacked messages per consumer successfully for ["+topicName+"]\n")

	getArgs := []string{"get-max-unacked-messages-per-consumer", topicName}
	task := func(args []string, obj interface{}) bool {
		getOut, execErr, _, _ := TestTopicCommands(GetMaxUnackMessagesPerConsumerCmd, args)
		if execErr != nil {
			return false
		}

		return getOut.String() == "20"
	}
	err := cmdutils.RunFuncWithTimeout(task, true, 30 * time.Second, getArgs, nil)
	if err != nil {
		t.Fatal(err)
	}

	setArgs = []string{"remove-max-unacked-messages-per-consumer", topicName}
	setOut, execErr, _, _ = TestTopicCommands(RemoveMaxUnackMessagesPerConsumerCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Remove max unacked messages per consumer successfully for ["+topicName+"]\n")

	task = func(args []string, obj interface{}) bool {
		getOut, execErr, _, _ := TestTopicCommands(GetMaxUnackMessagesPerConsumerCmd, args)
		if execErr != nil {
			return false
		}

		return getOut.String() == "0"
	}
	err = cmdutils.RunFuncWithTimeout(task, true, 30 * time.Second, getArgs, nil)
	if err != nil {
		t.Fatal(err)
	}

	// test negative value
	setArgs = []string{"set-max-unacked-messages-per-consumer", topicName, "-m", "-2"}
	_, execErr, _, _ = TestTopicCommands(SetMaxUnackMessagesPerConsumerCmd, setArgs)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 412 reason: maxUnackedNum must be 0 or more")
}
