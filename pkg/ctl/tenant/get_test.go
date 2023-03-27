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
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "admin-get-test", "--allowed-clusters", "standalone", "get-tenant-test"}
	_, _, _, _ = TestTenantCommands(createTenantCmd, args)

	args = []string{"get", "get-tenant-test"}
	out, _, _, _ := TestTenantCommands(getTenantCmd, args)

	var tenantData utils.TenantData
	err := json.Unmarshal(out.Bytes(), &tenantData)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(tenantData.AdminRoles))
	assert.Equal(t, "admin-get-test", tenantData.AdminRoles[0])
	assert.Equal(t, 1, len(tenantData.AllowedClusters))
	assert.Equal(t, "standalone", tenantData.AllowedClusters[0])
}

func TestGetTenantArgsError(t *testing.T) {
	args := []string{"get"}
	_, _, nameErr, _ := TestTenantCommands(getTenantCmd, args)
	assert.Equal(t, "the tenant name is not specified or the tenant name is specified more than one", nameErr.Error())
}

func TestGetNonExistTenant(t *testing.T) {
	args := []string{"get", "non-existent-tenant"}
	_, execErr, _, _ := TestTenantCommands(getTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}
