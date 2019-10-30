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

func TestOffloadCmd(t *testing.T) {
	args := []string{"create", "test-offload-topic", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"offload", "test-offload-topic", "10M"}
	out, execErr, _, _ := TestTopicCommands(OffloadCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "Nothing to offload\n", out.String())
}

func TestOffloadArgsError(t *testing.T) {
	args := []string{"offload", "test-offload-topic-args-error"}
	_, _, nameErr, _ := TestTopicCommands(OffloadCmd, args)
	assert.Equal(t, "only two arguments are allowed to be used as names", nameErr.Error())
}

func TestOffloadNonExistingTopicError(t *testing.T) {
	args := []string{"offload", "test-offload-non-existing-topic", "10m"}
	_, execErr, _, _ := TestTopicCommands(OffloadCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Topic not found", execErr.Error())
}
