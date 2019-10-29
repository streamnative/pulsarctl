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
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/cluster"
	"github.com/streamnative/pulsarctl/pkg/ctl/tenant"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestCreateNs(t *testing.T) {
	args := []string{"create", "public/test-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-namespace successfully\n")

	args = []string{"list", "public"}
	out, _, _, _ := TestNamespaceCommands(getNamespacesFromTenant, args)
	assert.True(t, strings.Contains(out.String(), "public/test-namespace"))

	policiesArgs := []string{"policies", "public/test-namespace"}
	out, execErr, _, _ := TestNamespaceCommands(getPolicies, policiesArgs)
	assert.Nil(t, execErr)

	var police pulsar.Policies
	err = json.Unmarshal(out.Bytes(), &police)
	assert.Nil(t, err)

	for cluster := range police.ClusterSubscribeRate {
		assert.Equal(t, cluster, "standalone")
	}

	assert.Equal(t, police.Bundles.NumBundles, 4)
}

func TestCreateNsForBundles(t *testing.T) {
	args := []string{"create", "public/test-namespace-bundles", "--bundles", "0"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	t.Log(createOut.String())

	policiesArgs := []string{"policies", "public/test-namespace-bundles"}
	out, execErr, _, _ := TestNamespaceCommands(getPolicies, policiesArgs)
	assert.Nil(t, execErr)

	var police pulsar.Policies
	err = json.Unmarshal(out.Bytes(), &police)
	assert.Nil(t, err)
	assert.Equal(t, 4, police.Bundles.NumBundles)
}

func TestCreateNsForNegativeBundles(t *testing.T) {
	args := []string{"create", "public/test-namespace-negative-bundles", "--bundles", "-1"}
	createOut, execErr, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	exceptedErr := "invalid number of bundles. Number of numBundles has to be in the range of (0, 2^32]"
	t.Log(createOut.String())
	assert.Equal(t, exceptedErr, execErr.Error())
}

func TestCreateNsForPositiveBundles(t *testing.T) {
	args := []string{"create", "public/test-namespace-positive-bundles", "--bundles", "12"}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	policiesArgs := []string{"policies", "public/test-namespace-positive-bundles"}
	out, execErr, _, _ := TestNamespaceCommands(getPolicies, policiesArgs)
	assert.Nil(t, execErr)

	var police pulsar.Policies
	err := json.Unmarshal(out.Bytes(), &police)
	assert.Nil(t, err)
	assert.Equal(t, 12, police.Bundles.NumBundles)
}

func TestCreateNsAlreadyExistError(t *testing.T) {
	args := []string{"create", "public/test-ns-duplicate"}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"create", "public/test-ns-duplicate"}
	_, execErr, _, _ = TestNamespaceCommands(createNs, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Namespace already exists", execErr.Error())
}

func TestCreateNsForCluster(t *testing.T) {
	clusterArgs := []string{"create", "test-cluster", "--url", "192.168.12.11"}
	_, _, _, err := cluster.TestClusterCommands(cluster.CreateClusterCmd, clusterArgs)
	assert.Nil(t, err)

	updateTenantArgs := []string{"update", "--allowed-clusters", "test-cluster",
		"--allowed-clusters", "standalone", "public"}
	_, execErr, _, err := tenant.TestTenantCommands(tenant.UpdateTenantCmd, updateTenantArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	nsArgs := []string{"create", "public/test-namespace-cluster", "--clusters", "test-cluster"}
	nsOut, execErr, _, _ := TestNamespaceCommands(createNs, nsArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, "Created public/test-namespace-cluster successfully\n", nsOut.String())

	policiesArgs := []string{"policies", "public/test-namespace-cluster"}
	out, execErr, _, _ := TestNamespaceCommands(getPolicies, policiesArgs)
	assert.Nil(t, execErr)

	var police pulsar.Policies
	err = json.Unmarshal(out.Bytes(), &police)
	assert.Nil(t, err)
	assert.Equal(t, "test-cluster", police.ReplicationClusters[0])

	// reset namespace clusters for other test case
	updateTenantArgs = []string{"update", "--allowed-clusters", "standalone", "public"}
	_, execErr, _, err = tenant.TestTenantCommands(tenant.UpdateTenantCmd, updateTenantArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

}

func TestCreateNsArgsError(t *testing.T) {
	args := []string{"create"}
	_, _, nameErr, _ := TestNamespaceCommands(createNs, args)

	assert.Equal(t, "the namespace name is not specified or the namespace name is "+
		"specified more than one", nameErr.Error())
}
