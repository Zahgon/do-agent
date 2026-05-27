package collector

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/digitalocean/do-agent/internal/log"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

var defaultScrapeTimeout = 5 * time.Second

type scraperOpts struct {
	timeout         time.Duration
	logLevel        log.Level
	bearerToken     string
	bearerTokenFile string
}

// Option is used to configure optional scraper options.
type Option func(o *scraperOpts)

// WithBearerToken configures a scraper to use a bearer token
func WithBearerToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBearerTokenFile configures a scraper to use a bearer token read from a file
func WithBearerTokenFile(tokenFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout configures a scraper with a timeout for scraping metrics.
func WithTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogLevel configures a custom log level for scraping.
func WithLogLevel(l log.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewScraper creates a new scraper to scrape metrics from the provided host
func NewScraper(name, metricsEndpoint string, extraMetricLabels []*dto.LabelPair, whitelist map[string]bool, opts ...Option) (*Scraper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setup http client, add auth roundtrippers

// Scraper is a remote metric scraper that scrapes HTTP endpoints
type Scraper struct {
	timeout            time.Duration
	logLevel           log.Level
	req                *http.Request
	client             *http.Client
	name               string
	whitelist          map[string]bool
	extraMetricLabels  []*dto.LabelPair
	scrapeDurationDesc *prometheus.Desc
	scrapeSuccessDesc  *prometheus.Desc
}

// log emits log messages respecting the scraper's log level.
func (s *Scraper) log(msg string, params ...interface{}) { _ = "STUB: not implemented"; return }

// readStream makes an HTTP request to the remote and returns the response body
// upon successful response
func (s *Scraper) readStream(ctx context.Context) (r io.ReadCloser, outerr error) {
	_ = "STUB: not implemented"
	// close the reader if we return an error
	return *new(io.ReadCloser), nil
}

// This should not happen, but if it does it'll be nice
// to know why we have a bunch of unclosed messages

// Describe describes this collector
func (s *Scraper) Describe(ch chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

// Collect collectrs metrics from the remote endpoint and reports them to ch
func (s *Scraper) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (s *Scraper) scrape(ctx context.Context, ch chan<- prometheus.Metric) (outerr error) {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of this scraper
func (s *Scraper) Name() string {
	_ = "STUB: not implemented"

	// FilterMetric returns true if the metric should be skipped (filtered out)
	return ""
}

func (s *Scraper) FilterMetric(metricFamily *dto.MetricFamily) bool {
	_ = "STUB: not implemented"
	return false
	// if no whitelist treat all metrics as valid
}

// convertMetricFamily converts the dto metrics parsed from the expfmt package
// into the prometheus.Metrics required to pass over the channel
//
// this was copied and extended/refactored from github.com/prometheus/node_exporter
// see https://github.com/prometheus/node_exporter/blob/f56e8fcdf48ead56f1f149dbf1301ac028ef589b/collector/textfile.go#L63
// for more details
func convertMetricFamily(metricFamily *dto.MetricFamily, ch chan<- prometheus.Metric, extraLabels []*dto.LabelPair) {
	_ = "STUB: not implemented"
	return
}

// getLabelNamesAndValues returns a slice of label names and a slice of label values from the metric and extra labels.
func getLabelNamesAndValues(metric *dto.Metric, extraLabels []*dto.LabelPair, allLabelNames map[string]struct{}) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAllLabelNames returns the map of all label names from the metric family including any extra labels provided.
func getAllLabelNames(metricFamily *dto.MetricFamily, extraLabels []*dto.LabelPair) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
