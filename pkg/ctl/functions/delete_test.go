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
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDeleteFunctions(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-delete",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	deleteArgs := []string{"delete",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-delete",
	}

	_, err = TestFunctionsCommands(deleteFunctionsCmd, deleteArgs)
	assert.Nil(t, err)

	argsFqfn := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-delete-fqfn",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, err = TestFunctionsCommands(createFunctionsCmd, argsFqfn)
	assert.Nil(t, err)

	deleteArgsFqfn := []string{"delete",
		"--fqfn", "public/default/test-functions-delete-fqfn",
	}

	_, err = TestFunctionsCommands(deleteFunctionsCmd, deleteArgsFqfn)
	assert.Nil(t, err)
}
