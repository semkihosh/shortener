package shortener

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"net/url"
	"strings"
)

// UniqueCode generates unique short code. Safe for concurrent use.
func UniqueCode() string {
	uid := ulid.Make()
	return strings.ToLower(uid.String())
}

// ShortURL creates short unique URL for the given base.
// If provided base is not invalid, it returns error.
func ShortURL(base string) (string, error) {
	baseUrl, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("base is not a valid URL: %w", err)
	}

	if baseUrl.Host == "" {
		return "", fmt.Errorf("host cannot be empty: %s", baseUrl.String())
	}

	return fmt.Sprintf("%s/%s", baseUrl.Host, UniqueCode()), nil
}
