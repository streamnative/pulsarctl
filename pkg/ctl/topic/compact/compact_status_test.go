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

package compact

import (
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"

	"github.com/stretchr/testify/assert"
)

func TestCompactStatusArgsError(t *testing.T) {
	args := []string{"compact-status"}
	_, _, nameErr, _ := test.TestTopicCommands(StatusCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestCompactStatusNonExistingTopicError(t *testing.T) {
	args := []string{"compact-status", "test-non-existing-compact-status"}
	_, execErr, _, _ := test.TestTopicCommands(StatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestCompactStatusNonPersistentTopicError(t *testing.T) {
	args := []string{"compact-status", "non-persistent://public/default/test-non-persistent-topic-compact-status"}
	_, execErr, _, _ := test.TestTopicCommands(StatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "need to provide a persistent topic", execErr.Error())
}
