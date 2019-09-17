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
	"github.com/streamnative/pulsarctl/pkg/ctl/cluster"
	"github.com/streamnative/pulsarctl/pkg/ctl/tenant"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestClusters(t *testing.T) {
	clusterArgs := []string{"create", "test-replication-cluster", "--url", "192.168.12.11"}
	_, _, _, err := cluster.TestClusterCommands(cluster.CreateClusterCmd, clusterArgs)
	assert.Nil(t, err)

	updateTenantArgs := []string{"update", "--allowed-clusters", "test-replication-cluster", "--allowed-clusters", "standalone", "public"}
	_, execErr, _, err := tenant.TestTenantCommands(tenant.UpdateTenantCmd, updateTenantArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	setArgs := []string{"set-clusters", "public/default", "--clusters", "test-replication-cluster"}
	setOut, execErr, _, _ := TestNamespaceCommands(setReplicationClusters, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set replication clusters successfully for public/default")

	getArgs := []string{"get-clusters", "public/default"}
	getOut, execErr, _, _ := TestNamespaceCommands(getReplicationClusters, getArgs)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(getOut.String(), "test-replication-cluster"))

	// reset namespace clusters for other test case
	updateTenantArgs = []string{"update", "--allowed-clusters", "standalone", "public"}
	_, execErr, _, err = tenant.TestTenantCommands(tenant.UpdateTenantCmd, updateTenantArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	setArgs = []string{"set-clusters", "public/default", "--clusters", "standalone"}
	setOut, execErr, _, _ = TestNamespaceCommands(setReplicationClusters, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set replication clusters successfully for public/default")
}

func TestFailureCluster(t *testing.T) {
	setArgs := []string{"set-clusters", "public/default", "--clusters", "invalid-cluster"}
	_, execErr, _, _ := TestNamespaceCommands(setReplicationClusters, setArgs)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 403 reason: Invalid cluster id: invalid-cluster")
}
