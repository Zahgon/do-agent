package decorate

import (
	dto "github.com/prometheus/client_model/go"
)

// LabelAppender is a list of label pairs that need to be added on all metrics
type LabelAppender []*dto.LabelPair

// Decorate adds metric labels from its list
func (l LabelAppender) Decorate(mfs []*dto.MetricFamily) { _ = "STUB: not implemented"; return }

// Name is the name of this decorator
func (LabelAppender) Name() string { _ = "STUB: not implemented"; return "" }
