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

package cluster

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestGetFailureDomainSuccess(t *testing.T) {
	args := []string{"create", "failure-broker-A"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "failure-broker-B"}
	_, _, _, err = TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain",
		"-b", "failure-broker-A", "-b", "failure-broker-B", "standalone", "failure-domain"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"get-failure-domain", "standalone", "failure-domain"}
	out, _, _, err := TestClusterCommands(getFailureDomainCmd, args)
	assert.Nil(t, err)

	var brokers pulsar.FailureDomainData
	err = json.Unmarshal(out.Bytes(), &brokers)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "failure-broker-A", brokers.BrokerList[0])
	assert.Equal(t, "failure-broker-B", brokers.BrokerList[1])
}

func TestGetFailureDomainArgsError(t *testing.T) {
	args := []string{"get-failure-domain", "standalone"}
	_, _, nameErr, _ := TestClusterCommands(getFailureDomainCmd, args)
	assert.Equal(t, "need to specified the cluster name and the failure domain name", nameErr.Error())
}

func TestGetNonExistFailureDomain(t *testing.T) {
	args := []string{"get-failure-domain", "standalone", "non-exist"}
	_, execErr, _, _ := TestClusterCommands(getFailureDomainCmd, args)
	assert.NotNil(t, execErr)
}
