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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func TestLeaderBroker(t *testing.T) {
	oldURL := cmdutils.PulsarCtlConfig.WebServiceURL
	defer func() {
		cmdutils.PulsarCtlConfig.WebServiceURL = oldURL
	}()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/admin/v2/brokers/leaderBroker", r.URL.Path)
		_, _ = w.Write([]byte(`{"brokerId":"broker-1","serviceUrl":"http://127.0.0.1:8080"}`))
	}))
	defer srv.Close()

	cmdutils.PulsarCtlConfig.WebServiceURL = srv.URL

	args := []string{"leader-broker"}
	out, execErr, _, err := TestBrokersCommands(leaderBrokerCmd, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)

	var info utils.BrokerInfo
	err = json.Unmarshal(out.Bytes(), &info)
	assert.Nil(t, err)
	assert.Equal(t, "broker-1", info.BrokerID)
	assert.Equal(t, "http://127.0.0.1:8080", info.ServiceURL)
}
