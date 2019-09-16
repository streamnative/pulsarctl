package messages

import (
	"testing"
	"time"

	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/test"
	"github.com/stretchr/testify/assert"
)

func TestResetCursorCmd(t *testing.T) {
	args := []string{"create", "test-reset-cursor-topic", "test-reset-cursor-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"reset-cursor", "--time", "1m", "test-reset-cursor-topic", "test-reset-cursor-sub"}
	out, execErr, _, _ := TestSubCommands(ResetCursorCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Reset the cursor to 1m successfully", out.String())

	args = []string{"reset-cursor", "--message-id", "-1:-1", "test-reset-cursor-topic", "test-reset-cursor-sub"}
	out, execErr, _, _ = TestSubCommands(ResetCursorCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Reset the cursor to -1:-1 successfully", out.String())
}

func TestResetCursorArgsError(t *testing.T) {
	args := []string{"reset-cursor"}
	_, _, nameErr, _ := TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestResetCursorFlagError(t *testing.T) {
	args := []string{"reset-cursor", "test-reset-cursor-flag-topic", "flag-sub"}
	_, execErr, _, _ := TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "The reset position must be specified", execErr.Error())
}

func TestResetCursorNonExistingTopic(t *testing.T) {
	args := []string{"reset-cursor", "--time", "1m", "test-reset-cursor-non-existing-topic",
		"test-reset-cursor-non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestResetCursorNonExistingSub(t *testing.T) {
	args := []string{"create", "test-reset-cursor-non-existing-sub-topic", "test-reset-cursor-existing-sub"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"reset-cursor", "--time", "1m", "test-reset-cursor-non-existing-sub-topic",
		"test-reset-cursor-non-existing-sub"}
	_, execErr, _, _ = TestSubCommands(ResetCursorCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())
}

func TestParseTime(t *testing.T) {
	d, err := parseRelativeTimeInSeconds("1s")
	assert.Nil(t, err)
	assert.Equal(t, 1*time.Second, d)

	d, err = parseRelativeTimeInSeconds("1m")
	assert.Nil(t, err)
	assert.Equal(t, 1*time.Minute, d)

	d, err = parseRelativeTimeInSeconds("1h")
	assert.Nil(t, err)
	assert.Equal(t, 1*time.Hour, d)

	d, err = parseRelativeTimeInSeconds("1d")
	assert.Nil(t, err)
	assert.Equal(t, 24*time.Hour, d)

	d, err = parseRelativeTimeInSeconds("1w")
	assert.Nil(t, err)
	assert.Equal(t, 7*24*time.Hour, d)

	d, err = parseRelativeTimeInSeconds("1y")
	assert.Nil(t, err)
	assert.Equal(t, 365*7*24*time.Hour, d)
}
