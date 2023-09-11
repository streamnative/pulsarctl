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

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSchemaAutoUpdateStrategyCmd(t *testing.T) {
	ns := "public/test-schema-autoupdate-strategy"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-schema-autoupdate-strategy", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema auto-update strategy of the namespace %s is %s\n", ns, utils.Full.String()),
		out.String())

	args = []string{"set-schema-autoupdate-strategy", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the schema auto-update strategy of the namespace %s to %s\n",
			ns, utils.AutoUpdateDisabled.String()),
		out.String())

	args = []string{"get-schema-autoupdate-strategy", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema auto-update strategy of the namespace %s is %s\n",
			ns, utils.AutoUpdateDisabled.String()),
		out.String())

	args = []string{"set-schema-autoupdate-strategy", "--compatibility", "BackwardTransitive", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the schema auto-update strategy of the namespace %s to %s\n",
			ns, utils.BackwardTransitive.String()),
		out.String())

	args = []string{"get-schema-autoupdate-strategy", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The schema auto-update strategy of the namespace %s is %s\n",
			ns, utils.BackwardTransitive.String()),
		out.String())
}

func TestSetSchemaAutoUpdateStrategyOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-schema-autoupdate-strategy", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSchemaAutoUpdateStrategyCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestGetSchemaAutoUpdateStrategyOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-schema-autoupdate-strategy", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetSchemaAutoUpdateStrategyCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
