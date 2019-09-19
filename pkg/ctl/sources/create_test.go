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
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestCreateSources(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	// $ ./pulsarctl source create
	// --archive ./pulsar-io-kafka-2.4.0.nar
	// --classname org.apache.pulsar.io.kafka.KafkaBytesSource
	// --tenant public
	// --namespace default
	// --name kafka
	// --destination-topic-name my-topic
	// --source-config-file ./conf/kafkaSourceConfig.yaml
	// --parallelism 1
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-create",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	_, _, err = TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)
}

func TestFailureCreateSources(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	narName := "dummy-pulsar-io-kafka.nar"
	_, err = os.Create(narName)
	assert.Nil(t, err)

	defer os.Remove(narName)

	failArgs := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-create",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
	}

	exceptedErr := "Source test-source-create already exists"
	out, execErr, _ := TestSourcesCommands(createSourcesCmd, failArgs)
	assert.True(t, strings.Contains(out.String(), exceptedErr))
	assert.NotNil(t, execErr)

	narFailArgs := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-create-nar-fail",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", narName,
	}

	narErrInfo := "Source class org.apache.pulsar.io.kafka.KafkaBytesSource must be in class path"
	narOut, execErr, _ := TestSourcesCommands(createSourcesCmd, narFailArgs)
	assert.True(t, strings.Contains(narOut.String(), narErrInfo))
	assert.NotNil(t, execErr)
}
