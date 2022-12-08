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

package resourcequotas

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestResourceQuota(t *testing.T) {
	t.Skip("https://github.com/apache/pulsar/pull/18755")

	getDefaultArgs := []string{"get"}
	getDefaultOut, execErr, _, _ := TestResourceQuotaCommands(getResourceQuota, getDefaultArgs)
	assert.Nil(t, execErr)
	var quota utils.ResourceQuota
	err := json.Unmarshal(getDefaultOut.Bytes(), &quota)
	assert.Nil(t, err)

	assert.Equal(t, float64(40), quota.MsgRateIn)
	assert.Equal(t, float64(120), quota.MsgRateOut)
	assert.Equal(t, float64(100000), quota.BandwidthIn)
	assert.Equal(t, float64(300000), quota.BandwidthOut)
	assert.Equal(t, float64(80), quota.Memory)
	assert.Equal(t, true, quota.Dynamic)

	setDefaultArgs := []string{"set", "--bandwidthIn", "10", "--bandwidthOut", "20",
		"--memory", "30", "--msgRateIn", "40", "--msgRateOut", "50"}
	setDefaultOut, execErr, _, _ := TestResourceQuotaCommands(setResourceQuota, setDefaultArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, "Set default resource quota successful\n", setDefaultOut.String())

	getDefaultAgainArgs := []string{"get"}
	getDefaultOut, execErr, _, _ = TestResourceQuotaCommands(getResourceQuota, getDefaultAgainArgs)
	assert.Nil(t, execErr)
	err = json.Unmarshal(getDefaultOut.Bytes(), &quota)
	assert.Nil(t, err)

	assert.Equal(t, float64(10), quota.BandwidthIn)
	assert.Equal(t, float64(20), quota.BandwidthOut)
	assert.Equal(t, float64(30), quota.Memory)
	assert.Equal(t, float64(40), quota.MsgRateIn)
	assert.Equal(t, float64(50), quota.MsgRateOut)

	setArgs := []string{"set", "--bandwidthIn", "100", "--bandwidthOut", "200", "--memory", "300",
		"--msgRateIn", "400", "--msgRateOut", "500", "--namespace", "public/default",
		"--bundle", "0x80000000_0xc0000000"}
	setOut, execErr, _, _ := TestResourceQuotaCommands(setResourceQuota, setArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, "Set resource quota successful\n", setOut.String())

	getArgs := []string{"get", "public/default", "0x80000000_0xc0000000"}
	getOut, execErr, _, _ := TestResourceQuotaCommands(getResourceQuota, getArgs)
	assert.Nil(t, execErr)
	err = json.Unmarshal(getOut.Bytes(), &quota)
	assert.Nil(t, err)

	assert.Equal(t, float64(100), quota.BandwidthIn)
	assert.Equal(t, float64(200), quota.BandwidthOut)
	assert.Equal(t, float64(300), quota.Memory)
	assert.Equal(t, float64(400), quota.MsgRateIn)
	assert.Equal(t, float64(500), quota.MsgRateOut)

	resetArgs := []string{"reset",
		"public/default",
		"0x80000000_0xc0000000"}
	resetOut, execErr, _, _ := TestResourceQuotaCommands(resetNamespaceBundleResourceQuota, resetArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, "Reset resource quota successful\n", resetOut.String())
}
