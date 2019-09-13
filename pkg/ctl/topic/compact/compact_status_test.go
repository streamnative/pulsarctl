package compact

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestCompactStatusArgsError(t *testing.T) {
	args := []string{"compact-status"}
	_, _, nameErr, _ := TestTopicCommands(CompactStatusCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestCompactStatusNonExistingTopicError(t *testing.T) {
	args := []string{"compact-status", "test-non-existing-compact-status"}
	_, execErr, _, _ := TestTopicCommands(CompactStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestCompactStatusNonPersistentTopicError(t *testing.T) {
	args := []string{"compact-status", "non-persistent://public/default/test-non-persistent-topic-compact-status"}
	_, execErr, _, _ := TestTopicCommands(CompactStatusCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Need to provide a persistent topic.", execErr.Error())
}
