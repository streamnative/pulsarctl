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

package schemas

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	fileName := "avro-schema"
	f, err := os.Create(fileName)
	assert.Nil(t, err)
	defer func(name string) {
		_ = os.Remove(name)
	}(fileName)

	_, err = f.WriteString("{\n" +
		"   \"type\": \"AVRO\",\n" +
		"   \"schema\": " +
		"\"{\\\"type\\\":\\\"record\\\"," +
		"\\\"name\\\":\\\"Test\\\"," +
		"\\\"fields\\\":[{" +
		"\\\"name\\\":\\\"id\\\"," +
		"\\\"type\\\":[\\\"null\\\",\\\"int\\\"]}," +
		"{\\\"name\\\":\\\"name\\\",\\\"type\\\":[\\\"null\\\",\\\"string\\\"]}]}\",\n" +
		"   \"properties\": {}\n" +
		"}\n")
	assert.Nil(t, err)

	args := []string{"upload", "test-schema", "-f", fileName}
	out, execErr, err := TestSchemasCommands(uploadSchema, args)
	assert.NoError(t, err)
	assert.NoError(t, execErr)
	assert.Equal(t, "Upload test-schema successfully\n", out.String())

	getArgs := []string{"get", "test-schema"}
	getOut, execErr, err := TestSchemasCommands(getSchema, getArgs)
	assert.NoError(t, err)
	assert.NoError(t, execErr)
	assert.True(t, strings.Contains(getOut.String(), "AVRO"))
	assert.True(t, strings.Contains(getOut.String(), "test-schema"))

	allArgs := []string{"get", "test-schema", "--all-version"}
	allOut, execErr, err := TestSchemasCommands(getSchema, allArgs)
	assert.NoError(t, err)
	assert.NoError(t, execErr)
	assert.True(t, strings.Contains(allOut.String(), "version"))

	compatibilityArgs := []string{"compatibility", "test-schema", "-f", fileName}
	compatibilityOut, execErr, err := TestSchemasCommands(testCompatibility, compatibilityArgs)
	assert.NoError(t, err)
	assert.NoError(t, execErr)
	assert.True(t, strings.Contains(compatibilityOut.String(), "compatibility"))

	delArgs := []string{"delete", "test-schema"}
	delOut, execErr, err := TestSchemasCommands(deleteSchema, delArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, delOut.String(), "Deleted test-schema successfully\n")

	forceDeleteArgs := []string{"delete", "test-schema", "--force"}
	forceDeleteOut, execErr, err := TestSchemasCommands(deleteSchema, forceDeleteArgs)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, forceDeleteOut.String(), "Deleted test-schema successfully\n")
}

func TestFailSchema(t *testing.T) {
	getArgs := []string{"get", "fail-schema"}
	_, execErr, _ := TestSchemasCommands(getSchema, getArgs)
	assert.NotNil(t, execErr)

	uploadErr := "open avro-schema: no such file or directory"
	args := []string{"upload", "test-schema", "-f", "avro-schema"}
	_, execErr, _ = TestSchemasCommands(uploadSchema, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), uploadErr)
}
