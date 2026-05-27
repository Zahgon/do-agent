package writer

import (
	"io"
	"sync"

	"github.com/digitalocean/do-agent/pkg/aggregate"
	"github.com/prometheus/client_golang/prometheus"
)

// File writes metrics to an io.Writer
type File struct {
	w io.Writer
	m *sync.Mutex
	c *prometheus.CounterVec
}

// NewFile creates a new File writer with the provided writer
func NewFile(w io.Writer, c *prometheus.CounterVec) *File { _ = "STUB: not implemented"; return nil }

// Write writes metrics to the file
func (w *File) Write(mets []aggregate.MetricWithValue) error { _ = "STUB: not implemented"; return nil }

// Name is the name of this writer
func (w *File) Name() string { _ = "STUB: not implemented"; return "" }
