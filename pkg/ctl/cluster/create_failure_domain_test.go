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

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFailureDomainCmdSuccess(t *testing.T) {
	args := []string{"create-failure-domain", "-b", "cluster-A", "standalone", "standalone-failure-domain"}
	_, execErr, NameErr, err := TestClusterCommands(createFailureDomainCmd, args)
	assert.Nil(t, execErr)
	assert.Nil(t, NameErr)
	assert.Nil(t, err)
}

func TestCreateFailureDomainCmdBrokerListError(t *testing.T) {
	args := []string{"create-failure-domain", "standalone", "standalone-failure-domain"}
	_, execErr, _, _ := TestClusterCommands(createFailureDomainCmd, args)
	assert.Equal(t, "broker list must be specified", execErr.Error())
}
