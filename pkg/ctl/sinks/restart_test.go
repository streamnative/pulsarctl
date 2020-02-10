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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestartSink(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-restart",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, _, err := TestSinksCommands(createSinksCmd, args)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, createOut.String(), "Created test-sink-restart successfully\n")

	restartArgs := []string{"restart",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-restart",
	}

	out, execErr, err := TestSinksCommands(restartSinksCmd, restartArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, execErr)
	assert.NotEmpty(t, out.String())

	notExistInstanceIDArgs := []string{"restart",
		"--name", "test-sink-restart",
		"--instance-id", "12345678",
	}
	_, execErr, err = TestSinksCommands(restartSinksCmd, notExistInstanceIDArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, execErr)
}

func TestRestartFailed(t *testing.T) {
	// test failure case
	failureArgs := []string{"restart",
		"--name", "not-exist",
	}
	_, execErr, err := TestSinksCommands(restartSinksCmd, failureArgs)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, execErr)
}
