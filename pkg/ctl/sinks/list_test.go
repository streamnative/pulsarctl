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
	"context"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestListSinks(t *testing.T) {
	ctx := context.Background()
	c := pulsar.DefaultStandalone()
	c.WaitForLog("Function worker service started")
	c.Start(ctx)
	defer c.Stop(ctx)

	requestURL, err := c.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}

	args := []string{"--admin-service-url", requestURL, "create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-list",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, execErr, err := TestSinksCommands(createSinksCmd, args)
	if err != nil {
		t.Fatal(err)
	}
	if execErr != nil {
		t.Fatal(execErr)
	}
	assert.Equal(t, "Created test-sink-list successfully\n", createOut.String())

	listArgs := []string{"--admin-service-url", requestURL, "list",
		"--tenant", "public",
		"--namespace", "default",
	}
	listOut, execErr, err := TestSinksCommands(listSinksCmd, listArgs)
	if err != nil {
		t.Fatal(err)
	}
	if execErr != nil {
		t.Fatal(execErr)
	}
	assert.True(t, strings.Contains(listOut.String(), "test-sink-list"))

	deleteArgs := []string{"--admin-service-url", requestURL, "delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-list",
	}

	deleteOut, execErr, err := TestSinksCommands(deleteSinksCmd, deleteArgs)
	if err != nil {
		t.Fatal(err)
	}
	if execErr != nil {
		t.Fatal(execErr)
	}
	assert.Equal(t, deleteOut.String(), "Deleted test-sink-list successfully\n")

	listArgsAgain := []string{"--admin-service-url", requestURL, "list"}
	sinks, execErr, err := TestSinksCommands(listSinksCmd, listArgsAgain)
	if err != nil {
		t.Fatal(err)
	}
	if execErr != nil {
		t.Fatal(execErr)
	}
	assert.False(t, strings.Contains(sinks.String(), "test-sink-list"))
}
