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

func TestListFailureDomainsCmd(t *testing.T) {
	args := []string{"create", "list-failure-test"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "list-failure-broker-A"}
	_, _, _, err = TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create", "list-failure-broker-B"}
	_, _, _, err = TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--brokers", "list-failure-broker-A", "list-failure-test", "list-failure-A"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "--brokers", "list-failure-broker-B", "list-failure-test", "list-failure-B"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"list-failure-domains", "list-failure-test"}
	out, _, _, err := TestClusterCommands(listFailureDomainCmd, args)
	assert.Nil(t, err)

	var brokerMap pulsar.FailureDomainMap
	err = json.Unmarshal(out.Bytes(), &brokerMap)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, brokerMap["list-failure-A"])
	assert.Equal(t, "list-failure-broker-A", brokerMap["list-failure-A"].BrokerList[0])
	assert.NotNil(t, brokerMap["list-failure-B"])
	assert.Equal(t, "list-failure-broker-B", brokerMap["list-failure-B"].BrokerList[0])
}

func TestListFailureArgsError(t *testing.T) {
	args := []string{"list-failure-domains"}
	_, _, nameErr, _ := TestClusterCommands(listFailureDomainCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}
