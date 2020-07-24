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

	"github.com/stretchr/testify/assert"
)

func TestSetRetentionCmd(t *testing.T) {
	ns := "public/test-retention-3"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"set-retention", ns, "--size", "1G", "--time", "-1"}
	out, execErr, _, _ := TestNamespaceCommands(setRetention, args)
	if execErr != nil {
		assert.FailNow(t, "set retention failed: %s", execErr.Error())
	}

	assert.Equal(t, fmt.Sprintf("Set retention successfully for [%s]." +
		" The retention policy is: time = %d min, size = %d MB\n",ns, -1, 1024), out.String())
}
