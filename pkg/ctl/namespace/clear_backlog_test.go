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
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic"

	"github.com/stretchr/testify/assert"
)

func TestClearBacklogCmd(t *testing.T) {
	ns := "public/test-clear-backlog-test"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"clear-backlog", "-f", ns}
	out, execErr, _, _ := TestNamespaceCommands(ClearBacklogCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully clear backlog for all topics of the namespace %s\n", ns),
		out.String())

	args = []string{"clear-backlog", "-f", "--sub", "test-clear-sub", ns}
	out, execErr, _, _ = TestNamespaceCommands(ClearBacklogCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully clear backlog for all topics of the namespace %s\n", ns),
		out.String())

	args = []string{"create", ns + "/test-clear-with-bundle", "0"}
	_, execErr, _, _ = topic.TestTopicCommands(topic.CreateTopicCmd, args)
	assert.Nil(t, execErr)

	args = []string{"bundle-range", ns + "/test-clear-with-bundle"}
	out, execErr, _, _ = topic.TestTopicCommands(topic.GetBundleRangeCmd, args)
	assert.Nil(t, execErr)
	t.Logf(out.String())
	bundle := strings.Split(out.String(), ":")[2]
	t.Logf(strings.TrimSpace(bundle))
	assert.True(t, strings.HasPrefix(strings.TrimSpace(bundle), "0x"))

	args = []string{"clear-backlog", "-f", "--bundle", strings.TrimSpace(bundle), ns}
	out, execErr, _, _ = TestNamespaceCommands(ClearBacklogCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Successfully clear backlog for all topics of the namespace %s\n", ns),
		out.String())
}

func TestClearBacklogOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"clear-backlog", "-f", "--bundle", "0xc0000000_0xffffffff", ns}
	_, execErr, _, _ := TestNamespaceCommands(ClearBacklogCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())

	args = []string{"clear-backlog", "-f", "--sub", "sub", "--bundle", "0xc0000000_0xffffffff", ns}
	_, execErr, _, _ = TestNamespaceCommands(ClearBacklogCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
