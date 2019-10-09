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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCmd(t *testing.T) {
	args := []string{"create", "test-sub-topic", "test-sub-default-arg"}
	out, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Create subscription test-sub-default-arg on topic "+
		"persistent://public/default/test-sub-topic starting from latest successfully\n", out.String())

	args = []string{"create", "--messageId", "earliest", "test-sub-topic", "test-sub-earliest"}
	out, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Create subscription test-sub-earliest on topic "+
		"persistent://public/default/test-sub-topic starting from earliest successfully\n", out.String())

	args = []string{"create", "--messageId", "-1:-1", "test-sub-topic", "test-sub-messageId"}
	out, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Create subscription test-sub-messageId on topic "+
		"persistent://public/default/test-sub-topic starting from -1:-1 successfully\n", out.String())

	args = []string{"create", "--messageId", "-1", "test-sub-topic", "test-sub-invalid-messageId"}
	_, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Invalid message id string. -1", execErr.Error())

	args = []string{"list", "test-sub-topic"}
	out, execErr, _, _ = TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-sub-default-arg"))
	assert.True(t, strings.Contains(out.String(), "test-sub-earliest"))
	assert.True(t, strings.Contains(out.String(), "test-sub-messageId"))
}

func TestCreateArgsError(t *testing.T) {
	args := []string{"create"}
	_, _, nameErr, _ := TestSubCommands(CreateCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}
