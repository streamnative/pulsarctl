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
	"path"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestCreateJavaFunctions(t *testing.T) {
	jarName := path.Join(ResourceDir(), "api-examples.jar")
	fName := "jf" + test.RandomSuffix()
	args := []string{
		"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--classname", "org.apache.pulsar.functions.api.examples.ExclamationFunction",
		"--jar", jarName,
		"--processing-guarantees", "EFFECTIVELY_ONCE"}
	_, execErr, err := TestFunctionsCommands(createFunctionsCmd, args)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
}

func TestCreatePyFunction(t *testing.T) {
	pyName := path.Join(ResourceDir(), "logging_function.py")
	fName := "function-py" + test.RandomSuffix()
	argsWithFileURLPy := []string{
		"create",
		"--tenant", "public",
		"--namespace", "default",
		"--name", fName,
		"--inputs", "test-input-topic",
		"--output", "persistent://public/default/test-output-topic",
		"--py", pyName,
		"--classname", "LoggingFunction",
	}
	_, execErr, err := TestFunctionsCommands(createFunctionsCmd, argsWithFileURLPy)
	FailImmediatelyIfErrorNotNil(t, execErr, err)
}
