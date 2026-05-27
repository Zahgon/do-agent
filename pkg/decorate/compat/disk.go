package compat

import (
	dto "github.com/prometheus/client_model/go"
)

const diskSectorSize = float64(512)

// Disk converts node_exporter disk metrics from bytes to sectors
type Disk struct{}

// Name is the name of this decorator
func (d Disk) Name() string { _ = "STUB: not implemented"; return "" }

// Decorate converts bytes to sectors
func (Disk) Decorate(mfs []*dto.MetricFamily) { _ = "STUB: not implemented"; return }

func bytesToSector(val *float64) *float64 { _ = "STUB: not implemented"; return nil }
