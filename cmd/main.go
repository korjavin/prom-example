package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	listen = "172.17.0.1:6060"
	max    = 1_000_000
	system = "myapp"
)

type Metrics struct {
	calls             prometheus.Counter
	sliceSize         prometheus.Gauge
	duration          prometheus.Summary
	durationHistogram prometheus.Histogram
}

var metrics = &Metrics{
	calls:     promauto.NewCounter(prometheus.CounterOpts{Namespace: system, Name: "superjob_calls"}),
	sliceSize: promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   system,
		Name:        "superjob_slice_size",
	}),
	duration: promauto.NewSummary(prometheus.SummaryOpts{
		Namespace:  system,
		Name:       "superjob_calls_duration_summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}),
	durationHistogram: promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: system,
		Name:      "superjob_calls_duration_histogram",
		Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
	}),
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/superjob", superhandler)
	log.Printf("Server listening %s", listen)
	http.ListenAndServe(listen, nil)
}

// TODO measure me
func superhandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))

	duration := time.Since(start)

	metrics.calls.Add(1)
	metrics.sliceSize.Set(float64(size))
	metrics.duration.Observe(float64(duration.Milliseconds()))
	metrics.durationHistogram.Observe(float64(duration.Milliseconds()))
}
