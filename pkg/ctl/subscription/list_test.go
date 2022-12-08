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

package subscription

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListArgError(t *testing.T) {
	args := []string{"list"}
	_, _, nameErr, _ := TestSubCommands(ListCmd, args)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the topic name is not specified or the topic name is specified more than one",
		nameErr.Error())
}

func TestListNonExistingTopicSub(t *testing.T) {
	args := []string{"list", "non-existing-topic"}
	_, execErr, _, _ := TestSubCommands(ListCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "code: 404 reason: Topic")
}
