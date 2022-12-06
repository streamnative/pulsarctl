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

func TestDeleteCmd(t *testing.T) {
	args := []string{"create", "test-delete-sub-topic", "test-delete-sub-1"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"create", "test-delete-sub-topic", "test-delete-sub-2"}
	_, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "test-delete-sub-topic"}
	out, execErr, _, _ := TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-sub-1"))
	assert.True(t, strings.Contains(out.String(), "test-delete-sub-2"))

	args = []string{"delete", "test-delete-sub-topic", "test-delete-sub-1"}
	out, execErr, _, _ = TestSubCommands(DeleteCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Delete the subscription test-delete-sub-1 of the topic "+
		"persistent://public/default/test-delete-sub-topic successfully\n", out.String())

	args = []string{"list", "test-delete-sub-topic"}
	out, execErr, _, _ = TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-sub-1"))
	assert.True(t, strings.Contains(out.String(), "test-delete-sub-2"))
}

func TestDeleteArgsError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())

	args = []string{"delete", "arg-error"}
	_, _, nameErr, _ = TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestDeleteNonExistingSub(t *testing.T) {
	args := []string{"create", "test-delete-non-existing-sub-topic", "test-delete-non-existing-sub-existing"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "test-delete-non-existing-sub-topic"}
	out, execErr, _, _ := TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-non-existing-sub-existing"))

	args = []string{"delete", "test-delete-non-existing-sub-topic", "test-delete-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "code: 404 reason: Subscription")

	args = []string{"list", "test-delete-non-existing-sub-topic"}
	out, execErr, _, _ = TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-non-existing-sub-existing"))
}

func TestDeleteNonExistingTopicSub(t *testing.T) {
	args := []string{"delete", "non-existing-topic", "non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "code: 404 reason: Topic")
}
