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

func TestGrantSubPermissionsCmd(t *testing.T) {
	ns := "public/test-grant-sub-permissions-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"grant-subscription-permission", "--role", "test-permissions", ns, "test-grant-sub"}
	_, execErr, _, _ = TestNamespaceCommands(GrantSubPermissionsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 501 reason: Authorization is not enabled", execErr.Error())
}

func TestGrantSubPermissionsArgsError(t *testing.T) {
	ns := "public/grant-sub-permissions-args-tests"

	args := []string{"grant-subscription-permission", ns}
	_, _, _, err := TestNamespaceCommands(GrantSubPermissionsCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"role\" not set", err.Error())

	args = []string{"grant-subscription-permission", "--role", "test-role"}
	_, _, nameErr, _ := TestNamespaceCommands(GrantSubPermissionsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified namespace name and subscription name", nameErr.Error())

	args = []string{"grant-subscription-permission", "--role", "test-role", ns}
	_, execErr, _, _ := TestNamespaceCommands(GrantSubPermissionsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "need to specified namespace name and subscription name", execErr.Error())
}
