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

package sinks

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/streamnative/pulsarctl/pkg/test/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestGetSink(t *testing.T) {
	ctx := context.Background()
	c := pulsar.DefaultStandalone()
	c.WaitForLog("Function worker service started")
	c.Start(ctx)
	defer c.Stop(ctx)

	requestURL, err := c.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}

	args := []string{"--admin-service-url", requestURL, "create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-get",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	_, _, err = TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)

	getArgs := []string{"--admin-service-url", requestURL, "get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-get",
	}

	out, execErr, _ := TestSinksCommands(getSinksCmd, getArgs)
	assert.Nil(t, execErr)

	var sinkConf utils.SinkConfig
	err = json.Unmarshal(out.Bytes(), &sinkConf)
	assert.Nil(t, err)

	assert.Equal(t, sinkConf.Tenant, "public")
	assert.Equal(t, sinkConf.Namespace, "default")
	assert.Equal(t, sinkConf.Name, "test-sink-get")

	// check configs
	sinkConfMap := map[string]interface{}{
		"userName":  "root",
		"password":  "jdbc",
		"jdbcUrl":   "jdbc:mysql://127.0.0.1:3306/test_jdbc",
		"tableName": "test_jdbc",
	}
	assert.Equal(t, sinkConf.Configs, sinkConfMap)
}

func TestGetFailureSink(t *testing.T) {
	ctx := context.Background()
	c := pulsar.DefaultStandalone()
	c.WaitForLog("Function worker service started")
	c.Start(ctx)
	defer c.Stop(ctx)

	requestURL, err := c.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	failureGetArgs := []string{"--admin-service-url", requestURL, "get",
		"--name", "test-sink-get",
	}
	_, execErr, err := TestSinksCommands(getSinksCmd, failureGetArgs)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, execErr)
	exceptedErr := "code: 404 reason: Sink test-sink-get doesn't exist"
	assert.Equal(t, exceptedErr, execErr.Error())
}
