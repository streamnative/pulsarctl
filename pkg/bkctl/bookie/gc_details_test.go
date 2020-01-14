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

package bookie

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper/bkdata"
	"github.com/stretchr/testify/assert"
)

func TestGCDetailsCmd(t *testing.T) {
	args := []string{"gc-details"}
	out, execErr, nameErr, err := testBookieCommands(gcDetailsCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, nameErr)
	assert.Nil(t, execErr)

	var result []bkdata.GCStatus
	json.Unmarshal(out.Bytes(), &result)
	assert.Equal(t, 1, len(result))
	assert.False(t, result[0].ForceCompacting)
	assert.False(t, result[0].MajorCompacting)
	assert.False(t, result[0].MinorCompacting)
	assert.NotEqual(t, 0, result[0].LastMajorCompactionTime)
	assert.NotEqual(t, 0, result[0].LastMinorCompactionTime)
	assert.Equal(t, int64(0), result[0].MajorCompactionCounter)
	assert.Equal(t, int64(0), result[0].MinorCompactionCounter)
}
