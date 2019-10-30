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

package subscription

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipCmd(t *testing.T) {
	args := []string{"create", "test-skip-messages-topic", "test-skip-messages-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"skip", "--count", "1", "test-skip-messages-topic", "test-skip-messages-sub"}
	out, execErr, _, _ := TestSubCommands(SkipCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "The subscription test-skip-messages-sub skips 1 messages of the topic "+
		"persistent://public/default/test-skip-messages-topic successfully\n", out.String())

	args = []string{"skip", "--all", "test-skip-messages-topic", "test-skip-messages-sub"}
	out, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "The subscription test-skip-messages-sub skips -1 messages of the topic "+
		"persistent://public/default/test-skip-messages-topic successfully\n", out.String())
}

func TestSkipArgsError(t *testing.T) {
	args := []string{"skip"}
	_, _, nameErr, _ := TestSubCommands(SkipCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())

	args = []string{"skip", "test-topic", "test-sub"}
	_, execErr, _, _ := TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the skip message number is not specified", execErr.Error())
}

func TestSkipNonExistingTopic(t *testing.T) {
	args := []string{"skip", "--count", "1", "test-skip-messages-non-existing-topic",
		"test-skip-messages-non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())

	args = []string{"skip", "--all", "test-skip-messages-non-existing-topic",
		"test-skip-messages-non-existing-topic-sub"}
	_, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestSkipNonExistingSub(t *testing.T) {
	args := []string{"create", "test-skip-messages-non-existing-sub-topic",
		"test-skip-messages-non-existing-sub-existing"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"skip", "--count", "1", "test-skip-messages-non-existing-sub-topic",
		"test-skip-messages-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())

	args = []string{"skip", "--all", "test-skip-messages-non-existing-sub-topic",
		"test-skip-messages-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())
}
