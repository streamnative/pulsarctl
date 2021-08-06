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

package sources

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartAndStopSource(t *testing.T) {
	// test failure case
	failureStartArgs := []string{"start",
		"--name", "not-exist",
	}
	_, err, _ := TestSourcesCommands(startSourcesCmd, failureStartArgs)
	assert.NotNil(t, err)
	failMsg := "Source not-exist doesn't exist"
	assert.True(t, strings.ContainsAny(err.Error(), failMsg))

	// test the --name args not exist
	notExistNameOrFqfnArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
	}
	_, err, _ = TestSourcesCommands(startSourcesCmd, notExistNameOrFqfnArgs)
	assert.NotNil(t, err)
	failNameMsg := "You must specify a name for the source"
	assert.True(t, strings.ContainsAny(err.Error(), failNameMsg))

	// test the instance id not exist
	notExistInstanceIDArgs := []string{"start",
		"--tenant", "public",
		"--namespace", "default",
		"--name", "test-source-start",
		"--instance-id", "12345678",
	}
	_, err, _ = TestSourcesCommands(startSourcesCmd, notExistInstanceIDArgs)
	assert.NotNil(t, err)
	failInstanceIDMsg := "Operation not permitted"
	assert.True(t, strings.ContainsAny(err.Error(), failInstanceIDMsg))
}
