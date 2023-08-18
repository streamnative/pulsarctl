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
	"github.com/stretchr/testify/assert"
)

func TestSubscribeRateCmd(t *testing.T) {
	ns := "public/test-subscribe-rate-ns"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"get-subscribe-rate", ns}
	out, execErr, _, _ := TestNamespaceCommands(GetSubscribeRateCmd, args)
	assert.Nil(t, execErr)

	var rate utils.SubscribeRate
	err := json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, rate.SubscribeThrottlingRatePerConsumer)
	assert.Equal(t, 0, rate.RatePeriodInSecond)

	args = []string{"set-subscribe-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSubscribeRateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Success set the default subscribe rate "+
			"of the namespace %s to %+v", ns,
			utils.SubscribeRate{
				SubscribeThrottlingRatePerConsumer: -1,
				RatePeriodInSecond:                 30,
			}),
		out.String())

	args = []string{"get-subscribe-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSubscribeRateCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, -1, rate.SubscribeThrottlingRatePerConsumer)
	assert.Equal(t, 30, rate.RatePeriodInSecond)

	args = []string{"set-subscribe-rate", "--subscribe-rate", "10", "--period", "10", ns}
	out, execErr, _, _ = TestNamespaceCommands(SetSubscribeRateCmd, args)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("Success set the default subscribe rate of the namespace %s to %+v", ns,
			utils.SubscribeRate{
				SubscribeThrottlingRatePerConsumer: 10,
				RatePeriodInSecond:                 10,
			}),
		out.String())

	args = []string{"get-subscribe-rate", ns}
	out, execErr, _, _ = TestNamespaceCommands(GetSubscribeRateCmd, args)
	assert.Nil(t, execErr)
	err = json.Unmarshal(out.Bytes(), &rate)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 10, rate.SubscribeThrottlingRatePerConsumer)
	assert.Equal(t, 10, rate.RatePeriodInSecond)
}

func TestSetSubscribeRateOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"set-subscribe-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(SetSubscribeRateCmd, args)
	assert.NotNil(t, execErr)
	assert.Contains(t, execErr.Error(), "404")
}

func TestGetSubscribeRateOnNonExistingNs(t *testing.T) {
	ns := "public/non-existing-ns"

	args := []string{"get-subscribe-rate", ns}
	_, execErr, _, _ := TestNamespaceCommands(GetSubscribeRateCmd, args)
	assert.NotNil(t, execErr)
	assert.Equal(t, "code: 404 reason: Namespace does not exist", execErr.Error())
}
