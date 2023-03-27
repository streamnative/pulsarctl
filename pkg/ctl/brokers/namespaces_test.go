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

//go:build namespace
// +build namespace

package brokers

import (
	"encoding/json"
	"testing"

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetOwnedNamespaces(t *testing.T) {
	args := []string{"namespaces", "standalone", "--url", "127.0.0.1:8080"}
	listOut, execErr, _, _ := TestBrokersCommands(getOwnedNamespacesCmd, args)
	assert.Nil(t, execErr)

	var tmpMap map[string]utils.NamespaceOwnershipStatus
	err := json.Unmarshal(listOut.Bytes(), &tmpMap)
	assert.Nil(t, err)

	key := "pulsar/standalone/127.0.0.1:8080/0x00000000_0xffffffff"
	assert.Equal(t, 2, len(tmpMap))
	assert.Equal(t, false, tmpMap[key].IsControlled)
	assert.Equal(t, true, tmpMap[key].IsActive)

	failArgs := []string{"namespaces", "--url", "127.0.0.1:8080"}
	_, _, nameErr, _ := TestBrokersCommands(getOwnedNamespacesCmd, failArgs)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the cluster name is not specified or the cluster name is specified more than one",
		nameErr.Error())
}
