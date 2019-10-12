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
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestartFunctions(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-restart",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	restartArgs := []string{"restart",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-restart",
	}

	_, _, err = TestFunctionsCommands(restartFunctionsCmd, restartArgs)
	assert.Nil(t, err)

	argsFqfn := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-restart-fqfn",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, argsFqfn)
	assert.Nil(t, err)

	restartArgsFqfn := []string{"restart",
		"--fqfn", "public/default/test-functions-restart-fqfn",
	}

	_, _, err = TestFunctionsCommands(restartFunctionsCmd, restartArgsFqfn)
	assert.Nil(t, err)
}

func TestRestartFunctionsFailure(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-restart-failure",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

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
		"--name", "test-functions-restart-failure",
		"--instance-id", "12345678",
	}
	_, err, _ = TestFunctionsCommands(restartFunctionsCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
