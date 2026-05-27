package roundtrippers

import (
	"net/http"
)

type bearerTokenFileRoundTripper struct {
	tokenFile string
	rt        http.RoundTripper
}

// RoundTrip implements http.RoundTripper's interface
func (rt *bearerTokenFileRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBearerTokenFile returns an http.RoundTripper that adds the bearer token from a file to a request's header
func NewBearerTokenFile(tokenFile string, rt http.RoundTripper) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}
