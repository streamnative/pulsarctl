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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestRetentionCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	topic := fmt.Sprintf("test-retention-topic-%s", test.RandomSuffix())

	args := []string{"create", topic, "0"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	args = []string{"set-retention", topic, "--time", "12h", "--size", "100g"}
	out, execErr, nameErr, cmdErr := TestTopicCommands(SetRetentionCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(nameErr).Should(gomega.BeNil())
	g.Expect(cmdErr).Should(gomega.BeNil())
	g.Expect(out).ShouldNot(gomega.BeNil())
	g.Expect(out.String()).ShouldNot(gomega.BeEmpty())

	args = []string{"get-retention", topic}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, nameErr, cmdErr = TestTopicCommands(GetRetentionCmd, args)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(nameErr).Should(gomega.BeNil())
		g.Expect(cmdErr).Should(gomega.BeNil())
		g.Expect(out).ShouldNot(gomega.BeNil())
		g.Expect(out.String()).ShouldNot(gomega.BeEmpty())
		var data utils.RetentionPolicies
		err := json.Unmarshal(out.Bytes(), &data)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(data.RetentionTimeInMinutes).Should(gomega.Equal(720))
		g.Expect(data.RetentionSizeInMB).Should(gomega.Equal(int64(102400)))
	}).Should(gomega.Succeed())

	args = []string{"remove-retention", topic}
	out, execErr, nameErr, cmdErr = TestTopicCommands(RemoveRetentionCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(nameErr).Should(gomega.BeNil())
	g.Expect(cmdErr).Should(gomega.BeNil())
	g.Expect(out).ShouldNot(gomega.BeNil())
	g.Expect(out.String()).ShouldNot(gomega.BeEmpty())

	args = []string{"get-retention", topic}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, nameErr, cmdErr = TestTopicCommands(GetRetentionCmd, args)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(nameErr).Should(gomega.BeNil())
		g.Expect(cmdErr).Should(gomega.BeNil())
		g.Expect(out).ShouldNot(gomega.BeNil())
		g.Expect(out.String()).ShouldNot(gomega.BeEmpty())
		var data utils.RetentionPolicies
		err := json.Unmarshal(out.Bytes(), &data)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(data.RetentionTimeInMinutes).Should(gomega.Equal(0))
		g.Expect(data.RetentionSizeInMB).Should(gomega.Equal(int64(0)))
	}).Should(gomega.Succeed())
}
