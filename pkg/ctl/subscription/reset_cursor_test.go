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

func TestResetCursorCmd(t *testing.T) {
	args := []string{"create", "test-reset-cursor-topic", "test-reset-cursor-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"seek", "--time", "1m", "test-reset-cursor-topic", "test-reset-cursor-sub"}
	out, execErr, _, _ := TestSubCommands(ResetCursorCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Reset the cursor of the subscription test-reset-cursor-sub to 1m successfully\n",
		out.String())

	args = []string{"seek", "--message-id", "-1:-1", "test-reset-cursor-topic", "test-reset-cursor-sub"}
	out, execErr, _, _ = TestSubCommands(ResetCursorCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Reset the cursor of the subscription test-reset-cursor-sub to -1:-1 successfully\n",
		out.String())
}

func TestResetCursorArgsError(t *testing.T) {
	args := []string{"seek"}
	_, _, nameErr, _ := TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestResetCursorFlagError(t *testing.T) {
	args := []string{"seek", "test-reset-cursor-flag-topic", "flag-sub"}
	_, execErr, _, _ := TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "either timestamp or message-id should be specified", execErr.Error())
}

func TestResetCursorNonExistingTopic(t *testing.T) {
	args := []string{"seek", "--time", "1m", "test-reset-cursor-non-existing-topic",
		"test-reset-cursor-non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestResetCursorNonExistingSub(t *testing.T) {
	args := []string{"create", "test-reset-cursor-non-existing-sub-topic", "test-reset-cursor-existing-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"seek", "--time", "1m", "test-reset-cursor-non-existing-sub-topic",
		"test-reset-cursor-non-existing-sub"}
	_, execErr, _, _ = TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())
}
