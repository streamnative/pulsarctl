package tenant

import (
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "update-role", "--allowed-clusters", "standalone", "update-tenant-test"}
	_, execErr, _, _ := TestTenantCommands(createTenantCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "update-tenant-test"}
	out, execErr, _, _ := TestTenantCommands(getTenantCmd, args)
	assert.Nil(t, execErr)

	t.Log(out.String())
	var tenantData pulsar.TenantData
	err := json.Unmarshal(out.Bytes(), &tenantData)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(tenantData.AdminRoles))
	assert.Equal(t, "update-role", tenantData.AdminRoles[0])
	assert.Equal(t, 1, len(tenantData.AllowedClusters))
	assert.Equal(t, "standalone", tenantData.AllowedClusters[0])

	args = []string{"update", "--admin-roles", "new-role", "--allowed-clusters", "standalone", "update-tenant-test"}
	_, execErr, _, _ = TestTenantCommands(updateTenantCmd, args)
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
	_, _, nameErr, _ := TestTenantCommands(updateTenantCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestUpdateNonExistTenantError(t *testing.T) {
	args := []string{"update", "--admin-roles", "update-role", "--allowed-clusters", "standalone", "non-existent-topic"}
	_, execErr, _, _ := TestTenantCommands(updateTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}
