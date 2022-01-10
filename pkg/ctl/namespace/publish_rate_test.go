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

	"github.com/onsi/gomega"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
	"github.com/streamnative/pulsarctl/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestPublishRateCmd(t *testing.T) {
	g := gomega.NewWithT(t)

	ns := "public/test-publish-rate-ns" + test.RandomSuffix()

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"set-publish-rate", "--msg-rate", "10", "--byte-rate", "10", ns}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ := TestNamespaceCommands(SetPublishRateCmd, args)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(out).ShouldNot(gomega.BeNil())
		g.Expect(out.String()).Should(gomega.Equal(fmt.Sprintf("Success set the default message publish rate "+
			"of the namespace %s to %+v\n", ns,
			utils.PublishRate{
				PublishThrottlingRateInMsg:  10,
				PublishThrottlingRateInByte: 10,
			})))
	}).Should(gomega.Succeed())

	args = []string{"get-publish-rate", ns}
	g.Eventually(func(g gomega.Gomega) {
		out, execErr, _, _ := TestNamespaceCommands(GetPublishRateCmd, args)
		g.Expect(execErr).Should(gomega.BeNil())
		g.Expect(out).ShouldNot(gomega.BeNil())

		var rate utils.PublishRate
		err := json.Unmarshal(out.Bytes(), &rate)
		g.Expect(err).Should(gomega.BeNil())

		g.Expect(rate.PublishThrottlingRateInMsg).Should(gomega.Equal(10))
		g.Expect(rate.PublishThrottlingRateInByte).Should(gomega.Equal(int64(10)))
	}).Should(gomega.Succeed())
}

func TestSetPublishRateOnNonExistingNs(t *testing.T) {
	g := gomega.NewWithT(t)

	ns := "public/non-existing-ns"

	args := []string{"set-publish-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetPublishRateCmd, args)
	g.Expect(execErr).ShouldNot(gomega.BeNil())
	g.Expect(execErr.Error()).Should(gomega.ContainSubstring("404"))
}

func TestGetPublishRateOnNonExistingNs(t *testing.T) {
	g := gomega.NewWithT(t)

	ns := "public/non-existing-ns"

	args := []string{"get-publish-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetPublishRateCmd, args)
	g.Expect(execErr).ShouldNot(gomega.BeNil())
	g.Expect(execErr.Error()).Should(gomega.Equal("code: 404 reason: Namespace does not exist"))
}
