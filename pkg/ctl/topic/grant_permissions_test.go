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
	"encoding/json"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGrantPermissionToNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-grant-permission-non-partitioned-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-non-partitioned-topic"}
	out, execErr, _, _ := TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var permissions map[string][]utils.AuthAction
	err := json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, map[string][]utils.AuthAction{}, permissions)

	args = []string{"grant-permissions",
		"--role", "grant-non-partitioned-role",
		"--actions", "produce",
		"test-grant-permission-non-partitioned-topic",
	}
	_, execErr, _, _ = TestTopicCommands(GrantPermissionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-non-partitioned-topic"}
	out, execErr, _, _ = TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []utils.AuthAction{"produce"}, permissions["grant-non-partitioned-role"])
}

func TestGrantPermissionToPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-grant-permission-partitioned-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-partitioned-topic"}
	out, execErr, _, _ := TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var permissions map[string][]utils.AuthAction
	err := json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, map[string][]utils.AuthAction{}, permissions)

	args = []string{"grant-permissions",
		"--role", "grant-partitioned-role",
		"--actions", "consume",
		"test-grant-permission-partitioned-topic",
	}
	_, execErr, _, _ = TestTopicCommands(GrantPermissionCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-permissions", "test-grant-permission-partitioned-topic"}
	out, execErr, _, _ = TestTopicCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []utils.AuthAction{"consume"}, permissions["grant-partitioned-role"])
}

func TestGrantPermissionArgError(t *testing.T) {
	args := []string{"grant-permissions", "--role", "test-arg-error-role", "--actions", "produce"}
	_, _, nameErr, _ := TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())

	args = []string{"grant-permissions", "args-error-topic"}
	_, _, _, err := TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"actions\", \"role\" not set", err.Error())

	args = []string{"grant-permissions", "--role", "", "--actions", "produce", "role-empty-topic"}
	_, execErr, _, _ := TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Invalid role name", execErr.Error())

	args = []string{"grant-permissions",
		"--role", "args-error-role",
		"--actions", "args-error-action",
		"invalid-actions-topic",
	}
	_, execErr, _, _ = TestTopicCommands(GrantPermissionCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "The auth action only can be specified as 'produce', "+
		"'consume', 'sources', 'sinks', 'packages', or 'functions'. Invalid auth action 'args-error-action'", execErr.Error())
}
