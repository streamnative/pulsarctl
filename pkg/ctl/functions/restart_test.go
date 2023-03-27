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
	"path"
	"strings"
	"testing"
	"time"

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestRestartFunctions(t *testing.T) {
	fName := "restart-f" + test.RandomSuffix()
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
	_, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

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
	err = cmdutils.RunFuncWithTimeout(task, true, 1*time.Minute, statusArgs, &status)
	if err != nil {
		t.Fatal(err)
	}

	restartArgs := []string{"restart",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
	}
	_, execErr, err = TestFunctionsCommands(restartFunctionsCmd, restartArgs)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
}

func TestRestartFunctionWithFQFN(t *testing.T) {
	fName := "restart-f" + test.RandomSuffix()
	jarName := path.Join(ResourceDir(), "api-examples.jar")
	argsFqfn := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, execErr, err := TestFunctionsCommands(createFunctionsCmd, argsFqfn)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

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
	err = cmdutils.RunFuncWithTimeout(task, true, 1*time.Minute, statusArgs, &status)
	if err != nil {
		t.Fatal(err)
	}

	restartArgsFqfn := []string{"restart",
		"--fqfn", "public/default/" + fName,
	}
	_, execErr, err = TestFunctionsCommands(restartFunctionsCmd, restartArgsFqfn)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
}

func TestRestartFunctionsFailure(t *testing.T) {
	fName := "restart-fail-function" + test.RandomSuffix()
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

	_, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)

	// test the function name not exist
	failureDeleteArgs := []string{"restart",
		"--name", "not-exist",
	}
	_, err, _ = TestFunctionsCommands(restartFunctionsCmd, failureDeleteArgs)
	assert.NotNil(t, err)
	failMsg := "Function not-exist doesn't exist"
	assert.True(t, strings.ContainsAny(err.Error(), failMsg))

	// test the --name args not exist
	notExistNameOrFqfnArgs := []string{"restart",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, err, _ = TestFunctionsCommands(restartFunctionsCmd, notExistNameOrFqfnArgs)
	assert.NotNil(t, err)
	failNameMsg := "you must specify a name for the function or a Fully Qualified Function Name (FQFN)"
	assert.True(t, strings.ContainsAny(err.Error(), failNameMsg))

	// test the instance id not exist
	notExistInstanceIDArgs := []string{"restart",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--instance-id", "12345678",
	}
	_, err, _ = TestFunctionsCommands(restartFunctionsCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
