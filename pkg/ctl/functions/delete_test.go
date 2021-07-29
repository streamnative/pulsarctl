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
	"path"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestDeleteFunctions(t *testing.T) {
	fName := "df" + test.RandomSuffix()
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

	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
	}

	_, execErr, err = TestFunctionsCommands(deleteFunctionsCmd, deleteArgs)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
}

func TestDeleteFunctionsWithFQFN(t *testing.T) {
	fName := "df" + test.RandomSuffix()
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

	deleteArgsFqfn := []string{"delete",
		"--fqfn", "public/default/" + fName,
	}

	_, execErr, err = TestFunctionsCommands(deleteFunctionsCmd, deleteArgsFqfn)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
}

func TestDeleteFunctionsWithFailure(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-delete-failure",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)

	assert.Nil(t, err)

	failureDeleteArgs := []string{"delete",
		"--name", "not-exist",
	}

	_, execErrMsg, _ := TestFunctionsCommands(deleteFunctionsCmd, failureDeleteArgs)
	assert.NotNil(t, execErrMsg)
	exceptMsg := "Function not-exist doesn't exist"
	assert.True(t, strings.ContainsAny(execErrMsg.Error(), exceptMsg))

	notExistNameOrFqfnArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, execErrMsg, _ = TestFunctionsCommands(deleteFunctionsCmd, notExistNameOrFqfnArgs)
	failMsg := "you must specify a name for the function or a Fully Qualified Function Name (FQFN)"
	assert.NotNil(t, execErrMsg)
	assert.True(t, strings.ContainsAny(execErrMsg.Error(), failMsg))
}
