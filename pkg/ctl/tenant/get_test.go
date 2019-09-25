package tenant

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestGetTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "admin-get-test", "--allowed-clusters", "standalone", "get-tenant-test"}
	_, _, _, _ = TestTenantCommands(createTenantCmd, args)

	args = []string{"get", "get-tenant-test"}
	out, _, _, _ := TestTenantCommands(getTenantCmd, args)

	var tenantData pulsar.TenantData
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
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetNonExistTenant(t *testing.T) {
	args := []string{"get", "non-existent-tenant"}
	_, execErr, _, _ := TestTenantCommands(getTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}
