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

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTopicCmd(t *testing.T) {
	args := []string{"create", "test-update-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "test-update-topic"}
	out, execErr, _, _ := TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	var partitions utils.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, partitions.Partitions)

	args = []string{"update", "test-update-topic", "3"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "test-update-topic"}
	out, execErr, _, _ = TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, partitions.Partitions)
}

func TestUpdateTopicArgsError(t *testing.T) {
	args := []string{"update", "test-topic"}
	_, _, nameErr, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the partitions", nameErr.Error())
}

func TestUpdateTopicWithInvalidPartitions(t *testing.T) {
	args := []string{"update", "test-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number '0'", execErr.Error())

	args = []string{"update", "test-topic", "a"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number 'a'", execErr.Error())

	args = []string{"update", "test-topic", "--", "-1"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number '-1'", execErr.Error())
}

func TestUpdateTopicNotExist(t *testing.T) {
	args := []string{"update", "non-exist-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "code: 409")
}

func TestUpdateNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-update-non-partitioned-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"update", "test-update-non-partitioned-topic", "3"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "code: 409")
}
