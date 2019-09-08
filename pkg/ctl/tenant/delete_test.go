package tenant

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
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
	args :=  []string{"delete", "public"}
	_, execErr, _, _ := TestTenantCommands(deleteTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: The tenant still has active namespaces", execErr.Error())
}
