package crud

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestListTopicsCmd(t *testing.T) {
	args := []string{"list", "public/default"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
}

func TestListTopicArgError(t *testing.T) {
	args := []string{"list"}
	_, _, nameErr, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestListNonExistNamespace(t *testing.T) {
	args := []string{"list", "public/non-exist-namespace"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestListNonExistTenant(t *testing.T) {
	args := []string{"list", "non-exist-tenant/default"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}
