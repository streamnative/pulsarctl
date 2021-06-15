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
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTLSWithJsonConfiguration(t *testing.T)  {
	basePath, err := os.Getwd()
	if err != nil {
		t.Fatal(err.Error())
	}
	args := []string{
		"--auth-plugin", "tls",
		"--auth-params",
		"{\"tlsCertFile\":\"" + basePath + "/test/auth/certs/client-cert.pem\"" +
			",\"tlsKeyFile\":\"" + basePath + "/test/auth/certs/client-key.pem\"}",
		"clusters", "list",
	}
	_, execErr, _ = TestTLSHelp(CreateClusterCmd, args)
	assert.NotNil(t, execErr)
}

func TestTLS(t *testing.T) {
	// There is no tls configuration, the execErr should not nil
	args := []string{"cluster", "add", "tls"}
	_, execErr, _ := TestTLSHelp(CreateClusterCmd, args)
	assert.NotNil(t, execErr)

	basePath, err := os.Getwd()
	if err != nil {
		t.Fatal(err.Error())
	}

	// The allowInsecureConnection is not specified. So the test should failed with 'doesn't contain any IP SANs'
	args = []string{
		"--auth-params",
		"{\"tlsCertFile\":\"" + basePath + "/test/auth/certs/client-cert.pem\"" +
			",\"tlsKeyFile\":\"" + basePath + "/test/auth/certs/client-key.pem\"}",
		"cluster", "add", "tls"}
	_, execErr, _ = TestTLSHelp(CreateClusterCmd, args)
	assert.NotNil(t, execErr)

	allArgs := []string{
		"--auth-params",
		"{\"tlsCertFile\":\"" + basePath + "/test/auth/certs/client-cert.pem\"" +
			",\"tlsKeyFile\":\"" + basePath + "/test/auth/certs/client-key.pem\"}",
		"--tls-trust-cert-path", basePath + "/test/auth/certs/cacert.pem",
		"--tls-allow-insecure",
	}
	args = append(allArgs, []string{"clusters", "add", "tls"}...)
	_, _, err = TestTLSHelp(CreateClusterCmd, args)
	assert.Nil(t, err)

	args = append(allArgs, []string{"clusters", "list"}...)
	out, _, err := TestTLSHelp(ListClustersCmd, args)
	assert.Nil(t, err)
	clusters := out.String()
	assert.True(t, strings.Contains(clusters, "tls"))

	args = append(allArgs, []string{"clusters", "delete", "tls"}...)
	_, _, err = TestTLSHelp(deleteClusterCmd, args)
	assert.Nil(t, err)

	args = append(allArgs, []string{"clusters", "list"}...)
	out, _, err = TestTLSHelp(ListClustersCmd, args)
	assert.Nil(t, err)
	clusters = out.String()
	assert.False(t, strings.Contains(clusters, "tls"))
}
