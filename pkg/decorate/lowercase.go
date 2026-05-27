package decorate

import (
	dto "github.com/prometheus/client_model/go"
)

// LowercaseNames decorates metrics to be have all lowercase label names
type LowercaseNames struct{}

// Decorate decorates the provided metrics for compatibility
func (LowercaseNames) Decorate(mfs []*dto.MetricFamily) {
	_ = "STUB: not implemented"
	// names come back with varying cases like some_TCP_connection
	// and we want consistency so we lowercase them
	return
}

// Name is the name of this decorator
func (LowercaseNames) Name() string { _ = "STUB: not implemented"; return "" }
