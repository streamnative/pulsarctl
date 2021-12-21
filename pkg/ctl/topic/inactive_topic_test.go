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
	"fmt"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestInactiveTopicCmd(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-inactive-topic-%s",
		test.RandomSuffix())
	createArgs := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, createArgs)
	assert.Nil(t, execErr)

	<-time.After(5 * time.Second)

	setArgs := []string{"set-inactive-topic-policies", topicName,
		"-e=true",
		"-t", "1h",
		"-m", "delete_when_no_subscriptions"}
	out, execErr, _, _ := TestTopicCommands(SetInactiveTopicCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(),
		fmt.Sprintf("Set inactive topic policies successfully for [%s]", topicName))

	<-time.After(5 * time.Second)

	getArgs := []string{"get-inactive-topic-policies", topicName}
	var inactiveTopic utils.InactiveTopicPolicies
	task := func(args []string, obj interface{}) bool {
		out, execErr, _, _ = TestTopicCommands(GetInactiveTopicCmd, args)
		if execErr != nil {
			return false
		}
		err := json.Unmarshal(out.Bytes(), &inactiveTopic)
		if err != nil {
			return false
		}

		return inactiveTopic.DeleteWhileInactive &&
			inactiveTopic.MaxInactiveDurationSeconds == 3600 &&
			inactiveTopic.InactiveTopicDeleteMode.String() == "delete_when_no_subscriptions"
	}
	err := cmdutils.RunFuncWithTimeout(task, true, 30 * time.Second, getArgs, nil)
	if err != nil {
		t.Fatal(err)
	}

	removeArgs := []string{"remove-inactive-topic-policies", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveInactiveTopicCmd, removeArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(),
		fmt.Sprintf("Remove inactive topic policies successfully from [%s]", topicName))

	<-time.After(5 * time.Second)

	out, execErr, _, _ = TestTopicCommands(GetInactiveTopicCmd, getArgs)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &inactiveTopic)
	assert.NoError(t, err)
	assert.Equal(t, false, inactiveTopic.DeleteWhileInactive)
	assert.Equal(t, 0, inactiveTopic.MaxInactiveDurationSeconds)
	assert.Nil(t, inactiveTopic.InactiveTopicDeleteMode)
}
