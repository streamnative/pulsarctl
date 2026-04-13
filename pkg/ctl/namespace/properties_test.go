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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamespacePropertiesCmd(t *testing.T) {
	ns := "public/test-namespace-properties"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"clear-properties", ns}
	_, execErr, _, _ = TestNamespaceCommands(ClearPropertiesCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-properties", ns, "-p", "k1=v1,k2=v2"}
	setPropertiesOut, execErr, _, _ := TestNamespaceCommands(SetPropertiesCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Updated properties successfully for [%s]\n", ns), setPropertiesOut.String())

	args = []string{"get-properties", ns}
	getPropertiesOut, execErr, _, _ := TestNamespaceCommands(GetPropertiesCmd, args)
	assert.Nil(t, execErr)
	properties := map[string]string{}
	err := json.Unmarshal(getPropertiesOut.Bytes(), &properties)
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"k1": "v1", "k2": "v2"}, properties)

	args = []string{"get-property", ns, "-k", "k1"}
	getPropertyOut, execErr, _, _ := TestNamespaceCommands(GetPropertyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "v1\n", getPropertyOut.String())

	args = []string{"set-property", ns, "-k", "k3", "--value", "v3"}
	setPropertyOut, execErr, _, _ := TestNamespaceCommands(SetPropertyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Set property %q successfully for [%s]\n", "k3", ns), setPropertyOut.String())

	args = []string{"get-properties", ns}
	getPropertiesOut, execErr, _, _ = TestNamespaceCommands(GetPropertiesCmd, args)
	assert.Nil(t, execErr)
	properties = map[string]string{}
	err = json.Unmarshal(getPropertiesOut.Bytes(), &properties)
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}, properties)

	args = []string{"remove-property", ns, "-k", "k2"}
	removePropertyOut, execErr, _, _ := TestNamespaceCommands(RemovePropertyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "v2\n", removePropertyOut.String())

	args = []string{"remove-property", ns, "-k", "k1"}
	_, execErr, _, _ = TestNamespaceCommands(RemovePropertyCmd, args)
	assert.Nil(t, execErr)

	args = []string{"remove-property", ns, "-k", "k3"}
	_, execErr, _, _ = TestNamespaceCommands(RemovePropertyCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-properties", ns}
	getPropertiesOut, execErr, _, _ = TestNamespaceCommands(GetPropertiesCmd, args)
	assert.Nil(t, execErr)
	properties = map[string]string{}
	err = json.Unmarshal(getPropertiesOut.Bytes(), &properties)
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{}, properties)
}

func TestGetMissingPropertyCmd(t *testing.T) {
	ns := "public/test-missing-namespace-property"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"clear-properties", ns}
	_, execErr, _, _ = TestNamespaceCommands(ClearPropertiesCmd, args)
	assert.Nil(t, execErr)

	args = []string{"get-property", ns, "-k", "missing"}
	out, execErr, _, _ := TestNamespaceCommands(GetPropertyCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "null\n", out.String())
}

func TestRemovePropertyCmdUsesSinglePropertyEndpoint(t *testing.T) {
	ns := "public/test-namespace-properties"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t,
			"/admin/v2/namespaces/public/test-namespace-properties/property/k2",
			r.URL.Path)
		_, _ = w.Write([]byte("v2"))
	}))
	defer srv.Close()

	withNamespaceAdminURLForTest(t, srv.URL)

	out, execErr, _, err := TestNamespaceCommands(RemovePropertyCmd, []string{"remove-property", ns, "-k", "k2"})
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "v2\n", out.String())
}
