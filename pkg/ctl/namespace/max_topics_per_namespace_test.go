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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTopicsPerNamespaceCmd(t *testing.T) {
	ns := "public/test-max-topics-per-namespace"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-max-topics-per-namespace", ns}
	initialOut, execErr, _, _ := TestNamespaceCommands(GetMaxTopicsPerNamespaceCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-max-topics-per-namespace", "--max-topics-per-namespace", "11", ns}
	setOut, execErr, _, _ := TestNamespaceCommands(SetMaxTopicsPerNamespaceCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the max topics per namespace of the namespace %s to %d\n", ns, 11),
		setOut.String())

	args = []string{"get-max-topics-per-namespace", ns}
	getOut, execErr, _, _ := TestNamespaceCommands(GetMaxTopicsPerNamespaceCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max topics per namespace of the namespace %s is %d\n", ns, 11),
		getOut.String())

	args = []string{"remove-max-topics-per-namespace", ns}
	removeOut, execErr, _, _ := TestNamespaceCommands(RemoveMaxTopicsPerNamespaceCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully removed the max topics per namespace of the namespace %s\n", ns),
		removeOut.String())

	args = []string{"get-max-topics-per-namespace", ns}
	getOut, execErr, _, _ = TestNamespaceCommands(GetMaxTopicsPerNamespaceCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialOut.String(), getOut.String())
}

func TestSetMaxTopicsPerNamespaceWithInvalidSize(t *testing.T) {
	args := []string{"set-max-topics-per-namespace", "--max-topics-per-namespace", "-1", "public/invalid-max-topics"}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxTopicsPerNamespaceCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the specified max topics value must bigger than 0", execErr.Error())
}
