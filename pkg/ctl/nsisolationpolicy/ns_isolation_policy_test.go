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

package nsisolationpolicy

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestNsIsolationPolicy(t *testing.T) {
	setPolicyArgs := []string{"set", "standalone", "test-policy-2",
		"--auto-failover-policy-params", "min_limit=3,usage_threshold=100",
		"--auto-failover-policy-type", "min_available",
		"--namespaces", "default",
		"--primary", "test-primary", "--secondary", "test-secondary"}
	setPolicyOut, execErr, _, _ := TestNsIsolationPolicyCommands(setPolicy, setPolicyArgs)
	assert.Nil(t, execErr)
	expectedOut := "Create/Update namespaces isolation policy:test-policy-2 successful\n"
	assert.Equal(t, expectedOut, setPolicyOut.String())

	getArgs := []string{"get", "standalone", "test-policy-2"}
	getOut, execErr, _, _ := TestNsIsolationPolicyCommands(getNsIsolationPolicy, getArgs)
	assert.Nil(t, execErr)

	var getData pulsar.NamespaceIsolationData
	err := json.Unmarshal(getOut.Bytes(), &getData)
	assert.Nil(t, err)
	assert.Equal(t, "default", getData.Namespaces[0])
	assert.Equal(t, "test-primary", getData.Primary[0])
	assert.Equal(t, "test-secondary", getData.Secondary[0])

	listArgs := []string{"list", "standalone"}
	listOut, execErr, _, _ := TestNsIsolationPolicyCommands(getNsIsolationPolicies, listArgs)
	assert.Nil(t, execErr)

	var listData map[string]pulsar.NamespaceIsolationData
	err = json.Unmarshal(listOut.Bytes(), &listData)
	assert.Nil(t, err)

	_, ok := listData["test-policy-2"]
	assert.Equal(t, true, ok)

	deleteArgs := []string{"delete", "standalone", "test-policy-2"}
	deleteOut, execErr, _, _ := TestNsIsolationPolicyCommands(deleteNsIsolationPolicy, deleteArgs)
	assert.Nil(t, execErr)
	expectedDelOut := "Delete namespaces isolation policy: test-policy-2 successful\n"
	assert.Equal(t, expectedDelOut, deleteOut.String())
}
