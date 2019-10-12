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

func TestOffloadThresholdCmd(t *testing.T) {
	ns := "public/test-offload-threshold"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-offload-threshold", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetOffloadThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload threshold of the namespace %s is %d byte(s)\n", ns, -1),
		out.String())

	args = []string{"set-offload-threshold", "--size", "10m", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetOffloadThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the offload threshold of the namespace %s to %s\n", ns, "10m"),
		out.String())

	args = []string{"get-offload-threshold", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetOffloadThresholdCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload threshold of the namespace %s is %d byte(s)\n", ns, 10*1024*1024),
		out.String())
}

func TestSetOffloadThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-offload-threshold", "--size", "10m", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetOffloadThresholdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetOffloadThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-offload-threshold", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetOffloadThresholdCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
