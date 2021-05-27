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

func TestExpireCmd(t *testing.T) {
	args := []string{"create", "test-expire-messages-topic", "test-expire-messages-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	// expired message on the topic without any messages should be failed.
	// That was introduce by https://github.com/apache/pulsar/pull/9561
	args = []string{"expire", "--expire-time", "1", "test-expire-messages-topic", "test-expire-messages-sub"}
	_, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)

	args = []string{"expire", "--expire-time", "1", "--all", "test-expire-messages-topic"}
	_, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
}

func TestExpireArgsError(t *testing.T) {
	args := []string{"expire"}
	_, _, _, err := TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"expire-time\" not set", err.Error())

	args = []string{"expire", "--expire-time", "1"}
	_, _, nameErr, _ := TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestExpireNonExistingTopic(t *testing.T) {
	args := []string{"expire", "--expire-time", "1", "non-existing-topic",
		"test-expire-messages-non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())

	args = []string{"expire", "--expire-time", "1", "--all", "non-existing-topic"}
	_, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestExpireNonExistingSub(t *testing.T) {
	args := []string{"create", "test-expire-messages-non-existing-sub-topic",
		"test-expire-messages-non-existing-sub-existing"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"expire", "--expire-time", "1", "test-expire-messages-non-existing-sub-topic",
		"test-expire-messages-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())
}
