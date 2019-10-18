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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteClusterCmd(t *testing.T) {
	args := []string{"add", "delete-test"}
	_, _, _, err := TestClusterCommands(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, err := TestClusterCommands(ListClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "delete-test"))

	args = []string{"delete", "delete-test"}
	_, _, _, err = TestClusterCommands(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"list"}
	out, _, _, err = TestClusterCommands(ListClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "delete-test"))
}
