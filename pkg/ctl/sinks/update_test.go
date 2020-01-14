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
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/streamnative/pulsarctl/pkg/test/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSink(t *testing.T) {
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
		"--name", "test-sinks-update",
		"--inputs", "test-topic",
		"--archive", basePath + "/test/sinks/pulsar-io-jdbc-2.4.0.nar",
		"--sink-config-file", basePath + "/test/sinks/mysql-jdbc-sink.yaml",
	}

	createOut, _, err := TestSinksCommands(createSinksCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created test-sinks-update successfully\n")

	updateArgs := []string{"--admin-service-url", requestURL, "update",
		"--name", "test-sinks-update",
		"--parallelism", "3",
	}

	updateOut, _, err := TestSinksCommands(updateSinksCmd, updateArgs)
	fmt.Println(updateOut.String())
	assert.Nil(t, err)
	getArgs := []string{"--admin-service-url", requestURL, "get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-sinks-update",
	}

	out, _, err := TestSinksCommands(getSinksCmd, getArgs)
	assert.Nil(t, err)

	var sinkConf utils.SinkConfig
	err = json.Unmarshal(out.Bytes(), &sinkConf)
	assert.Nil(t, err)
	assert.Equal(t, sinkConf.Parallelism, 3)

	// test the sink name not exist
	failureUpdateArgs := []string{"--admin-service-url", requestURL, "update",
		"--name", "not-exist",
	}
	_, err, _ = TestSinksCommands(updateSinksCmd, failureUpdateArgs)
	assert.NotNil(t, err)
	failMsg := "Sink not-exist doesn't exist"
	assert.True(t, strings.Contains(err.Error(), failMsg))
}
