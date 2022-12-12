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

package sinks

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSink(t *testing.T) {
	// test the sink name not exist
	failureUpdateArgs := []string{"update",
		"--name", "not-exist",
	}
	_, err, _ := TestSinksCommands(updateSinksCmd, failureUpdateArgs)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "404"))
}
