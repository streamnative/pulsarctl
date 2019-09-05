package partitioned

import (
	"encoding/json"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateTopicCmd(t *testing.T) {
	args := []string{"create-partitioned-topic", "test-update-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-partitioned-topic-metadata", "test-update-topic"}
	out, execErr, _, _ := TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, partitions.Partitions)

	args = []string{"update-partitioned-topic", "test-update-topic", "3"}
	_, execErr, _, _ = TestTopicCommands(UpdateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-partitioned-topic-metadata", "test-update-topic"}
	out, execErr, _, _ = TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, partitions.Partitions)
}

func TestUpdateTopicArgsError(t *testing.T) {
	args := []string{"update-partitioned-topic", "test-topic"}
	_, _, nameErr, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the topic name and the partitions", nameErr.Error())
}

func TestUpdateTopicNotExist(t *testing.T) {
	args := []string{"update-partitioned-topic", "non-exist-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(UpdateTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Topic is not partitioned topic", execErr.Error())
}
