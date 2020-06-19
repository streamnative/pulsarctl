// Copyright (c) 2020 StreamNative, Inc.. All Rights Reserved.

package oauth2

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// Challenge holds challenge and verification data needed for the PKCE flow
type Challenge struct {
	Code     string
	Verifier string
	Method   string
}

// Challenger is used to generate a new Challenge
type Challenger func() Challenge

// DefaultChallengeGenerator generates a default Challenge
func DefaultChallengeGenerator() Challenge {
	return generateChallenge(32)
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func generateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(b)
}

// generateChallenge generates a new Challenge of a specific length
func generateChallenge(length int) Challenge {
	c := Challenge{}

	c.Verifier = generateRandomString(length)

	csum := sha256.Sum256([]byte(c.Verifier))
	c.Code = base64.RawURLEncoding.EncodeToString(csum[:])
	c.Method = "S256"

	return c
}

// State is used to generate a new state string
type State func() string

// DefaultStateGenerator generates a default State
func DefaultStateGenerator() string {
	return generateRandomString(32)
}
