// Package auth provides authentication for App Store Connect API.
//
// See https://developer.apple.com/documentation/appstoreconnectapi for more
// information.
//
// While App Store Connect API doesn't use OAuth 2.0, this package relies on
// oauth2 package to provide a convenient way to refresh tokens.
package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

var (
	// ErrAPIKeyNotPem is returned when the API key is not PEM file.
	ErrAPIKeyNotPem = errors.New("api key is not pem file")
	// ErrAPIKeyNotECDSA is returned when the API key is not ECDSA key.
	ErrAPIKeyNotECDSA = errors.New("api key is not ecdsa key")
)

const (
	// Audience is the audience of the JWT token.
	Audience = "appstoreconnect-v1"
)

// Guarantees *TokenSource implements oauth2.TokenSource.
var _ oauth2.TokenSource = (*TokenSource)(nil)

// TokenSource implements oauth2.TokenSource.
type TokenSource struct {
	issuerID string
	keyID    string
	apiKey   *ecdsa.PrivateKey

	apiKeyBytes      []byte
	apiKeyPath       string
	scopes           []string
	expirationPeriod time.Duration
}

// Option is an option for TokenSource.
type Option func(*TokenSource)

// WithAPIKeyBytes specifies API key bytes.
// This option is exclusive with WithAPIKeyPath.
func WithAPIKeyBytes(b []byte) Option {
	return func(s *TokenSource) {
		s.apiKeyBytes = b
	}
}

// WithAPIKeyPath specifies API key path.
// This option is exclusive with WithAPIKeyBytes.
func WithAPIKeyPath(p string) Option {
	return func(s *TokenSource) {
		s.apiKeyPath = p
	}
}

// WithScopes specifies scopes.
func WithScopes(scopes ...string) Option {
	return func(s *TokenSource) {
		s.scopes = scopes
	}
}

// WithExpirationPeriod specifies expiration period.
func WithExpirationPeriod(d time.Duration) Option {
	return func(s *TokenSource) {
		s.expirationPeriod = d
	}
}

// NewTokenSource returns a new TokenSource.
func NewTokenSource(issuerID, keyID string, opts ...Option) (s *TokenSource, err error) {
	s = &TokenSource{
		issuerID:         issuerID,
		keyID:            keyID,
		expirationPeriod: 10 * time.Minute, //nolint:gomnd // Default value.
	}

	for _, opt := range opts {
		opt(s)
	}

	if s.apiKeyBytes == nil && s.apiKeyPath == "" {
		return nil, errors.New("api key is not specified")
	}

	if s.apiKeyBytes != nil && s.apiKeyPath != "" {
		return nil, errors.New("api key must be specified either by bytes or path")
	}

	key := s.apiKeyBytes

	if p := s.apiKeyPath; p != "" {
		key, err = os.ReadFile(p)
		if err != nil {
			return nil, err
		}
	}

	pk, err := privateKey(key)
	if err != nil {
		return nil, err
	}

	s.apiKey = pk

	return s, nil
}

func privateKey(pk []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(pk)
	if block == nil {
		return nil, ErrAPIKeyNotPem
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	if pk, ok := key.(*ecdsa.PrivateKey); ok {
		return pk, nil
	}

	return nil, ErrAPIKeyNotECDSA
}

var timeNow = time.Now

// Token returns a new token.
func (s *TokenSource) Token() (*oauth2.Token, error) {
	exp := timeNow().Add(s.expirationPeriod)

	t := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss":   s.issuerID,
		"iap":   jwt.NewNumericDate(timeNow()),
		"exp":   jwt.NewNumericDate(exp),
		"aud":   jwt.ClaimStrings{Audience},
		"scope": s.scopes,
	})

	t.Header["kid"] = s.keyID

	st, err := t.SignedString(s.apiKey)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: st,
		Expiry:      exp,
	}, nil
}
