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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/onsi/gomega"

	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestInactiveTopicCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	nsName := fmt.Sprintf("public/test-inactive-topic-ns-%s", test.RandomSuffix())
	createArgs := []string{"create", nsName}
	_, execErr, _, _ := TestNamespaceCommands(createNs, createArgs)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-inactive-topic-policies", nsName,
		"-e=true",
		"-t", "1h",
		"-m", "delete_when_no_subscriptions"}
	out, execErr, _, _ := TestNamespaceCommands(SetInactiveTopicCmd, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(out.String()).Should(gomega.Equal(fmt.Sprintf("Set inactive topic policies successfully for [%s]",
		nsName)))

	getArgs := []string{"get-inactive-topic-policies", nsName}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ = TestNamespaceCommands(GetInactiveTopicCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		var inactiveTopic utils.InactiveTopicPolicies
		err := json.Unmarshal(out.Bytes(), &inactiveTopic)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(inactiveTopic.DeleteWhileInactive).Should(gomega.Equal(true))
		g.Expect(inactiveTopic.MaxInactiveDurationSeconds).Should(gomega.Equal(3600))
		g.Expect(inactiveTopic.InactiveTopicDeleteMode.String()).Should(gomega.Equal("delete_when_no_subscriptions"))
	}).Should(gomega.Succeed())

	removeArgs := []string{"remove-inactive-topic-policies", nsName}
	out, execErr, _, _ = TestNamespaceCommands(RemoveInactiveTopicCmd, removeArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(out.String()).Should(gomega.Equal(fmt.Sprintf("Remove inactive topic policies successfully from [%s]",
		nsName)))

	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ = TestNamespaceCommands(GetInactiveTopicCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		var inactiveTopic utils.InactiveTopicPolicies
		err := json.Unmarshal(out.Bytes(), &inactiveTopic)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(inactiveTopic.DeleteWhileInactive).Should(gomega.Equal(false))
		g.Expect(inactiveTopic.MaxInactiveDurationSeconds).Should(gomega.Equal(0))
		g.Expect(inactiveTopic.InactiveTopicDeleteMode).Should(gomega.BeNil())
	}).Should(gomega.Succeed())
}
