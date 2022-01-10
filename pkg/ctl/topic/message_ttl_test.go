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

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestMessageTTL(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := "persistent://public/default/test-message-ttl-topic" + test.RandomSuffix()
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	setTTLArgs := []string{"set-message-ttl", topicName, "-t", "20"}
	g.Eventually(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(SetMessageTTLCmd, setTTLArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(gomega.Equal("Set message TTL successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	getTTLArgs := []string{"get-message-ttl", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetMessageTTLCmd, getTTLArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(getOut.String()).Should(gomega.Equal("20"))
	}).Should(gomega.Succeed())

	removeTTLArgs := []string{"remove-message-ttl", topicName}
	g.Eventually(func(g gomega.Gomega) {
		removeOut, execErr, _, _ := TestTopicCommands(RemoveMessageTTLCmd, removeTTLArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(removeOut.String()).Should(gomega.Equal("Remove message TTL successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	getTTLArgs = []string{"get-message-ttl", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetMessageTTLCmd, getTTLArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(getOut.String()).Should(gomega.Equal("0"))
	}).Should(gomega.Succeed())

	// test negative value
	setTTLArgs = []string{"set-message-ttl", topicName, "-t", "-2"}
	g.Eventually(func(g gomega.Gomega) {
		_, execErr, _, _ = TestTopicCommands(SetMessageTTLCmd, setTTLArgs)
		g.Expect(execErr).ShouldNot(gomega.BeNil())
		g.Expect(execErr.Error()).Should(gomega.Equal("code: 412 reason: Invalid value for message TTL"))
	}).Should(gomega.Succeed())
}
