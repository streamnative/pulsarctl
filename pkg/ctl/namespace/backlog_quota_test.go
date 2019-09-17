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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBacklogQuota(t *testing.T) {
	args := []string{"set-backlog-quota", "public/default", "--limit", "12M", "--policy", "consumer_backlog_eviction"}
	setOut, execErr, _, _ := TestNamespaceCommands(setBacklogQuota, args)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set backlog quota successfully for [public/default]")

	getArgs := []string{"get-backlog-quotas", "public/default"}
	getOut, execErr, _, _ := TestNamespaceCommands(getBacklogQuota, getArgs)
	assert.Nil(t, execErr)
	var backlogQuotaMap map[pulsar.BacklogQuotaType]pulsar.BacklogQuota
	err := json.Unmarshal(getOut.Bytes(), &backlogQuotaMap)
	assert.Nil(t, err)

	for key, value := range backlogQuotaMap {
		assert.Equal(t, key, pulsar.DestinationStorage)
		assert.Equal(t, value.Limit, int64(12582912))
		assert.Equal(t, value.Policy, pulsar.ConsumerBacklogEviction)
	}

	delArgs := []string{"remove-backlog-quota", "public/default"}
	delOut, execErr, _, _ := TestNamespaceCommands(removeBacklogQuota, delArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, delOut.String(), "Remove backlog quota successfully for [public/default]")
}

func TestFailureBacklogQuota(t *testing.T) {
	args := []string{"set-backlog-quota", "public/default", "--limit", "12M", "--policy", "no-support-policy"}
	_, execErr, _, _ := TestNamespaceCommands(setBacklogQuota, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "invalid retention policy type: no-support-policy")
}
