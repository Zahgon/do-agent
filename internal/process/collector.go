package process

import (
	"github.com/prometheus/client_golang/prometheus"
)

type processCollector struct {
	collectFn func(chan<- prometheus.Metric)
	rss       *prometheus.Desc
	cpuTime   *prometheus.Desc
}

// NewProcessCollector returns a collector which exports the current state of
// process metrics including CPU, memory and file descriptor usage as well as
// the process start time.
func NewProcessCollector() prometheus.Collector {
	_ = "STUB: not implemented"
	return *new(prometheus.Collector)
}

// nop

// Describe returns all descriptions of the collector.
func (c *processCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"

	// Collect returns the current state of all metrics of the collector.
	return
}

func (c *processCollector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (c *processCollector) processCollect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}
