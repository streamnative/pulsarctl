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

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func TestDispatchRate(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := "persistent://public/default/test-dispatch-rate-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-dispatch-rate", topicName, "--msg-dispatch-rate", "5", "--byte-dispatch-rate", "4",
		"--dispatch-rate-period", "3", "--relative-to-publish-rate"}
	g.Eventually(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(SetDispatchRateCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(
			gomega.Equal("Set message dispatch rate successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	getArgs := []string{"get-dispatch-rate", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetDispatchRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var dispatchRateData utils.DispatchRateData
		err := json.Unmarshal(getOut.Bytes(), &dispatchRateData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(dispatchRateData.DispatchThrottlingRateInMsg).Should(gomega.Equal(int64(5)))
		g.Expect(dispatchRateData.DispatchThrottlingRateInByte).Should(gomega.Equal(int64(4)))
		g.Expect(dispatchRateData.RatePeriodInSecond).Should(gomega.Equal(int64(3)))
		g.Expect(dispatchRateData.RelativeToPublishRate).Should(gomega.Equal(true))
	}).Should(gomega.Succeed())

	removeArgs := []string{"remove-dispatch-rate", topicName}
	g.Expect(func(g gomega.Gomega) {
		removeOut, execErr, _, _ := TestTopicCommands(RemoveDispatchRateCmd, removeArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(removeOut.String()).Should(
			gomega.Equal("Remove message dispatch rate successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetDispatchRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var dispatchRateData utils.DispatchRateData
		err := json.Unmarshal(getOut.Bytes(), &dispatchRateData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(dispatchRateData.DispatchThrottlingRateInMsg).Should(gomega.Equal(int64(0)))
		g.Expect(dispatchRateData.DispatchThrottlingRateInByte).Should(gomega.Equal(int64(0)))
		g.Expect(dispatchRateData.RatePeriodInSecond).Should(gomega.Equal(int64(0)))
		g.Expect(dispatchRateData.RelativeToPublishRate).Should(gomega.Equal(false))
	}).Should(gomega.Succeed())

	setArgs = []string{"set-dispatch-rate", topicName, "--msg-dispatch-rate", "5", "--byte-dispatch-rate", "4",
		"--dispatch-rate-period", "3"}
	g.Eventually(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(SetDispatchRateCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(
			gomega.Equal("Set message dispatch rate successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	getArgs = []string{"get-dispatch-rate", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetDispatchRateCmd, getArgs)
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
