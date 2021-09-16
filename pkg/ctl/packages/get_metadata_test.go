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
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestPackagesGetMetadata(t *testing.T) {
	randomVersion := test.RandomSuffix()
	packageURL := fmt.Sprintf("function://public/default/get-metadata@%s", randomVersion)
	jarName := path.Join(ResourceDir(), "api-examples.jar")

	args := []string{"upload",
		packageURL,
		"--description", randomVersion,
		"--path", jarName,
	}

	output, execErr, err := TestPackagesCommands(uploadPackagesCmd, args)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Equal(t, output.String(),
		fmt.Sprintf("The package '%s' uploaded from path '%s' successfully\n", packageURL, jarName))

	args = []string{"get-metadata",
		packageURL,
	}

	output, execErr, err = TestPackagesCommands(getPackageMetadataCmd, args)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Contains(t, output.String(), "description")
}
