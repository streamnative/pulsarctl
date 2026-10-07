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

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionExpirationTimeCmd(t *testing.T) {
	ns := "public/test-subscription-expiration-time"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-subscription-expiration-time", ns}
	initialOut, execErr, _, _ := TestNamespaceCommands(GetSubscriptionExpirationTimeCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-subscription-expiration-time", "--time", "60", ns}
	setOut, execErr, _, _ := TestNamespaceCommands(SetSubscriptionExpirationTimeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the subscription expiration time of the namespace %s to %d\n", ns, 60),
		setOut.String())

	args = []string{"get-subscription-expiration-time", ns}
	getOut, execErr, _, _ := TestNamespaceCommands(GetSubscriptionExpirationTimeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The subscription expiration time of the namespace %s is %d\n", ns, 60),
		getOut.String())

	args = []string{"remove-subscription-expiration-time", ns}
	removeOut, execErr, _, _ := TestNamespaceCommands(RemoveSubscriptionExpirationTimeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully removed the subscription expiration time of the namespace %s\n", ns),
		removeOut.String())

	args = []string{"get-subscription-expiration-time", ns}
	getOut, execErr, _, _ = TestNamespaceCommands(GetSubscriptionExpirationTimeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, initialOut.String(), getOut.String())
}

func TestSetSubscriptionExpirationTimeWithInvalidValue(t *testing.T) {
	args := []string{"set-subscription-expiration-time", "--time", "-1", "public/test-invalid-expiration-time"}
	_, execErr, _, _ := TestNamespaceCommands(SetSubscriptionExpirationTimeCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the specified subscription expiration time must bigger than or equal to 0", execErr.Error())
}
