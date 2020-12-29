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
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/stretchr/testify/assert"
)

func TestStatusSink(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-status",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, _, err := TestSinksCommands(createSinksCmd, args)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, createOut.String(), "Created test-sink-status successfully\n")

	statusArgs := []string{"status",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-status",
	}

	var status utils.SinkStatus

	task := func(args []string, obj interface{}) bool {
		outStatus, execErr, _ := TestSinksCommands(statusSinksCmd, args)
		if execErr != nil {
			return false
		}

		err = json.Unmarshal(outStatus.Bytes(), &obj)
		if err != nil {
			return false
		}

		s := obj.(*utils.SinkStatus)
		return len(s.Instances) == 1 && s.Instances[0].Status.Running
	}

	err = cmdutils.RunFuncWithTimeout(task, true, 3*time.Minute, statusArgs, &status)
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

	_, execErr, err := TestSinksCommands(statusSinksCmd, statusArgs)
	if err != nil {
		t.Fatal(err)
	}

	errMsg := "Sink not-exist doesn't exist"
	assert.True(t, strings.Contains(execErr.Error(), errMsg))
}
