package messages

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/test"
	"github.com/stretchr/testify/assert"
)

func TestExpireCmd(t *testing.T) {
	args := []string{"create", "test-expire-messages-topic", "test-expire-messages-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"expire-messages", "--expire-time", "1", "test-expire-messages-topic", "test-expire-messages-sub"}
	out, execErr, _, _ := TestSubCommands(ExpireCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Expire messages after 1(s) for subscription test-expire-messages-sub of topic "+
		"persistent://public/default/test-expire-messages-topic successfully", out.String())

	args = []string{"expire-messages", "--expire-time", "1", "test-expire-messages-topic", "all"}
	out, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Expire messages after 1(s) for subscription all of topic "+
		"persistent://public/default/test-expire-messages-topic successfully", out.String())
}

func TestExpireArgsError(t *testing.T) {
	args := []string{"expire-messages"}
	_, _, _, err := TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"expire-time\" not set", err.Error())

	args = []string{"expire-messages", "--expire-time", "1"}
	_, _, nameErr, _ := TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestExpireNonExistingTopic(t *testing.T) {
	args := []string{"expire-messages", "--expire-time", "1", "non-existing-topic",
		"test-expire-messages-non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())

	args = []string{"expire-messages", "--expire-time", "1", "non-existing-topic", "all"}
	_, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestExpireNonExistingSub(t *testing.T) {
	args := []string{"create", "test-expire-messages-non-existing-sub-topic",
		"test-expire-messages-non-existing-sub-existing"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"expire-messages", "--expire-time", "1", "test-expire-messages-non-existing-sub-topic",
		"test-expire-messages-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(ExpireCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())
}
