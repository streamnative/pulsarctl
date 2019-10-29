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

func TestDeduplicationStatus(t *testing.T) {
	args := []string{"create", "public/test-dedup-namespace"}
	createOut, _, _, err := TestNamespaceCommands(createNs, args)
	assert.Nil(t, err)
	assert.Equal(t, createOut.String(), "Created public/test-dedup-namespace successfully\n")

	args = []string{"set-deduplication", "public/test-dedup-namespace"}
	out, execErr, _, _ := TestNamespaceCommands(setDeduplication, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Set deduplication is [false] successfully for public/test-dedup-namespace\n")

	args = []string{"set-deduplication", "public/test-dedup-namespace", "--enable"}
	out, execErr, _, _ = TestNamespaceCommands(setDeduplication, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Set deduplication is [true] successfully for public/test-dedup-namespace\n")
}
