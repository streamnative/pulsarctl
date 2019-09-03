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

func TestStopFunctions(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-stop",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	stopArgs := []string{"stop",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-stop",
	}

	_, err = TestFunctionsCommands(stopFunctionsCmd, stopArgs)
	assert.Nil(t, err)

    argsFqfn := []string{"create",
        "--tenant", "public",
        "--namespace", "default",
        "--name", "test-functions-stop-fqfn",
        "--inputs", "test-input-topic",
        "--output", "persistent://public/default/test-output-topic",
        "--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
        "--jar", jarName,
    }

    _, err = TestFunctionsCommands(createFunctionsCmd, argsFqfn)
    assert.Nil(t, err)

    stopArgsFqfn := []string{"stop",
        "--fqfn", "public/default/test-functions-stop-fqfn",
    }

    _, err = TestFunctionsCommands(stopFunctionsCmd, stopArgsFqfn)
    assert.Nil(t, err)
}
