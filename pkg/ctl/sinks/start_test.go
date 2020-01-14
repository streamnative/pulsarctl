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

func TestStartAndStopSink(t *testing.T) {
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
		"--name", "test-sink-start",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, _, err := TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-sink-start successfully\n")

	stopArgs := []string{"--admin-service-url", requestURL, "stop",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-start",
	}

	_, _, err = TestSinksCommands(stopSinksCmd, stopArgs)
	assert.Nil(t, err)

	startArgs := []string{"--admin-service-url", requestURL, "start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-start",
	}
	_, _, err = TestSinksCommands(startSinksCmd, startArgs)
	assert.Nil(t, err)

	// test failure case
	failureStartArgs := []string{"--admin-service-url", requestURL, "start",
		"--name", "not-exist",
	}
	_, err, _ = TestSinksCommands(startSinksCmd, failureStartArgs)
	assert.NotNil(t, err)
	failMsg := "Sink not-exist doesn't exist"
	assert.True(t, strings.ContainsAny(err.Error(), failMsg))

	// test the --name args not exist
	notExistNameOrFqfnArgs := []string{"--admin-service-url", requestURL, "start",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, err, _ = TestSinksCommands(startSinksCmd, notExistNameOrFqfnArgs)
	assert.NotNil(t, err)
	failNameMsg := "You must specify a name for the sink"
	assert.True(t, strings.ContainsAny(err.Error(), failNameMsg))

	// test the instance id not exist
	notExistInstanceIDArgs := []string{"--admin-service-url", requestURL, "start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-start",
		"--instance-id", "12345678",
	}
	_, err, _ = TestSinksCommands(startSinksCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
