package crud

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeletePartitionedTopic(t *testing.T) {
	args := []string{"create", "test-delete-partitioned-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-partitioned-topic"))

	args = []string{"delete", "test-delete-partitioned-topic"}
	_, execErr, _, _ = TestTopicCommands(DeleteTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ = TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-partitioned-topic"))
}

func TestDeleteNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-delete-non-partitioned-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-non-partitioned-topic"))

	args = []string{"delete", "--non-partitioned", "test-delete-non-partitioned-topic"}
	_, execErr, _, _ = TestTopicCommands(DeleteTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "public/default"}
	out, execErr, _, _ = TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-non-partitioned-topic"))
}

func TestDeleteTopicArgError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestDeleteNonExistPartitionedTopic(t *testing.T) {
	args := []string{"delete", "non-existent-partitioned-topic"}
	_, execErr, _, _ := TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Partitioned topic does not exist", execErr.Error())
}

func TestDeleteNonExistNonPartitionedTopic(t *testing.T) {
	args := []string{"delete", "--non-partitioned", "non-existent-non-partitioned-topic"}
	_, execErr, _, _ := TestTopicCommands(DeleteTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
