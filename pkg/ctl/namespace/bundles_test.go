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
	"encoding/json"
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBundlesCommandArgsError(t *testing.T) {
	args := []string{"bundles"}
	_, _, nameErr, _ := TestNamespaceCommands(getBundles, args)
	assert.Equal(t, "the namespace name is not specified or the namespace name is specified more than one",
		nameErr.Error())
}

func TestBundlesCommand(t *testing.T) {
	ns := "public/test-bundles-namespace"

	args := []string{"create", ns}
	_, execErr, _, _ := TestNamespaceCommands(createNs, args)
	assert.Nil(t, execErr)

	args = []string{"bundles", ns}
	out, execErr, _, _ := TestNamespaceCommands(getBundles, args)
	assert.Nil(t, execErr)

	var bundles utils.BundlesData
	err := json.Unmarshal(out.Bytes(), &bundles)
	assert.Nil(t, err)
	assert.True(t, bundles.NumBundles > 0)
	assert.True(t, len(bundles.Boundaries) >= 2)
}
