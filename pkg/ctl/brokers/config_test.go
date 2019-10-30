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

package brokers

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetInternalConfig(t *testing.T) {
	args := []string{"get-internal-config"}
	internalOut, execErr, _, _ := TestBrokersCommands(getInternalConfigCmd, args)
	assert.Nil(t, execErr)

	var internalData utils.InternalConfigurationData
	err := json.Unmarshal(internalOut.Bytes(), &internalData)
	assert.Nil(t, err)

	assert.Equal(t, "127.0.0.1:2181", internalData.ZookeeperServers)
	assert.Equal(t, "127.0.0.1:2181", internalData.ConfigurationStoreServers)
	assert.Equal(t, "/ledgers", internalData.LedgersRootPath)
	assert.Equal(t, "bk://127.0.0.1:4181", internalData.StateStorageServiceURL)
}

func TestGetRuntimeConfig(t *testing.T) {
	args := []string{"get-runtime-config"}
	runtimeOut, execErr, _, _ := TestBrokersCommands(getRuntimeConfigCmd, args)
	assert.Nil(t, execErr)

	var runtimeConf map[string]string
	err := json.Unmarshal(runtimeOut.Bytes(), &runtimeConf)
	assert.Nil(t, err)

	assert.Equal(t, "false", runtimeConf["authenticateOriginalAuthData"])
	assert.Equal(t, "true", runtimeConf["backlogQuotaCheckEnabled"])
	assert.Equal(t, "0.0.0.0", runtimeConf["bindAddress"])
	assert.Equal(t, "CRC32", runtimeConf["managedLedgerDigestType"])
	assert.Equal(t, "127.0.0.1:2181", runtimeConf["zookeeperServers"])
	assert.Equal(t, "30000", runtimeConf["zooKeeperSessionTimeoutMillis"])
	assert.Equal(t, "300000", runtimeConf["webSocketSessionIdleTimeoutMillis"])
	assert.Equal(t, "30", runtimeConf["keepAliveIntervalSeconds"])
}
