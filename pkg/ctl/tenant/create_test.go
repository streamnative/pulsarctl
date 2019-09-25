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

	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestCreateTenantAlreadyExistError(t *testing.T) {
	args := []string{"create", "--allowed-clusters", "standalone", "create-tenant-duplicate"}
	_, execErr, _, _ := TestTenantCommands(createTenantCmd, args)
	assert.Nil(t, execErr)

	args = []string{"create", "--allowed-clusters", "standalone", "create-tenant-duplicate"}
	_, execErr, _, _ = TestTenantCommands(createTenantCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Tenant already exists", execErr.Error())
}
