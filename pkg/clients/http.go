package clients

import (
	"net/http"
	"time"
)

// HTTPClient is can make HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHTTP creates a new HTTP client with the provided timeout
func NewHTTP(timeout time.Duration) *http.Client { _ = "STUB: not implemented"; return nil }

// FakeHTTPClient is used for testing
type FakeHTTPClient struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

// Do an HTTP request for testing
func (c *FakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDebug creates a new DebugHTTPClient
func NewDebug(timeout time.Duration) *DebugHTTPClient { _ = "STUB: not implemented"; return nil }

// DebugHTTPClient is an *http.Client that prints Headers and Body to log
type DebugHTTPClient struct {
	*http.Client
}

// Do sends the http request and logs headers and body to DEBUG
func (c *DebugHTTPClient) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
