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

func TestMessageTTL(t *testing.T) {
	args := []string{"create", "public/test-ttl-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-ttl-namespace successfully\n")

	setTTLArgs := []string{"set-message-ttl", "public/test-ttl-namespace", "-t", "20"}
	setOut, execErr, _, _ := TestNamespaceCommands(setMessageTTL, setTTLArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, setOut.String(), "Set message TTL successfully for [public/test-ttl-namespace]\n")

	getTTLArgs := []string{"get-message-ttl", "public/test-ttl-namespace"}
	getOut, execErr, _, _ := TestNamespaceCommands(getMessageTTL, getTTLArgs)
	assert.Nil(t, execErr)
	assert.Equal(t, getOut.String(), "20")

	// test negative value for ttl arg
	setTTLArgs = []string{"set-message-ttl", "public/test-ttl-namespace", "-t", "-2"}
	_, execErr, _, _ = TestNamespaceCommands(setMessageTTL, setTTLArgs)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 412 reason: Invalid value for message TTL")
}
