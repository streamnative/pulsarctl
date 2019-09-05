package partitioned

import (
	"encoding/json"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopicCmd(t *testing.T) {
	args := []string{"create-partitioned-topic", "test-get-topic", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-partitioned-topic-metadata", "test-get-topic"}
	out, execErr, _, _ := TestTopicCommands(GetTopicCmd, args)
	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, execErr)
	assert.Equal(t, 2, partitions.Partitions)
}

func TestGetTopicArgsError(t *testing.T) {
	args := []string{"get-partitioned-topic-metadata"}
	_, _, nameErr, _ := TestTopicCommands(GetTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetNonExistTopic(t *testing.T) {
	args := []string{"get-partitioned-topic-metadata", "non-exist-topic"}
	out, execErr, _, _ := TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)

	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, partitions.Partitions)
}
