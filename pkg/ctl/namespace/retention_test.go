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

package namespace

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestRetention(t *testing.T) {
	args := []string{"create", "public/test-retention"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-retention successfully\n")

	getArgs := []string{"get-retention", "public/test-retention"}
	getOut, execErr, _, _ := TestNamespaceCommands(getRetention, getArgs)
	assert.Nil(t, execErr)

	var retention pulsar.RetentionPolicies
	err = json.Unmarshal(getOut.Bytes(), &retention)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), retention.RetentionSizeInMB)
	assert.Equal(t, 0, retention.RetentionTimeInMinutes)

	setArgs := []string{"set-retention", "public/test-retention", "--time", "10m", "--size", "10M"}
	setOut, execErr, _, _ := TestNamespaceCommands(setRetention, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set retention successfully for [public/test-retention]\n")

	getArgs = []string{"get-retention", "public/test-retention"}
	getOut, execErr, _, _ = TestNamespaceCommands(getRetention, getArgs)
	assert.Nil(t, execErr)

	err = json.Unmarshal(getOut.Bytes(), &retention)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), retention.RetentionSizeInMB)
	assert.Equal(t, 10, retention.RetentionTimeInMinutes)

	// test negative value for time arg
	setArgWithTime := []string{"set-retention", "public/test-retention", "--time", "-10m", "--size", "10M"}
	_, execErr, _, _ = TestNamespaceCommands(setRetention, setArgWithTime)
	assert.Nil(t, execErr)

	getArgs = []string{"get-retention", "public/test-retention"}
	getOut, execErr, _, _ = TestNamespaceCommands(getRetention, getArgs)
	assert.Nil(t, execErr)

	err = json.Unmarshal(getOut.Bytes(), &retention)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), retention.RetentionSizeInMB)
	assert.Equal(t, -10, retention.RetentionTimeInMinutes)
}
