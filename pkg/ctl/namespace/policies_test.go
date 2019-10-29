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
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestPolicesCommand(t *testing.T) {
	args := []string{"create", "public/test-policy-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-policy-namespace successfully\n")

	args = []string{"policies", "public/test-policy-namespace"}
	out, execErr, _, _ := TestNamespaceCommands(getPolicies, args)
	assert.Nil(t, execErr)

	var police pulsar.Policies
	err = json.Unmarshal(out.Bytes(), &police)
	assert.Nil(t, err)

	assert.Equal(t, police.DeduplicationEnabled, false)
	assert.Equal(t, police.Deleted, false)
	for key, value := range police.ClusterSubscribeRate {
		exceptedValue := pulsar.SubscribeRate{
			SubscribeThrottlingRatePerConsumer: 0,
			RatePeriodInSecond:                 30,
		}
		assert.Equal(t, key, "standalone")
		assert.Equal(t, exceptedValue, value)
	}
}

func TestPolicesNsArgsError(t *testing.T) {
	args := []string{"policies"}
	_, _, nameErr, _ := TestNamespaceCommands(getPolicies, args)
	assert.Equal(t, "the namespace name is not specified or the namespace name is "+
		"specified more than one", nameErr.Error())
}

func TestPolicesNonExistTenant(t *testing.T) {
	args := []string{"policies", "non-existent-tenant/default"}
	_, execErr, _, _ := TestNamespaceCommands(getPolicies, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}

func TestPolicesNonExistNs(t *testing.T) {
	args := []string{"policies", "public/test-not-exist-ns"}
	_, execErr, _, _ := TestNamespaceCommands(getPolicies, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
