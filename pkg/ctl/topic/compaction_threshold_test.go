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
	"fmt"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestCompactionThresholdCmd(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-compaction-threshold-topic-%s",
		test.RandomSuffix())
	createArgs := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, createArgs)
	assert.Nil(t, execErr)

	setArgs := []string{"set-compaction-threshold", topicName, "--threshold", "1G"}
	out, execErr, _, _ := TestTopicCommands(SetCompactionThresholdCmd, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(),
		fmt.Sprintf("Successfully set compaction threshold to %d for topic %s", 1024*1024*1024, topicName))

	<-time.After(5 * time.Second)

	getArgs := []string{"get-compaction-threshold", topicName}
	out, execErr, _, _ = TestTopicCommands(GetCompactionThresholdCmd, getArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(),
		fmt.Sprintf("The compaction threshold of the topic %s is %d byte(s)", topicName, 1024*1024*1024))

	removeArgs := []string{"remove-compaction-threshold", topicName}
	out, execErr, _, _ = TestTopicCommands(RemoveCompactionThresholdCmd, removeArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(),
		fmt.Sprintf("Successfully remove compaction threshold for topic %s", topicName))

	<-time.After(5 * time.Second)

	out, execErr, _, _ = TestTopicCommands(GetCompactionThresholdCmd, getArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(),
		fmt.Sprintf("The compaction threshold of the topic %s is %d byte(s)", topicName, 0))
}
