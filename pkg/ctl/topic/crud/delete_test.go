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

package crud

import (
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"

	"github.com/stretchr/testify/assert"
)

func TestDeletePartitionedTopic(t *testing.T) {
	args := []string{"create", "test-delete-partitioned-topic", "2"}
	_, execErr, _, _ := test.TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ := test.TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-partitioned-topic"))

	args = []string{"delete", "test-delete-partitioned-topic"}
	_, execErr, _, _ = test.TestTopicCommands(DeleteTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ = test.TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-partitioned-topic"))
}

func TestDeleteNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-delete-non-partitioned-topic", "0"}
	_, execErr, _, _ := test.TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ := test.TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-non-partitioned-topic"))

	args = []string{"delete", "--non-partitioned", "test-delete-non-partitioned-topic"}
	_, execErr, _, _ = test.TestTopicCommands(DeleteTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ = test.TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-non-partitioned-topic"))
}

func TestDeleteTopicArgError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := test.TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestDeleteNonExistPartitionedTopic(t *testing.T) {
	args := []string{"delete", "non-existent-partitioned-topic"}
	_, execErr, _, _ := test.TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Partitioned topic does not exist", execErr.Error())
}

func TestDeleteNonExistNonPartitionedTopic(t *testing.T) {
	args := []string{"delete", "--non-partitioned", "non-existent-non-partitioned-topic"}
	_, execErr, _, _ := test.TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
