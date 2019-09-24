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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteFailureDomainCmd(t *testing.T) {
	args := []string{"create", "delete-failure-test"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"create-failure-domain", "-b", "127.0.0.1:6650", "delete-failure-test", "delete-failure-domain"}
	_, _, _, err = TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, err)

	args = []string{"delete-failure-domain", "delete-failure-test", "delete-failure-domain"}
	_, _, _, err = TestClusterCommands(deleteFailureDomainCmd, args)
	assert.Nil(t, err)
}

func TestDeleteFailureDomainArgsError(t *testing.T) {
	args := []string{"delete-failure-domain", "standalone"}
	_, _, nameErr, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified the cluster name and the failure domain name", nameErr.Error())
}

// delete a non-existent failure domain in an existing cluster
func TestDeleteNonExistentFailureDomain(t *testing.T) {
	args := []string{"delete-failure-domain", "standalone", "non-existent-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Domain-name non-existent-failure-domain"+
		" or cluster standalone does not exist", execErr.Error())
}

// delete a non-existent failure domain in a non-existent cluster
func TestDeleteNonExistentFailureDomainInNonExistentCluster(t *testing.T) {
	args := []string{"delete-failure-domain", "non-existent-cluster", "non-existent-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(deleteFailureDomainCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 412 reason: Cluster non-existent-cluster does not exist.", execErr.Error())
}
