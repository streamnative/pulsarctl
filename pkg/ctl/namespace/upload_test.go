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

func TestUpload(t *testing.T) {
	args := []string{"create", "public/test-unload-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-unload-namespace successfully\n")

	args = []string{"unload", "public/test-unload-namespace"}
	unloadOut, execErr, _, _ := TestNamespaceCommands(unload, args)
	assert.Nil(t, execErr)
	assert.Equal(t, unloadOut.String(), "Unload namespace public/test-unload-namespace successfully\n")

	argsWithBundle := []string{"unload", "public/test-unload-namespace", "--bundle", "0x40000000_0x80000000"}
	unloadOut, execErr, _, _ = TestNamespaceCommands(unload, argsWithBundle)
	assert.Nil(t, execErr)
	assert.Equal(t, unloadOut.String(),
		"Unload namespace public/test-unload-namespace with bundle 0x40000000_0x80000000 successfully\n")

	// test invalid upper boundary for bundle
	argsWithInvalidBundle := []string{"unload", "public/test-unload-namespace", "--bundle", "0x00000000_0x60000000"}
	_, execErr, _, _ = TestNamespaceCommands(unload, argsWithInvalidBundle)
	assert.NotNil(t, execErr)
}
