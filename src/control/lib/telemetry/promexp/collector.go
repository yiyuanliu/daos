//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//
// +build linux,amd64
//

package promexp

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/daos-stack/daos/src/control/lib/telemetry"
	"github.com/daos-stack/daos/src/control/logging"
)

type (
	Collector struct {
		log            logging.Logger
		summary        *prometheus.SummaryVec
		ignoredMetrics []*regexp.Regexp
		sources        []*EngineSource
	}

	CollectorOpts struct {
		Ignores []string
	}

	EngineSource struct {
		ctx   context.Context
		Index uint32
	}

	labelMap map[string]string
)

func NewEngineSource(parent context.Context, idx uint32) (*EngineSource, error) {
	ctx, err := telemetry.Init(parent, idx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init telemetry")
	}

	return &EngineSource{
		ctx:   ctx,
		Index: idx,
	}, nil
}

func defaultCollectorOpts() *CollectorOpts {
	return &CollectorOpts{}
}

func NewCollector(log logging.Logger, opts *CollectorOpts, sources ...*EngineSource) (*Collector, error) {
	if len(sources) == 0 {
		return nil, errors.New("Collector must have > 0 sources")
	}

	if opts == nil {
		opts = defaultCollectorOpts()
	}

	c := &Collector{
		log:     log,
		sources: sources,
		summary: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Namespace: "engine",
				Subsystem: "exporter",
				Name:      "scrape_duration_seconds",
				Help:      "daos_exporter: Duration of a scrape job.",
			},
			[]string{"source", "result"},
		),
	}

	for _, pat := range opts.Ignores {
		re, err := regexp.Compile(pat)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to compile %q", pat)
		}
		c.ignoredMetrics = append(c.ignoredMetrics, re)
	}

	return c, nil
}

func fixPath(in string) (labels labelMap, name string) {
	name = strings.Map(func(r rune) rune {
		if !unicode.IsPrint(r) || unicode.IsSpace(r) {
			return '_'
		}
		switch r {
		case '/', ':':
			return '_'
		}

		return r
	}, strings.TrimLeft(in, "/"))

	labels = make(labelMap)
	ID_re := regexp.MustCompile(`ID_(\d+)_`)
	io_re := regexp.MustCompile(`io_(\d+)_`)

	name = ID_re.ReplaceAllString(name, "")
	io_matches := io_re.FindStringSubmatch(name)
	if len(io_matches) > 0 {
		labels["target"] = io_matches[1]
		name = io_re.ReplaceAllString(name, "io_")
	}

	return
}

func (es *EngineSource) Collect(log logging.Logger, ch chan<- *rankMetric) {
	metrics := make(chan telemetry.Metric)
	go func() {
		//log.Debugf("starting collection for idx %d", es.Index)
		startedAt := time.Now()
		if err := telemetry.CollectMetrics(es.ctx, "", metrics); err != nil {
			log.Errorf("failed to collect metrics for engine idx %d: %s", es.Index, err)
			return
		}
		log.Debugf("finished collection for idx %d in %s", es.Index, time.Now().Sub(startedAt))
	}()

	rank, err := telemetry.GetRank(es.ctx)
	if err != nil {
		log.Errorf("failed to get rank for engine idx %d: %s", es.Index, err)
		return
	}

	//log.Debugf("reading metrics for rank %d\n", rank)
	for metric := range metrics {
		ch <- &rankMetric{
			r: rank,
			m: metric,
		}
	}
}

type rankMetric struct {
	r uint32
	m telemetry.Metric
}

func (c *Collector) isIgnored(name string) bool {
	for _, re := range c.ignoredMetrics {
		if re.MatchString(name) {
			return true
		}
	}

	return false
}

func (lm labelMap) keys() (keys []string) {
	for label := range lm {
		keys = append(keys, label)
	}

	return
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	rankMetrics := make(chan *rankMetric)
	go func(sources []*EngineSource) {
		for _, source := range c.sources {
			source.Collect(c.log, rankMetrics)
		}
		close(rankMetrics)
	}(c.sources)

	gauges := make(map[string]*prometheus.GaugeVec)
	counters := make(map[string]*prometheus.CounterVec)

	startedAt := time.Now()
	for rm := range rankMetrics {
		//c.log.Debugf("%s: %.02f\n", rm.m.Name(), rm.m.FloatValue())

		labels, path := fixPath(rm.m.Path())
		labels["rank"] = fmt.Sprintf("%d", rm.r)

		baseName := prometheus.BuildFQName("engine", path, rm.m.Name())
		shortDesc := rm.m.ShortDesc()

		var found bool
		switch rm.m.Type() {
		case telemetry.MetricTypeGauge:
			if c.isIgnored(baseName) {
				continue
			}
			var gv *prometheus.GaugeVec
			gv, found = gauges[baseName]
			if !found {
				gv = prometheus.NewGaugeVec(prometheus.GaugeOpts{
					Name: baseName,
					Help: shortDesc,
				}, labels.keys())
				gauges[baseName] = gv
			}
			g := gv.With(prometheus.Labels(labels))
			g.Set(rm.m.FloatValue())
		case telemetry.MetricTypeCounter:
			if c.isIgnored(baseName) {
				break
			}
			var cv *prometheus.CounterVec
			cv, found = counters[baseName]
			if !found {
				cv = prometheus.NewCounterVec(prometheus.CounterOpts{
					Name: baseName,
					Help: shortDesc,
				}, labels.keys())
				counters[baseName] = cv
			}
			c := cv.With(prometheus.Labels(labels))
			c.Add(rm.m.FloatValue())
		default:
			c.log.Errorf("metric type %d not supported", rm.m.Type())
		}
	}

	for _, gv := range gauges {
		gv.Collect(ch)
	}
	for _, cv := range counters {
		cv.Collect(ch)
	}
	c.log.Debugf("collected all metrics in %s", time.Now().Sub(startedAt))
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.summary.Describe(ch)
}
