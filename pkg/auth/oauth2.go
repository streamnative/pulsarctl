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

package auth

import (
	"fmt"
	"net/http"

	"github.com/streamnative/pulsarctl/pkg/auth/oauth2"
	"github.com/streamnative/pulsarctl/pkg/auth/oauth2/cache"
	"github.com/streamnative/pulsarctl/pkg/auth/oauth2/store"
	xoauth2 "golang.org/x/oauth2"
	clock2 "k8s.io/utils/clock"
)

const (
	TypeClientCredential = "client_credentials"
	TypeDeviceCode       = "device_code"
)

type OAuth2Provider struct {
	isCommandLineTools bool
	clock              clock2.RealClock
	issuer             oauth2.Issuer
	store              store.Store
	source             cache.CachingTokenSource
	T                  http.RoundTripper
}

func NewAuthenticationOauth2(issuer oauth2.Issuer, store store.Store, types, keyFile string) (*OAuth2Provider, error) {
	p := &OAuth2Provider{
		clock:  clock2.RealClock{},
		issuer: issuer,
		store:  store,
	}
	err := p.Init(types, keyFile)
	return p, err
}

func NewAuthenticationOAuth2WithParams(
	commandLineTool bool,
	types,
	issueEndpoint,
	clientID,
	audience,
	keyFile string) (*OAuth2Provider, error) {

	issuer := oauth2.Issuer{
		IssuerEndpoint: issueEndpoint,
		ClientID:       clientID,
		Audience:       audience,
	}

	var st store.Store
	if commandLineTool {
		keyringStore, err := store.MakeKeyringStore()
		if err != nil {
			return nil, err
		}
		st = keyringStore
	} else {
		st = store.NewMemoryStore()
	}

	return NewAuthenticationOauth2(issuer, st, types, keyFile)
}

func (o *OAuth2Provider) Load() {

}

func (o *OAuth2Provider) Init(types, keyFile string) error {
	var grant *oauth2.AuthorizationGrant
	switch types {
	case TypeClientCredential:
		flow, err := oauth2.NewDefaultClientCredentialsFlow(o.issuer, keyFile)
		if err != nil {
			return err
		}
		if o.isCommandLineTools {
			grant, err = o.store.LoadGrant(o.issuer.Audience)
			if err != nil {
				return err
			}
		} else {
			grant, err := flow.Authorize()
			if err != nil {
				return err
			}
			err = o.store.SaveGrant(o.issuer.Audience, *grant)
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported authentication type: %s", types)
	}

	refresher, err := o.getRefresher(o.issuer, grant.Type)
	if err != nil {
		return err
	}

	source, err := cache.NewDefaultTokenCache(o.store, o.issuer.Audience, refresher)
	if err != nil {
		return err
	}
	o.source = source
	return nil
}

func (o *OAuth2Provider) RoundTrip(req *http.Request) (*http.Response, error) {
	return o.Transport().RoundTrip(req)
}

func (o *OAuth2Provider) Transport() http.RoundTripper {
	return &transport{
		source: o.source,
		wrapped: &xoauth2.Transport{
			Source: o.source,
			Base:   o.T,
		},
	}
}

func (o *OAuth2Provider) getRefresher(issuerData oauth2.Issuer,
	t oauth2.AuthorizationGrantType) (oauth2.AuthorizationGrantRefresher, error) {
	switch t {
	case oauth2.GrantTypeClientCredentials:
		return oauth2.NewDefaultClientCredentialsGrantRefresher(issuerData, o.clock)
	case oauth2.GrantTypeDeviceCode:
		return oauth2.NewDefaultDeviceAuthorizationGrantRefresher(issuerData, o.clock)
	default:
		return nil, store.ErrUnsupportedAuthData
	}
}

type transport struct {
	source  cache.CachingTokenSource
	wrapped *xoauth2.Transport
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return t.wrapped.Base.RoundTrip(req)
	}

	res, err := t.wrapped.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 401 {
		err := t.source.InvalidateToken()
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (t *transport) WrappedRoundTripper() http.RoundTripper { return t.wrapped.Base }
