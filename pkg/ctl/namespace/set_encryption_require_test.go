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
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestSetEncryptionRequiredCmd(t *testing.T) {
	ns := "public/test-set-message-encryption-test"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"messages-encryption", ns}
	out, execErr, _, _ := TestNamespaceCommands(SetEncryptionRequiredCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("%s messages encryption of the namespace %s", "Enable", ns),
		out.String())

	args = []string{"policies", ns}
	out, execErr, _, _ = TestNamespaceCommands(getPolicies, args)
	assert.Nil(t, execErr)

	var policies pulsar.Policies
	err := json.Unmarshal(out.Bytes(), &policies)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, policies.EncryptionRequired)

	args = []string{"messages-encryption", "--disable", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetEncryptionRequiredCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("%s messages encryption of the namespace %s", "Disable", ns),
		out.String())

	args = []string{"policies", ns}
	out, execErr, _, _ = TestNamespaceCommands(getPolicies, args)
	assert.Nil(t, execErr)

	err = json.Unmarshal(out.Bytes(), &policies)
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, policies.EncryptionRequired)
}
