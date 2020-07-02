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

package plugin

import (
	ctloauth2 "github.com/streamnative/pulsarctl/pkg/auth/oauth2"
	"golang.org/x/oauth2"
)

const (
	ProviderName = "streamnative"
)

type GrantType string

const (
	// GrantTypeClientCredentials represents a client credentials grant
	GrantTypeClientCredentials GrantType = "client_credentials"

	// GrantTypeDeviceCode represents a device authorization grant
	GrantTypeDeviceCode GrantType = "device_code"
)

// Item represents an item stored in the keyring
type Item struct {
	Audience string
	UserName string
	Grant    Grant
}

// grant represents a persisted authorization grant
type Grant struct {
	// Type describes the type of authorization grant represented by this structure
	Type GrantType `json:"type"`

	// ClientCredentials is credentials data for the client credentials grant type
	ClientCredentials *ctloauth2.KeyFile `json:"client_credentials,omitempty"`

	// Token contains an access token in the client credentials grant type,
	// and a refresh token in the device authorization grant type
	Token *oauth2.Token `json:"token,omitempty"`
}
