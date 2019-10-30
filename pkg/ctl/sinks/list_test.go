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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSinks(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-list",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, _, err := TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-sink-list successfully\n")

	listArgs := []string{"list",
		"--tenant", "public",
		"--namespace", "default",
	}
	listOut, _, _ := TestSinksCommands(listSinksCmd, listArgs)
	t.Logf("pulsar sink name:%s", listOut.String())
	assert.True(t, strings.Contains(listOut.String(), "test-sink-list"))

	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-list",
	}

	deleteOut, _, _ := TestSinksCommands(deleteSinksCmd, deleteArgs)
	assert.Equal(t, deleteOut.String(), "Deleted test-sink-list successfully\n")

	listArgsAgain := []string{"list"}
	sinks, _, _ := TestSinksCommands(listSinksCmd, listArgsAgain)
	assert.False(t, strings.Contains(sinks.String(), "test-sink-list"))
}
