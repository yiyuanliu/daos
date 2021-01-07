//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package telemetry

import (
	"testing"

	"github.com/daos-stack/daos/src/control/common"
)

func TestTelemetry_Basics(t *testing.T) {
	ctx := setupTestMetrics(t)
	defer cleanupTestMetrics(t)

	var m Metric
	var err error

	for mt, mv := range testMetrics {
		switch mt {
		case MetricTypeGauge:
			m, err = GetGauge(ctx, mv.c.name)
			if err != nil {
				t.Fatal(err)
			}

			common.AssertEqual(t, mv.c.name, m.Name(), "Name() failed")
			common.AssertEqual(t, mv.c.short, m.ShortDesc(), "ShortDesc() failed")
			common.AssertEqual(t, mv.c.long, m.LongDesc(), "LongDesc() failed")
			common.AssertEqual(t, mv.c.cur, m.FloatValue(), "FloatValue() failed")
			common.AssertEqual(t, mv.c.str, m.String(), "String() failed")

			if sm, ok := m.(StatsMetric); ok {
				common.AssertEqual(t, mv.c.lo, sm.FloatLow(), "FloatLow() failed")
				common.AssertEqual(t, mv.c.hi, sm.FloatHigh(), "FloatHigh() failed")
				common.AssertEqual(t, mv.c.sum, sm.FloatSum(), "FloatSum() failed")
				common.AssertEqual(t, mv.c.mean, sm.Mean(), "Mean() failed")
				common.AssertEqual(t, mv.c.stddev, sm.StdDev(), "StdDev() failed")
			}
		default:
			t.Fatalf("%d not tested", mt)
		}
	}
}
