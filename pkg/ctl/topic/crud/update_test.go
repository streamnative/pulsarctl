package crud

import (
	"encoding/json"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTopicCmd(t *testing.T) {
	args := []string{"create", "test-update-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "test-update-topic"}
	out, execErr, _, _ := TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, partitions.Partitions)

	args = []string{"update", "test-update-topic", "3"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "test-update-topic"}
	out, execErr, _, _ = TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, partitions.Partitions)
}

func TestUpdateTopicArgsError(t *testing.T) {
	args := []string{"update", "test-topic"}
	_, _, nameErr, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the partitions", nameErr.Error())
}

func TestUpdateTopicWithInvalidPartitions(t *testing.T) {
	args := []string{"update", "test-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number '0'", execErr.Error())

	args = []string{"update", "test-topic", "a"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number 'a'", execErr.Error())

	args = []string{"update", "test-topic", "--", "-1"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "invalid partition number '-1'", execErr.Error())
}

func TestUpdateTopicNotExist(t *testing.T) {
	args := []string{"update", "non-exist-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Topic is not partitioned topic", execErr.Error())
}

func TestUpdateNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-update-non-partitioned-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"update", "test-update-non-partitioned-topic", "3"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Topic is not partitioned topic", execErr.Error())
}
