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

func TestListSources(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-list",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	createOut, _, err := TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-source-list successfully\n")

	listArgs := []string{"list",
		"--tenant", "public",
		"--namespace", "default",
	}
	listOut, _, _ := TestSourcesCommands(listSourcesCmd, listArgs)
	t.Logf("pulsar source name:%s", listOut.String())
	assert.True(t, strings.Contains(listOut.String(), "test-source-list"))

	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-list",
	}

	deleteOut, _, _ := TestSourcesCommands(deleteSourcesCmd, deleteArgs)
	assert.Equal(t, deleteOut.String(), "Deleted test-source-list successfully\n")

	listArgsAgain := []string{"list"}
	sources, _, _ := TestSourcesCommands(listSourcesCmd, listArgsAgain)
	assert.False(t, strings.Contains(sources.String(), "test-source-list"))
}
