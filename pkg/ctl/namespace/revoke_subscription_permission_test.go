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

func TestRevokeSubPermissionsArgsError(t *testing.T) {
	ns := "public/revoke-sub-permissions-args-tests"

	args := []string{"revoke-subscription-permission", ns}
	_, _, _, err := TestNamespaceCommands(RevokeSubPermissionsCmd, args)
	assert.NotNil(t, err)
	assert.Equal(t, "required flag(s) \"role\" not set", err.Error())

	args = []string{"revoke-subscription-permission", "--role", "test-role"}
	_, _, nameErr, _ := TestNamespaceCommands(RevokeSubPermissionsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "need to specified namespace name and subscription name", nameErr.Error())
}
