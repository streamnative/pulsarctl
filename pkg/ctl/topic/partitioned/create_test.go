package partitioned

import (
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTopicCmd(t *testing.T) {
	args := []string{"create-partitioned-topic", "test-create-topic", "2"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateNonPersistentTopic(t *testing.T) {
	args := []string{"create-partitioned-topic", "non-persistent://public/default/test-create-topic", "2"}
	_, execErr, argsErr, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)
}

func TestCreateTopicDuplicate(t *testing.T) {
	args := []string{"create-partitioned-topic", "test-duplicate-topic", "2"}
	_, _, _, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, err)

	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Partitioned topic already exists", execErr.Error())
}

func TestCreateTopicArgsError(t *testing.T) {
	args := []string{"create-partitioned-topic", "topic"}
	_, _, nameErr, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the partitions", nameErr.Error())
}
