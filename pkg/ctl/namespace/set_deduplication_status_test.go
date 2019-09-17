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

func TestDeduplicationStatus(t *testing.T) {
	args := []string{"set-deduplication", "public/default"}
	out, execErr, _, _ := TestNamespaceCommands(setDeduplication, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Set deduplication is [false] successfully for public/default")

	args = []string{"set-deduplication", "public/default", "--enable"}
	out, execErr, _, _ = TestNamespaceCommands(setDeduplication, args)
	assert.Nil(t, execErr)
	assert.Equal(t, out.String(), "Set deduplication is [true] successfully for public/default")
}
