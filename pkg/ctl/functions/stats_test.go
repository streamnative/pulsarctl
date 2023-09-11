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
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestStatsFunctions(t *testing.T) {
	fName := "status-f" + test.RandomSuffix()
	jarName := path.Join(ResourceDir(), "api-examples.jar")

	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	out, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Equal(t, out.String(), fmt.Sprintf("Created %s successfully\n", fName))

	statsArgs := []string{"stats",
		"--name", fName,
	}

	outStats, execErr, err := TestFunctionsCommands(statsFunctionsCmd, statsArgs)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

	var stats utils.FunctionStats
	err = json.Unmarshal(outStats.Bytes(), &stats)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

	assert.Equal(t, int64(0), stats.ReceivedTotal)
	assert.Equal(t, int64(0), stats.ProcessedSuccessfullyTotal)
}

func TestFailureStats(t *testing.T) {
	statsArgs := []string{"stats",
		"--name", "test-functions-stats-failure",
	}

	_, execErr, err := TestFunctionsCommands(statsFunctionsCmd, statsArgs)
	if err != nil {
		t.Fatal(err)
	}

	errMsg := "Function test-functions-stats-failure doesn't exist"
	assert.True(t, strings.Contains(execErr.Error(), errMsg))
}
