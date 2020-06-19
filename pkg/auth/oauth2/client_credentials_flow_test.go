// Copyright (c) 2020 StreamNative, Inc.. All Rights Reserved.

package oauth2

import (
	"errors"
	"time"

	"k8s.io/utils/clock"
	"k8s.io/utils/clock/testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MockClientCredentialsProvider struct {
	Called                  bool
	ClientCredentialsResult *KeyFile
	ReturnsError            error
}

func (m *MockClientCredentialsProvider) GetClientCredentials() (*KeyFile, error) {
	m.Called = true
	return m.ClientCredentialsResult, m.ReturnsError
}

var _ ClientCredentialsProvider = &MockClientCredentialsProvider{}

var _ = Describe("ClientCredentialsFlow", func() {
	issuer := Issuer{
		IssuerEndpoint: "http://issuer",
		ClientID:       "",
		Audience:       "test_audience",
	}

	Describe("Authorize", func() {

		var mockClock clock.Clock
		var mockCredsProvider *MockClientCredentialsProvider
		var mockTokenExchanger *MockTokenExchanger

		BeforeEach(func() {
			mockClock = testing.NewFakeClock(time.Unix(0, 0))

			mockCredsProvider = &MockClientCredentialsProvider{
				ClientCredentialsResult: &KeyFile{
					Type:         KeyFileTypeServiceAccount,
					ClientID:     "test_clientID",
					ClientSecret: "test_clientSecret",
					ClientEmail:  "test_clientEmail",
				},
			}

			expectedTokens := TokenResult{AccessToken: "accessToken", RefreshToken: "refreshToken", ExpiresIn: 1234}
			mockTokenExchanger = &MockTokenExchanger{
				ReturnsTokens: &expectedTokens,
			}
		})

		It("invokes TokenExchanger with credentials", func() {
			provider := NewClientCredentialsFlow(
				issuer,
				mockCredsProvider,
				mockTokenExchanger,
				mockClock,
			)

			_, _, err := provider.Authorize()
			Expect(err).ToNot(HaveOccurred())
			Expect(mockCredsProvider.Called).To(BeTrue())
			Expect(mockTokenExchanger.CalledWithRequest).To(Equal(&ClientCredentialsExchangeRequest{
				ClientID:     mockCredsProvider.ClientCredentialsResult.ClientID,
				ClientSecret: mockCredsProvider.ClientCredentialsResult.ClientSecret,
				Audience:     issuer.Audience,
			}))
		})

		It("returns TokensResult from TokenExchanger", func() {
			provider := NewClientCredentialsFlow(
				issuer,
				mockCredsProvider,
				mockTokenExchanger,
				mockClock,
			)

			_, token, err := provider.Authorize()
			Expect(err).ToNot(HaveOccurred())
			expected := convertToOAuth2Token(mockTokenExchanger.ReturnsTokens, mockClock)
			Expect(*token).To(Equal(expected))
		})

		It("returns an error if client credentials request errors", func() {
			mockCredsProvider.ReturnsError = errors.New("someerror")

			provider := NewClientCredentialsFlow(
				issuer,
				mockCredsProvider,
				mockTokenExchanger,
				mockClock,
			)

			_, _, err := provider.Authorize()
			Expect(err.Error()).To(Equal("could not get client credentials: someerror"))
		})

		It("returns an error if token exchanger errors", func() {
			mockTokenExchanger.ReturnsError = errors.New("someerror")
			mockTokenExchanger.ReturnsTokens = nil

			provider := NewClientCredentialsFlow(
				issuer,
				mockCredsProvider,
				mockTokenExchanger,
				mockClock,
			)

			_, _, err := provider.Authorize()
			Expect(err.Error()).To(Equal("authentication failed using client credentials: " +
				"could not exchange client credentials: someerror"))
		})
	})
})
