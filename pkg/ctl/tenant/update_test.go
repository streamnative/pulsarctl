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

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "update-role", "--allowed-clusters", "standalone", "update-tenant-test"}
	_, execErr, _, _ := TestTenantCommands(createTenantCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "update-tenant-test"}
	out, execErr, _, _ := TestTenantCommands(getTenantCmd, args)
	assert.Nil(t, execErr)

	var tenantData utils.TenantData
	err := json.Unmarshal(out.Bytes(), &tenantData)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(tenantData.AdminRoles))
	assert.Equal(t, "update-role", tenantData.AdminRoles[0])
	assert.Equal(t, 1, len(tenantData.AllowedClusters))
	assert.Equal(t, "standalone", tenantData.AllowedClusters[0])

	args = []string{"update", "--admin-roles", "new-role", "--allowed-clusters", "standalone", "update-tenant-test"}
	_, execErr, _, _ = TestTenantCommands(UpdateTenantCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "update-tenant-test"}
	out, _, _, _ = TestTenantCommands(getTenantCmd, args)

	err = json.Unmarshal(out.Bytes(), &tenantData)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(tenantData.AdminRoles))
	assert.Equal(t, "new-role", tenantData.AdminRoles[0])
	assert.Equal(t, 1, len(tenantData.AllowedClusters))
	assert.Equal(t, "standalone", tenantData.AllowedClusters[0])
}

func TestUpdateArgsError(t *testing.T) {
	args := []string{"update"}
	_, _, nameErr, _ := TestTenantCommands(UpdateTenantCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the tenant name is not specified or the tenant name is specified more than one", nameErr.Error())
}

func TestUpdateNonExistTenantError(t *testing.T) {
	args := []string{"update", "--admin-roles", "update-role", "--allowed-clusters", "standalone", "non-existent-topic"}
	_, execErr, _, _ := TestTenantCommands(UpdateTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "404")
}
