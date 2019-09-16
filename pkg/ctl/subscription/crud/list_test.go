package crud

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/test"
	"github.com/stretchr/testify/assert"
)

func TestListArgError(t *testing.T) {
	args := []string{"list"}
	_, _, nameErr, _ := TestSubCommands(ListCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestListNonExistingTopicSub(t *testing.T) {
	args := []string{"list", "non-existing-topic"}
	_, execErr, _, _ := TestSubCommands(ListCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
