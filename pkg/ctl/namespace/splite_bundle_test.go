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
	"github.com/streamnative/pulsarctl/pkg/ctl/topic/crud"
	topic "github.com/streamnative/pulsarctl/pkg/ctl/topic/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitBundle(t *testing.T) {
	args := []string{"split-bundle", "public/default", "--bundle", "0x00000000_0x40000000"}
	_, execErr, _, _ := TestNamespaceCommands(splitBundle, args)
	assert.NotNil(t, execErr)
	errMsg := "code: 412 reason: Failed to find ownership for ServiceUnit:public/default/0x00000000_0x40000000"
	assert.Equal(t, execErr.Error(), errMsg)

	topicArgs := []string{"create", "test-topic", "0"}
	_, _, argsErr, err := topic.TestTopicCommands(crud.CreateTopicCmd, topicArgs)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)

	args = []string{"split-bundle", "public/default", "--bundle", "0x00000000_0x40000000"}
	splitOut, execErr, _, _ := TestNamespaceCommands(splitBundle, args)
	assert.Nil(t, execErr)
	assert.Equal(t, splitOut.String(), "Split a namespace bundle: 0x00000000_0x40000000 successfully")
}
