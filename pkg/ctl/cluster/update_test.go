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

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCluster(t *testing.T) {
	args := []string{
		"update",
		"--url", "http://example:8080",
		"--url-tls", "https://example:8080",
		"--broker-url", "pulsar://example:6650",
		"--broker-url-tls", "pulsar+ssl://example:6650",
		"-p", "cluster-a",
		"-p", "cluster-b",
		"standalone",
	}

	_, _, _, err := TestClusterCommands(UpdateClusterCmd, args)
	if err != nil {
		t.Error(err)
	}

	args = []string{"get", "standalone"}
	out, execErr, _, _ := TestClusterCommands(getClusterDataCmd, args)
	assert.Nil(t, execErr)

	var data utils.ClusterData
	err = json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "http://example:8080", data.ServiceURL)
	assert.Equal(t, "https://example:8080", data.ServiceURLTls)
	assert.Equal(t, "pulsar://example:6650", data.BrokerServiceURL)
	assert.Equal(t, "pulsar+ssl://example:6650", data.BrokerServiceURLTls)
}
