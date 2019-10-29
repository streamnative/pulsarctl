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

package brokers

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateDynamicConfig(t *testing.T) {
	args := []string{"update-dynamic-config", "--config", "dispatchThrottlingRatePerTopicInMsg", "--value", "true"}
	listOut, execErr, _, _ := TestBrokersCommands(updateDynamicConfig, args)
	assert.Nil(t, execErr)
	expectedOut := "Update dynamic config: dispatchThrottlingRatePerTopicInMsg successful\n"
	assert.Equal(t, expectedOut, listOut.String())

	failArgs := []string{"update-dynamic-config", "--config", "errorName", "--value", "true"}
	_, nameErr, _, _ := TestBrokersCommands(updateDynamicConfig, failArgs)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "code: 412 reason:  Can't update non-dynamic configuration", nameErr.Error())

	getArgs := []string{"get-all-dynamic-config"}
	getOut, execErr, _, _ := TestBrokersCommands(getAllDynamicConfigsCmd, getArgs)
	assert.Nil(t, execErr)
	var tmpMap map[string]string
	err := json.Unmarshal(getOut.Bytes(), &tmpMap)
	assert.Nil(t, err)
	assert.Equal(t, "true", tmpMap["dispatchThrottlingRatePerTopicInMsg"])
}
