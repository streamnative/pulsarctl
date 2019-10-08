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
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestStatusSink(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-status",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, _, err := TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-sink-status successfully")

	statusArgs := []string{"status",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-status",
	}

	var outStatus *bytes.Buffer
	var status pulsar.SinkStatus

	for {
		outStatus, _, _ = TestSinksCommands(statusSinksCmd, statusArgs)
		if strings.Contains(outStatus.String(), "true") {
			break
		}
	}

	t.Log(outStatus.String())
	err = json.Unmarshal(outStatus.Bytes(), &status)
	assert.Nil(t, err)

	assert.Equal(t, 1, status.NumRunning)
	assert.Equal(t, 1, status.NumInstances)
}

func TestFailureStatus(t *testing.T) {
	statusArgs := []string{"status",
		"--name", "not-exist",
	}

	out, _, err := TestSinksCommands(statusSinksCmd, statusArgs)
	assert.Nil(t, err)

	errMsg := "Sink not-exist doesn't exist"
	t.Logf(out.String())
	assert.True(t, strings.Contains(out.String(), errMsg))
}
