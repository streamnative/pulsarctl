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
	"regexp"
	"testing"
)


func TestGetClusterData(t *testing.T) {
	args := []string{"get", "standalone"}
	out, _, _, err := TestClusterCommands(getClusterDataCmd, args)
	if err != nil {
		t.Error(err)
	}
	c := pulsar.ClusterData{}
	err = json.Unmarshal(out.Bytes(), &c)
	if err != nil {
		t.Error(err)
	}

	pulsarUrl, err := regexp.Compile("^pulsar://[a-z-A-Z0-9]*:6650$")
	if err != nil {
		t.Error(err)
	}

	res := pulsarUrl.MatchString(c.BrokerServiceURL)
	assert.True(t, res)

	httpUrl, err := regexp.Compile("^http://[a-z-A-Z0-9]*:8080$")
	if err != nil {
		t.Error(err)
	}

	res = httpUrl.MatchString(c.ServiceURL)
	assert.True(t, res)
}

