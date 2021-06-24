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

func TestBacklogQuota(t *testing.T) {
	topicName := "persistent://public/default/test-backlog-quotas-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-backlog-quota", topicName, "-l", "1k", "-p", "producer_exception"}
	out, execErr, _, _ := TestTopicCommands(SetBacklogQuotaCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Set backlog quota policy successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(10) * time.Second)
	args = []string{"get-backlog-quotas", topicName}
	out, execErr, _, _ = TestTopicCommands(GetBacklogQuotasCmd, args)
	backlogQuotaMap := make(map[string]utils.BacklogQuotaData)
	print(out)
	err := json.Unmarshal(out.Bytes(), &backlogQuotaMap)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.NotNil(t, backlogQuotaMap["destination_storage"])
	assert.Equal(t, backlogQuotaMap["destination_storage"].Limit, int64(1024))
	assert.Equal(t, backlogQuotaMap["destination_storage"].Policy, "producer_exception")

	args = []string{"remove-backlog-quota", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveBacklogQuotaCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Remove backlog quota successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(2) * time.Second)
	args = []string{"get-backlog-quotas", topicName}
	out, execErr, _, _ = TestTopicCommands(GetBacklogQuotasCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "{}")

	// test error policy type
	args = []string{"set-backlog-quota", topicName, "-l", "1k", "-p", "error"}
	_, execErr, _, _ = TestTopicCommands(SetBacklogQuotaCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "Invalid retention policy type 'error'. Valid options are: "+
		"[producer_request_hold, producer_exception, consumer_backlog_eviction]")
}
