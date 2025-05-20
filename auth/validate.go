package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

var (
	ErrMissingAuthHeader     = errors.New("authorization header is missing")
	ErrInvalidAuthScheme     = errors.New("authorization scheme must be Basic")
	ErrInvalidBase64Encoding = errors.New("authorization credentials are not valid base64")
	ErrInvalidCredentials    = errors.New("invalid authorization credentials format")
	ErrInvalidUserOrPassword = errors.New("invalid username or password")
)

func ValidateBasicAuth(r *http.Request, expectedUser, expectedSecret string) error {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return ErrMissingAuthHeader
	}

	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return ErrInvalidAuthScheme
	}

	encodedCredentials := strings.TrimPrefix(auth, prefix)
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return ErrInvalidBase64Encoding
	}

	credentials := string(decodedBytes)
	sepIndex := strings.IndexByte(credentials, ':')
	if sepIndex < 0 {
		return ErrInvalidCredentials
	}

	user := credentials[:sepIndex]
	secret := credentials[sepIndex+1:]

	if user != expectedUser || secret != expectedSecret {
		return ErrInvalidUserOrPassword
	}

	return nil
}
