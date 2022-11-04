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

package oauth2

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

type ClientCredentials struct {
	IssuerUrl  string `json:"issuerUrl,omitempty"`
	Audience   string `json:"audience,omitempty"`
	Scope      string `json:"scope,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	ClientId   string `json:"clientId,omitempty"`
}

func Command(grouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"oauth2",
		"Operations about oauth2",
		"Login as a user or activate a service account using OAuth 2.0 authentication",
		"o")

	cmdutils.AddVerbCmd(grouping, resourceCmd, activateCmd)
	cmdutils.AddVerbCmd(grouping, resourceCmd, loginCmd)

	return resourceCmd
}

func applyClientCredentialsToConfig(config *cmdutils.ClusterConfig) (*cmdutils.ClusterConfig, error) {
	if config.AuthParams != "" && config.KeyFile == "" &&
		config.IssuerEndpoint == "" && config.Audience == "" && config.Scope == "" {
		var paramsJSON ClientCredentials
		err := json.Unmarshal([]byte(config.AuthParams), &paramsJSON)
		if err != nil {
			return config, err
		}
		config.IssuerEndpoint = paramsJSON.IssuerUrl
		config.Audience = paramsJSON.Audience
		config.Scope = paramsJSON.Scope
		config.KeyFile = paramsJSON.PrivateKey
	}
	return config, nil
}
