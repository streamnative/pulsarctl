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

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestDelayedDelivery(t *testing.T) {
	topicName := "persistent://public/default/test-delayed-delivery-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-delayed-delivery", topicName, "-t", "10s", "-e"}
	out, execErr, _, _ := TestTopicCommands(SetDelayedDeliveryCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Set delayed delivery policy successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	args = []string{"get-delayed-delivery", topicName}
	out, execErr, _, _ = TestTopicCommands(GetDelayedDeliveryCmd, args)
	var delayedDeliveryData utils.DelayedDeliveryData
	err := json.Unmarshal(out.Bytes(), &delayedDeliveryData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, delayedDeliveryData.Active, true)
	assert.Equal(t, delayedDeliveryData.TickTime, float64(10))

	args = []string{"remove-delayed-delivery", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveDelayedDeliveryCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Remove delayed delivery policy successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	args = []string{"get-delayed-delivery", topicName}
	out, execErr, _, _ = TestTopicCommands(GetDelayedDeliveryCmd, args)
	err = json.Unmarshal(out.Bytes(), &delayedDeliveryData)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.Equal(t, delayedDeliveryData.Active, false)
	assert.Equal(t, delayedDeliveryData.TickTime, float64(0))

	// test specify either --enable or --disable
	args = []string{"set-delayed-delivery", topicName, "-t", "10s", "-e", "-d"}
	_, execErr, _, _ = TestTopicCommands(SetDelayedDeliveryCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "Need to specify either --enable or --disable")
}
