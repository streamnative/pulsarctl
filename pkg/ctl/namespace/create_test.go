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
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreateNs(t *testing.T) {
	args := []string{"create", "public/test-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-namespace successfully")

	args = []string{"list", "public"}
	out, _, _, _ := TestNamespaceCommands(getNamespacesPerProperty, args)
	fmt.Println(out.String())
	assert.True(t, strings.Contains(out.String(), "public/test-namespace"))
}

func TestCreateNsArgsError(t *testing.T) {
	args := []string{"create"}
	_, _, nameErr, _ := TestNamespaceCommands(createNs, args)

	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}

func TestCreateNsAlreadyExistError(t *testing.T) {
	args := []string{"create", "public/test-ns-duplicate"}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"create", "public/test-ns-duplicate"}
	_, execErr, _, _ = TestNamespaceCommands(createNs, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 409 reason: Namespace already exists", execErr.Error())
}
