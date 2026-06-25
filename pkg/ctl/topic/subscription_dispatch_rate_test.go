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
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/onsi/gomega"
)

func TestSubscriptionDispatchRate(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := "persistent://public/default/test-subscription-dispatch-rate-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	getArgs := []string{"get-subscription-dispatch-rate", topicName}
	getOut, execErr, _, _ := TestTopicCommands(GetSubscriptionDispatchRateCmd, getArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(getOut.String()).Should(gomega.Equal("null"))

	setArgs := []string{"set-subscription-dispatch-rate", topicName, "--msg-dispatch-rate", "5",
		"--byte-dispatch-rate", "4", "--dispatch-rate-period", "3", "--relative-to-publish-rate"}
	setOut, execErr, _, _ := TestTopicCommands(SetSubscriptionDispatchRateCmd, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(setOut.String()).Should(gomega.Equal(
		"Set subscription message dispatch rate successfully for [" + topicName + "]\n"))

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetSubscriptionDispatchRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		var dispatchRateData utils.DispatchRateData
		err := json.Unmarshal(getOut.Bytes(), &dispatchRateData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(dispatchRateData.DispatchThrottlingRateInMsg).Should(gomega.Equal(int64(5)))
		g.Expect(dispatchRateData.DispatchThrottlingRateInByte).Should(gomega.Equal(int64(4)))
		g.Expect(dispatchRateData.RatePeriodInSecond).Should(gomega.Equal(int64(3)))
		g.Expect(dispatchRateData.RelativeToPublishRate).Should(gomega.Equal(true))
	}).Should(gomega.Succeed())

	setArgs = []string{"remove-subscription-dispatch-rate", topicName}
	setOut, execErr, _, _ = TestTopicCommands(RemoveSubscriptionDispatchRateCmd, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(setOut.String()).Should(gomega.Equal(
		"Remove subscription message dispatch rate successfully for [" + topicName + "]\n"))

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetSubscriptionDispatchRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(getOut.String()).Should(gomega.Equal("null"))
	}).Should(gomega.Succeed())

	setArgs = []string{"set-subscription-dispatch-rate", topicName, "--msg-dispatch-rate", "5",
		"--byte-dispatch-rate", "4", "--dispatch-rate-period", "3"}
	setOut, execErr, _, _ = TestTopicCommands(SetSubscriptionDispatchRateCmd, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(setOut.String()).Should(gomega.Equal(
		"Set subscription message dispatch rate successfully for [" + topicName + "]\n"))

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetSubscriptionDispatchRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		var dispatchRateData utils.DispatchRateData
		err := json.Unmarshal(getOut.Bytes(), &dispatchRateData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(dispatchRateData.DispatchThrottlingRateInMsg).Should(gomega.Equal(int64(5)))
		g.Expect(dispatchRateData.DispatchThrottlingRateInByte).Should(gomega.Equal(int64(4)))
		g.Expect(dispatchRateData.RatePeriodInSecond).Should(gomega.Equal(int64(3)))
		g.Expect(dispatchRateData.RelativeToPublishRate).Should(gomega.Equal(false))
	}).Should(gomega.Succeed())
}

func TestSetSubscriptionDispatchRateOnNonExistingTopic(t *testing.T) {
	g := gomega.NewWithT(t)

	args := []string{"set-subscription-dispatch-rate", "persistent://public/default/non-existing-subscription-dispatch-rate-topic"}
	_, execErr, _, _ := TestTopicCommands(SetSubscriptionDispatchRateCmd, args)
	g.Expect(execErr).ShouldNot(gomega.BeNil())
}

func TestGetSubscriptionDispatchRateOnNonExistingTopic(t *testing.T) {
	g := gomega.NewWithT(t)

	args := []string{"get-subscription-dispatch-rate", "persistent://public/default/non-existing-subscription-dispatch-rate-topic"}
	_, execErr, _, _ := TestTopicCommands(GetSubscriptionDispatchRateCmd, args)
	g.Expect(execErr).ShouldNot(gomega.BeNil())
}
