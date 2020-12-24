package main

import (
	"github.com/korjavin/prom-example/internal/pkg/handlers/superjob"
	"github.com/korjavin/prom-example/internal/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	appPort = "172.17.0.1:6060"
	system  = "myapp"
)

func main() {
	httpMetrics := metrics.New(system, "handlers")
	prometheus.MustRegister(
		httpMetrics.CallsCounter,
		httpMetrics.CallsGauge,
		httpMetrics.CallHistogram,
	)
	http.Handle("/metrics", promhttp.Handler())

	http.Handle("/superjob", superjob.New(httpMetrics))

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		if err := http.ListenAndServe(appPort, nil); err != nil {
			log.Println(err)
		}

		wg.Done()
	}()
	wg.Wait()

}
