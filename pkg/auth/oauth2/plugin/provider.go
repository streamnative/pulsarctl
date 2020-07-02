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
	"fmt"
	"net/http"
	"sync"

	ctloauth2 "github.com/streamnative/pulsarctl/pkg/auth/oauth2"
	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/util/net"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
)

// region anonymousProvider

type anonymousProvider struct {
}

var _ rest.AuthProvider = &anonymousProvider{}

func (a *anonymousProvider) Login() error {
	return nil
}

func (a *anonymousProvider) WrapTransport(rt http.RoundTripper) http.RoundTripper {
	return rt
}

//endregion

// region defaultProvider

// defaultProvider implements a rest.AuthProvider based on an OAuth 2.0 authorization grant.
// An authorization grant represents a permission given by a user to access a protected resource
// on their behalf.
type defaultProvider struct {
	grant ctloauth2.AuthorizationGrant
	cache TokenCache
	lock  sync.Mutex
}

// newDefaultProvider creates a new provider for the given authorization grant.
// A token cache is provided to read and write short-term access tokens which
// are obtained using the grant.
func newDefaultProvider(grant ctloauth2.AuthorizationGrant, cache TokenCache) *defaultProvider {
	return &defaultProvider{
		grant: grant,
		cache: cache,
	}
}

var _ rest.AuthProvider = &defaultProvider{}

func (p *defaultProvider) Login() error {
	return nil
}

func (p *defaultProvider) WrapTransport(rt http.RoundTripper) http.RoundTripper {
	return &conditionalTransport{&oauth2.Transport{Source: p, Base: rt}, p.cache}
}

var _ oauth2.TokenSource = &defaultProvider{}

func (p *defaultProvider) Token() (*oauth2.Token, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	// get current access token from the cache, if not expired, use it
	token, err := p.cache.GetToken()
	if err != nil {
		return nil, err
	}
	if token != nil {
		return token, nil
	}

	// obtain and cache a fresh access token
	token, err = p.grant.Refresh()
	if err != nil {
		return nil, fmt.Errorf("authentication failure: %v", err)
	}

	err = p.cache.UpdateToken(token)
	if err != nil {
		// TODO log rather than throw
		return nil, fmt.Errorf("unable to update the token cache: %v", err)
	}

	return token, nil
}

type conditionalTransport struct {
	oauthTransport *oauth2.Transport
	cache          TokenCache
}

var _ net.RoundTripperWrapper = &conditionalTransport{}

func (t *conditionalTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return t.oauthTransport.Base.RoundTrip(req)
	}

	res, err := t.oauthTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 401 {
		klog.V(4).Infof("The credentials that were supplied are invalid for the target audience")
		err := t.cache.InvalidateToken()
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (t *conditionalTransport) WrappedRoundTripper() http.RoundTripper { return t.oauthTransport.Base }
