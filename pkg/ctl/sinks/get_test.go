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
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
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

	var sinkConf pulsar.SinkConfig
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
		"--name", "test-sink-get",
	}

	deleteOut, _, _ := TestSinksCommands(deleteSinksCmd, deleteArgs)
	assert.Equal(t, deleteOut.String(), "Deleted test-sink-get successfully\n")

	failureGetArgs := []string{"get",
		"--name", "test-sink-get",
	}
	getOut, execErr, _ := TestSinksCommands(getSinksCmd, failureGetArgs)
	assert.NotNil(t, execErr)
	exceptedErr := "Sink test-sink-get doesn't exist"
	assert.True(t, strings.Contains(getOut.String(), exceptedErr))
}
