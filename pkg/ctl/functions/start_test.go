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

func TestStartFunctions(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-start",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	stopArgs := []string{"stop",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-start",
	}

	_, _, err = TestFunctionsCommands(stopFunctionsCmd, stopArgs)
	assert.Nil(t, err)

	startArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-start",
	}

	_, _, err = TestFunctionsCommands(startFunctionsCmd, startArgs)
	assert.Nil(t, err)

	argsFqfn := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-start-fqfn",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, argsFqfn)
	assert.Nil(t, err)

	stopArgsFqfn := []string{"stop",
		"--fqfn", "public/default/test-functions-start-fqfn",
	}

	_, _, err = TestFunctionsCommands(stopFunctionsCmd, stopArgsFqfn)
	assert.Nil(t, err)

	startArgsFqfn := []string{"start",
		"--fqfn", "public/default/test-functions-start-fqfn",
	}

	_, _, err = TestFunctionsCommands(startFunctionsCmd, startArgsFqfn)
	assert.Nil(t, err)

	// test failure cases

	stopArgsFqfnAgain := []string{"stop",
		"--fqfn", "public/default/test-functions-start-fqfn",
	}
	_, _, err = TestFunctionsCommands(stopFunctionsCmd, stopArgsFqfnAgain)
	assert.Nil(t, err)

	// test the function name not exist
	failureStartArgs := []string{"start",
		"--name", "not-exist",
	}
	_, err, _ = TestFunctionsCommands(startFunctionsCmd, failureStartArgs)
	assert.NotNil(t, err)
	failMsg := "Function not-exist doesn't exist"
	assert.True(t, strings.ContainsAny(err.Error(), failMsg))

	// test the --name args not exist
	notExistNameOrFqfnArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, err, _ = TestFunctionsCommands(startFunctionsCmd, notExistNameOrFqfnArgs)
	assert.NotNil(t, err)
	failNameMsg := "you must specify a name for the function or a Fully Qualified Function Name (FQFN)"
	assert.True(t, strings.ContainsAny(err.Error(), failNameMsg))

	// test the instance id not exist
	notExistInstanceIDArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-start-fqfn",
		"--instance-id", "12345678",
	}
	_, err, _ = TestFunctionsCommands(startFunctionsCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
