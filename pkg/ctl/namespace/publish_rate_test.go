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

	"github.com/streamnative/pulsar-admin-go/pkg/utils"
	"github.com/stretchr/testify/assert"

	"github.com/streamnative/pulsarctl/pkg/test"
)

func TestPublishRateCmd(t *testing.T) {
	ns := "public/test-publish-rate-ns" + test.RandomSuffix()

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-publish-rate", ns}
	_, execErr, _, _ = TestNamespaceCommands(GetPublishRateCmd, args)
	assert.Nil(t, execErr)

	args = []string{"set-publish-rate", ns}
	out, execErr, _, _ := TestNamespaceCommands(SetPublishRateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Success set the default message publish rate "+
			"of the namespace %s to %+v\n", ns,
			utils.PublishRate{
				PublishThrottlingRateInMsg:  -1,
				PublishThrottlingRateInByte: -1,
			}),
		out.String())

	args = []string{"get-publish-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetPublishRateCmd, args)
	assert.Nil(t, execErr)

	var rate utils.PublishRate

	err := json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, -1, rate.PublishThrottlingRateInMsg)
	assert.Equal(t, int64(-1), rate.PublishThrottlingRateInByte)

	args = []string{"set-publish-rate", "--msg-rate", "10", "--byte-rate", "10", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetPublishRateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Success set the default message publish rate "+
			"of the namespace %s to %+v\n", ns,
			utils.PublishRate{
				PublishThrottlingRateInMsg:  10,
				PublishThrottlingRateInByte: 10,
			}),
		out.String())

	args = []string{"get-publish-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetPublishRateCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 10, rate.PublishThrottlingRateInMsg)
	assert.Equal(t, int64(10), rate.PublishThrottlingRateInByte)
}

func TestSetPublishRateOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"set-publish-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetPublishRateCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "404")
}

func TestGetPublishRateOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"get-publish-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetPublishRateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
