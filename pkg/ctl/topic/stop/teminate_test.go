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

package stop

import (
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"

	"github.com/stretchr/testify/assert"
)

func TestTerminateCmd(t *testing.T) {
	args := []string{"create", "test-terminate-topic", "0"}
	_, execErr, _, _ := test.TestTopicCommands(crud.CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"terminate", "test-terminate-topic"}
	out, execErr, _, _ := test.TestTopicCommands(TopicTerminateCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.HasPrefix(out.String(),
		"Topic persistent://public/default/test-terminate-topic is successfully terminated at"))
}

func TestTerminatePartitionedTopicError(t *testing.T) {
	args := []string{"create", "test-terminate-partitioned-topic", "2"}
	_, execErr, _, _ := test.TestTopicCommands(crud.CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"terminate", "test-terminate-partitioned-topic"}
	_, execErr, _, _ = test.TestTopicCommands(TopicTerminateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 405 reason: Termination of a partitioned topic is not allowed", execErr.Error())
}

func TestTerminateArgError(t *testing.T) {
	args := []string{"terminate"}
	_, _, nameErr, _ := test.TestTopicCommands(TopicTerminateCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestTerminateNonExistingTopic(t *testing.T) {
	args := []string{"terminate", "non-existing-topic"}
	_, execErr, _, _ := test.TestTopicCommands(TopicTerminateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
