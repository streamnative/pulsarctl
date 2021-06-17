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

func TestCreateTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "super-user", "--allowed-clusters", "standalone", "create-tenant-test"}
	_, _, _, err := TestTenantCommands(createTenantCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, _ := TestTenantCommands(listTenantCmd, args)

	assert.True(t, strings.Contains(out.String(), "create-tenant-test"))
}

func TestCreateTenantArgsError(t *testing.T) {
	args := []string{"create"}
	_, _, nameErr, _ := TestTenantCommands(createTenantCmd, args)

	assert.Equal(t, "the tenant name is not specified or the tenant name is specified more than one", nameErr.Error())
}

func TestCreateTenantAlreadyExistError(t *testing.T) {
	args := []string{"create", "--allowed-clusters", "standalone", "create-tenant-duplicate"}
	_, execErr, _, _ := TestTenantCommands(createTenantCmd, args)
	assert.Nil(t, execErr)

	args = []string{"create", "--allowed-clusters", "standalone", "create-tenant-duplicate"}
	_, execErr, _, _ = TestTenantCommands(createTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Tenant already exist", execErr.Error())
}
