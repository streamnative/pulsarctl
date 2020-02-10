// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package topic

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

var defaultStats = utils.TopicStats{
	MsgRateIn:           0,
	MsgRateOut:          0,
	MsgThroughputIn:     0,
	MsgThroughputOut:    0,
	AverageMsgSize:      0,
	StorageSize:         0,
	Publishers:          []utils.PublisherStats{},
	Subscriptions:       map[string]utils.SubscriptionStats{},
	Replication:         map[string]utils.ReplicatorStats{},
	DeDuplicationStatus: "Disabled",
}

func TestGetStatsCmd(t *testing.T) {
	args := []string{"create", "test-non-partitioned-topic-stats", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "test-non-partitioned-topic-stats"}
	out, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.Nil(t, execErr)

	var stats utils.TopicStats
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
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
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

	var stats utils.PartitionedTopicStats
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
	args := []string{"create", "test-topic-per-partitioned-stats", "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"stats", "--partitioned-topic", "--per-partition", "test-topic-per-partitioned-stats"}
	out, execErr, _, _ := TestTopicCommands(GetStatsCmd, args)
	assert.Nil(t, execErr)

	var stats utils.PartitionedTopicStats
	err := json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}

	defaultStats := utils.PartitionedTopicStats{
		MsgRateIn:           0,
		MsgRateOut:          0,
		MsgThroughputIn:     0,
		MsgThroughputOut:    0,
		AverageMsgSize:      0,
		StorageSize:         0,
		Publishers:          []utils.PublisherStats{},
		Subscriptions:       map[string]utils.SubscriptionStats{},
		Replication:         map[string]utils.ReplicatorStats{},
		DeDuplicationStatus: "",
		Metadata:            utils.PartitionedTopicMetadata{Partitions: 1},
		Partitions: map[string]utils.TopicStats{
			"persistent://public/default/test-topic-per-partitioned-stats-partition-0": {
				MsgRateIn:           0,
				MsgRateOut:          0,
				MsgThroughputIn:     0,
				MsgThroughputOut:    0,
				AverageMsgSize:      0,
				StorageSize:         0,
				Publishers:          []utils.PublisherStats{},
				Subscriptions:       map[string]utils.SubscriptionStats{},
				Replication:         map[string]utils.ReplicatorStats{},
				DeDuplicationStatus: "Disabled",
			},
		},
	}

	assert.Equal(t, defaultStats, stats)
}

func TestGetPartitionedStatsArgError(t *testing.T) {
	args := []string{"stats", "--partitioned-topic"}
	_, _, nameErr, _ := TestTopicCommands(GetStatsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
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
