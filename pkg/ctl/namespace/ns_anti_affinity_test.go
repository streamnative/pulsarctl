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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNsAntiAffinityGroup(t *testing.T) {
	args := []string{"create", "public/test-anti-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-anti-namespace successfully\n")

	setArgs := []string{"set-anti-affinity-group", "public/test-anti-namespace", "--group", "test-ns"}
	setOut, execErr, _, _ := TestNamespaceCommands(setAntiAffinityGroup, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set the anti-affinity group: test-ns successfully for public/test-anti-namespace\n")

	getArgs := []string{"get-anti-affinity-group", "public/test-anti-namespace"}
	getOut, execErr, _, _ := TestNamespaceCommands(getAntiAffinityGroup, getArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, getOut.String(), "\"test-ns\"\n")

	delArgs := []string{"delete-anti-affinity-group", "public/test-anti-namespace"}
	delOut, execErr, _, _ := TestNamespaceCommands(deleteAntiAffinityGroup, delArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, delOut.String(), "Delete the anti-affinity group successfully for [public/test-anti-namespace]\n")

	getArgs = []string{"get-anti-affinity-group", "public/test-anti-namespace"}
	getOut, execErr, _, _ = TestNamespaceCommands(getAntiAffinityGroup, getArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, getOut.String(), "")
}
