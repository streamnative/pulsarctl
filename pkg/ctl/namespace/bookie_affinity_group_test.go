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

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBookieAffinityGroupCmd(t *testing.T) {
	ns := "public/test-bookie-affinity-group"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-bookie-affinity-group", ns}
	getOut, execErr, _, _ := TestNamespaceCommands(GetBookieAffinityGroupCmd, args)
	assert.Nil(t, execErr)
	var initialGroup *utils.BookieAffinityGroupData
	err := json.Unmarshal(getOut.Bytes(), &initialGroup)
	assert.Nil(t, err)

	args = []string{"set-bookie-affinity-group", ns, "--primary-group", "primary", "--secondary-group", "secondary"}
	setOut, execErr, _, _ := TestNamespaceCommands(SetBookieAffinityGroupCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Set bookie affinity group successfully for [%s]\n", ns), setOut.String())

	args = []string{"get-bookie-affinity-group", ns}
	getOut, execErr, _, _ = TestNamespaceCommands(GetBookieAffinityGroupCmd, args)
	assert.Nil(t, execErr)
	var currentGroup *utils.BookieAffinityGroupData
	err = json.Unmarshal(getOut.Bytes(), &currentGroup)
	assert.Nil(t, err)
	if !assert.NotNil(t, currentGroup) {
		return
	}
	assert.Equal(t, "primary", currentGroup.BookkeeperAffinityGroupPrimary)
	assert.Equal(t, "secondary", currentGroup.BookkeeperAffinityGroupSecondary)

	args = []string{"delete-bookie-affinity-group", ns}
	delOut, execErr, _, _ := TestNamespaceCommands(DeleteBookieAffinityGroupCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("Deleted bookie affinity group successfully for [%s]\n", ns), delOut.String())

	args = []string{"get-bookie-affinity-group", ns}
	getOut, execErr, _, _ = TestNamespaceCommands(GetBookieAffinityGroupCmd, args)
	assert.Nil(t, execErr)
	var afterDeleteGroup *utils.BookieAffinityGroupData
	err = json.Unmarshal(getOut.Bytes(), &afterDeleteGroup)
	assert.Nil(t, err)
	assert.Equal(t, initialGroup, afterDeleteGroup)
}

func TestGetBookieAffinityGroupCmdLocalPoliciesMissing(t *testing.T) {
	ns := "public/test-bookie-affinity-group"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t,
			"/admin/v2/namespaces/public/test-bookie-affinity-group/persistence/bookieAffinity",
			r.URL.Path)

		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"reason":"Namespace local-policies does not exist"}`))
	}))
	defer srv.Close()

	withNamespaceAdminURLForTest(t, srv.URL)

	out, execErr, _, err := TestNamespaceCommands(GetBookieAffinityGroupCmd, []string{"get-bookie-affinity-group", ns})
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, "null", out.String())
}
