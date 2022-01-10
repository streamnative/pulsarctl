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
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestPersistence(t *testing.T) {
	g := gomega.NewWithT(t)

	topicName := "persistent://public/default/test-persistence-topic-" + test.RandomSuffix()
	args := []string{"create", topicName, "1"}
	_, execErr, _, _ := TestTopicCommands(CreateTopicCmd, args)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-persistence", topicName, "-e", "5", "-w", "4", "-a", "3", "-r", "2.2"}
	g.Eventually(func(g gomega.Gomega) {
		setOut, execErr, _, _ := TestTopicCommands(SetPersistenceCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(setOut.String()).Should(
			gomega.Equal("Set persistence successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	getArgs := []string{"get-persistence", topicName}
	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetPersistenceCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var persistenceData utils.PersistenceData
		err := json.Unmarshal(getOut.Bytes(), &persistenceData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(persistenceData.BookkeeperEnsemble).Should(gomega.Equal(int64(5)))
		g.Expect(persistenceData.BookkeeperWriteQuorum).Should(gomega.Equal(int64(4)))
		g.Expect(persistenceData.BookkeeperAckQuorum).Should(gomega.Equal(int64(3)))
		g.Expect(persistenceData.ManagedLedgerMaxMarkDeleteRate).Should(gomega.Equal(2.2))
	}).Should(gomega.Succeed())

	removeArgs := []string{"remove-persistence", topicName}
	g.Eventually(func(g gomega.Gomega) {
		removeOut, execErr, _, _ := TestTopicCommands(RemovePersistenceCmd, removeArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(removeOut.String()).Should(
			gomega.Equal("Remove persistence successfully for [" + topicName + "]\n"))
	}).Should(gomega.Succeed())

	g.Eventually(func(g gomega.Gomega) {
		getOut, execErr, _, _ := TestTopicCommands(GetPersistenceCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var persistenceData utils.PersistenceData
		err := json.Unmarshal(getOut.Bytes(), &persistenceData)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(persistenceData.BookkeeperEnsemble).Should(gomega.Equal(int64(0)))
		g.Expect(persistenceData.BookkeeperWriteQuorum).Should(gomega.Equal(int64(0)))
		g.Expect(persistenceData.BookkeeperAckQuorum).Should(gomega.Equal(int64(0)))
		g.Expect(persistenceData.ManagedLedgerMaxMarkDeleteRate).Should(gomega.Equal(float64(0)))
	}).Should(gomega.Succeed())

	// test value
	setArgs = []string{"set-persistence", topicName, "-e", "1", "-w", "4", "-a", "3", "-r", "2.2"}
	g.Eventually(func(g gomega.Gomega) {
		_, execErr, _, _ = TestTopicCommands(SetPersistenceCmd, setArgs)
		g.Expect(execErr).ShouldNot(gomega.BeNil())
		g.Expect(execErr.Error()).Should(
			gomega.Equal("code: 400 reason: Bookkeeper Ensemble (1) >= WriteQuorum (4) >= AckQuoru (3)"))
	}).Should(gomega.Succeed())
}
