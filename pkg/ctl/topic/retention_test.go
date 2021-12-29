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

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
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

	args = []string{"get-retention", topic}
	var data utils.RetentionPolicies
	task := func(args []string, obj interface{}) bool {
		out, execErr, _, _= TestTopicCommands(GetRetentionCmd, args)
		if execErr != nil {
			return false
		}
		err := json.Unmarshal(out.Bytes(), obj)
		if err != nil {
			return false
		}
		d := obj.(*utils.RetentionPolicies)
		return d.RetentionTimeInMinutes == 720 && d.RetentionSizeInMB == int64(102400)
	}
	err := cmdutils.RunFuncWithTimeout(task, true, 30 * time.Second, args, &data)
	if err != nil {
		t.Fatal(err)
	}

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
