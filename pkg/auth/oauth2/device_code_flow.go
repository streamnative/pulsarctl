// Copyright (c) 2020 StreamNative, Inc.. All Rights Reserved.

package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"k8s.io/utils/clock"

	"github.com/pkg/errors"
)

// DeviceCodeFlow takes care of the mechanics needed for getting an access
// token using the OAuth 2.0 "Device Code Flow"
type DeviceCodeFlow struct {
	issuerData   Issuer
	options      DeviceCodeFlowOptions
	codeProvider DeviceCodeProvider
	exchanger    DeviceTokenExchanger
	callback     DeviceCodeCallback
	clock        clock.Clock
}

// AuthorizationCodeProvider abstracts getting an authorization code
type DeviceCodeProvider interface {
	GetCode(additionalScopes ...string) (*DeviceCodeResult, error)
}

// DeviceTokenExchanger abstracts exchanging for tokens
type DeviceTokenExchanger interface {
	ExchangeDeviceCode(ctx context.Context, req DeviceCodeExchangeRequest) (*TokenResult, error)
	ExchangeRefreshToken(req RefreshTokenExchangeRequest) (*TokenResult, error)
}

type DeviceCodeCallback func(code *DeviceCodeResult) error

type DeviceCodeFlowOptions struct {
	AdditionalScopes []string
	AllowRefresh     bool
}

func NewDeviceCodeFlow(
	issuerData Issuer,
	options DeviceCodeFlowOptions,
	codeProvider DeviceCodeProvider,
	exchanger DeviceTokenExchanger,
	callback DeviceCodeCallback,
	clock clock.Clock) *DeviceCodeFlow {
	return &DeviceCodeFlow{
		options:      options,
		issuerData:   issuerData,
		codeProvider: codeProvider,
		exchanger:    exchanger,
		callback:     callback,
		clock:        clock,
	}
}

// NewDefaultDeviceCodeFlow provides an easy way to build up a default
// device code flow with all the correct configuration. If refresh tokens should
// be allowed pass in true for <allowRefresh>
func NewDefaultDeviceCodeFlow(issuerData Issuer, options DeviceCodeFlowOptions,
	callback DeviceCodeCallback) (*DeviceCodeFlow, error) {
	wellKnownEndpoints, err := GetOIDCWellKnownEndpointsFromIssuerURL(issuerData.IssuerEndpoint)
	if err != nil {
		return nil, err
	}

	codeProvider := NewLocalDeviceCodeProvider(
		issuerData,
		*wellKnownEndpoints,
		&http.Client{},
	)

	tokenRetriever := NewTokenRetriever(
		*wellKnownEndpoints,
		&http.Client{})

	return NewDeviceCodeFlow(
		issuerData,
		options,
		codeProvider,
		tokenRetriever,
		callback,
		clock.RealClock{}), nil
}

var _ Flow = &DeviceCodeFlow{}

func (p *DeviceCodeFlow) Authorize() (AuthorizationGrant, *oauth2.Token, error) {
	var additionalScopes []string
	additionalScopes = append(additionalScopes, p.options.AdditionalScopes...)
	if p.options.AllowRefresh {
		additionalScopes = append(additionalScopes, "offline_access")
	}

	codeResult, err := p.codeProvider.GetCode(additionalScopes...)
	if err != nil {
		return nil, nil, err
	}

	if p.callback != nil {
		err := p.callback(codeResult)
		if err != nil {
			return nil, nil, err
		}
	}

	exchangeRequest := DeviceCodeExchangeRequest{
		ClientID:     p.issuerData.ClientID,
		DeviceCode:   codeResult.DeviceCode,
		PollInterval: time.Duration(codeResult.Interval) * time.Second,
	}

	tr, err := p.exchanger.ExchangeDeviceCode(context.Background(), exchangeRequest)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not exchange code")
	}

	token := convertToOAuth2Token(tr, p.clock)
	grant := &DeviceCodeGrant{
		issuerData: p.issuerData,
		exchanger:  p.exchanger,
		token:      token,
		clock:      p.clock,
	}

	return grant, &token, nil
}

type DeviceCodeGrant struct {
	issuerData Issuer
	exchanger  DeviceTokenExchanger
	token      oauth2.Token
	clock      clock.Clock
}

// NewDefaultDeviceCodeGrant constructs a device authorization grant based on the result
// of the device authorization flow.
func NewDefaultDeviceCodeGrant(issuerData Issuer, token oauth2.Token, clock clock.Clock) (*DeviceCodeGrant, error) {
	wellKnownEndpoints, err := GetOIDCWellKnownEndpointsFromIssuerURL(issuerData.IssuerEndpoint)
	if err != nil {
		return nil, err
	}

	tokenRetriever := NewTokenRetriever(
		*wellKnownEndpoints,
		&http.Client{})

	return &DeviceCodeGrant{
		issuerData: issuerData,
		exchanger:  tokenRetriever,
		token:      token,
		clock:      clock,
	}, nil
}

var _ AuthorizationGrant = &DeviceCodeGrant{}

func (g *DeviceCodeGrant) Refresh() (*oauth2.Token, error) {

	if g.token.RefreshToken == "" {
		return nil, fmt.Errorf("the authorization grant has expired (no refresh token); please re-login")
	}

	exchangeRequest := RefreshTokenExchangeRequest{
		ClientID:     g.issuerData.ClientID,
		RefreshToken: g.token.RefreshToken,
	}

	tr, err := g.exchanger.ExchangeRefreshToken(exchangeRequest)
	if err != nil {
		return nil, err
	}

	// RFC 6749 Section 1.5 - token exchange MAY issue a new refresh token (otherwise the result is blank).
	// also see: https://tools.ietf.org/html/draft-ietf-oauth-security-topics-13#section-4.12
	if tr.RefreshToken == "" {
		tr.RefreshToken = g.token.RefreshToken
	}

	g.token = convertToOAuth2Token(tr, g.clock)

	return &g.token, nil
}
