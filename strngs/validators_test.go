package strngs_test

import (
	"testing"

	"github.com/Nivl/go-types/strngs"
	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		uri         string
		shouldFail  bool
	}{
		{"ftp should fail", "ftp://google.com", shouldFail},
		{"file should fail", "file:///dev/urandom", shouldFail},
		{"http should work", "http://google.com", !shouldFail},
		{"https should work", "https://google.com", !shouldFail},
		{"uri should fail", "/dev/random", shouldFail},
		{"not a url should fail", "not a url", shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, !tc.shouldFail, strngs.IsValidURL(tc.uri))
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		email       string
		shouldFail  bool
	}{
		{"valid email should work", "email@domain.tld", !shouldFail},
		{"email with a + should work", "email+filter@domain.tld", !shouldFail},
		{"email without @ should fail", "emaildomain.tld", shouldFail},
		{"email without . should fail", "email@domaintld", shouldFail},
		{"email with nothing before @ should fail", "@domain.tld", shouldFail},
		{"email with nothing after . should fail", "email@domain.", shouldFail},
		{"email with nothing before . should fail", "email@.tld", shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, !tc.shouldFail, strngs.IsValidEmail(tc.email))
		})
	}
}

func TestIsValidUUID(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		input       string
		shouldFail  bool
	}{
		{"valid uuid", "676a5135-897e-40fc-b37a-95b6b0bcf09e", !shouldFail},
		{"invalid uuid", "not a uuid", shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, !tc.shouldFail, strngs.IsValidUUID(tc.input))
		})
	}
}

func TestIsValidSlug(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		input       string
		shouldFail  bool
	}{
		{"uuid should pass", "676a5135-897e-40fc-b37a-95b6b0bcf09e", !shouldFail},
		{"uppercase should fail", "ABCD", shouldFail},
		{"spaces should fail", "a b c d", shouldFail},
		{"underscore at the begining should fail", "-abcd", shouldFail},
		{"dash at the end should fail", "abcd-", shouldFail},
		{"# should fail", "ab#cd", shouldFail},
		{"? should fail", "ab?cd", shouldFail},
		{"/ should fail", "a/b/cd", shouldFail},
		{"\\ should fail", "a\\b\\cd", shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, !tc.shouldFail, strngs.IsValidSlug(tc.input))
		})
	}
}
