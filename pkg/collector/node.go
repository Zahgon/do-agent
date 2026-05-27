package collector

import (
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/node_exporter/collector"
)

var networkNameWhitelistReg = regexp.MustCompile(`^(eno|eth|ens)\d+`)

var metricWhitelist = []string{
	"node_network_receive_bytes_total",
	"node_network_transmit_bytes_total",

	"node_memory_memtotal_bytes",
	"node_memory_memavailable_bytes",
	"node_memory_memfree_bytes",
	"node_memory_cached_bytes",
	"node_memory_swapcached_bytes",
	"node_memory_swapfree_bytes",
	"node_memory_swaptotal_bytes",

	"node_filesystem_size_bytes",
	"node_filesystem_free_bytes",

	"node_disk_read_bytes_total",
	"node_disk_written_bytes_total",
	"node_disk_reads_completed_total",
	"node_disk_read_time_seconds_total",
	"node_disk_writes_completed_total",
	"node_disk_write_time_seconds_total",
	"node_disk_discards_completed_total",
	"node_disk_discarded_sectors_total",
	"node_disk_discard_time_seconds_total",

	"node_cpu_seconds_total",
	"node_load1",
	"node_load5",
	"node_load15",
}

// NewNodeCollector creates a new prometheus NodeCollector
func NewNodeCollector() (*NodeCollector, error) { _ = "STUB: not implemented"; return nil, nil }

// NodeCollector is a collector that collects data using
// prometheus/node_exporter. Since prometheus returns an internal type we have
// to wrap it with our own type
type NodeCollector struct {
	collectFunc    func(ch chan<- prometheus.Metric)
	describeFunc   func(ch chan<- *prometheus.Desc)
	collectorsFunc func() map[string]collector.Collector
}

// Collectors returns the list of collectors registered
func (n *NodeCollector) Collectors() map[string]collector.Collector {
	_ = "STUB: not implemented"
	return nil

	// Name returns the name of this collector
}

func (n *NodeCollector) Name() string {
	_ = "STUB: not implemented"

	// Collect collects metrics using prometheus/node_exporter
	return ""
}

func (n *NodeCollector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

// Desc doesn't allow access to underlying fields like fqName. The String() output contains
// Desc{fqName: "node_network_transmit_bytes_total", help: "Network device statistic transmit_bytes.", constLabels: {}, variableLabels: [device]}
// this is ugly but currently all we can do

// validNetwork checks that the network name for this metric is whitelisted. If the metric is not in the whitelist
func validNetwork(m prometheus.Metric) bool { _ = "STUB: not implemented"; return false }

// Describe describes the metrics collected using prometheus/node_exporter
func (n *NodeCollector) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }
