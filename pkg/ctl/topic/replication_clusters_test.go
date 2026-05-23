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

package topic

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/cluster"
	"github.com/streamnative/pulsarctl/pkg/ctl/tenant"
	"github.com/streamnative/pulsarctl/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestReplicationClustersCmd(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-replication-clusters-topic-%s", test.RandomSuffix())

	// Create the topic first
	args := []string{"create", topicName, "0"}
	_, execErr, _, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	// Create a test cluster for replication
	clusterArgs := []string{"create", "test-replication-cluster", "--url", "http://192.168.12.11:8080"}
	_, execErr, _, err = cluster.TestClusterCommands(cluster.CreateClusterCmd, clusterArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	// Update tenant to allow the new cluster
	updateTenantArgs := []string{"update", "--allowed-clusters", "test-replication-cluster",
		"--allowed-clusters", "standalone", "public"}
	_, execErr, _, err = tenant.TestTenantCommands(tenant.UpdateTenantCmd, updateTenantArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	// Set replication clusters for the topic
	setArgs := []string{"set-replication-clusters", topicName, "--clusters", "test-replication-cluster"}
	setOut, execErr, _, _ := TestTopicCommands(SetReplicationClustersCmd, setArgs)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(setOut.String(), "Set replication clusters successfully"))

	// Get replication clusters for the topic
	getArgs := []string{"get-replication-clusters", topicName}
	getOut, execErr, _, _ := TestTopicCommands(GetReplicationClustersCmd, getArgs)
	assert.Nil(t, execErr)

	var clusters []string
	err = json.Unmarshal(getOut.Bytes(), &clusters)
	assert.Nil(t, err)
	assert.Contains(t, clusters, "test-replication-cluster")

	// Reset tenant clusters for other test cases
	updateTenantArgs = []string{"update", "--allowed-clusters", "standalone", "public"}
	_, execErr, _, err = tenant.TestTenantCommands(tenant.UpdateTenantCmd, updateTenantArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
}

func TestSetReplicationClustersArgError(t *testing.T) {
	// Test with no topic name
	args := []string{"set-replication-clusters", "--clusters", "standalone"}
	_, _, nameErr, _ := TestTopicCommands(SetReplicationClustersCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestGetReplicationClustersArgError(t *testing.T) {
	// Test with no topic name
	args := []string{"get-replication-clusters"}
	_, _, nameErr, _ := TestTopicCommands(GetReplicationClustersCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestSetReplicationClustersInvalidCluster(t *testing.T) {
	topicName := fmt.Sprintf("persistent://public/default/test-invalid-cluster-topic-%s", test.RandomSuffix())

	// Create the topic first
	args := []string{"create", topicName, "0"}
	_, execErr, _, err := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	// Try to set an invalid cluster
	setArgs := []string{"set-replication-clusters", topicName, "--clusters", "invalid-cluster"}
	_, execErr, _, _ = TestTopicCommands(SetReplicationClustersCmd, setArgs)
	assert.NotNil(t, execErr)
}
