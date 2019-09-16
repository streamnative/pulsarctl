package messages

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/test"
	"github.com/stretchr/testify/assert"
)

func TestSkipCmd(t *testing.T) {
	args := []string{"create", "test-skip-messages-topic", "test-skip-messages-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"skip-messages", "--count", "1", "test-skip-messages-topic", "test-skip-messages-sub"}
	out, execErr, _, _ := TestSubCommands(SkipCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Subscription test-skip-messages-sub skip 1 messages on topic "+
		"persistent://public/default/test-skip-messages-topic successfully", out.String())

	args = []string{"skip-messages", "--count", "-1", "test-skip-messages-topic", "test-skip-messages-sub"}
	out, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Subscription test-skip-messages-sub skip -1 messages on topic "+
		"persistent://public/default/test-skip-messages-topic successfully", out.String())
}

func TestSkipArgsError(t *testing.T) {
	args := []string{"create"}
	_, _, nameErr, _ := TestSubCommands(CreateCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestSkipNonExistingTopic(t *testing.T) {
	args := []string{"skip-messages", "--count", "1", "test-skip-messages-non-existing-topic",
		"test-skip-messages-non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())

	args = []string{"skip-messages", "--count", "-1", "test-skip-messages-non-existing-topic",
		"test-skip-messages-non-existing-topic-sub"}
	_, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestSkipNonExistingSub(t *testing.T) {
	args := []string{"create", "test-skip-messages-non-existing-sub-topic",
		"test-skip-messages-non-existing-sub-existing"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"skip-messages", "--count", "1", "test-skip-messages-non-existing-sub-topic",
		"test-skip-messages-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())

	args = []string{"skip-messages", "--count", "-1", "test-skip-messages-non-existing-sub-topic",
		"test-skip-messages-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(SkipCmd, args)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())
}
