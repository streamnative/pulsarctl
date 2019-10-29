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

func TestGetBundleRangeCmd(t *testing.T) {
	args := []string{"create", "test-get-topic-bundle-range", "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"bundle-range", "test-get-topic-bundle-range"}
	out, execErr, _, _ := TestTopicCommands(GetBundleRangeCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t, "The bundle range of the topic "+
		"persistent://public/default/test-get-topic-bundle-range is: 0xc0000000_0xffffffff\n", out.String())
}

func TestGetBundleRangeArgError(t *testing.T) {
	args := []string{"bundle-range"}
	_, _, nameErr, _ := TestTopicCommands(GetBundleRangeCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one", nameErr.Error())
}
