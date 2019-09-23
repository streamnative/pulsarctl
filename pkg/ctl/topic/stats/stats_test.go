package stats

import (
	"encoding/json"
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
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

func TestGetPartitionedStatsCmd(t *testing.T) {
	args := []string{"create", "test-topic-partitioned-stats", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "--partitioned-topic", "test-topic-partitioned-stats"}
	out, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.Nil(t, execErr)

	var stats pulsar.PartitionedTopicStats
	err := json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stats)

	assert.Equal(t, float64(0), stats.MsgRateIn)
	assert.Equal(t, float64(0), stats.MsgRateOut)
	assert.Equal(t, float64(0), stats.MsgThroughputIn)
	assert.Equal(t, float64(0), stats.MsgThroughputOut)
	assert.Equal(t, float64(0), stats.AverageMsgSize)
	assert.Equal(t, int64(0), stats.StorageSize)
	assert.Equal(t, 0, len(stats.Publishers))
	assert.Equal(t, 0, len(stats.Subscriptions))
	assert.Equal(t, 0, len(stats.Replication))
	assert.Equal(t, "", stats.DeDuplicationStatus)
	assert.Equal(t, 2, stats.Metadata.Partitions)
	assert.Equal(t, 0, len(stats.Partitions))
}

func TestGetPerPartitionedStatsCmd(t *testing.T) {
	args := []string{"create", "test-topic-per-partitioned-stats", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "--partitioned-topic", "--per-partition", "test-topic-per-partitioned-stats"}
	out, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.Nil(t, execErr)

	var stats pulsar.PartitionedTopicStats
	err := json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}

	defaultStats := pulsar.PartitionedTopicStats{
		MsgRateIn:           0,
		MsgRateOut:          0,
		MsgThroughputIn:     0,
		MsgThroughputOut:    0,
		AverageMsgSize:      0,
		StorageSize:         0,
		Publishers:          []pulsar.PublisherStats{},
		Subscriptions:       map[string]pulsar.SubscriptionStats{},
		Replication:         map[string]pulsar.ReplicatorStats{},
		DeDuplicationStatus: "",
		Metadata:            pulsar.PartitionedTopicMetadata{Partitions: 2},
		Partitions: map[string]pulsar.TopicStats{
			"persistent://public/default/test-topic-per-partitioned-stats": {
				MsgRateIn:           0,
				MsgRateOut:          0,
				MsgThroughputIn:     0,
				MsgThroughputOut:    0,
				AverageMsgSize:      0,
				StorageSize:         0,
				Publishers:          []pulsar.PublisherStats{},
				Subscriptions:       map[string]pulsar.SubscriptionStats{},
				Replication:         map[string]pulsar.ReplicatorStats{},
				DeDuplicationStatus: "",
			},
		},
	}

	assert.Equal(t, defaultStats, stats)
}

func TestGetPartitionedStatsArgError(t *testing.T) {
	args := []string{"stats", "--partitioned-topic"}
	_, _, nameErr, _ := TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetNonExistingTopicStatsError(t *testing.T) {
	args := []string{"stats", "--partitioned-topic", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Partitioned Topic not found", execErr.Error())
}

func TestGetNonPartitionedTopicStatsError(t *testing.T) {
	args := []string{"create", "test-non-partitioned-topic-partitioned-stats", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "--partitioned-topic", "test-non-partitioned-topic-partitioned-stats"}
	_, execErr, _, _ = TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Partitioned Topic not found", execErr.Error())
}
