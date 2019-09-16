package crud

import (
	"strings"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/test"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCmd(t *testing.T) {
	args := []string{"create", "test-delete-sub-topic", "test-delete-sub-1"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"create", "test-delete-sub-topic", "test-delete-sub-2"}
	_, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "test-delete-sub-topic"}
	out, execErr, _, _ := TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-sub-1"))
	assert.True(t, strings.Contains(out.String(), "test-delete-sub-2"))

	args = []string{"delete", "test-delete-sub-topic", "test-delete-sub-1"}
	out, execErr, _, _ = TestSubCommands(DeleteCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Delete subscription test-delete-sub-1 on the topic "+
		"persistent://public/default/test-delete-sub-topic successfully", out.String())

	args = []string{"list", "test-delete-sub-topic"}
	out, execErr, _, _ = TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.False(t, strings.Contains(out.String(), "test-delete-sub-1"))
	assert.True(t, strings.Contains(out.String(), "test-delete-sub-2"))
}

func TestDeleteArgsError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}

func TestDeleteNonExistingSub(t *testing.T) {
	args := []string{"create", "test-delete-non-existing-sub-topic", "test-delete-non-existing-sub-existing"}
	_, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"list", "test-delete-non-existing-sub-topic"}
	out, execErr, _, _ := TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-non-existing-sub-existing"))

	args = []string{"delete", "test-delete-non-existing-sub-topic", "test-delete-non-existing-sub-non-existing"}
	_, execErr, _, _ = TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Subscription not found", execErr.Error())

	args = []string{"list", "test-delete-non-existing-sub-topic"}
	out, execErr, _, _ = TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-delete-non-existing-sub-existing"))
}

func TestDeleteNonExistingTopicSub(t *testing.T) {
	args := []string{"delete", "non-existing-topic", "non-existing-topic-sub"}
	_, execErr, _, _ := TestSubCommands(DeleteCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
