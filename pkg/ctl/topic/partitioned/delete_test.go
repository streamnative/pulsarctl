package partitioned

import (
	"fmt"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeleteTopicCmd(t *testing.T) {
	args := []string{"create-partitioned-topic", "test-delete-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list-partitioned-topics", "public/default"}
	out, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	fmt.Println(out.String())
	assert.True(t, strings.Contains(out.String(), "test-delete-topic"))

	args = []string{"delete-partitioned-topic", "test-delete-topic"}
	_, execErr, _, _ = TestTopicCommands(DeleteTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list-partitioned-topics", "public/default"}
	out, execErr, _, _ = TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-topic"))
}

func TestDeleteTopicArgError(t *testing.T) {
	args := []string{"delete-partitioned-topic"}
	_, _, nameErr, _ := TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestDeleteNonExistTopic(t *testing.T) {
	args := []string{"delete-partitioned-topic", "non-exist-topic"}
	_, execErr, _, _ := TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Partitioned topic does not exist", execErr.Error())
}
