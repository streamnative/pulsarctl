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

package info

import (
	"testing"

	. "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
)

func TestGetInternalInfoArgError(t *testing.T) {
	args := []string{"internal-info"}
	_, _, nameErr, _ := TestTopicCommands(GetInternalInfoCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}

func TestGetNonExistingTopicInternalInfo(t *testing.T) {
	args := []string{"internal-info", "non-existing-topic"}
	_, execErr, _, _ := TestTopicCommands(GetInternalInfoCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 500 reason: Unknown pulsar error", execErr.Error())
}
