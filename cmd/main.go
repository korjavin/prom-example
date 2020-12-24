package main

import (
	"github.com/korjavin/prom-example/internal/pkg/handlers/superjob"
	"github.com/korjavin/prom-example/internal/pkg/metrics"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	appPort     = "172.17.0.1:6060"
	metricsPort = "172.17.0.1:6061"
	system      = "myapp"
)

func main() {

	httpMetrics := metrics.New(system, "handlers")
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())

	httpMux := http.NewServeMux()
	httpMux.Handle("/superjob", superjob.New(httpMetrics))

	var wg sync.WaitGroup

	for port, handler := range map[string]*http.ServeMux{appPort: httpMux, metricsPort: metricsMux} {
		wg.Add(1)

		go func(port string, handler *http.ServeMux) {
			if err := http.ListenAndServe(port, handler); err != nil {
				log.Println(err)
			}

			wg.Done()
		}(port, handler)
	}
	wg.Wait()

}
