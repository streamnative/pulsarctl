package crud

import (
	"strings"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/subscription/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateCmd(t *testing.T) {
	args := []string{"create", "test-sub-topic", "test-sub-default-arg"}
	out, execErr, _, _ := TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Create subscription test-sub-default-arg on topic persistent://public/default/test-sub-topic from latest successfully", out.String())

	args = []string{"create", "--messageId", "earliest", "test-sub-topic", "test-sub-earliest"}
	out, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Create subscription test-sub-earliest on topic persistent://public/default/test-sub-topic from earliest successfully", out.String())

	args = []string{"create", "--messageId", "-1:-1", "test-sub-topic", "test-sub-messageId"}
	out, execErr, _, _ = TestSubCommands(CreateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Create subscription test-sub-messageId on topic persistent://public/default/test-sub-topic from -1:-1 successfully", out.String())

	args = []string{"list", "test-sub-topic"}
	out, execErr, _, _ = TestSubCommands(ListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(out.String(), "test-sub-default-arg"))
	assert.True(t, strings.Contains(out.String(), "test-sub-earliest"))
	assert.True(t, strings.Contains(out.String(), "test-sub-messageId"))
}

func TestCreateArgsError(t *testing.T) {
	args := []string{"create"}
	_, _, nameErr, _ := TestSubCommands(CreateCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the subscription name", nameErr.Error())
}
