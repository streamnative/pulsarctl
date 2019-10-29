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

	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionDispatchRateCmd(t *testing.T) {
	ns := "public/test-subscription-dispatch-rate-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-subscription-dispatch-rate", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSubscriptionDispatchRateCmd, args)
	assert.Nil(t, execErr)

	var rate pulsar.DispatchRate
	err := json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, rate.DispatchThrottlingRateInMsg)
	assert.Equal(t, int64(0), rate.DispatchThrottlingRateInByte)
	assert.Equal(t, 1, rate.RatePeriodInSecond)

	args = []string{"set-subscription-dispatch-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSubscriptionDispatchRateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Success set the default subscription message dispatch rate "+
			"of the namespace %s to %+v\n", ns,
			pulsar.DispatchRate{
				DispatchThrottlingRateInMsg:  -1,
				DispatchThrottlingRateInByte: -1,
				RatePeriodInSecond:           1,
			}),
		out.String())

	args = []string{"get-subscription-dispatch-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSubscriptionDispatchRateCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, -1, rate.DispatchThrottlingRateInMsg)
	assert.Equal(t, int64(-1), rate.DispatchThrottlingRateInByte)
	assert.Equal(t, 1, rate.RatePeriodInSecond)

	args = []string{"set-subscription-dispatch-rate", "--msg-rate", "10", "--byte-rate", "10", "--period", "10", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSubscriptionDispatchRateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Success set the default subscription message dispatch rate "+
			"of the namespace %s to %+v\n", ns,
			pulsar.DispatchRate{
				DispatchThrottlingRateInByte: 10,
				DispatchThrottlingRateInMsg:  10,
				RatePeriodInSecond:           10,
			}),
		out.String())

	args = []string{"get-subscription-dispatch-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSubscriptionDispatchRateCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 10, rate.DispatchThrottlingRateInMsg)
	assert.Equal(t, int64(10), rate.DispatchThrottlingRateInByte)
	assert.Equal(t, 10, rate.RatePeriodInSecond)
}

func TestSetSubscriptionDispatchRateOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"set-subscription-dispatch-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSubscriptionDispatchRateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace public/non-existing-ns does not exist", execErr.Error())
}

func TestGetSubscriptionDispatchRateOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"get-subscription-dispatch-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetSubscriptionDispatchRateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
