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

func TestBrokerAndBrokers(t *testing.T) {
	brokersFailArgs := []string{"brokers", "standalone"}
	_, execErr, _, _ := TestNsIsolationPolicyCommands(getAllBrokersWithPolicies, brokersFailArgs)
	assert.NotNil(t, execErr)
	exceptedErr := "code: 404 reason: namespace-isolation policies not found for standalone"
	assert.Equal(t, exceptedErr, execErr.Error())

	brokerFailArgs := []string{"broker", "standalone", "--broker", "127.0.0.1:8080"}
	_, execErr, _, _ = TestNsIsolationPolicyCommands(getBrokerWithPolicies, brokerFailArgs)
	assert.NotNil(t, execErr)
	exceptedErr = "code: 404 reason: namespace-isolation policies not found for standalone"
	assert.Equal(t, exceptedErr, execErr.Error())

	setPolicyArgs := []string{"set", "standalone", "test-policy-1",
		"--auto-failover-policy-params", "min_limit=3,usage_threshold=100",
		"--auto-failover-policy-type", "min_available",
		"--namespaces", "default",
		"--primary", "test-primary", "--secondary", "test-secondary"}
	setPolicyOut, execErr, _, _ := TestNsIsolationPolicyCommands(setPolicy, setPolicyArgs)
	assert.Nil(t, execErr)
	expectedOut := "Create/Update namespaces isolation policy:test-policy-1 successful."
	assert.Equal(t, expectedOut, setPolicyOut.String())

	brokersArgs := []string{"brokers", "standalone"}
	brokersOut, execErr, _, _ := TestNsIsolationPolicyCommands(getAllBrokersWithPolicies, brokersArgs)
	assert.Nil(t, execErr)

	var brokersData []pulsar.BrokerNamespaceIsolationData
	err := json.Unmarshal(brokersOut.Bytes(), &brokersData)
	assert.Nil(t, err)
	assert.Equal(t, "127.0.0.1:8080", brokersData[0].BrokerName)
	assert.Equal(t, false, brokersData[0].IsPrimary)

	brokerArgs := []string{"broker", "standalone", "--broker", "127.0.0.1:8080"}
	brokerOut, execErr, _, _ := TestNsIsolationPolicyCommands(getBrokerWithPolicies, brokerArgs)
	assert.Nil(t, execErr)

	var brokerData pulsar.BrokerNamespaceIsolationData
	err = json.Unmarshal(brokerOut.Bytes(), &brokerData)
	assert.Nil(t, err)
	assert.Equal(t, "127.0.0.1:8080", brokerData.BrokerName)
	assert.Equal(t, false, brokerData.IsPrimary)

	deleteArgs := []string{"delete", "standalone", "test-policy-1"}
	deleteOut, execErr, _, _ := TestNsIsolationPolicyCommands(deleteNsIsolationPolicy, deleteArgs)
	assert.Nil(t, execErr)
	expectedDelOut := "Delete namespaces isolation policy: test-policy-1 successful."
	assert.Equal(t, expectedDelOut, deleteOut.String())
}
