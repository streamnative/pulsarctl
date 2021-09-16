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

package packages

import (
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestDeletePackages(t *testing.T) {
	randomVersion := test.RandomSuffix()
	packageURL := fmt.Sprintf("function://public/default/api-examples@%s", randomVersion)
	jarName := path.Join(ResourceDir(), "api-examples.jar")

	args := []string{"upload",
		packageURL,
		"--description", "examples",
		"--path", jarName,
	}

	_, execErr, err := TestPackagesCommands(uploadPackagesCmd, args)
	failImmediatelyIfErrorNotNil(t, execErr, err)

	deleteArgs := []string{"delete",
		packageURL,
	}

	_, execErr, err = TestPackagesCommands(deletePackagesCmd, deleteArgs)
	failImmediatelyIfErrorNotNil(t, execErr, err)
}

func TestDeletePackagesWithFailure(t *testing.T) {
	failureDeleteArgs := []string{"delete",
		"function://public/default/api-examples@non-exist",
	}

	_, execErrMsg, _ := TestPackagesCommands(deletePackagesCmd, failureDeleteArgs)
	assert.NotNil(t, execErrMsg)
	exceptMsg := "Package 'function://public/default/api-examples@non-exist' doesn't exist"
	assert.True(t, strings.ContainsAny(execErrMsg.Error(), exceptMsg))

}
