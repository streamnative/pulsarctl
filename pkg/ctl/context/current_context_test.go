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

package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentContextCmd(t *testing.T) {
	currentArgs := []string{"current"}
	out, execErr, err := TestConfigCommands(currentContextCmd, currentArgs)
	assert.Nil(t, err)
	assert.Equal(t, "test-set-context\n", out.String())
	assert.Nil(t, execErr)

	setArgs := []string{"set", "test-current-context"}
	out, execErr, err = TestConfigCommands(setContextCmd, setArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "Context \"test-current-context\" created.\n", out.String())

	useArgs := []string{"use", "test-current-context"}
	out, execErr, err = TestConfigCommands(useContextCmd, useArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "Switched to context \"test-current-context\".\n", out.String())

	out, execErr, err = TestConfigCommands(currentContextCmd, currentArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "test-current-context\n", out.String())

	getArgs := []string{"get"}
	out, execErr, err = TestConfigCommands(getContextsCmd, getArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	expectedOut := "+---------+----------------------+--------------------+--------------------+\n| CURRENT |    " +
		"     NAME         | BROKER SERVICE URL | BOOKIE SERVICE URL |\n+---------+----------------------+-----" +
		"---------------+--------------------+\n| *       | test-current-context |                    |          " +
		"          |\n+---------+----------------------+--------------------+--------------------+\n"
	assert.Equal(t, expectedOut, out.String())
}
