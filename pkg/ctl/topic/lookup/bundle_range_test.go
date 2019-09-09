package lookup

import (
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBundleRangeCmd(t *testing.T) {
	args := []string{"create", "test-get-topic-bundle-range", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"bundle-range", "test-get-topic-bundle-range"}
	out, execErr, _, _ := TestTopicCommands(GetBundleRangeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "The bundle range of the topic "+
		"persistent://public/default/test-get-topic-bundle-range is: 0xc0000000_0xffffffff", out.String())
}

func TestGetBundleRangeArgError(t *testing.T) {
	args  := []string{"bundle-range"}
	_, _, nameErr, _ := TestTopicCommands(GetBundleRangeCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t,  "only one argument is allowed to be used as a name", nameErr.Error())
}
