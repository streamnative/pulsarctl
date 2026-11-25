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
	"testing"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/stretchr/testify/assert"
)

func withNamespaceAdminURLForTest(t *testing.T, webServiceURL string) {
	t.Helper()

	oldURL := cmdutils.PulsarCtlConfig.WebServiceURL
	cmdutils.PulsarCtlConfig.WebServiceURL = webServiceURL
	t.Cleanup(func() {
		cmdutils.PulsarCtlConfig.WebServiceURL = oldURL
	})
}

func TestNamespaceAdminEndpointEscapesNamespaceSegments(t *testing.T) {
	ns, err := utils.GetNamespaceName("public/test-namespace-properties")
	assert.NoError(t, err)

	assert.Equal(t,
		"/admin/v2/namespaces/public/test-namespace-properties/property/k%2F2",
		namespaceAdminEndpoint(*ns, "property", "k/2"))
}
