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

package sources

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestGetSource(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-get",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	_, _, err = TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)

	getArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-get",
	}

	out, _, _ := TestSourcesCommands(getSourcesCmd, getArgs)

	var sourceConf pulsar.SourceConfig
	err = json.Unmarshal(out.Bytes(), &sourceConf)
	assert.Nil(t, err)

	assert.Equal(t, sourceConf.Tenant, "public")
	assert.Equal(t, sourceConf.Namespace, "default")
	assert.Equal(t, sourceConf.Name, "test-source-get")
	assert.Equal(t, sourceConf.ClassName, "org.apache.pulsar.io.kafka.KafkaBytesSource")

	// check configs
	sourceConfMap := map[string]interface{}{
		"autoCommitEnabled": "false",
		"bootstrapServers":  "pulsar-kafka:9092",
		"groupId":           "test-pulsar-io",
		"sessionTimeoutMs":  "10000",
		"topic":             "my-topic",
	}
	assert.Equal(t, sourceConf.Configs, sourceConfMap)
	t.Logf("get source value:%s", out.String())
}

func TestGetFailureSource(t *testing.T) {
	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-get",
	}

	deleteOut, _, _ := TestSourcesCommands(deleteSourcesCmd, deleteArgs)
	assert.Equal(t, deleteOut.String(), "Deleted test-source-get successfully\n")

	failureGetArgs := []string{"get",
		"--name", "test-source-get",
	}
	getOut, execErr, _ := TestSourcesCommands(getSourcesCmd, failureGetArgs)
	assert.NotNil(t, execErr)
	exceptedErr := "error: Source test-source-get doesn't exist\n"
	assert.Equal(t, getOut.String(), exceptedErr)
}
