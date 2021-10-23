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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestBacklogQuotaCmd(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-backlog-quotas-topic-%s", test.RandomSuffix())
	createArgs := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, createArgs)
	assert.Nil(t, execErr)

	setArgs := []string{"set-backlog-quota", topicName,
		"--limit-size", "1k",
		"--limit-time", "20",
		"-p", "producer_exception"}
	out, execErr, _, _ := TestTopicCommands(SetBacklogQuotaCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), fmt.Sprintf("Set backlog quota successfully for [%s]\n", topicName))

	<-time.After(5 * time.Second)

	getArgs := []string{"get-backlog-quotas", topicName}
	out, execErr, _, _ = TestTopicCommands(GetBacklogQuotaCmd, getArgs)
	assert.Nil(t, execErr)

	var backlogQuotaMap map[string]utils.BacklogQuota
	err := json.Unmarshal(out.Bytes(), &backlogQuotaMap)
	assert.Nil(t, err)
	assert.NotNil(t, backlogQuotaMap)
	assert.NotNil(t, backlogQuotaMap["destination_storage"].LimitTime, int64(20))
	assert.NotNil(t, backlogQuotaMap["destination_storage"].LimitSize, int64(1024))
	assert.Equal(t, backlogQuotaMap["destination_storage"].Policy, utils.ProducerException)

	removeArgs := []string{"remove-backlog-quota", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveBacklogQuotaCmd, removeArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Remove backlog quota successfully for ["+topicName+"]\n")

	<-time.After(5 * time.Second)

	out, execErr, _, _ = TestTopicCommands(GetBacklogQuotaCmd, getArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "{}\n")
}
