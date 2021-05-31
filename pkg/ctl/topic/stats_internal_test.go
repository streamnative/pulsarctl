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

func TestGetInternalStatsCmd(t *testing.T) {
	args := []string{"create", "test-topic-internal-stats", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"internal-stats", "test-topic-internal-stats"}
	out, execErr, _, _ := TestTopicCommands(GetInternalStatsCmd, args)
	assert.Nil(t, execErr)

	var stats utils.PersistentTopicInternalStats
	err := json.Unmarshal(out.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}

	defaultLedgerInfo := []utils.LedgerInfo{
		{
			Entries:   0,
			Size:      0,
			Timestamp: 0,
			Offloaded: false,
		},
	}

	defaultStats := utils.PersistentTopicInternalStats{
		EntriesAddedCounter:    0,
		NumberOfEntries:        0,
		TotalSize:              0,
		CurrentLedgerEntries:   0,
		CurrentLedgerSize:      0,
		WaitingCursorsCount:    0,
		PendingAddEntriesCount: 0,
		State:                  "LedgerOpened",
		Ledgers:                defaultLedgerInfo,
	}

	assert.Equal(t, defaultStats.EntriesAddedCounter, stats.EntriesAddedCounter)
	assert.Equal(t, defaultStats.NumberOfEntries, stats.NumberOfEntries)
	assert.Equal(t, defaultStats.TotalSize, stats.TotalSize)
	assert.Equal(t, defaultStats.CurrentLedgerEntries, stats.CurrentLedgerEntries)
	assert.Equal(t, defaultStats.CurrentLedgerSize, stats.CurrentLedgerSize)
	assert.Equal(t, defaultStats.WaitingCursorsCount, stats.WaitingCursorsCount)
	assert.Equal(t, defaultStats.PendingAddEntriesCount, stats.PendingAddEntriesCount)
	assert.Equal(t, defaultStats.State, stats.State)
	assert.Equal(t, defaultStats.Ledgers[0].Entries, stats.Ledgers[0].Entries)
	assert.Equal(t, defaultStats.Ledgers[0].Size, stats.Ledgers[0].Size)
	assert.Equal(t, defaultStats.Ledgers[0].Timestamp, stats.Ledgers[0].Timestamp)
	assert.Equal(t, defaultStats.Ledgers[0].Offloaded, stats.Ledgers[0].Offloaded)
}

func TestGetPartitionedTopicInternalStatsError(t *testing.T) {
	args := []string{"create", "test-partitioned-topic-internal-stats-error", "2"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"internal-stats", "test-partitioned-topic-internal-stats-error"}
	_, execErr, _, _ = TestTopicCommands(GetInternalStatsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}

func TestGetInternalStatsArgsError(t *testing.T) {
	args := []string{"internal-stats"}
	_, _, nameErr, _ := TestTopicCommands(GetInternalStatsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestGetNonExistingTopic(t *testing.T) {
	args := []string{"internal-stats", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(GetInternalStatsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
