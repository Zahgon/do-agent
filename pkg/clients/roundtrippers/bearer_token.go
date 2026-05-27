package roundtrippers

import (
	"net/http"
)

type bearerTokenRoundTripper struct {
	token string
	rt    http.RoundTripper
}

// RoundTrip implements http.RoundTripper's interface
func (rt *bearerTokenRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBearerToken returns an http.RoundTripper that adds the bearer token to a request's header
func NewBearerToken(token string, rt http.RoundTripper) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}
