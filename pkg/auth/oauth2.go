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
	"os"
	"path"
	"path/filepath"

	"github.com/99designs/keyring"
	"github.com/pkg/errors"
	"github.com/streamnative/pulsarctl/pkg/auth/oauth2"
	oauth2os "github.com/streamnative/pulsarctl/pkg/auth/oauth2/os"
	"github.com/streamnative/pulsarctl/pkg/auth/oauth2/plugin"
)

const (
	serviceName  = "pulsar"
	keychainName = "pulsarctl"
)

type OAuth2Configuration struct {
	SkipOpen  bool
	NoRefresh bool
}

type OAuth2Provider struct {
	conf    *OAuth2Configuration
	issuer  *oauth2.Issuer
	keyFile string
	T       http.RoundTripper
	plugin.Factory
	osInteractor oauth2os.Interactor
}

func NewAuthenticationOAuth2(
	conf *OAuth2Configuration,
	issuerEndpoint,
	clientID,
	audience,
	keyFile string,
	transport http.RoundTripper) (*OAuth2Provider, error) {

	issuer := &oauth2.Issuer{
		IssuerEndpoint: issuerEndpoint,
		ClientID:       clientID,
		Audience:       audience,
	}

	kr, err := makeKeyring()
	if err != nil {
		return nil, err
	}
	factory, err := plugin.NewDefaultFactory(kr, func() (issuer oauth2.Issuer, err error) {
		return issuer, nil
	})
	if err != nil {
		return nil, err
	}

	return &OAuth2Provider{
		conf:         conf,
		issuer:       issuer,
		keyFile:      keyFile,
		T:            transport,
		Factory:      factory,
		osInteractor: &oauth2os.DefaultInteractor{},
	}, nil
}

func (o *OAuth2Provider) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := o.getToken()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return o.T.RoundTrip(req)
}

func (o *OAuth2Provider) Transport() http.RoundTripper {
	return o.T
}

func (o *OAuth2Provider) getToken() (string, error) {
	if o.keyFile != "" {
		return o.getTokenWithClientCredentialsFlow()
	} else {
		return o.getTokenWithDeviceCodeFlow()
	}
}

func makeKeyring() (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{
		ServiceName:              serviceName,
		KeychainName:             keychainName,
		KeychainTrustApplication: true,
		AllowedBackends:          keyring.AvailableBackends(),
		FileDir:                  filepath.Join(credentialDir(), "credentials"),
		FilePasswordFunc:         keyringPrompt,
	})
}

func keyringPrompt(prompt string) (string, error) {
	return "", nil
}

func credentialDir() string {
	return path.Join(os.Getenv("HOME"), ".config/pulsar")
}

func (o *OAuth2Provider) getTokenWithClientCredentialsFlow() (string, error) {
	flow, err := oauth2.NewDefaultClientCredentialsFlow(*o.issuer, o.keyFile)
	if err != nil {
		return "", err
	}
	grant, token, err := flow.Authorize()
	if err != nil {
		return "", err
	}

	keyFile := grant.(*oauth2.ClientCredentialsGrant).KeyFile
	if err = o.UseClientCredentialsGrant(o.issuer.Audience, keyFile, token); err != nil {
		return "", errors.Wrap(err, "unable to store the authorization data")
	}
	return token.AccessToken, nil
}

func (o *OAuth2Provider) getTokenWithDeviceCodeFlow() (string, error) {
	flow, err := o.newDeviceCodeFlow(o.issuer)
	if err != nil {
		return "", err
	}
	_, token, err := flow.Authorize()
	if err != nil {
		return "", err
	}

	if err = o.UseDeviceAuthorizationGrant(o.issuer.Audience, *token); err != nil {
		return "", errors.Wrap(err, "unable to store the authorization data")
	}
	return token.AccessToken, nil
}

func (o *OAuth2Provider) newDeviceCodeFlow(issuer *oauth2.Issuer) (*oauth2.DeviceCodeFlow, error) {
	options := oauth2.DeviceCodeFlowOptions{
		AdditionalScopes: []string{"admin"},
		AllowRefresh:     false,
	}
	flow, err := oauth2.NewDefaultDeviceCodeFlow(*issuer, options, o.Prompt)
	if err != nil {
		return nil, err
	}
	return flow, nil
}

func (o *OAuth2Provider) Prompt(code *oauth2.DeviceCodeResult) error {
	if !o.conf.SkipOpen {
		err := o.osInteractor.OpenURL(code.VerificationURIComplete)
		if err == nil {
			fmt.Printf(`We've launched your web browser to complete the login process.
Verification code: %s

Waiting for login to complete...
`, code.UserCode)
			return nil
		}
	}
	fmt.Printf(`Please follow these steps to complete the login procedure:
1. Using your web browser, go to: %s
2. Enter the following code: %s

Waiting for login to complete...
`, code.VerificationURI, code.UserCode)

	return nil

}
