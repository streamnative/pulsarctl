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

package cmdutils

import (
	"net/http"
	"net/url"
	"path"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/auth"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/rest"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
)

func NewPulsarRESTClient() (*rest.Client, error) {
	return NewPulsarRESTClientWithAPIVersion(config.V2)
}

func NewPulsarRESTClientWithAPIVersion(version config.APIVersion) (*rest.Client, error) {
	cfg := config.Config(*PulsarCtlConfig)
	cfg.PulsarAPIVersion = version
	if len(cfg.WebServiceURL) == 0 {
		cfg.WebServiceURL = admin.DefaultWebServiceURL
	}

	authProvider, err := auth.GetAuthProvider(&cfg)
	if err != nil {
		return nil, err
	}

	return &rest.Client{
		ServiceURL:  cfg.WebServiceURL,
		VersionInfo: admin.ReleaseVersion,
		HTTPClient: &http.Client{
			Timeout:   admin.DefaultHTTPTimeOutDuration,
			Transport: authProvider,
		},
	}, nil
}

func BuildAdminEndpoint(version config.APIVersion, componentPath string, parts ...string) string {
	escapedParts := make([]string, len(parts))
	for i, part := range parts {
		escapedParts[i] = url.PathEscape(part)
	}

	return path.Join(
		utils.MakeHTTPPath(version.String(), componentPath),
		path.Join(escapedParts...),
	)
}
