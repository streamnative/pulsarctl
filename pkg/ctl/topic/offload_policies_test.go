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
	"path"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestOffloadPoliciesCmd(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-offload-topic-%s", test.RandomSuffix())
	createArgs := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, createArgs)
	assert.Nil(t, execErr)

	<-time.After(5 * time.Second)
	fileSystemProfilePath := path.Join(TestResourceDir(), "policies", "filesystem_offload_core_site")
	setArgs := []string{"set-offload-policies", topicName,
		"--driver", "filesystem",
		"--file-system-uri", "hdfs://127.0.0.1:9000",
		"--file-system-profile-path", fileSystemProfilePath,
		"--threshold", "1048576",
	}

	out, execErr, _, _ := TestTopicCommands(SetOffloadPoliciesCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Set the offload policies successfully for [%s]", topicName), out.String())

	<-time.After(5 * time.Second)
	getArgs := []string{"get-offload-policies", topicName}
	out, execErr, _, _ = TestTopicCommands(GetOffloadPoliciesCmd, getArgs)
	assert.Nil(t, execErr)
	var offloadPolicies utils.OffloadPolicies
	err := json.Unmarshal(out.Bytes(), &offloadPolicies)
	assert.Nil(t, err)
	assert.Equal(t, "filesystem", offloadPolicies.ManagedLedgerOffloadDriver)
	assert.Equal(t, int64(1048576), offloadPolicies.ManagedLedgerOffloadThresholdInBytes)

	<-time.After(5 * time.Second)
	removeArgs := []string{"remove-offload-policies", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveOffloadPoliciesCmd, removeArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Remove the offload policies successfully for [%s]", topicName), out.String())

	<-time.After(5 * time.Second)
	out, execErr, _, _ = TestTopicCommands(GetOffloadPoliciesCmd, getArgs)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &offloadPolicies)
	assert.Nil(t, err)
	assert.Equal(t, "", offloadPolicies.ManagedLedgerOffloadDriver)
	assert.Equal(t, int64(0), offloadPolicies.ManagedLedgerOffloadThresholdInBytes)
}
