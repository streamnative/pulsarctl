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
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/ctl/topic"
	"github.com/stretchr/testify/assert"
)

func TestSplitBundle(t *testing.T) {
	ns := "public/test-split-bundle-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"split-bundle", ns, "--bundle", "0x80000000_0xc0000000"}
	_, execErr, _, _ = TestNamespaceCommands(splitBundle, args)
	assert.NotNil(t, execErr)
	errMsg := "code: 412 reason: Failed to find ownership for ServiceUnit:" + ns + "/0x80000000_0xc0000000"
	assert.Equal(t, execErr.Error(), errMsg)

	args = []string{"create", ns + "/test-topic", "0"}
	_, _, argsErr, err := topic.TestTopicCommands(topic.CreateTopicCmd, args)
	assert.Nil(t, argsErr)
	assert.Nil(t, err)

	args = []string{"bundle-range", ns + "/test-topic"}
	out, execErr, _, _ := topic.TestTopicCommands(topic.GetBundleRangeCmd, args)
	assert.Nil(t, execErr)

	bundle := strings.Split(out.String(), ":")[2]
	bundle = strings.TrimSpace(bundle)

	args = []string{"split-bundle", ns, "--bundle", bundle}
	splitOut, execErr, _, _ := TestNamespaceCommands(splitBundle, args)
	assert.Nil(t, execErr)
	assert.Equal(t, splitOut.String(), "Split a namespace bundle: "+bundle+" successfully\n")
}
