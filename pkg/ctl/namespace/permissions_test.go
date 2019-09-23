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
	"encoding/json"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermissionsCmd(t *testing.T) {
	ns := "public/test-permissions-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"permissions", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetPermissionsCmd, args)
	assert.Nil(t, execErr)

	var permissions map[string][]pulsar.AuthAction
	err := json.Unmarshal(out.Bytes(), &permissions)
	if err != nil {
		t.Fatal(err)
	}

	empty := make(map[string][]pulsar.AuthAction)

	assert.Equal(t, empty, permissions)
}

func TestGetPermissionsArgsError(t *testing.T)  {
	args := []string{"permissions"}
	_, _, nameErr, _ := TestNamespaceCommands(GetPermissionsCmd, args)
	assert.NotNil(t, nameErr.Error())
	assert.Equal(t, "only one argument is allowed to be used as a name", nameErr.Error())
}
