package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	CallsCounter *prometheus.CounterVec
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
	}
}
