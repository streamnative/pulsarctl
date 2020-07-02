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
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/99designs/keyring"
	xoauth2 "golang.org/x/oauth2"
	ctloauth2 "github.com/streamnative/pulsarctl/pkg/auth/oauth2"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
	clock "k8s.io/utils/clock"
)

// ErrNoAuthenticationData is returned by the plugin when authentication data is not configured
var ErrNoAuthenticationData = errors.New("authentication data is not available")

var ErrUnsupportedAuthData = errors.New("authentication data is not usable")

// Factory acts as an authorization provider DefaultFactory
type Factory interface {
	// Register registers the auth plugin for use with client-go
	Register() error

	// UseClientCredentialsGrant configures the factory to use a client credentials grant
	UseClientCredentialsGrant(audience string, keyFile ctloauth2.KeyFile, token *xoauth2.Token) error

	// UseDeviceAuthorizationGrant configures the factory to use a device authorization grant
	UseDeviceAuthorizationGrant(audience string, token xoauth2.Token) error

	// WhoAmI returns the current user name (or an error if nobody is logged in)
	WhoAmI(audience string) (string, error)

	// Logout deletes all stored credentials
	Logout() error
}

type IssuerDataGetter func() (ctloauth2.Issuer, error)

const (
	// expiryDelta adjusts the token TTL to avoid using tokens which are almost expired
	expiryDelta = time.Duration(60) * time.Second
)

type DefaultFactory struct {
	kr               keyring.Keyring
	issuerDataGetter IssuerDataGetter
	clock            clock.Clock
}

func NewDefaultFactory(kr keyring.Keyring, issuerDataGetter IssuerDataGetter) (*DefaultFactory, error) {
	return &DefaultFactory{
		kr:               kr,
		issuerDataGetter: issuerDataGetter,
		clock:            clock.RealClock{},
	}, nil
}

var _ Factory = &DefaultFactory{}

func (f *DefaultFactory) Register() error {
	return restclient.RegisterAuthProviderPlugin(ProviderName, f.GetAuthProvider)
}

func (f *DefaultFactory) Logout() error {
	var err error
	keys, err := f.kr.Keys()
	if err != nil {
		return fmt.Errorf("unable to get information from the keyring: %v", err)
	}
	for _, key := range keys {
		err = f.kr.Remove(key)
	}
	if err != nil {
		return fmt.Errorf("unable to update the keyring: %v", err)
	}
	return nil
}

func (f *DefaultFactory) UseClientCredentialsGrant(audience string, keyFile ctloauth2.KeyFile, token *xoauth2.Token) error {
	item := Item{
		Audience: audience,
		UserName: keyFile.ClientEmail,
		Grant: Grant{
			Type:              GrantTypeClientCredentials,
			ClientCredentials: &keyFile,
			Token:             token,
		},
	}
	err := f.setItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (f *DefaultFactory) UseDeviceAuthorizationGrant(audience string, token xoauth2.Token) error {
	userName, err := ctloauth2.ExtractUserName(token)
	if err != nil {
		return err
	}
	item := Item{
		Audience: audience,
		UserName: userName,
		Grant: Grant{
			Type:  GrantTypeDeviceCode,
			Token: &token,
		},
	}
	err = f.setItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (f *DefaultFactory) WhoAmI(audience string) (string, error) {
	key := hashKeyringKey(audience)
	authItem, err := f.kr.GetMetadata(key)
	if err != nil {
		if err == keyring.ErrKeyNotFound {
			return "", ErrNoAuthenticationData
		}
		return "", fmt.Errorf("unable to get information from the keyring: %v", err)
	}
	return authItem.Label, nil
}

// GetAuthProvider implements the rest.Factory plugin interface
func (f *DefaultFactory) GetAuthProvider(_ string, _ map[string]string,
	_ rest.AuthProviderConfigPersister) (rest.AuthProvider, error) {
	issuerData, err := f.issuerDataGetter()
	if err != nil {
		return nil, err
	}

	item, err := f.getItem(issuerData.Audience)
	if err != nil {
		if err == keyring.ErrKeyNotFound {
			return &anonymousProvider{}, nil
		}
		return nil, err
	}

	var grant ctloauth2.AuthorizationGrant
	switch item.Grant.Type {
	case GrantTypeClientCredentials:
		if item.Grant.ClientCredentials == nil {
			return nil, ErrUnsupportedAuthData
		}
		grant, err = ctloauth2.NewDefaultClientCredentialsGrant(issuerData, *item.Grant.ClientCredentials, f.clock)
	default:
		return nil, ErrUnsupportedAuthData
	}
	if err != nil {
		return nil, err
	}

	cache := f.newTokenCache(item)
	provider := newDefaultProvider(grant, cache)
	return provider, nil
}

func (f *DefaultFactory) getItem(audience string) (Item, error) {
	key := hashKeyringKey(audience)
	i, err := f.kr.Get(key)
	if err != nil {
		return Item{}, err
	}
	var data Grant
	err = json.Unmarshal(i.Data, &data)
	if err != nil {
		// the auth data appears to be invalid
		return Item{}, ErrUnsupportedAuthData
	}
	return Item{
		Audience: audience,
		UserName: i.Label,
		Grant:    data,
	}, nil
}

func (f *DefaultFactory) setItem(item Item) error {
	key := hashKeyringKey(item.Audience)
	data, err := json.Marshal(item.Grant)
	if err != nil {
		return err
	}
	i := keyring.Item{
		Key:                         key,
		Data:                        data,
		Label:                       item.UserName,
		Description:                 "authorization grant",
		KeychainNotTrustApplication: false,
		KeychainNotSynchronizable:   false,
	}
	err = f.kr.Set(i)
	if err != nil {
		return fmt.Errorf("unable to update the keyring: %v", err)
	}
	return nil
}

// hashKeyringKey creates a safe key based on the given string
func hashKeyringKey(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// region TokenCache

func (f *DefaultFactory) newTokenCache(item Item) *tokenCache {
	return &tokenCache{
		f:    f,
		item: item,
	}
}

// tokenCache provides a write-through cache for the token associated with a specific audience
type tokenCache struct {
	f    *DefaultFactory
	item Item
	lock sync.Mutex
}

var _ TokenCache = &tokenCache{}

// GetToken returns a valid access token, if available.
func (t *tokenCache) GetToken() (*xoauth2.Token, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	// check token validity
	if t.item.Grant.Token == nil || t.item.Grant.Token.AccessToken == "" {
		return nil, nil
	}
	token := t.item.Grant.Token
	if !token.Expiry.IsZero() && t.f.clock.Now().After(token.Expiry.Round(0).Add(-expiryDelta)) {
		return nil, nil
	}

	return token, nil
}

// UpdateToken persists a new token, overwriting the existing refresh and access token.
func (t *tokenCache) UpdateToken(token *xoauth2.Token) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.item.Grant.Token = token
	err := t.f.setItem(t.item)
	return err
}

// InvalidateToken clears the access token (likely due to a response from the resource server).
// Note that the token within the persisted grant may contain a refresh token which should survive.
func (t *tokenCache) InvalidateToken() error {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.item.Grant.Token != nil && t.item.Grant.Token.AccessToken != "" {
		t.item.Grant.Token.AccessToken = ""
		err := t.f.setItem(t.item)
		if err != nil {
			return err
		}
	}
	return nil
}

// endregion
