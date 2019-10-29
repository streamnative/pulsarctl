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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNsTopicsCmd(t *testing.T) {
	args := []string{"create", "public/test-topics-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-topics-namespace successfully\n")

	args = []string{"topics", "public/test-topics-namespace"}
	_, execErr, _, _ := TestNamespaceCommands(getTopics, args)
	assert.Nil(t, execErr)
}

func TestListTopicArgError(t *testing.T) {
	args := []string{"topics"}
	_, _, nameErr, _ := TestNamespaceCommands(getTopics, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the namespace name is not specified or the namespace name is "+
		"specified more than one", nameErr.Error())
}

func TestListNonExistNamespace(t *testing.T) {
	args := []string{"topics", "public/non-exist-namespace"}
	_, execErr, _, _ := TestNamespaceCommands(getTopics, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestListNonExistTenant(t *testing.T) {
	args := []string{"topics", "non-exist-tenant/default"}
	_, execErr, _, _ := TestNamespaceCommands(getTopics, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Tenant does not exist", execErr.Error())
}
