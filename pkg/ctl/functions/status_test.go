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
	"time"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestStatusFunctions(t *testing.T) {
	fName := "status-function" + test.RandomSuffix()
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
	assert.Equal(t, fmt.Sprintf("Created %s successfully\n", fName), out.String())

	getArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
	}
	outGet, execErr, err := TestFunctionsCommands(getFunctionsCmd, getArgs)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

	var functionConfig utils.FunctionConfig
	err = json.Unmarshal(outGet.Bytes(), &functionConfig)
	FailImmediatelyIfErrorNotNil(t, err)

	assert.Equal(t, functionConfig.Tenant, "public")
	assert.Equal(t, functionConfig.Namespace, "default")
	assert.Equal(t, functionConfig.Name, fName)

	statusArgs := []string{"status",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
	}

	var status utils.FunctionStatus

	task := func(args []string, obj interface{}) bool {
		outStatus, execErr, _ := TestFunctionsCommands(statusFunctionsCmd, args)
		if execErr != nil {
			return false
		}

		err = json.Unmarshal(outStatus.Bytes(), obj)
		if err != nil {
			return false
		}

		s := obj.(*utils.FunctionStatus)
		return len(s.Instances) == 1 && s.Instances[0].Status.Running
	}
	err = cmdutils.RunFuncWithTimeout(task, true, 30*time.Second, statusArgs, &status)
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

	_, execErr, err := TestFunctionsCommands(statusFunctionsCmd, statusArgs)
	if err != nil {
		t.Fatal(err)
	}

	errMsg := "Function test-functions-status-failure doesn't exist"
	assert.True(t, strings.Contains(execErr.Error(), errMsg))
}
