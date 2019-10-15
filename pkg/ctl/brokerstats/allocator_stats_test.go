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

	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestDumpAllocatorStats(t *testing.T) {
	args := []string{"allocator-stats"}
	_, _, nameErr, _ := TestBrokerStatsCommands(dumpAllocatorStats, args)
	failOut := "the namespace name is not specified or the namespace name is specified more than one"
	assert.Equal(t, failOut, nameErr.Error())

	successArgs := []string{"allocator-stats", "default"}
	statsOut, _, nameErr, _ := TestBrokerStatsCommands(dumpAllocatorStats, successArgs)
	assert.Nil(t, nameErr)

	var allocatorStats pulsar.AllocatorStats
	err := json.Unmarshal(statsOut.Bytes(), &allocatorStats)
	assert.Nil(t, err)
	assert.Equal(t, 512, allocatorStats.TinyCacheSize)
	assert.Equal(t, 256, allocatorStats.SmallCacheSize)
	assert.Equal(t, 64, allocatorStats.NormalCacheSize)
	assert.Equal(t, 24, allocatorStats.NumDirectArenas)
}
