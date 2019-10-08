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

package crud

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestGetTopicCmd(t *testing.T) {
	args := []string{"create", "test-get-topic", "2"}
	_, execErr, _, _ := test.TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "test-get-topic"}
	out, execErr, _, _ := test.TestTopicCommands(GetTopicCmd, args)
	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, execErr)
	assert.Equal(t, 2, partitions.Partitions)
}

func TestBetNonPartitionedTopic(t *testing.T) {
	args := []string{"create", "test-get-non-partitioned-topic", "0"}
	_, execErr, _, _ := test.TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get", "test-get-non-partitioned-topic"}
	out, execErr, _, _ := test.TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)
	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, partitions.Partitions)
}

func TestGetTopicArgsError(t *testing.T) {
	args := []string{"get"}
	_, _, nameErr, _ := test.TestTopicCommands(GetTopicCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestGetNonExistTopic(t *testing.T) {
	args := []string{"get", "non-exist-topic"}
	out, execErr, _, _ := test.TestTopicCommands(GetTopicCmd, args)
	assert.Nil(t, execErr)

	var partitions pulsar.PartitionedTopicMetadata
	err := json.Unmarshal(out.Bytes(), &partitions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, partitions.Partitions)
}
