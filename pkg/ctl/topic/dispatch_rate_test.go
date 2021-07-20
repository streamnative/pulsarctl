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

func TestDispatchRate(t *testing.T) {
	topicName := "persistent://public/default/test-dispatch-rate-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	setArgs := []string{"set-dispatch-rate", topicName, "--msg-dispatch-rate", "5", "--byte-dispatch-rate", "4",
		"--dispatch-rate-period", "3", "--relative-to-publish-rate"}
	setOut, execErr, _, _ := TestTopicCommands(SetDispatchRateCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set message dispatch rate successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs := []string{"get-dispatch-rate", topicName}
	getOut, execErr, _, _ := TestTopicCommands(GetDispatchRateCmd, getArgs)
	var dispatchRateData utils.DispatchRateData
	err := json.Unmarshal(getOut.Bytes(), &dispatchRateData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, dispatchRateData.DispatchThrottlingRateInMsg, int64(5))
	assert.Equal(t, dispatchRateData.DispatchThrottlingRateInByte, int64(4))
	assert.Equal(t, dispatchRateData.RatePeriodInSecond, int64(3))
	assert.Equal(t, dispatchRateData.RelativeToPublishRate, true)

	setArgs = []string{"remove-dispatch-rate", topicName}
	setOut, execErr, _, _ = TestTopicCommands(RemoveDispatchRateCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Remove message dispatch rate successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs = []string{"get-dispatch-rate", topicName}
	getOut, execErr, _, _ = TestTopicCommands(GetDispatchRateCmd, getArgs)
	err = json.Unmarshal(getOut.Bytes(), &dispatchRateData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, dispatchRateData.DispatchThrottlingRateInMsg, int64(0))
	assert.Equal(t, dispatchRateData.DispatchThrottlingRateInByte, int64(0))
	assert.Equal(t, dispatchRateData.RatePeriodInSecond, int64(0))
	assert.Equal(t, dispatchRateData.RelativeToPublishRate, false)

	setArgs = []string{"set-dispatch-rate", topicName, "--msg-dispatch-rate", "5", "--byte-dispatch-rate", "4",
		"--dispatch-rate-period", "3"}
	setOut, execErr, _, _ = TestTopicCommands(SetDispatchRateCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set message dispatch rate successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	getArgs = []string{"get-dispatch-rate", topicName}
	getOut, execErr, _, _ = TestTopicCommands(GetDispatchRateCmd, getArgs)
	err = json.Unmarshal(getOut.Bytes(), &dispatchRateData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, dispatchRateData.DispatchThrottlingRateInMsg, int64(5))
	assert.Equal(t, dispatchRateData.DispatchThrottlingRateInByte, int64(4))
	assert.Equal(t, dispatchRateData.RatePeriodInSecond, int64(3))
	assert.Equal(t, dispatchRateData.RelativeToPublishRate, false)
}
