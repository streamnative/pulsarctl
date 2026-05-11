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

func TestTopicAutoCreationCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	nsName := fmt.Sprintf("public/test-topic-auto-creation-ns-%s", test.RandomSuffix())
	createArgs := []string{"create", nsName}
	_, execErr, _, _ := TestNamespaceCommands(createNs, createArgs)
	g.Expect(execErr).Should(gomega.BeNil())

	setArgs := []string{"set-topic-auto-creation", nsName, "--type", "partitioned", "--partitions", "2"}
	out, execErr, _, _ := TestNamespaceCommands(setTopicAutoCreation, setArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(out.String()).Should(gomega.Equal(
		fmt.Sprintf("Set topic auto-creation config successfully for [%s]\n", nsName)))

	getArgs := []string{"get-topic-auto-creation", nsName}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ = TestNamespaceCommands(GetTopicAutoCreationCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())

		var cfg utils.TopicAutoCreationConfig
		err := json.Unmarshal(out.Bytes(), &cfg)
		g.Expect(err).Should(gomega.BeNil())
		g.Expect(cfg.Allow).Should(gomega.BeTrue())
		g.Expect(cfg.Type.String()).Should(gomega.Equal("partitioned"))
		g.Expect(cfg.Partitions).ShouldNot(gomega.BeNil())
		g.Expect(*cfg.Partitions).Should(gomega.Equal(2))
	}).Should(gomega.Succeed())

	removeArgs := []string{"remove-topic-auto-creation", nsName}
	out, execErr, _, _ = TestNamespaceCommands(removeTopicAutoCreation, removeArgs)
	g.Expect(execErr).Should(gomega.BeNil())
	g.Expect(out.String()).Should(gomega.Equal(
		fmt.Sprintf("Removed topic auto-creation config successfully for [%s]\n", nsName)))
}
