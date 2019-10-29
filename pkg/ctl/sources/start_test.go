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

func TestStartAndStopSource(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-start",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	createOut, _, err := TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-source-start successfully\n")

	stopArgs := []string{"stop",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-start",
	}

	_, _, err = TestSourcesCommands(stopSourcesCmd, stopArgs)
	assert.Nil(t, err)

	startArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-start",
	}
	_, _, err = TestSourcesCommands(startSourcesCmd, startArgs)
	assert.Nil(t, err)

	// test failure case
	failureStartArgs := []string{"start",
		"--name", "not-exist",
	}
	_, err, _ = TestSourcesCommands(startSourcesCmd, failureStartArgs)
	assert.NotNil(t, err)
	failMsg := "Source not-exist doesn't exist"
	assert.True(t, strings.ContainsAny(err.Error(), failMsg))

	// test the --name args not exist
	notExistNameOrFqfnArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, err, _ = TestSourcesCommands(startSourcesCmd, notExistNameOrFqfnArgs)
	assert.NotNil(t, err)
	failNameMsg := "You must specify a name for the source"
	assert.True(t, strings.ContainsAny(err.Error(), failNameMsg))

	// test the instance id not exist
	notExistInstanceIDArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-start",
		"--instance-id", "12345678",
	}
	_, err, _ = TestSourcesCommands(startSourcesCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
