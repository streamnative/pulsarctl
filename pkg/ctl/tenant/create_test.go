package tenant

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreateTenantCmd(t *testing.T) {
	args := []string{"create", "--admin-roles", "super-user", "create-tenant-test"}
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
