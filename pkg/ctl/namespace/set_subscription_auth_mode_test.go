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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestSetSubscriptionAuthModeCmd(t *testing.T) {
	ns := "public/test-subscription-auth-mode-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"set-subscription-auth-mode", "--mode", "Prefix", ns}
	out, execErr, _, _ := TestNamespaceCommands(SetSubscriptionAuthModeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the default subscription auth mode of namespace %s to %s", ns, "Prefix"),
		out.String())

	args = []string{"policies", ns}
	out, execErr, _, _ = TestNamespaceCommands(getPolicies, args)
	assert.Nil(t, execErr)

	var policies utils.Policies
	err := json.Unmarshal(out.Bytes(), &policies)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Prefix", policies.SubscriptionAuthMode.String())

	args = []string{"set-subscription-auth-mode", "--mode", "None", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSubscriptionAuthModeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the default subscription auth mode of namespace %s to %s", ns, "None"),
		out.String())

	args = []string{"policies", ns}
	out, execErr, _, _ = TestNamespaceCommands(getPolicies, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &policies)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "None", policies.SubscriptionAuthMode.String())
}

func TestSetSubscriptionAuthModeOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"set-subscription-auth-mode", "--mode", "Prefix", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSubscriptionAuthModeCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestSetSubscriptionAuthModeWithInvalidMode(t *testing.T) {
	ns := "public/test-invalid-subscription-auth-mode-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"set-subscription-auth-mode", "--mode", "Invalid", ns}
	_, execErr, _, _ = TestNamespaceCommands(SetSubscriptionAuthModeCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Invalid subscription auth mode", execErr.Error())
}
