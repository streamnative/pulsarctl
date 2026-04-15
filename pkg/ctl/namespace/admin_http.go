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
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	pulsaradmin "github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/auth"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/rest"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func newNamespaceRESTClient() (*rest.Client, error) {
	cfg := config.Config(*cmdutils.PulsarCtlConfig)
	cfg.PulsarAPIVersion = config.V2

	authProvider, err := auth.GetAuthProvider(&cfg)
	if err != nil {
		return nil, err
	}

	return &rest.Client{
		ServiceURL:  cfg.WebServiceURL,
		VersionInfo: pulsaradmin.ReleaseVersion,
		HTTPClient: &http.Client{
			Timeout:   pulsaradmin.DefaultHTTPTimeOutDuration,
			Transport: authProvider,
		},
	}, nil
}

func namespaceAdminEndpoint(ns utils.NameSpaceName, parts ...string) string {
	escapedParts := make([]string, 0, len(parts)+2)
	for _, segment := range strings.Split(ns.String(), "/") {
		escapedParts = append(escapedParts, url.PathEscape(segment))
	}
	for _, part := range parts {
		escapedParts = append(escapedParts, url.PathEscape(part))
	}

	return path.Join(
		utils.MakeHTTPPath(config.V2.String(), "/namespaces"),
		path.Join(escapedParts...),
	)
}

func readOptionalStringResponse(resp *http.Response) (*string, error) {
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := strings.TrimSpace(string(body))
	if value == "" || value == "null" {
		return nil, nil
	}

	if unquoted, err := strconv.Unquote(value); err == nil {
		value = unquoted
	}

	return &value, nil
}

func removeNamespaceProperty(ns utils.NameSpaceName, key string) (*string, error) {
	client, err := newNamespaceRESTClient()
	if err != nil {
		return nil, err
	}

	endpoint := namespaceAdminEndpoint(ns, "property", key)
	resp, err := client.MakeRequest(http.MethodDelete, endpoint)
	if err != nil {
		return nil, err
	}

	return readOptionalStringResponse(resp)
}
