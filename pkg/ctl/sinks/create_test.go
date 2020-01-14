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
	"os"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/streamnative/pulsarctl/pkg/test/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestCreateSinks(t *testing.T) {
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
		"--name", "test-sink-create",
		"--inputs", "persistent://public/default/my-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
		"--parallelism", "1",
	}
	out, _, err := TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, "Created test-sink-create successfully\n", out.String())

	// create sink again
	_, execErr, err := TestSinksCommands(createSinksCmd, args)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 400 reason: Sink test-sink-create already exists", execErr.Error())
}

func TestFailureCreateSinks(t *testing.T) {
	ctx := context.Background()
	c := pulsar.DefaultStandalone()
	c.WaitForLog("Function worker service started")
	c.Start(ctx)
	defer c.Stop(ctx)

	requestURL, err := c.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	out, err := test.ExecCmd(c.GetContainerID(), []string{"bin/pulsar-admin", "namespaces", "list", "public"})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, strings.Contains(out, "public/default"))

	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}

	narName := "dummy-pulsar-io-mysql.nar"
	_, err = os.Create(narName)
	defer os.Remove(narName)
	if err != nil {
		t.Fatal(err)
	}

	narFailArgs := []string{"--admin-service-url", requestURL, "create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-create-nar-fail",
		"--inputs", "my-topic",
		"--archive", narName,
	}

	_, execErr, err := TestSinksCommands(createSinksCmd, narFailArgs)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 400 reason: Sink package does not have the correct format. "+
		"Pulsar cannot determine if the package is a NAR package or JAR package.Sink classname "+
		"is not provided and attempts to load it as a NAR package produced error: zip file is empty",
		execErr.Error())
}
