package info

import (
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"

	"github.com/stretchr/testify/assert"
)

func TestGetLastMessageIdArgsError(t *testing.T) {
	args := []string{"last-message-id"}
	_, _, nameErr, _ := test.TestTopicCommands(GetLastMessageIDCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetLastMessageIdTopicNotExistError(t *testing.T) {
	args := []string{"last-message-id", "not-existent-topic"}
	_, execErr, _, _ := test.TestTopicCommands(GetLastMessageIDCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestGetLastMessageIdNotAllowedError(t *testing.T) {
	args := []string{"create",
		"non-persistent://public/default/last-message-id-non-persistent-topic", "0"}
	_, execErr, _, _ := test.TestTopicCommands(crud.CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"last-message-id", "non-persistent://public/default/last-message-id-non-persistent-topic"}
	_, execErr, _, _ = test.TestTopicCommands(GetLastMessageIDCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t,
		"code: 405 reason: GetLastMessageID on a non-persistent topic is not allowed",
		execErr.Error())
}
