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
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestStatusFunctions(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}
	t.Logf("base path: %s", basePath)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-status",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", basePath + "/test/functions/api-examples.jar",
	}

	out, _, err := TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)
	assert.Equal(t, "Created test-functions-status successfully\n", out.String())

	getArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-status",
	}

	outGet, _, _ := TestFunctionsCommands(getFunctionsCmd, getArgs)
	assert.Nil(t, err)

	var functionConfig pulsar.FunctionConfig
	err = json.Unmarshal(outGet.Bytes(), &functionConfig)
	assert.Nil(t, err)

	assert.Equal(t, functionConfig.Tenant, "public")
	assert.Equal(t, functionConfig.Namespace, "default")
	assert.Equal(t, functionConfig.Name, "test-functions-status")

	statusArgs := []string{"status",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-status",
	}

	var status pulsar.FunctionStatus

	task := func(args []string, obj interface{}) bool {
		outStatus, execErr, _ := TestFunctionsCommands(statusFunctionsCmd, args)
		if execErr != nil {
			return false
		}

		err = json.Unmarshal(outStatus.Bytes(), obj)
		if err != nil {
			return false
		}

		s := obj.(*pulsar.FunctionStatus)
		return len(s.Instances) == 1 && s.Instances[0].Status.Running
	}

	err = cmdutils.RunFuncWithTimeout(task, true, 1*time.Minute, statusArgs, &status)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, status.NumRunning)
	assert.Equal(t, 1, status.NumInstances)
}

func TestFailureStatus(t *testing.T) {
	statusArgs := []string{"status",
		"--name", "test-functions-status-failure",
	}

	out, _, err := TestFunctionsCommands(statusFunctionsCmd, statusArgs)
	assert.Nil(t, err)

	errMsg := "Function test-functions-status-failure doesn't exist"
	assert.True(t, strings.Contains(out.String(), errMsg))
}
