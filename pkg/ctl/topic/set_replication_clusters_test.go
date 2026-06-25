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
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestSetReplicationClustersCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	topic := fmt.Sprintf("test-replication-clusters-topic-%s", test.RandomSuffix())

	args := []string{"create", topic, "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	args = []string{"set-replication-clusters", topic, "--clusters", "standalone"}
	out, execErr, nameErr, cmdErr := TestTopicCommands(SetReplicationClustersCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(nameErr).Should(gomega.BeNil())
	g.Expect(cmdErr).Should(gomega.BeNil())
	g.Expect(out).ShouldNot(gomega.BeNil())
	g.Expect(out.String()).ShouldNot(gomega.BeEmpty())

	// Since there is no get-replication-clusters command in this PR, we only test the set command success.
	// In a real scenario, we might want to verify using the client or adding a get command.
	// The set command output verification implies the call was successful.
}
