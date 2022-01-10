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

func TestInactiveTopicCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := fmt.Sprintf("persistent://public/default/test-inactive-topic-%s",
		test.RandomSuffix())
	createArgs := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, createArgs)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-inactive-topic-policies", topicName,
		"-e=true",
		"-t", "1h",
		"-m", "delete_when_no_subscriptions"}
	g.Eventually(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(SetInactiveTopicCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(
			gomega.Equal(fmt.Sprintf("Set inactive topic policies successfully for [%s]", topicName)))
	}).Should(gomega.Succeed())

	getArgs := []string{"get-inactive-topic-policies", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetInactiveTopicCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var inactiveTopic utils.InactiveTopicPolicies
		err := json.Unmarshal(getOut.Bytes(), &inactiveTopic)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(inactiveTopic.DeleteWhileInactive).Should(gomega.Equal(true))
		g.Expect(inactiveTopic.MaxInactiveDurationSeconds).Should(gomega.Equal(3600))
		g.Expect(inactiveTopic.InactiveTopicDeleteMode.String()).Should(gomega.Equal("delete_when_no_subscriptions"))
	}).Should(gomega.Succeed())

	removeArgs := []string{"remove-inactive-topic-policies", topicName}
	g.Expect(func(g gomega.Gomega) {
		removeOut, execErr, _, _ := TestTopicCommands(RemoveInactiveTopicCmd, removeArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(removeOut.String()).Should(
			gomega.Equal(fmt.Sprintf("Remove inactive topic policies successfully from [%s]", topicName)))
	}).Should(gomega.Succeed())

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetInactiveTopicCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var inactiveTopic utils.InactiveTopicPolicies
		err := json.Unmarshal(getOut.Bytes(), &inactiveTopic)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(inactiveTopic.DeleteWhileInactive).Should(gomega.Equal(false))
		g.Expect(inactiveTopic.MaxInactiveDurationSeconds).Should(gomega.Equal(0))
		g.Expect(inactiveTopic.InactiveTopicDeleteMode).Should(gomega.BeNil())
	}).Should(gomega.Succeed())
}
