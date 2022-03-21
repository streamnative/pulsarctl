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
	"github.com/stretchr/testify/require"
)

func TestSplitBundle(t *testing.T) {
	ns := "public/test-split-bundle-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	require.Nil(t, execErr)

	args = []string{"split-bundle", ns, "--bundle", "0x80000000_0xc0000000"}
	_, execErr, _, _ = TestNamespaceCommands(splitBundle, args)
	require.Nil(t, execErr)

	args = []string{"create", ns + "/test-topic", "0"}
	_, _, argsErr, err := topic.TestTopicCommands(topic.CreateTopicCmd, args)
	require.Nil(t, argsErr)
	require.Nil(t, err)

	args = []string{"bundle-range", ns + "/test-topic"}
	out, execErr, _, _ := topic.TestTopicCommands(topic.GetBundleRangeCmd, args)
	require.Nil(t, execErr)

	bundle := strings.Split(out.String(), ":")[2]
	bundle = strings.TrimSpace(bundle)

	args = []string{"split-bundle", ns, "--bundle", bundle}
	splitOut, execErr, _, _ := TestNamespaceCommands(splitBundle, args)
	require.Nil(t, execErr)
	require.Equal(t, splitOut.String(), "Split a namespace bundle: "+bundle+" successfully\n")
}
