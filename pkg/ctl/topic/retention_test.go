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
	"fmt"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/test"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestRetentionCmd(t *testing.T) {
	topic := fmt.Sprintf("test-retention-topic-%s", test.RandomSuffix())

	args := []string{"create", topic, "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-retention", topic, "--time", "12h", "--size", "100g"}
	out, execErr, nameErr, cmdErr := TestTopicCommands(SetRetentionCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, nameErr)
	assert.Nil(t, cmdErr)
	assert.NotNil(t, out)
	assert.NotEmpty(t, out.String())

	// waiting for the pulsar to be configured
	<-time.After(5 * time.Second)

	args = []string{"get-retention", topic}
	out, execErr, nameErr, cmdErr = TestTopicCommands(GetRetentionCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, nameErr)
	assert.Nil(t, cmdErr)
	assert.NotNil(t, out)
	assert.NotEmpty(t, out.String())

	var data utils.RetentionPolicies
	err := json.Unmarshal(out.Bytes(), &data)
	assert.Nil(t, err)
	assert.Equal(t, 720, data.RetentionTimeInMinutes)
	assert.Equal(t, int64(102400), data.RetentionSizeInMB)

	args = []string{"remove-retention", topic}
	out, execErr, nameErr, cmdErr = TestTopicCommands(RemoveRetentionCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, nameErr)
	assert.Nil(t, cmdErr)
	assert.NotNil(t, out)
	assert.NotEmpty(t, out.String())

	args = []string{"get-retention", topic}
	out, execErr, nameErr, cmdErr = TestTopicCommands(GetRetentionCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, nameErr)
	assert.Nil(t, cmdErr)
	assert.NotNil(t, out)
	assert.NotEmpty(t, out.String())

	data = utils.RetentionPolicies{}
	err = json.Unmarshal(out.Bytes(), &data)
	assert.Nil(t, err)
	assert.Equal(t, 0, data.RetentionTimeInMinutes)
	assert.Equal(t, int64(0), data.RetentionSizeInMB)
}
