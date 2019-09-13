package offload

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestOffloadStatusCmd(t *testing.T) {
	args := []string{"create", "test-offload-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"offload-status", "test-offload-topic"}
	out, execErr, _, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Offload has not been run for " +
		"persistent://public/default/test-offload-topic since broker startup/n", out.String())
}

func TestOffloadStatusArgsError(t *testing.T) {
	args := []string{"offload-status"}
	_, _, nameErr, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestOffloadStatusNonPersistentTopicError(t *testing.T)  {
	args := []string{"offload-status", "non-persistent://public/default/test-offload-status-non-persistent-topic"}
	_, execErr, _, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Need to provide a persistent topic.", execErr.Error())
}

func TestOffloadStatusNonExistingTopicError(t *testing.T)  {
	args := []string{"offload-status", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(OffloadStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
