package offload

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestOffloadCmd(t *testing.T) {
	args := []string{"create", "test-offload-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"offload", "test-offload-topic", "10M"}
	out, execErr, _, _ := TestTopicCommands(OffloadCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Nothing to offload", out.String())
}

func TestOffloadArgsError(t *testing.T)  {
	args :=  []string{"offload", "test-offload-topic-args-error"}
	_, _, nameErr, _ := TestTopicCommands(OffloadCmd, args)
	assert.Equal(t, "only two argument is allowed to be used as names", nameErr.Error())
}

func TestOffloadNonExistingTopicError(t *testing.T) {
	args := []string{"offload", "test-offload-non-existing-topic", "10m"}
	_, execErr, _, _ := TestTopicCommands(OffloadCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestValidateSize(t *testing.T) {
	s := "10k"
	size, err := validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024), size)

	s = "10K"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024), size)

	s = "10m"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024*1024), size)

	s = "10M"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024*1024), size)

	s = "10g"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024*1024*1024), size)

	s = "10G"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024*1024*1024), size)

	s = "10t"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024*1024*1024*1024), size)

	s = "10T"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(10*1024*1024*1024*1024), size)

	s = "2048000"
	size, err = validateSize(s)
	assert.Nil(t, err)
	assert.Equal(t, int64(2048000), size)
}
