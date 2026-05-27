package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/digitalocean/do-agent/pkg/clients/tsclient"
	"github.com/digitalocean/do-agent/pkg/decorate"
)

var (
	config struct {
		targets                map[string]string
		metadataURL            *url.URL
		authURL                *url.URL
		bearerToken            string
		bearerTokenFile        string
		sonarEndpoint          string
		stdoutOnly             bool
		debug                  bool
		syslog                 bool
		noProcesses            bool
		noNode                 bool
		kubernetes             string
		dbaas                  string
		mongodb                string
		webListenAddress       string
		webListen              bool
		additionalLabels       []string
		defaultMaxBatchSize    int
		defaultMaxMetricLength int
		promAddr               string
		diMetricsPath          string
		gpuMetricsPath         string
		topK                   int
		scrapeTimeout          time.Duration
	}

	// additionalParams is a list of extra command line flags to append
	// this is mostly needed for appending node_exporter flags when necessary.
	additionalParams []string

	// disabledCollectors is a hash used by disableCollectors to prevent
	// duplicate entries
	disabledCollectors = map[string]interface{}{}
)

const internalProxyURL = "http://169.254.169.254"

const (
	defaultAuthURL          = internalProxyURL
	defaultSonarURL         = ""
	defaultWebListenAddress = "127.0.0.1:9100"

	processScrapingDropletPct = 50
)

var defaultMetadataURL = fmt.Sprintf("%s/metadata", internalProxyURL)

func init() {
	kingpin.CommandLine.Name = "do-agent"

	kingpin.Flag("auth-host", "Endpoint to use for obtaining droplet app key").
		Default(defaultAuthURL).
		Envar("DO_AGENT_AUTH_URL").
		URLVar(&config.authURL)

	kingpin.Flag("metadata-host", "Endpoint to use for obtaining droplet metadata").
		Default(defaultMetadataURL).
		URLVar(&config.metadataURL)

	kingpin.Flag("sonar-host", "Endpoint to use for delivering metrics").
		Default(defaultSonarURL).
		Envar("DO_AGENT_SONAR_HOST").
		StringVar(&config.sonarEndpoint)

	kingpin.Flag("stdout-only", "write all metrics to stdout only").
		BoolVar(&config.stdoutOnly)

	kingpin.Flag("debug", "display debug information to stdout").
		BoolVar(&config.debug)

	kingpin.Flag("syslog", "enable logging to syslog").
		BoolVar(&config.syslog)

	kingpin.Flag("k8s-metrics-path", "enable DO Kubernetes metrics collection (this must be a DOKS metrics endpoint)").
		StringVar(&config.kubernetes)

	kingpin.Flag("bearer-token", "sets the `Authorization` header on every scrape request with the configured bearer token (mutually exclusive with `bearer-token-file`)").
		StringVar(&config.bearerToken)

	kingpin.Flag("bearer-token-file", "sets the `Authorization` header on every scrape request with the bearer token read from the configured file (mutually exclusive with `bearer-token`)").
		StringVar(&config.bearerTokenFile)

	kingpin.Flag("no-collector.processes", "disable processes cpu/memory collection").
		Default("true").
		BoolVar(&config.noProcesses)

	kingpin.Flag("no-collector.node", "disable processes node collection").
		Default("false").
		BoolVar(&config.noNode)

	kingpin.Flag("dbaas-metrics-path", "enable DO DBAAS metrics collection (this must be a DO DBAAS metrics endpoint)").
		StringVar(&config.dbaas)

	kingpin.Flag("mongodb-metrics-path", "enable DO DBAAS MongoDB metrics collection (this must be a DO DBAAS metrics endpoint)").
		StringVar(&config.mongodb)

	kingpin.Flag("metrics-path", "enable metrics collection from a prometheus endpoint").
		StringVar(&config.promAddr)

	kingpin.Flag("gpu-metrics-path", "enable GPU metrics collection from a prometheus endpoint (e.g., AMD device-metrics-exporter)").
		StringVar(&config.gpuMetricsPath)

	kingpin.Flag("di-metrics-path", "enable Dedicated Inference (DI) metrics collection from a prometheus endpoint").
		StringVar(&config.diMetricsPath)

	kingpin.Flag("web.listen", "enable a local endpoint for scrapeable prometheus metrics as well").
		Default("false").
		BoolVar(&config.webListen)

	kingpin.Flag("web.listen-address", `write prometheus metrics to the specified port (ex. ":9100")`).
		Default(defaultWebListenAddress).
		StringVar(&config.webListenAddress)

	kingpin.Flag("additional-label", "key value pairs for labels to add to all metrics (ex: user_id:1234)").StringsVar(&config.additionalLabels)

	kingpin.Flag("max-batch-size", "default max batch size for sending metrics. This will be overridden after first write").
		IntVar(&config.defaultMaxBatchSize)
	kingpin.Flag("max-metric-length", "default max metric length for metrics. This will be overridden after first write").
		IntVar(&config.defaultMaxMetricLength)

	kingpin.Flag("process-topk", "number of top processes to scrape").Default("30").IntVar(&config.topK)

	kingpin.Flag("scrape-timeout", "timeout for scraping metrics").
		Default("30s").
		DurationVar(&config.scrapeTimeout)

}

func initConfig() { _ = "STUB: not implemented"; return }

// read flags from cli directly first so we have access to them

// parse all command line flags which are defined across the app

func checkConfig() error { _ = "STUB: not implemented"; return nil }

func toggleGradualRollouts() { _ = "STUB: not implemented"; return }

func initWriter(wc *prometheus.CounterVec) (metricWriter, limiter) {
	_ = "STUB: not implemented"
	return *new(metricWriter), *new(limiter)
}

func initDecorator() decorate.Chain { _ = "STUB: not implemented"; return *new(decorate.Chain) }

// TopK sonar processes

// If additionalLabels provided convert into decorator

// initAggregatorSpecs initializes the field aggregation specifications.
// The map's key is the prometheus metric name to aggregate over, and the value is the label to aggregate away.
// The metric name should be in the format expected after the decorators are applied, e.g., lowercase.
func initAggregatorSpecs() map[string][]string { _ = "STUB: not implemented"; return nil }

// WrappedTSClient wraps the tsClient and adds a Name method to it
type WrappedTSClient struct {
	tsclient.Client
}

// Name returns the name of the client
func (m *WrappedTSClient) Name() string { _ = "STUB: not implemented"; return "" }

func newTimeseriesClient() *WrappedTSClient { _ = "STUB: not implemented"; return nil }

// initCollectors initializes the prometheus collectors. By default this
// includes node_exporter and buildInfo for each remote target
func initCollectors() []prometheus.Collector {
	_ = "STUB: not implemented"
	// buildInfo provides build information for tracking metrics internally
	return nil
}

// Top process collection

// create the default DO agent to collect metrics about
// this device

// appendKubernetesCollectors appends a kubernetes metrics collector if it can be initialized successfully
func appendKubernetesCollectors(cols []prometheus.Collector) []prometheus.Collector {
	_ = "STUB: not implemented"
	return nil
}

// disableCollectors disables collectors by names by adding a list of
// --no-collector.<name> flags to additionalParams
func disableCollectors(names ...string) { _ = "STUB: not implemented"; return }

// already disabled

// disableCollectorFlag creates the correct cli flag for the given collector name
func disableCollectorFlag(name string) string { _ = "STUB: not implemented"; return "" }

func convertToLabelPairs(s []string) []*dto.LabelPair { _ = "STUB: not implemented"; return nil }

// require a key value pair
