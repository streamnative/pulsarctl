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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubscriptionAuthModeArgsError(t *testing.T) {
	args := []string{"get-subscription-auth-mode"}
	_, _, nameErr, _ := TestNamespaceCommands(GetSubscriptionAuthModeCmd, args)
	assert.Equal(t, "the namespace name is not specified or the namespace name is specified more than one",
		nameErr.Error())
}

func TestGetSubscriptionAuthModeCmd(t *testing.T) {
	ns := "public/test-get-subscription-auth-mode-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-subscription-auth-mode", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSubscriptionAuthModeCmd, args)
	assert.Nil(t, execErr)

	var mode string
	err := json.Unmarshal(out.Bytes(), &mode)
	assert.Nil(t, err)
	assert.Equal(t, "None", mode)
}
