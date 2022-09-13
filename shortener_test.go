package shortener_test

import (
	"github.com/semkihosh/shortener"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestUniqueCode(t *testing.T) {
	t.Parallel()

	code := shortener.UniqueCode()
	assert.NotEmpty(t, code)
}

func TestShortURL(t *testing.T) {
	cases := []struct {
		name    string
		base    string
		wantUrl *regexp.Regexp
		wantErr bool
	}{
		{"it can generate url", "https://example.com", regexp.MustCompile(`example.com/\w+`), false},
		{"it require scheme", "example.com", &regexp.Regexp{}, true},
		{"it will fail with invalid base", "https://%%2.com", &regexp.Regexp{}, true},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := shortener.ShortURL(tc.base)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Truef(t, tc.wantUrl.MatchString(got), "`%s` do not match regex `%s`", got, tc.wantUrl.String())
			}
		})
	}
}
