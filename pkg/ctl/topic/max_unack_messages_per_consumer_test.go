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

func TestMaxUnackMessagesPerConsumer(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := "persistent://public/default/test-max-unacked-messages-per-consumer-topic" + test.RandomSuffix()
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-max-unacked-messages-per-consumer", topicName, "-m", "20"}
	g.Eventually(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(SetMaxUnackMessagesPerConsumerCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(
			gomega.Equal("Set max unacked messages per consumer successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	getArgs := []string{"get-max-unacked-messages-per-consumer", topicName}
	g.Expect(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetMaxUnackMessagesPerConsumerCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(getOut.String()).Should(gomega.Equal("20"))
	}).Should(gomega.Succeed())

	removeArgs := []string{"remove-max-unacked-messages-per-consumer", topicName}
	g.Expect(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(RemoveMaxUnackMessagesPerConsumerCmd, removeArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(
			gomega.Equal("Remove max unacked messages per consumer successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	g.Expect(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetMaxUnackMessagesPerConsumerCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(getOut.String()).Should(gomega.Equal("0"))
	}).Should(gomega.Succeed())

	// test negative value
	setArgs = []string{"set-max-unacked-messages-per-consumer", topicName, "-m", "-2"}
	g.Expect(func(g gomega.Gomega) {
		_, execErr, _, _ = TestTopicCommands(SetMaxUnackMessagesPerConsumerCmd, setArgs)
		g.Expect(execErr).ShouldNot(gomega.BeNil())
		g.Expect(execErr.Error()).Should(gomega.Equal("code: 412 reason: maxUnackedNum must be 0 or more"))
	}).Should(gomega.Succeed())
}
