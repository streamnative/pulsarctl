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

package topic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTopicsCmd(t *testing.T) {
	args := []string{"list", "public/default"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.Nil(t, execErr)
}

func TestListTopicArgError(t *testing.T) {
	args := []string{"list"}
	_, _, nameErr, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the namespace name is not specified or the namespace name is "+
		"specified more than one", nameErr.Error())
}

func TestListNonExistNamespace(t *testing.T) {
	args := []string{"list", "public/non-exist-namespace"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestListNonExistTenant(t *testing.T) {
	args := []string{"list", "non-exist-tenant/default"}
	_, execErr, _, _ := TestTopicCommands(ListTopicsCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "404")
}
