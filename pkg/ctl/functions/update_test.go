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
	"os"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/stretchr/testify/assert"
)

func TestUpdateFunctions(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-update",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	updateArgs := []string{"update",
		"--name", "test-functions-update",
		"--output", "update-output-topic",
		"--cpu", "5.0",
	}

	_, _, err = TestFunctionsCommands(updateFunctionsCmd, updateArgs)
	assert.Nil(t, err)

	getArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-update",
	}

	out, _, err := TestFunctionsCommands(getFunctionsCmd, getArgs)
	assert.Nil(t, err)

	var functionConfig pulsar.FunctionConfig
	err = json.Unmarshal(out.Bytes(), &functionConfig)
	assert.Nil(t, err)

	assert.Equal(t, functionConfig.Tenant, "public")
	assert.Equal(t, functionConfig.Namespace, "default")
	assert.Equal(t, functionConfig.Name, "test-functions-update")
	assert.Equal(t, functionConfig.Output, "update-output-topic")
	assert.Equal(t, functionConfig.Resources.CPU, 5.0)
}

func TestUpdateFunctionsFailure(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-update-failure",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	// test the function name not exist
	failureUpdateArgs := []string{"update",
		"--name", "not-exist",
	}
	_, err, _ = TestFunctionsCommands(updateFunctionsCmd, failureUpdateArgs)
	assert.NotNil(t, err)
	failMsg := "Function not-exist doesn't exist"
	assert.True(t, strings.Contains(err.Error(), failMsg))

	//test no change for update
	noChangeArgs := []string{"update",
		"--name", "test-functions-update-failure",
		"--output", "persistent://public/default/test-output-topic",
	}

	_, err, _ = TestFunctionsCommands(updateFunctionsCmd, noChangeArgs)
	assert.NotNil(t, err)
	failNoChangeMsg := "Update contains no change"
	assert.True(t, strings.Contains(err.Error(), failNoChangeMsg))
}
