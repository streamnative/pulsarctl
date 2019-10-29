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
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSource(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-update",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	createOut, _, err := TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-source-update successfully\n")

	updateArgs := []string{"update",
		"--name", "test-source-update",
		"--cpu", "5.0",
	}

	_, _, err = TestSourcesCommands(updateSourcesCmd, updateArgs)
	assert.Nil(t, err)
	getArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-update",
	}

	out, _, err := TestSourcesCommands(getSourcesCmd, getArgs)
	assert.Nil(t, err)

	var sourceConfig pulsar.SourceConfig
	err = json.Unmarshal(out.Bytes(), &sourceConfig)
	assert.Nil(t, err)

	assert.Equal(t, sourceConfig.Resources.CPU, 5.0)

	// test the source name not exist
	failureUpdateArgs := []string{"update",
		"--name", "not-exist",
	}
	_, err, _ = TestSourcesCommands(updateSourcesCmd, failureUpdateArgs)
	assert.NotNil(t, err)
	failMsg := "Source not-exist doesn't exist"
	assert.True(t, strings.Contains(err.Error(), failMsg))
}
