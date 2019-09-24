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
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateFailureDomain(t *testing.T) {
	args := []string{"create-failure-domain", "-b", "127.0.0.1:6650", "standalone", "standalone-failure-domain"}
	_, execErr, NameErr, err := TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, NameErr)
	assert.Nil(t, err)

	args = []string{"get-failure-domain", "standalone", "standalone-failure-domain"}
	out,  execErr, _, _ := TestClusterCommands(getFailureDomainCmd, args)
	assert.Nil(t, execErr)

	var failureDomain  pulsar.FailureDomainData
	err = json.Unmarshal(out.Bytes(), &failureDomain)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(failureDomain.BrokerList))
	assert.Equal(t, "127.0.0.1:6650", failureDomain.BrokerList[0])

	args = []string{"update-failure-domain", "-b", "192.168.0.1:6650", "standalone",  "standalone-failure-domain"}
	_, execErr, _, _ = TestClusterCommands(updateFailureDomainCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-failure-domain", "standalone", "standalone-failure-domain"}
	out,  execErr, _, _ = TestClusterCommands(getFailureDomainCmd, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &failureDomain)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(failureDomain.BrokerList))
	assert.Equal(t, "192.168.0.1:6650", failureDomain.BrokerList[0])
}

func TestUpdateFailureDomainArgsError(t *testing.T) {
	args := []string{"update-failure-domain", "standalone", "standalone-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(updateFailureDomainCmd, args)
	assert.Equal(t, "broker list must be specified", execErr.Error())
}

func TestUpdateFailureDomainWithNonExistTopic(t *testing.T)  {
	args := []string{"update-failure-domain", "-b", "192.168.0.1:6650", "non-exist-cluster", "failure-domain"}
	_, execErr, _, _ := TestClusterCommands(updateFailureDomainCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 412 reason: Cluster non-exist-cluster does not exist.", execErr.Error())
}
