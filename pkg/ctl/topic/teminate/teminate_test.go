package teminate

import (
	"strings"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestTerminateCmd(t *testing.T) {
	args := []string{"create", "test-terminate-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"terminate", "test-terminate-topic"}
	out, execErr, _, _ := TestTopicCommands(TerminateCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.HasPrefix(out.String(),
		"Topic persistent://public/default/test-terminate-topic successfully terminated at"))
}

func TestTerminatePartitionedTopicError(t *testing.T)  {
	args := []string{"create",  "test-terminate-partitioned-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"terminate", "test-terminate-partitioned-topic"}
	_, execErr, _, _ = TestTopicCommands(TerminateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 405 reason: Termination of a partitioned topic is not allowed", execErr.Error())
}

func TestTerminateArgError(t *testing.T) {
	args := []string{"terminate"}
	_, _, nameErr, _ := TestTopicCommands(TerminateCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestTerminateNonExistingTopic(t *testing.T)  {
	args := []string{"terminate", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(TerminateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
