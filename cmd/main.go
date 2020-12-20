package main

import (
	"fmt"
	"math/rand"
	"net/http"

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
	calls prometheus.Counter
}

var metrics = &Metrics{
	calls: promauto.NewCounter(prometheus.CounterOpts{Namespace: system, Name: "superjob_calls"}),
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/superjob", superhandler)
	http.ListenAndServe(listen, nil)
}

// TODO measure me
func superhandler(w http.ResponseWriter, r *http.Request) {
	metrics.calls.Add(1)
	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))
}
