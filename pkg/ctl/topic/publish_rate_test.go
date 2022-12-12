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

func TestPublishRate(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := "persistent://public/default/test-publish-rate-topic"
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-publish-rate", topicName, "--msg-publish-rate", "5", "--byte-publish-rate", "4"}
	setOut, execErr, _, _ := TestTopicCommands(SetPublishRateCmd, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(setOut.String()).Should(gomega.Equal("Set message publish rate successfully for [" + topicName + "]\n"))

	getArgs := []string{"get-publish-rate", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetPublishRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		var publishRateData utils.PublishRateData
		err := json.Unmarshal(getOut.Bytes(), &publishRateData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(publishRateData.PublishThrottlingRateInMsg).Should(gomega.Equal(int64(5)))
		g.Expect(publishRateData.PublishThrottlingRateInByte).Should(gomega.Equal(int64(4)))
	}).Should(gomega.Succeed())

	setArgs = []string{"remove-publish-rate", topicName}
	setOut, execErr, _, _ = TestTopicCommands(RemovePublishRateCmd, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(setOut.String()).Should(gomega.Equal("Remove message publish rate successfully for [" + topicName + "]\n"))

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetPublishRateCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		var publishRateData utils.PublishRateData
		err := json.Unmarshal(getOut.Bytes(), &publishRateData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(publishRateData.PublishThrottlingRateInMsg).Should(gomega.Equal(int64(0)))
		g.Expect(publishRateData.PublishThrottlingRateInByte).Should(gomega.Equal(int64(0)))
	}).Should(gomega.Succeed())
}
