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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxConsumersPerTopicCmd(t *testing.T) {
	ns := "public/test-max-consumers-per-topic"
	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-max-consumers-per-topic", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max consumers per topic of the namespace %s is %d\n", ns, 0),
		out.String())

	args = []string{"set-max-consumers-per-topic", "--size", "10", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully set the max consumers per topic of the namespace %s to %d\n", ns, 10),
		out.String())

	args = []string{"get-max-consumers-per-topic", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetMaxConsumersPerTopicCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The max consumers per topic of the namespace %s is %d\n", ns, 10),
		out.String())
}

func TestSetMaxConsumersPerTopicOnNonExistingTopic(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"set-max-consumers-per-topic", "--size", "10", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxConsumersPerTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}

func TestSetMaxConsumersPerTopicWithInvalidSize(t *testing.T) {
	args := []string{"set-max-consumers-per-topic", "--size", "-1", "public/invalid-size"}
	_, execErr, _, _ := TestNamespaceCommands(SetMaxConsumersPerTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "the specified consumers value must bigger than 0", execErr.Error())
}

func TestGetMaxConsumersPerTopicOnNonExistingTopic(t *testing.T) {
	ns := "public/non-existing-namespace"
	args := []string{"get-max-consumers-per-topic", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetMaxConsumersPerTopicCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
