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
	. "github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGetFunction(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-get",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	getArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-get",
	}

	out, _, err := TestFunctionsCommands(getFunctionsCmd, getArgs)
	assert.Nil(t, err)

	var functionConfig FunctionConfig
	err = json.Unmarshal(out.Bytes(), &functionConfig)
	assert.Nil(t, err)

	assert.Equal(t, functionConfig.Tenant, "public")
	assert.Equal(t, functionConfig.Namespace, "default")
	assert.Equal(t, functionConfig.Name, "test-functions-get")
	assert.Equal(t, functionConfig.Output, "persistent://public/default/test-output-topic")
	assert.Equal(t, functionConfig.ClassName, "org.apache.pulsar.functions.api.examples.ExclamationFunction")
}

func TestGetFunctionsWithFailure(t *testing.T) {
	jarName := "dummyExample.jar"
	_, err := os.Create(jarName)
	assert.Nil(t, err)

	defer os.Remove(jarName)
	args := []string{"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-functions-get-failure",
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
	}

	_, _, err = TestFunctionsCommands(createFunctionsCmd, args)
	assert.Nil(t, err)

	failureDeleteArgs := []string{"get",
		"--name", "not-exist",
	}

	_, execErrMsg, _ := TestFunctionsCommands(getFunctionsCmd, failureDeleteArgs)
	assert.NotNil(t, execErrMsg)
	exceptMsg := "Function not-exist doesn't exist"
	assert.True(t, strings.Contains(execErrMsg.Error(), exceptMsg))

	notExistNameOrFqfnArgs := []string{"get",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, execErrMsg, _ = TestFunctionsCommands(getFunctionsCmd, notExistNameOrFqfnArgs)
	failMsg := "you must specify a name for the function or a Fully Qualified Function Name (FQFN)"
	assert.NotNil(t, execErrMsg)
	assert.True(t, strings.Contains(execErrMsg.Error(), failMsg))
}
