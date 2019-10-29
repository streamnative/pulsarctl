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
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSinks(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-create",
		"--inputs", "persistent://public/default/my-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
		"--parallelism", "1",
	}
	out, _, err := TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)
	fmt.Println(out.String())
	assert.Equal(t, out.String(), "Created test-sink-create successfully\n")
}

func TestFailureCreateSinks(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)

	narName := "dummy-pulsar-io-mysql.nar"
	_, err = os.Create(narName)
	assert.Nil(t, err)

	defer os.Remove(narName)

	failArgs := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-create",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	exceptedErr := "Sink test-sink-create already exists"
	out, execErr, _ := TestSinksCommands(createSinksCmd, failArgs)
	assert.True(t, strings.Contains(out.String(), exceptedErr))
	assert.NotNil(t, execErr)

	narFailArgs := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-create-nar-fail",
		"--inputs", "my-topic",
		"--archive", narName,
	}

	narErrInfo := "error: zip file is empty"
	narOut, execErr, _ := TestSinksCommands(createSinksCmd, narFailArgs)
	fmt.Println(narOut.String())
	assert.True(t, strings.Contains(narOut.String(), narErrInfo))
	assert.NotNil(t, execErr)
}
