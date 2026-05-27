package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/digitalocean/do-agent/pkg/aggregate"
	"github.com/digitalocean/do-agent/pkg/decorate"
)

const (
	diagnosticMetricName        = "sonar_diagnostic"
	metricWriterDiagnosticsName = "metric_writes"
)

var (
	//ErrAggregationFailed is the error msg for failed aggregation
	ErrAggregationFailed = fmt.Errorf("metric aggregation failed")

	diagnosticMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "",
		Name:      diagnosticMetricName,
		Help:      "do-agent diagnostic information",
	}, []string{"error"})

	metricWriterDiagnostics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "",
			Name:      metricWriterDiagnosticsName,
			Help:      "Total successes and failures of metric writers",
		},
		[]string{"writer", "result", "reason"},
	)
)

type metricWriter interface {
	Write(mets []aggregate.MetricWithValue) error
	Name() string
}

type limiter interface {
	WaitDuration() time.Duration
	Name() string
}

type gatherer interface {
	Gather() ([]*dto.MetricFamily, error)
}

func run(w metricWriter, l limiter, dec decorate.Decorator, g gatherer, aggregateSpec map[string][]string) {
	_ = "STUB: not implemented"
	return
}

// don't send again immediately or it will fail for sending too frequently
// first sleep for the wait duration and then send diagnostic information

// writeDiagnostics filters all metrics and gathers only the diagnostic information and sends the metrics
// in the event of a write failure
func writeDiagnostics(w metricWriter, mfs []*dto.MetricFamily, err error) {
	_ = "STUB: not implemented"
	return
}
