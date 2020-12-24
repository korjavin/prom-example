package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	CallsCounter  *prometheus.CounterVec
	CallsGauge    *prometheus.GaugeVec
	CallHistogram *prometheus.HistogramVec
}

func New(namespace, subsystem string) *Metric {
	return &Metric{
		CallsCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "calls_count",
				Help:        "HelpInfo",
				ConstLabels: prometheus.Labels{"labelKey": "labelValue"},
			},
			[]string{"handler"},
		),
		CallsGauge: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace:   namespace,
				Subsystem:   subsystem,
				Name:        "gauge",
				Help:        "Gauge HelpInfo",
				ConstLabels: prometheus.Labels{"labelKey": "labelValue"},
			},
			[]string{"handler"},
		),
		CallHistogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "histogram",
				Help:      "histogram HelpInfo",
			},
			[]string{"handler"},
		),
	}
}
