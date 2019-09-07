package topic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTopicCmd(t *testing.T) {
	args := []string{"create", "test-create-topic", "2"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateNonPersistentTopic(t *testing.T) {
	args := []string{"create", "non-persistent://public/default/test-create-topic", "2"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateTopicDuplicate(t *testing.T) {
	args := []string{"create", "test-duplicate-topic", "2"}
	_, _, _, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, err)

	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Partitioned topic already exists", execErr.Error())
}

func TestCreateTopicArgsError(t *testing.T) {
	args := []string{"create", "topic"}
	_, _, nameErr, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the partitions", nameErr.Error())
}

func TestCreateNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-create-non-partitioned-topic", "0"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateNonPersistentNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "non-persistent://public/default/test-create-non-partitioned-topic", "0"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}
