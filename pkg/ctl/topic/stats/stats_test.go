package stats

import (
	"encoding/json"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

var defaultStats = pulsar.TopicStats{
	MsgRateIn:           0,
	MsgRateOut:          0,
	MsgThroughputIn:     0,
	MsgThroughputOut:    0,
	AverageMsgSize:      0,
	StorageSize:         0,
	Publishers:          []pulsar.PublisherStats{},
	Subscriptions:       map[string]pulsar.SubscriptionStats{},
	Replication:         map[string]pulsar.ReplicatorStats{},
	DeDuplicationStatus: "Disabled",
}

func TestGetStatsCmd(t *testing.T) {
	args := []string{"create", "test-non-partitioned-topic-stats", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "test-non-partitioned-topic-stats"}
	out, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.Nil(t, execErr)

	var stats pulsar.TopicStats
	err := json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, defaultStats, stats)
}

func TestGetPartitionedStats(t *testing.T) {
	args := []string{"create", "test-partitioned-topic-stats", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "test-partitioned-topic-stats"}
	_, execErr, _, _ = TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestGetStatsArgsError(t *testing.T) {
	args := []string{"stats"}
	_, _, nameErr, _ := TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetNonExistingTopicStats(t *testing.T) {
	args := []string{"stats", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
