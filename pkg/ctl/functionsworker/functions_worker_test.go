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

package functionsworker

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestFunctionsWorker(t *testing.T) {
	metricsArgs := []string{"monitoring-metrics"}
	metricsOut, execErr, _, _ := TestFunctionsWorkerCmd(monitoringMetrics, metricsArgs)
	assert.Nil(t, execErr)

	var metrics []utils.Metrics
	err := json.Unmarshal(metricsOut.Bytes(), &metrics)
	assert.Nil(t, err)

	assert.Equal(t, "jvm_metrics", metrics[0].Dimensions["metric"])

	clustersArgs := []string{"get-cluster"}
	clusterOut, execErr, _, _ := TestFunctionsWorkerCmd(getCluster, clustersArgs)
	assert.Nil(t, execErr)

	var cluster []utils.WorkerInfo
	err = json.Unmarshal(clusterOut.Bytes(), &cluster)
	assert.Nil(t, err)

	assert.Equal(t, 8080, cluster[0].Port)

	clusterLeaderArgs := []string{"get-cluster-leader"}
	clusterLeaderOut, execErr, _, _ := TestFunctionsWorkerCmd(getClusterLeader, clusterLeaderArgs)
	assert.Nil(t, execErr)

	var clusterLeader utils.WorkerInfo
	err = json.Unmarshal(clusterLeaderOut.Bytes(), &clusterLeader)
	assert.Nil(t, err)

	assert.Equal(t, 8080, clusterLeader.Port)
}
