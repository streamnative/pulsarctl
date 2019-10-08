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

// +build tls

package cluster

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTLS(t *testing.T) {
	args := []string{"clusters", "add", "tls"}
	_, err := TestTLSHelp(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"clusters", "list"}
	out, err := TestTLSHelp(listClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "tls"))

	args = []string{"clusters", "delete", "tls"}
	_, err = TestTLSHelp(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = []string{"clusters", "list"}
	out, err = TestTLSHelp(listClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "tls"))
}
