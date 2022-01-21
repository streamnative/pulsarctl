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

func TestOffloadDeletionLagCmd(t *testing.T) {
	ns := "public/test-offload-deletion-lag"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-offload-deletion-lag", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload deletion lag of the namespace %s is %f minute(s)\n", ns, 0.000000),
		out.String())

	args = []string{"set-offload-deletion-lag", "--lag", "10m", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the offload deletion lag of the namespace %s to %s\n", ns, "10m"),
		out.String())

	args = []string{"get-offload-deletion-lag", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload deletion lag of the namespace %s is %f minute(s)\n", ns, 10.000000),
		out.String())

	args = []string{"clear-offload-deletion-lag", ns}
	out, execErr, _, _ = TestNamespaceCommands(ClearOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully clear the offload deletion lag of the namespace %s\n", ns),
		out.String())

	args = []string{"get-offload-deletion-lag", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The offload deletion lag of the namespace %s is %f minute(s)\n", ns, 0.000000),
		out.String())

}

func TestOffloadDeletionLagOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-offload-deletion-lag", "--lag", "10m", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetOffloadDeletionLagCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetOfloadThresholdOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-offload-deletion-lag", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetOffloadDeletionLagCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestClearOffloadDeletionLagOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"clear-offload-deletion-lag", ns}
	_, execErr, _, _ := TestNamespaceCommands(ClearOffloadDeletionLagCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
