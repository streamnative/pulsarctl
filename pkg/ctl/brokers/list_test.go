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

package brokers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBrokers(t *testing.T) {
	args := []string{"list", "standalone"}
	listOut, execErr, _, _ := TestBrokersCommands(getBrokerListCmd, args)
	assert.Nil(t, execErr)
	assert.True(t, strings.Contains(listOut.String(), "8080"))

	failArgs := []string{"list"}
	_, _, nameErr, _ := TestBrokersCommands(getBrokerListCmd, failArgs)
	assert.NotNil(t, nameErr)
	assert.Equal(t, "the cluster name is not specified or the cluster name is specified more than one",
		nameErr.Error())
}
