package unload

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestUnloadCmd(t *testing.T) {
	args := []string{"create", "test-unload-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"unload", "test-unload-topic"}
	out, execErr, _, _ := TestTopicCommands(UnloadCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Unload topic persistent://public/default/test-unload-topic successfully/n", out.String())
}

func TestUnloadArgError(t *testing.T) {
	args := []string{"unload"}
	_, _, nameErr, _ := TestTopicCommands(UnloadCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestUnloadNonExistingTopic(t *testing.T)  {
	args  := []string{"unload", "test-unload-non-existing-topic"}
	_, execErr, _, _ :=   TestTopicCommands(UnloadCmd,  args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
