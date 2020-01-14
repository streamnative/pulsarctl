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

func TestDeleteSinks(t *testing.T) {
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
		"--name", "test-sink-delete",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	_, _, err = TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)

	deleteArgs := []string{"--admin-service-url", requestURL, "delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sink-delete",
	}

	deleteOut, execErr, _ := TestSinksCommands(deleteSinksCmd, deleteArgs)
	delErr := "Deleted test-sink-delete successfully\n"
	assert.True(t, strings.Contains(deleteOut.String(), delErr))
	assert.Nil(t, execErr)
}

func TestFailureDeleteSink(t *testing.T) {
	ctx := context.Background()
	c := pulsar.DefaultStandalone()
	c.WaitForLog("Function worker service started")
	c.Start(ctx)
	defer c.Stop(ctx)

	requestURL, err := c.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	failureDeleteArgs := []string{"--admin-service-url", requestURL, "delete",
		"--name", "test-sink-delete",
	}

	exceptedErr := "Sink test-sink-delete doesn't exist"
	_, execErrMsg, _ := TestSinksCommands(deleteSinksCmd, failureDeleteArgs)
	assert.True(t, strings.Contains(execErrMsg.Error(), exceptedErr))
	assert.NotNil(t, execErrMsg)

	nameNotExist := []string{"--admin-service-url", requestURL, "delete",
		"--name", "not-exist",
	}
	_, execErrMsg, _ = TestSinksCommands(deleteSinksCmd, nameNotExist)
	nameErr := "Sink not-exist doesn't exist"
	assert.True(t, strings.Contains(execErrMsg.Error(), nameErr))
}
