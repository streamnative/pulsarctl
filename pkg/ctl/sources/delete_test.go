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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeleteSources(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-delete",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	_, _, err = TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)

	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-delete",
	}

	deleteOut, execErr, _ := TestSourcesCommands(deleteSourcesCmd, deleteArgs)
	delErr := "Deleted test-source-delete successfully\n"
	assert.True(t, strings.Contains(deleteOut.String(), delErr))
	assert.Nil(t, execErr)
}

func TestFailureDeleteSource(t *testing.T) {
	failureDeleteArgs := []string{"delete",
		"--name", "bad-delete-sources-" + time.Now().String(),
	}

	_, execErr, err := TestSourcesCommands(deleteSourcesCmd, failureDeleteArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, execErr)
}
