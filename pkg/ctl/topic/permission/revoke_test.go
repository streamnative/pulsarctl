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

package permission

import (
	"encoding/json"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestRevokePermissionsOnPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-revoke-partitioned-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	testRevokePermission(t, "test-revoke-partitioned-topic")
}

func TestRevokePermissionsOnNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-revoke-non-partitioned-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	testRevokePermission(t, "test-revoke-non-partitioned-topic")
}

func testRevokePermission(t *testing.T, topic string) {
	args := []string{"grant-permissions",
		"--role", "revoke-test-role",
		"--actions", "produce",
		topic,
	}
	_, execErr, _, _ := TestTopicCommands(GrantPermissionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", topic}
	out, execErr, _, _ := TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var permissions map[string][]pulsar.AuthAction
	err := json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(permissions["revoke-test-role"]))
	assert.Equal(t, "produce", permissions["revoke-test-role"][0].String())

	args = []string{"revoke-permissions", "--role", "revoke-test-role", topic}
	_, execErr, _, _ = TestTopicCommands(RevokePermissions, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", topic}
	out, execErr, _, _ = TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var emptyPermissions map[string][]pulsar.AuthAction
	err = json.Unmarshal(out.Bytes(), &emptyPermissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, map[string][]pulsar.AuthAction{}, emptyPermissions)
}

func TestRevokePermissionsArgError(t *testing.T) {
	args := []string{"revoke-permissions", "--role", "args-error-role"}
	_, _, nameErr, _ := TestTopicCommands(RevokePermissions, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())

	args = []string{"revoke-permissions", "--role", "", "empty-role-topic"}
	_, execErr, _, _ := TestTopicCommands(RevokePermissions, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Invalid role name", execErr.Error())

	args = []string{"revoke-permissions", "not-specified-role-topic"}
	_, _, _, err := TestTopicCommands(RevokePermissions, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"role\" not set", err.Error())
}
