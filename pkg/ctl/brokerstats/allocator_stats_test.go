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

package brokerstats

import (
	"encoding/json"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func TestDumpAllocatorStats(t *testing.T) {
	args := []string{"allocator-stats"}
	_, _, nameErr, _ := TestBrokerStatsCommands(dumpAllocatorStats, args)
	failOut := "the namespace name is not specified or the namespace name is specified more than one"
	assert.Equal(t, failOut, nameErr.Error())

	expectedStats, err := cmdutils.NewPulsarClient().BrokerStats().GetAllocatorStats("default")
	require.NoError(t, err)

	successArgs := []string{"allocator-stats", "default"}
	statsOut, execErr, nameErr, err := TestBrokerStatsCommands(dumpAllocatorStats, successArgs)
	require.NoError(t, err)
	require.NoError(t, nameErr)
	require.NoError(t, execErr)

	var allocatorStats utils.AllocatorStats
	err = json.Unmarshal(statsOut.Bytes(), &allocatorStats)
	require.NoError(t, err)
	assert.Equal(t, expectedStats.TinyCacheSize, allocatorStats.TinyCacheSize)
	assert.Equal(t, expectedStats.SmallCacheSize, allocatorStats.SmallCacheSize)
	assert.Equal(t, expectedStats.NormalCacheSize, allocatorStats.NormalCacheSize)
}
