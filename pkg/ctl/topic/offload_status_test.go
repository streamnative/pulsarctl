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

func TestOffloadStatusCmd(t *testing.T) {
	args := []string{"create", "test-offload-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"offload-status", "test-offload-topic"}
	out, execErr, _, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Offloading topic persistent://public/default/test-offload-topic is not running\n",
		out.String())
}

func TestOffloadStatusArgsError(t *testing.T) {
	args := []string{"offload-status"}
	_, _, nameErr, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestOffloadStatusNonPersistentTopicError(t *testing.T) {
	args := []string{"offload-status", "non-persistent://public/default/test-offload-status-non-persistent-topic"}
	_, execErr, _, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "need to provide a persistent topic", execErr.Error())
}

func TestOffloadStatusNonExistingTopicError(t *testing.T) {
	args := []string{"offload-status", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
