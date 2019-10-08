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

package tenant

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "super-user", "--allowed-clusters", "standalone", "delete-tenant-test"}
	_, _, _, _ = TestTenantCommands(createTenantCmd, args)

	args = []string{"list"}
	out, _, _, _ := TestTenantCommands(listTenantCmd, args)

	assert.True(t, strings.Contains(out.String(), "delete-tenant-test"))

	args = []string{"delete", "delete-tenant-test"}
	_, _, _, _ = TestTenantCommands(deleteTenantCmd, args)

	args = []string{"list"}
	out, _, _, _ = TestTenantCommands(listTenantCmd, args)

	assert.False(t, strings.Contains(out.String(), "delete-tenant-test"))
}

func TestDeleteTenantArgsError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := TestTenantCommands(deleteTenantCmd, args)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestDeleteNonExistTenant(t *testing.T) {
	args := []string{"delete", "non-existent-tenant"}
	_, execErr, _, _ := TestTenantCommands(deleteTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: The tenant does not exist", execErr.Error())
}

func TestDeleteNonEmptyTenant(t *testing.T) {
	args := []string{"delete", "public"}
	_, execErr, _, _ := TestTenantCommands(deleteTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: The tenant still has active namespaces", execErr.Error())
}
