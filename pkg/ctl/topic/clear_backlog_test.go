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

func TestClearBacklogCmd(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-clear-backlog-test-%s", test.RandomSuffix())

	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	<-time.After(3 * time.Second)
	args = []string{"clear-backlog", topicName, "--subscription", "app1"}
	out, execErr, _, _ := TestTopicCommands(ClearBacklogCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Clear backlog successfully for [%s]", topicName),
		out.String())
}
