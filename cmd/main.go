package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	listen = "172.17.0.1:6060"
	max    = 1_000_000
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/superjob", superhandler)
	http.ListenAndServe(listen, nil)
}

// TODO measure me
func superhandler(w http.ResponseWriter, r *http.Request) {
	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))
}
