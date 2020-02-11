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
	"encoding/json"
	"testing"
	"time"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetSink(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-get",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	_, _, err = TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)

	getArgs := []string{"get",
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
	t.Logf("get sink value:%s", out.String())
}

func TestGetFailureSink(t *testing.T) {
	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "bad-delete-sinks" + time.Now().String(),
	}

	_, execErr, err := TestSinksCommands(deleteSinksCmd, deleteArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, execErr)

	failureGetArgs := []string{"get",
		"--name", "bad-get-sinks" + time.Now().String(),
	}
	_, execErr, err = TestSinksCommands(getSinksCmd, failureGetArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, execErr)
}
