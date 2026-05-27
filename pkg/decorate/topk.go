package decorate

import (
	dto "github.com/prometheus/client_model/go"
)

// TopK is a decorator that removes metrics not in the top K by value
type TopK struct {
	K uint
	N string
}

// Decorate removes all but the top K metrics for a given metric name
func (t TopK) Decorate(mfs []*dto.MetricFamily) { _ = "STUB: not implemented"; return }

// Name is the name of this decorator
func (t TopK) Name() string { _ = "STUB: not implemented"; return "" }

type metricHeap []*dto.Metric

func (m metricHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (m metricHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }

// invert less function to create max heap
func (m metricHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (m *metricHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (m *metricHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (m *metricHeap) TopK(k uint) []*dto.Metric { _ = "STUB: not implemented"; return nil }
