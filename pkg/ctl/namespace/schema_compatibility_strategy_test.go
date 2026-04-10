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

package namespace

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSchemaCompatibilityStrategyCmd(t *testing.T) {
	ns := "public/test-schema-compatibility-strategy"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-schema-compatibility-strategy", ns}
	initialOut, execErr, _, _ := TestNamespaceCommands(GetSchemaCompatibilityStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Contains(t, initialOut.String(), ns)

	var setOut, getOut *bytes.Buffer
	args = []string{"set-schema-compatibility-strategy", "--compatibility", "FULL_TRANSITIVE", ns}
	setOut, execErr, _, _ = TestNamespaceCommands(SetSchemaCompatibilityStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the schema compatibility strategy of the namespace %s to %s\n",
			ns, utils.SchemaCompatibilityStrategyFullTransitive.String()),
		setOut.String())

	args = []string{"get-schema-compatibility-strategy", ns}
	getOut, execErr, _, _ = TestNamespaceCommands(GetSchemaCompatibilityStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema compatibility strategy of the namespace %s is %s\n",
			ns, utils.SchemaCompatibilityStrategyFullTransitive.String()),
		getOut.String())

	args = []string{"set-schema-compatibility-strategy", ns}
	_, execErr, _, _ = TestNamespaceCommands(SetSchemaCompatibilityStrategyCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "required flag(s) \"compatibility\" not set")

	args = []string{"set-schema-compatibility-strategy", "--compatibility", "INVALID", ns}
	_, execErr, _, _ = TestNamespaceCommands(SetSchemaCompatibilityStrategyCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "Invalid schema compatibility strategy INVALID", execErr.Error())

	args = []string{"get-schema-compatibility-strategy", ns}
	getOut, execErr, _, _ = TestNamespaceCommands(GetSchemaCompatibilityStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.NotEqual(t, initialOut.String(), "")
	assert.Equal(t,
		fmt.Sprintf("The schema compatibility strategy of the namespace %s is %s\n",
			ns, utils.SchemaCompatibilityStrategyFullTransitive.String()),
		getOut.String())
}
