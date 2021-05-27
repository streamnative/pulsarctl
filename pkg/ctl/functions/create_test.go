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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFunctions(t *testing.T) {
	basePath, err := getDirHelp()
	if basePath == "" || err != nil {
		t.Error(err)
	}

	jarName := "dummyExample.jar"
	_, err = os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)

	goName := "dummyExample.go"
	_, err = os.Create(goName)
	assert.Nil(t, err)

	defer os.Remove(goName)

	pyName := "dummyExample.py"
	_, err = os.Create(pyName)
	assert.Nil(t, err)

	defer os.Remove(pyName)

	// $ ./pulsarctl functions create
	// --tenant public
	// --namespace default
	// --name test-functions-create
	// --inputs test-input-topic
	// --output persistent://public/default/test-output-topic
	// --classname org.apache.pulsar.functions.api.examples.ExclamationFunction
	// --jar apache-pulsar-2.4.0/examples/api-examples.jar
	// --processing-guarantees EFFECTIVELY_ONCE
	args := []string{
		"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-create",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
		"--processing-guarantees", "EFFECTIVELY_ONCE",
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	// $ bin/pulsar-admin functions create
	// --function-config-file examples/example-function-config.yaml
	// --jar examples/api-examples.jar
	argsWithConf := []string{
		"create",
		"--function-config-file", basePath + "/test/functions/example-function-config.yaml",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, argsWithConf)
	assert.Nil(t, err)

	argsWithFileURL := []string{
		"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-create-file",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", "file:" + "/pulsar_test/" + jarName,
		"--processing-guarantees", "EFFECTIVELY_ONCE",
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, argsWithFileURL)
	assert.Nil(t, err)

	argsWithFileURLGo := []string{
		"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-create-file-go",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--go", "file:" + "/pulsar_test/" + goName,
		"--processing-guarantees", "EFFECTIVELY_ONCE",
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, argsWithFileURLGo)
	assert.Nil(t, err)

	argsWithFileURLPy := []string{
		"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-create-file-py",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--py", "file:" + "/pulsar_test/" + pyName,
		"--processing-guarantees", "EFFECTIVELY_ONCE",
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, argsWithFileURLPy)
	assert.Nil(t, err)
}
