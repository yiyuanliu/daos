//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//
// +build linux,amd64
//

package telemetry

import (
	"context"
	"testing"
)

/*
#cgo LDFLAGS: -lgurt

#include "gurt/telemetry_common.h"
#include "gurt/telemetry_consumer.h"
#include "gurt/telemetry_producer.h"

static int
set_gauge(struct d_tm_node_t **metric, uint64_t value, char *item)
{
	return d_tm_set_gauge(metric, value, item, NULL);
}
*/
import "C"

type (
	cMetric struct {
		name   string
		short  string
		long   string
		lo     float64
		hi     float64
		cur    float64
		sum    float64
		mean   float64
		stddev float64
		str    string
		node   *C.struct_d_tm_node_t
	}
	testMetric struct {
		c cMetric
		g Metric
	}
	testMetricsMap map[MetricType]testMetric
)

var testMetrics testMetricsMap

func setupTestMetrics(t *testing.T) context.Context {
	rc := C.d_tm_init(0, 8192, 0)
	if rc != 0 {
		t.Fatalf("failed to init telemetry: %d", rc)
	}

	ctx, err := Init(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}

	testMetrics = testMetricsMap{
		MetricTypeGauge: {
			c: cMetric{
				name:   "test_gauge",
				short:  "short gauge",
				long:   "long gauge",
				lo:     1,
				cur:    42,
				hi:     64,
				sum:    107,
				mean:   35.666666666666664,
				stddev: 31.973947728319903,
				str:    "Gauge: test_gauge = 42 min: 1 max: 64 mean: 35.666667 size: 3 std dev: 31.973948",
			},
		},
	}

	for mt, mv := range testMetrics {
		switch mt {
		case MetricTypeGauge:
			rc = C.d_tm_add_metric(&mv.c.node, C.CString(mv.c.name), C.D_TM_GAUGE, C.CString(mv.c.short), C.CString(mv.c.long))
			if rc != 0 {
				t.Fatalf("failed to add %s: %d", mv.c.name, rc)
			}
			for _, val := range []float64{mv.c.lo, mv.c.hi, mv.c.cur} {
				rc = C.set_gauge(&mv.c.node, C.uint64_t(val), nil)
				if rc != 0 {
					t.Fatalf("failed to set %s: %d", mv.c.name, rc)
				}
			}
		default:
			t.Fatalf("metric type %d not supported", mt)
		}
	}

	return ctx
}

func cleanupTestMetrics(t *testing.T) {
	C.d_tm_fini()
}
