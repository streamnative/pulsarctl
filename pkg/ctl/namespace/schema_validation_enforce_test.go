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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaValidationEnforcedCmd(t *testing.T) {
	ns := "public/test-schema-validation-enforced"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-schema-validation-enforced", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSchemaValidationEnforcedCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Namespace %s schema validation enforced is disabled\n", ns), out.String())

	args = []string{"set-schema-validation-enforced", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSchemaValidationEnforcedCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Enable the namespace %s schema validation enforced\n", ns), out.String())

	args = []string{"get-schema-validation-enforced", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSchemaValidationEnforcedCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Namespace %s schema validation enforced is enabled\n", ns), out.String())
}

func TestSetSchemaValidationEnforcedOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-schema-validation-enforced", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSchemaValidationEnforcedCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetSchemaValidationEnforcedOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-schema-validation-enforced", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetSchemaValidationEnforcedCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
