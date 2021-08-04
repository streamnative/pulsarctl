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
	"time"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestDeduplicationStatus(t *testing.T) {
	topicName := "persistent://public/default/test-deduplication-status-topic-" + test.RandomSuffix()
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-deduplication", topicName, "-e"}
	out, execErr, _, _ := TestTopicCommands(SetDeduplicationStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Enable the deduplication policy successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	args = []string{"get-deduplication", topicName}
	out, execErr, _, _ = TestTopicCommands(GetDeduplicationStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "true")

	args = []string{"remove-deduplication", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveDeduplicationStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Remove the deduplication policy successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	args = []string{"get-deduplication", topicName}
	out, execErr, _, _ = TestTopicCommands(GetDeduplicationStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "false")

	args = []string{"set-deduplication", topicName, "-d"}
	out, execErr, _, _ = TestTopicCommands(SetDeduplicationStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Disable the deduplication policy successfully for ["+topicName+"]\n")

	time.Sleep(time.Duration(1) * time.Second)
	args = []string{"get-deduplication", topicName}
	out, execErr, _, _ = TestTopicCommands(GetDeduplicationStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "false")

	args = []string{"set-deduplication", topicName, "-e", "-d"}
	_, execErr, _, _ = TestTopicCommands(SetDeduplicationStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "Need to specify either --enable or --disable")
}
