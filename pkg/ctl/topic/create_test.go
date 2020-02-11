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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTopicCmd(t *testing.T) {
	args := []string{"create", "test-create-topic", "2"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateNonPersistentTopic(t *testing.T) {
	args := []string{"create", "non-persistent://public/default/test-create-topic", "2"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateTopicAlreadyExists(t *testing.T) {
	args := []string{"create", "test-duplicate-topic", "2"}
	_, _, _, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, err)

	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, execErr)
}

func TestCreateTopicArgsError(t *testing.T) {
	args := []string{"create", "topic"}
	_, _, nameErr, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the partitions", nameErr.Error())
}

func TestCreateTopicWithInvalidPartitions(t *testing.T) {
	args := []string{"create", "topic", "a"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number 'a'", execErr.Error())

	args = []string{"create", "topic", "--", "-1"}
	_, execErr, _, _ = TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number '-1'", execErr.Error())
}

func TestCreateNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-create-non-partitioned-topic", "0"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateNonPersistentNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "non-persistent://public/default/test-create-non-partitioned-topic", "0"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}
