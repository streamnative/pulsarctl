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

package namespace

import (
	"encoding/json"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBacklogQuota(t *testing.T) {
	args := []string{"create", "public/test-backlog-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-backlog-namespace successfully\n")

	args = []string{"set-backlog-quota", "public/test-backlog-namespace",
		"--limit-size", "2G", "--policy", "producer_request_hold"}
	setOut, execErr, _, _ := TestNamespaceCommands(setBacklogQuota, args)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set backlog quota successfully for [public/test-backlog-namespace]\n")

	args = []string{"set-backlog-quota", "public/test-backlog-namespace",
		"--limit-size", "1G", "--limit-time", "12345", "--policy", "consumer_backlog_eviction", "--type", "message_age"}
	setOut, execErr, _, _ = TestNamespaceCommands(setBacklogQuota, args)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set backlog quota successfully for [public/test-backlog-namespace]\n")

	getArgs := []string{"get-backlog-quotas", "public/test-backlog-namespace"}
	getOut, execErr, _, _ := TestNamespaceCommands(getBacklogQuota, getArgs)
	assert.Nil(t, execErr)
	var backlogQuotaMap map[utils.BacklogQuotaType]utils.BacklogQuota
	err = json.Unmarshal(getOut.Bytes(), &backlogQuotaMap)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(backlogQuotaMap))
	assert.Equal(t, backlogQuotaMap[utils.DestinationStorage].LimitSize, int64(2147483648))
	assert.Equal(t, backlogQuotaMap[utils.DestinationStorage].Policy, utils.ProducerRequestHold)
	assert.Equal(t, backlogQuotaMap[utils.MessageAge].LimitSize, int64(1073741824))
	assert.Equal(t, backlogQuotaMap[utils.MessageAge].LimitTime, int64(12345))
	assert.Equal(t, backlogQuotaMap[utils.MessageAge].Policy, utils.ConsumerBacklogEviction)

	delArgs := []string{"remove-backlog-quota", "public/test-backlog-namespace"}
	delOut, execErr, _, _ := TestNamespaceCommands(removeBacklogQuota, delArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, delOut.String(), "Remove backlog quota successfully for [public/test-backlog-namespace]\n")
}

func TestFailureBacklogQuota(t *testing.T) {
	args := []string{"set-backlog-quota", "public/test-backlog-namespace",
		"--limit-size", "12M", "--policy", "no-support-policy"}
	_, execErr, _, _ := TestNamespaceCommands(setBacklogQuota, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "invalid retention policy type: no-support-policy")
}
