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

package brokerstats

import (
    "encoding/json"
    "testing"

    "github.com/streamnative/pulsarctl/pkg/pulsar"
    "github.com/stretchr/testify/assert"
)

func TestDumpMBeans(t *testing.T) {
	args := []string{"mbeans"}
	mbeansOut, execErr, _, _ := TestBrokerStatsCommands(dumpMBeans, args)
	assert.Nil(t, execErr)

	var out []pulsar.Metrics
	err := json.Unmarshal(mbeansOut.Bytes(), &out)
	assert.Nil(t, err)

	tmpMap := map[string]string{
		"MBean": "java.lang:type=MemoryPool,name=Metaspace",
	}

	assert.Equal(t, tmpMap, out[0].Dimensions)
}
