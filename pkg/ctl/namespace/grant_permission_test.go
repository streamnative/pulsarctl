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

package namespace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrantPermissionsCmd(t *testing.T) {
	ns := "public/test-grant-permissions-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)
	args = []string{"grant-permission", "--role", "test-permissions",
		"--actions", "produce", "--actions", "consume", "--actions", "functions", ns}
	_, execErr, _, _ = TestNamespaceCommands(GrantPermissionsCmd, args)
	assert.Nil(t, execErr)
}

func TestGrantPermissionsArgsError(t *testing.T) {
	ns := "public/grant-permissions-tests"

	args := []string{"grant-permission", ns}
	_, _, _, err := TestNamespaceCommands(GrantPermissionsCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"actions\", \"role\" not set", err.Error())

	args = []string{"grant-permission", "--role", "test-role", "--actions", "consume"}
	_, _, nameErr, _ := TestNamespaceCommands(GrantPermissionsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the namespace name is not specified or the namespace name is specified more than one",
		nameErr.Error())

	args = []string{"grant-permission", "--role", "test-role", "--actions", "fail", ns}
	_, execErr, _, _ := TestNamespaceCommands(GrantPermissionsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "The auth action only can be specified as 'produce', "+
		"'consume', or 'functions'. Invalid auth action 'fail'", execErr.Error())
}
