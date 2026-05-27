package aggregate

import (
	dto "github.com/prometheus/client_model/go"
)

// MetricWithValue is a representation of a label formatted metric with a value
type MetricWithValue struct {
	LFM   map[string]string
	Value float64
}

// Aggregate aggregates metric families according to the given aggregate spec.
// A spec with key: {"metricName": "aggregateLabel"} will remove the "aggregateLabel" from all
// "metricName" metric families
func Aggregate(metrics []*dto.MetricFamily, aggregateSpec map[string][]string) ([]MetricWithValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we currently don't support other types of metrics

// if the metric family is to be aggregated, aggregate away the specified labels
