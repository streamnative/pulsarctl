package permission

import (
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"

	"github.com/stretchr/testify/assert"
)

func TestGetPermissionsArgError(t *testing.T) {
	args := []string{"get-permissions"}
	_, _, nameErr, _ := test.TestTopicCommands(GetPermissionsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}
