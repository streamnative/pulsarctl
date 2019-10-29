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

package functions

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestStatsFunctions(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-stats",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", basePath + "/test/functions/api-examples.jar",
	}

	out, _, err := TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, out.String(), "Created test-functions-stats successfully\n")

	statsArgs := []string{"stats",
		"--name", "test-functions-stats",
	}

	outStats, _, err := TestFunctionsCommands(statsFunctionsCmd, statsArgs)
	assert.Nil(t, err)

	var stats pulsar.FunctionStats
	err = json.Unmarshal(outStats.Bytes(), &stats)
	assert.Nil(t, err)

	assert.Equal(t, int64(0), stats.ReceivedTotal)
	assert.Equal(t, int64(0), stats.ProcessedSuccessfullyTotal)
}

func TestFailureStats(t *testing.T) {
	statsArgs := []string{"stats",
		"--name", "test-functions-stats-failure",
	}

	out, _, err := TestFunctionsCommands(statsFunctionsCmd, statsArgs)
	assert.Nil(t, err)

	errMsg := "Function test-functions-stats-failure doesn't exist"
	assert.True(t, strings.Contains(out.String(), errMsg))
}
