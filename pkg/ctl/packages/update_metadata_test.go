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

func TestPackagesUpdateMetadata(t *testing.T) {
	randomVersion := test.RandomSuffix()
	packageURL := fmt.Sprintf("function://public/default/update-metadata@%s", randomVersion)
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

	args = []string{"update-metadata",
		packageURL,
		"--description", "update-description",
		"--contact", "pulsar@apache",
		"--properties", "foo=bar,abc=def",
	}

	output, execErr, err = TestPackagesCommands(putPackageMetadataCmd, args)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	exceptMsg := fmt.Sprintf("The metadata of the package '%s' updated successfully\n", packageURL)
	assert.True(t, strings.ContainsAny(output.String(), exceptMsg))

	args = []string{"get-metadata",
		packageURL,
	}

	output, execErr, err = TestPackagesCommands(getPackageMetadataCmd, args)
	failImmediatelyIfErrorNotNil(t, execErr, err)
	assert.Contains(t, output.String(), "update-description")
	assert.Contains(t, output.String(), "pulsar@apache")
	assert.Contains(t, output.String(), "foo")
	assert.Contains(t, output.String(), "bar")
	assert.Contains(t, output.String(), "abc")
	assert.Contains(t, output.String(), "def")
}
