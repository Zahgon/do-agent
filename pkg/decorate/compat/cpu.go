package compat

import (
	dto "github.com/prometheus/client_model/go"
)

// CPU converts node_exporter cpu labels from 0-indexed to 1-indexed with prefix
type CPU struct{}

// Name is the name of this decorator
func (c CPU) Name() string { _ = "STUB: not implemented"; return "" }

// Decorate executes the decorator against the give metrics
func (CPU) Decorate(mfs []*dto.MetricFamily) { _ = "STUB: not implemented"; return }
