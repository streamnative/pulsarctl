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
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestStatusSource(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-status",
		"--destination-topic-name", "my-topic",
		"--classname", "org.apache.pulsar.io.kafka.KafkaBytesSource",
		"--archive", basePath + "/test/sources/pulsar-io-kafka-2.4.0.nar",
		"--source-config-file", basePath + "/test/sources/kafkaSourceConfig.yaml",
	}

	createOut, _, err := TestSourcesCommands(createSourcesCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-source-status successfully\n")

	statusArgs := []string{"status",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-status",
	}

	var status pulsar.SourceStatus

	task := func(args []string, obj interface{}) bool {
		outStatus, execErr, _ := TestSourcesCommands(statusSourcesCmd, args)
		if execErr != nil {
			return false
		}
		err = json.Unmarshal(outStatus.Bytes(), &obj)
		if err != nil {
			return false
		}
		s := obj.(*pulsar.SourceStatus)
		return len(s.Instances) == 1 && s.Instances[0].Status.Running
	}

	err = cmdutils.RunFuncWithTimeout(task, true, 1*time.Minute, statusArgs, &status)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, status.NumRunning)
	assert.Equal(t, 1, status.NumInstances)
}

func TestFailureStatus(t *testing.T) {
	statusArgs := []string{"status",
		"--name", "not-exist",
	}

	out, _, err := TestSourcesCommands(statusSourcesCmd, statusArgs)
	assert.Nil(t, err)

	errMsg := "Source not-exist doesn't exist"
	t.Logf(out.String())
	assert.True(t, strings.Contains(out.String(), errMsg))
}
