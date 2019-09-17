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
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeleteNsCmd(t *testing.T) {
	args := []string{"create", "public/test-namespace-1"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-namespace-1 successfully")

	args = []string{"delete", "public/test-namespace-1"}
	delOut, _, _, _ := TestNamespaceCommands(deleteNs, args)
	assert.Equal(t, delOut.String(), "Deleted public/test-namespace-1 successfully")

	args = []string{"list", "public"}
	listOut, _, _, _ := TestNamespaceCommands(getNamespacesFromTenant, args)
	assert.False(t, strings.Contains(listOut.String(), "public/test-namespace-1"))
}

func TestDeleteNsArgsError(t *testing.T) {
	args := []string{"delete"}
	_, _, nameErr, _ := TestNamespaceCommands(deleteNs, args)
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestDeleteNonExistentTenant(t *testing.T) {
	args := []string{"delete", "non-existent-tenant/default"}
	_, execErr, _, _ := TestNamespaceCommands(deleteNs, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}

func TestDeleteNonExistentNamespace(t *testing.T) {
	args := []string{"delete", "public/non-existent-namespace"}
	_, execErr, _, _ := TestNamespaceCommands(deleteNs, args)
	assert.NotNil(t, execErr)
	nonExistentNsErr := "code: 404 reason: Namespace public/non-existent-namespace does not exist."
	assert.Equal(t, nonExistentNsErr, execErr.Error())
}
