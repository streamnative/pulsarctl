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

import "golang.org/x/oauth2"

// TokenCache provides an interface to work with access tokens at runtime.
// The auth provider implementation is expected to call GetToken when
// preparing an outgoing request.  If the token is rejected by the resource server,
// the provider is expected to call InvalidateToken.  The provider then obtains
// a new token by some means and is expected to call UpdateToken.
type TokenCache interface {
	// GetToken returns the cached token, if exists and is valid
	GetToken() (*oauth2.Token, error)

	// UpdateToken caches a token
	UpdateToken(token *oauth2.Token) error

	// InvalidateToken clears any cached token
	InvalidateToken() error
}
