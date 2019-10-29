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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestartSource(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-restart",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	createOut, _, err := TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-source-restart successfully\n")

	restartArgs := []string{"restart",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-restart",
	}

	_, _, err = TestSourcesCommands(restartSourcesCmd, restartArgs)
	assert.Nil(t, err)

	// test failure case
	failureArgs := []string{"restart",
		"--name", "not-exist",
	}
	_, execErr, _ := TestSourcesCommands(restartSourcesCmd, failureArgs)
	assert.Equal(t, execErr.Error(), "code: 404 reason: Source not-exist doesn't exist")

	notExistInstanceIDArgs := []string{"restart",
		"--name", "test-source-restart",
		"--instance-id", "12345678",
	}
	_, err, _ = TestSourcesCommands(restartSourcesCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
