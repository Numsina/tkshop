package metrcs

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func InitPrometheus() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9001", nil)
	}()
}
