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
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestIsAllowAutoUpdateSchemaCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	ns := "public/test-is-allow-auto-update-schema-" + test.RandomSuffix()
	createArgs := []string{"create", ns}
	g.Eventually(func(g gomega.Gomega) {
		_, execErr, _, _ := TestNamespaceCommands(createNs, createArgs)
		g.Expect(execErr).Should(gomega.BeNil())
	}).Should(gomega.Succeed())

	setArgs := []string{"set-is-allow-auto-update-schema", "--disable", ns}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ := TestNamespaceCommands(SetIsAllowAutoUpdateSchemaCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(out.String()).Should(gomega.Equal(
			fmt.Sprintf("Successfully disable auto update schema on a namespace %s\n", ns)))
	}).Should(gomega.Succeed())

	getArgs := []string{"get-is-allow-auto-update-schema", ns}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ := TestNamespaceCommands(GetIsAllowAutoUpdateSchemaCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(out.String()).Should(gomega.Equal("false\n"))
	}).Should(gomega.Succeed())

	setArgs = []string{"set-is-allow-auto-update-schema", "--enable", ns}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ := TestNamespaceCommands(SetIsAllowAutoUpdateSchemaCmd, setArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(out.String()).Should(gomega.Equal(
			fmt.Sprintf("Successfully enable auto update schema on a namespace %s\n", ns)))
	}).Should(gomega.Succeed())

	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ := TestNamespaceCommands(GetIsAllowAutoUpdateSchemaCmd, getArgs)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(out.String()).Should(gomega.Equal("true\n"))
	}).Should(gomega.Succeed())
}
