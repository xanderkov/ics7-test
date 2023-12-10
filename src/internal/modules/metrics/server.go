package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Listen(address string) error {
	//use separated ServeMux to prevent handling on the global Mux
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(address, mux)
}

func InvokeMetrics() {
	go func() {
		_ = Listen("0.0.0.0:8082")
	}()
}
