package superjob

import (
	"fmt"
	"github.com/korjavin/prom-example/internal/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
)

const max = 1_000_000

type Handler struct {
	metrics *metrics.Metric
}

func New(metrics *metrics.Metric) *Handler {
	return &Handler{
		metrics: metrics,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	size := rand.Int31n(max)
	slice := make([]int, size)
	for idx := range slice {
		slice[idx] = rand.Intn(max)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Len %d", len(slice))))

	h.metrics.CallsCounter.With(prometheus.Labels{"handler": "superjob"}).Inc()
}
