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
	"fmt"
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
	out, _, err := TestSchemasCommands(uploadSchema, args)
	assert.NoError(t, err)
	assert.Equal(t, "Upload test-schema successfully\n", out.String())

	getArgs := []string{"get", "test-schema"}
	getOut, _, err := TestSchemasCommands(getSchema, getArgs)

	fmt.Print(getOut.String())
	assert.NoError(t, err)
	assert.True(t, strings.Contains(getOut.String(), "AVRO"))
	assert.True(t, strings.Contains(getOut.String(), "test-schema"))

	delArgs := []string{"delete", "test-schema"}
	delOut, _, err := TestSchemasCommands(deleteSchema, delArgs)
	assert.Nil(t, err)
	assert.Equal(t, delOut.String(), "Deleted test-schema successfully\n")
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

func TestSchemaForceDelete(t *testing.T) {
	fileName := "avro-schema-force"
	f, err := os.Create(fileName)
	assert.Nil(t, err)
	defer func(name string) {
		_ = os.Remove(name)
	}(fileName)

	_, err = f.WriteString("{\n" +
		"   \"type\": \"AVRO\",\n" +
		"   \"schema\": " +
		"\"{\\\"type\\\":\\\"record\\\"," +
		"\\\"name\\\":\\\"TestForce\\\"," +
		"\\\"fields\\\":[{" +
		"\\\"name\\\":\\\"id\\\"," +
		"\\\"type\\\":[\\\"null\\\",\\\"int\\\"]}," +
		"{\\\"name\\\":\\\"name\\\",\\\"type\\\":[\\\"null\\\",\\\"string\\\"]}]}\",\n" +
		"   \"properties\": {}\n" +
		"}\n")
	assert.Nil(t, err)

	// Upload schema first
	args := []string{"upload", "test-schema-force", "-f", fileName}
	out, _, err := TestSchemasCommands(uploadSchema, args)
	assert.NoError(t, err)
	assert.Equal(t, "Upload test-schema-force successfully\n", out.String())

	// Test force delete with -f flag
	delArgs := []string{"delete", "test-schema-force", "-f"}
	delOut, _, err := TestSchemasCommands(deleteSchema, delArgs)
	assert.Nil(t, err)
	assert.Equal(t, delOut.String(), "Deleted test-schema-force successfully\n")
}

func TestSchemaForceDeleteLongFlag(t *testing.T) {
	fileName := "avro-schema-force-long"
	f, err := os.Create(fileName)
	assert.Nil(t, err)
	defer func(name string) {
		_ = os.Remove(name)
	}(fileName)

	_, err = f.WriteString("{\n" +
		"   \"type\": \"AVRO\",\n" +
		"   \"schema\": " +
		"\"{\\\"type\\\":\\\"record\\\"," +
		"\\\"name\\\":\\\"TestForceLong\\\"," +
		"\\\"fields\\\":[{" +
		"\\\"name\\\":\\\"id\\\"," +
		"\\\"type\\\":[\\\"null\\\",\\\"int\\\"]}," +
		"{\\\"name\\\":\\\"name\\\",\\\"type\\\":[\\\"null\\\",\\\"string\\\"]}]}\",\n" +
		"   \"properties\": {}\n" +
		"}\n")
	assert.Nil(t, err)

	// Upload schema first
	args := []string{"upload", "test-schema-force-long", "-f", fileName}
	out, _, err := TestSchemasCommands(uploadSchema, args)
	assert.NoError(t, err)
	assert.Equal(t, "Upload test-schema-force-long successfully\n", out.String())

	// Test force delete with --force flag
	delArgs := []string{"delete", "test-schema-force-long", "--force"}
	delOut, _, err := TestSchemasCommands(deleteSchema, delArgs)
	assert.Nil(t, err)
	assert.Equal(t, delOut.String(), "Deleted test-schema-force-long successfully\n")
}
