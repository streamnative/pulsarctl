package compact

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestCompactCmd(t *testing.T)  {
	args := []string{"create", "test-compact-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"compact-status", "test-compact-topic"}
	out, execErr, _, _ := TestTopicCommands(CompactStatusCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Compaction has not been run for " +
		"persistent://public/default/test-compact-topic since broker startup/n", out.String())

	args = []string{"compact", "test-compact-topic"}
	_, execErr, _, _ = TestTopicCommands(CompactCmd, args)
	assert.Nil(t, execErr)

	args = []string{"compact-status", "test-compact-topic"}
	out, execErr, _, _ = TestTopicCommands(CompactStatusCmd, args)
	assert.Nil(t, execErr)

	assert.Equal(t, "Compaction is currently running/n", out.String())
}

func TestCompactArgError(t *testing.T) {
	args := []string{"compact"}
	_, _, nameErr, _ := TestTopicCommands(CompactCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestCompactNonExistingTopic(t *testing.T) {
	args := []string{"compact", "test-compact-non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(CompactCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestCompactNonPersistentTopic(t *testing.T) {
	args  := []string{"compact", "non-persistent://public/default/test-compact-non-persistent-topic"}
	_, execErr, _, _ := TestTopicCommands(CompactCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Need to provide a persistent topic.", execErr.Error())
}
