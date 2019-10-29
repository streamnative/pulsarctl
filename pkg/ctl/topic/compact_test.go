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

func TestCompactCmd(t *testing.T) {
	args := []string{"create", "test-compact-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"compact-status", "test-compact-topic"}
	out, execErr, _, _ := TestTopicCommands(StatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Compacting the topic persistent://public/default/test-compact-topic is not running\n",
		out.String())

	args = []string{"compact", "test-compact-topic"}
	_, execErr, _, _ = TestTopicCommands(CompactCmd, args)
	assert.Nil(t, execErr)

	args = []string{"compact-status", "test-compact-topic"}
	out, execErr, _, _ = TestTopicCommands(StatusCmd, args)
	assert.Nil(t, execErr)

	assert.Equal(t, "Compacting the topic persistent://public/default/test-compact-topic is running\n",
		out.String())
}

func TestCompactArgError(t *testing.T) {
	args := []string{"compact"}
	_, _, nameErr, _ := TestTopicCommands(CompactCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestCompactNonExistingTopic(t *testing.T) {
	args := []string{"compact", "test-compact-non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(CompactCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestCompactNonPersistentTopic(t *testing.T) {
	args := []string{"compact", "non-persistent://public/default/test-compact-non-persistent-topic"}
	_, execErr, _, _ := TestTopicCommands(CompactCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "need to provide a persistent topic", execErr.Error())
}
