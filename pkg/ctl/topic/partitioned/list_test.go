package partitioned

import (
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListTopicsCmd(t *testing.T) {
	args := []string{"list-partitioned-topics", "public/default"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
}

func TestListTopicArgError(t *testing.T) {
	args := []string{"list-partitioned-topics"}
	_, _, nameErr, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestListNonExistNamespace(t *testing.T) {
	args := []string{"list-partitioned-topics", "public/non-exist-namespace"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
}

func TestListNonExistTenant(t *testing.T) {
	args := []string{"list-partitioned-topics", "non-exist-tenant/default"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}
