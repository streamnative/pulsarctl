package info

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestGetInternalInfoArgError(t *testing.T) {
	args := []string{"internal-info"}
	_, _, nameErr, _ := TestTopicCommands(GetInternalInfoCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetNonExistingTopicInternalInfo(t *testing.T)  {
	args := []string{"internal-info", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(GetInternalInfoCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 500 reason: Unknown pulsar error", execErr.Error())
}
