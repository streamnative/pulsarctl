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
	"fmt"
	"os"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteContextCmd(t *testing.T) {
	home := utils.HomeDir()
	path := fmt.Sprintf("%s/.config/pulsar/config", home)
	defer os.Remove(path)

	delArgs := []string{"delete", "test-delete-context"}
	out, execErr, err := TestConfigCommands(deleteContextCmd, delArgs)
	assert.Nil(t, err)

	expectedErr := fmt.Sprintf("cannot delete context test-delete-context, not in %s", path)
	assert.Equal(t, expectedErr, execErr.Error())
	assert.Equal(t, "", out.String())

	setArgs := []string{"set", "test-delete-context"}
	out, execErr, err = TestConfigCommands(setContextCmd, setArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "Context \"test-delete-context\" created.\n", out.String())

	out, execErr, err = TestConfigCommands(deleteContextCmd, delArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	warnOut := "warning: this removed your active context, " +
		"use \"pulsarctl context use\" to select a different one\n"
	expectedOut := fmt.Sprintf("deleted context test-delete-context from %s\n", path)
	assert.Equal(t, warnOut+expectedOut, out.String())
}
