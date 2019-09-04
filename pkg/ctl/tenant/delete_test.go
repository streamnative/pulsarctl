package tenant

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeleteTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "super-user", "delete-tenant-test"}
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
