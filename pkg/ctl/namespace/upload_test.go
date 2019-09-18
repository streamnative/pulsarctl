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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpload(t *testing.T) {
	args := []string{"unload", "public/default"}
	unloadOut, execErr, _, _ := TestNamespaceCommands(unload, args)
	assert.Nil(t, execErr)
	assert.Equal(t, unloadOut.String(), "Unload namespace public/default successfully")

	argsWithBundle := []string{"unload", "public/default", "--bundle", "0x40000000_0x80000000"}
	unloadOut, execErr, _, _ = TestNamespaceCommands(unload, argsWithBundle)
	assert.Nil(t, execErr)
	assert.Equal(t, unloadOut.String(), "Unload namespace public/default with bundle 0x40000000_0x80000000 successfully")

	// test invalid upper boundary for bundle
	argsWithInvalidBundle := []string{"unload", "public/default", "--bundle", "0x00000000_0x60000000"}
	_, execErr, _, _ = TestNamespaceCommands(unload, argsWithInvalidBundle)
	assert.NotNil(t, execErr)
	assert.Equal(t, execErr.Error(), "code: 500 reason: Unknown pulsar error")
}
