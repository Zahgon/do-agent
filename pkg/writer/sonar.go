package writer

import (
	"fmt"

	"github.com/digitalocean/do-agent/pkg/aggregate"
	"github.com/digitalocean/do-agent/pkg/clients/tsclient"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// ErrMetricTooLong is returned when trying to write a metric that exceeds the length limit
	// defined by client.MaxMetricLength
	ErrMetricTooLong = fmt.Errorf("metric length is too long to write")
	// ErrTooManyMetrics is returned when calling Write with too many metrics
	// defined by client.MaxBatchSize
	ErrTooManyMetrics = fmt.Errorf("too many metrics to send")

	// ErrFlushFailure is returned when Flush fails for any reason
	ErrFlushFailure = fmt.Errorf("flush failure")
)

// Sonar writes metrics to DigitalOcean sonar
type Sonar struct {
	client         tsclient.Client
	firstWriteSent bool
	c              *prometheus.CounterVec
}

// NewSonar creates a new Sonar writer
func NewSonar(client tsclient.Client, c *prometheus.CounterVec) *Sonar {
	_ = "STUB: not implemented"
	return nil
}

// Write writes the metrics to Sonar and returns the amount of time to wait
// before the next write
func (s *Sonar) Write(mets []aggregate.MetricWithValue) error {
	_ = "STUB: not implemented"
	return nil
}

// Name is the name of this writer
func (s *Sonar) Name() string { _ = "STUB: not implemented"; return "" }
