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
	"encoding/json"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestPublishRate(t *testing.T) {
	topicName := "persistent://public/default/test-publish-rate-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	setArgs := []string{"set-publish-rate", topicName, "--msg-publish-rate", "5", "--byte-publish-rate", "4"}
	setOut, execErr, _, _ := TestTopicCommands(SetPublishRateCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set message publish rate successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs := []string{"get-publish-rate", topicName}
	getOut, execErr, _, _ := TestTopicCommands(GetPublishRateCmd, getArgs)
	var publishRateData utils.PublishRateData
	err := json.Unmarshal(getOut.Bytes(), &publishRateData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, publishRateData.PublishThrottlingRateInMsg, int64(5))
	assert.Equal(t, publishRateData.PublishThrottlingRateInByte, int64(4))

	setArgs = []string{"remove-publish-rate", topicName}
	setOut, execErr, _, _ = TestTopicCommands(RemovePublishRateCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Remove message publish rate successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs = []string{"get-publish-rate", topicName}
	getOut, execErr, _, _ = TestTopicCommands(GetPublishRateCmd, getArgs)
	err = json.Unmarshal(getOut.Bytes(), &publishRateData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, publishRateData.PublishThrottlingRateInMsg, int64(0))
	assert.Equal(t, publishRateData.PublishThrottlingRateInByte, int64(0))
}
