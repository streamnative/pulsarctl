// Copyright (c) 2020 StreamNative, Inc.. All Rights Reserved.

package oauth2

import (
	"errors"
	"time"

	"golang.org/x/oauth2"
	"k8s.io/utils/clock"
	"k8s.io/utils/clock/testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockDeviceCodeProvider struct {
	Called                     bool
	CalledWithAdditionalScopes []string
	DeviceCodeResult           *DeviceCodeResult
	ReturnsError               error
}

func (cp *MockDeviceCodeProvider) GetCode(additionalScopes ...string) (*DeviceCodeResult, error) {
	cp.Called = true
	cp.CalledWithAdditionalScopes = additionalScopes
	return cp.DeviceCodeResult, cp.ReturnsError
}

type MockDeviceCodeCallback struct {
	Called           bool
	DeviceCodeResult *DeviceCodeResult
	ReturnsError     error
}

func (c *MockDeviceCodeCallback) Callback(code *DeviceCodeResult) error {
	c.Called = true
	c.DeviceCodeResult = code
	if c.ReturnsError != nil {
		return c.ReturnsError
	}
	return nil
}

var _ = Describe("DeviceCodeFlow", func() {
	issuer := Issuer{
		IssuerEndpoint: "http://issuer",
		ClientID:       "test_clientID",
		Audience:       "test_audience",
	}

	Describe("Authorize", func() {

		var mockClock clock.Clock
		var mockCodeProvider *MockDeviceCodeProvider
		var mockTokenExchanger *MockTokenExchanger
		var mockCallback *MockDeviceCodeCallback
		var flow *DeviceCodeFlow

		BeforeEach(func() {
			mockClock = testing.NewFakeClock(time.Unix(0, 0))

			mockCodeProvider = &MockDeviceCodeProvider{
				DeviceCodeResult: &DeviceCodeResult{
					DeviceCode:              "test_deviceCode",
					UserCode:                "test_userCode",
					VerificationURI:         "http://verification_uri",
					VerificationURIComplete: "http://verification_uri_complete",
					ExpiresIn:               10,
					Interval:                5,
				},
			}

			expectedTokens := TokenResult{AccessToken: "accessToken", RefreshToken: "refreshToken", ExpiresIn: 1234}
			mockTokenExchanger = &MockTokenExchanger{
				ReturnsTokens: &expectedTokens,
			}

			mockCallback = &MockDeviceCodeCallback{}

			opts := DeviceCodeFlowOptions{
				AdditionalScopes: nil,
				AllowRefresh:     true,
			}
			flow = NewDeviceCodeFlow(
				issuer,
				opts,
				mockCodeProvider,
				mockTokenExchanger,
				mockCallback.Callback,
				mockClock,
			)
		})

		It("invokes DeviceCodeProvider", func() {
			_, _, _ = flow.Authorize()
			Expect(mockCodeProvider.Called).To(BeTrue())
			Expect(mockCodeProvider.CalledWithAdditionalScopes).To(ContainElement("offline_access"))
		})

		It("invokes callback with returned code", func() {
			_, _, _ = flow.Authorize()
			Expect(mockCallback.Called).To(BeTrue())
			Expect(mockCallback.DeviceCodeResult).To(Equal(mockCodeProvider.DeviceCodeResult))
		})

		It("invokes TokenExchanger with returned code", func() {
			_, _, _ = flow.Authorize()
			Expect(mockTokenExchanger.CalledWithRequest).To(Equal(&DeviceCodeExchangeRequest{
				ClientID:     issuer.ClientID,
				PollInterval: time.Duration(5) * time.Second,
				DeviceCode:   "test_deviceCode",
			}))
		})

		It("returns TokensResult from TokenExchanger", func() {
			_, tokens, _ := flow.Authorize()
			expected := convertToOAuth2Token(mockTokenExchanger.ReturnsTokens, mockClock)
			Expect(*tokens).To(Equal(expected))
		})

		It("returns an authorization grant", func() {
			grant, _, _ := flow.Authorize()
			Expect(grant).ToNot(BeNil())
		})
	})

	Describe("Refresh", func() {
		var mockClock clock.Clock
		var mockTokenExchanger *MockTokenExchanger
		var grant *DeviceCodeGrant

		BeforeEach(func() {
			mockClock = testing.NewFakeClock(time.Unix(0, 0))

			mockTokenExchanger = &MockTokenExchanger{}

			token := oauth2.Token{AccessToken: "gat", RefreshToken: "grt", Expiry: time.Unix(1, 0)}
			grant = &DeviceCodeGrant{
				issuerData: issuer,
				exchanger:  mockTokenExchanger,
				token:      token,
				clock:      mockClock,
			}
		})

		It("invokes the token exchanger", func() {
			mockTokenExchanger.ReturnsTokens = &TokenResult{
				AccessToken: "new token",
			}

			_, _ = grant.Refresh()
			Expect(mockTokenExchanger.RefreshCalledWithRequest).To(Equal(&RefreshTokenExchangeRequest{
				ClientID:     issuer.ClientID,
				RefreshToken: "grt",
			}))
		})

		It("returns the refreshed access token from the TokenExchanger", func() {
			mockTokenExchanger.ReturnsTokens = &TokenResult{
				AccessToken: "new token",
			}

			token, _ := grant.Refresh()
			Expect(token.AccessToken).To(Equal(mockTokenExchanger.ReturnsTokens.AccessToken))
		})

		It("preserves the existing refresh token from the TokenExchanger", func() {
			mockTokenExchanger.ReturnsTokens = &TokenResult{
				AccessToken: "new token",
			}

			token, _ := grant.Refresh()
			Expect(token.RefreshToken).To(Equal("grt"))
		})

		It("returns the refreshed refresh token from the TokenExchanger", func() {
			mockTokenExchanger.ReturnsTokens = &TokenResult{
				AccessToken:  "new token",
				RefreshToken: "new token",
			}

			token, _ := grant.Refresh()
			Expect(token.RefreshToken).To(Equal("new token"))
		})

		It("returns a meaningful expiration time", func() {
			mockTokenExchanger.ReturnsTokens = &TokenResult{
				AccessToken: "new token",
				ExpiresIn:   60,
			}

			token, _ := grant.Refresh()
			Expect(token.Expiry).To(Equal(mockClock.Now().Add(time.Duration(60) * time.Second)))
		})

		It("returns an error when TokenExchanger does", func() {
			mockTokenExchanger.ReturnsError = errors.New("someerror")

			token, err := grant.Refresh()
			Expect(token).To(BeNil())
			Expect(err.Error()).To(Equal("someerror"))
		})
	})
})
