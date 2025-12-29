package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

var opsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "myapp",
	Help: "the total number of processed events",
})

// Define a Counter
var currentUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "myapp_current_users",
	Help: "Current number of active users",
})

var requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "myapp_request_duration_seconds",
	Help:    "Histogram of response time for handler",
	Buckets: prometheus.DefBuckets,
})

func init() {
	prometheus.MustRegister(opsProcessed)
	prometheus.MustRegister(currentUsers)
	prometheus.MustRegister(requestDuration)
}

func process() {
	start := time.Now()

	opsProcessed.Inc()

	currentUsers.Set(float64(rand.Intn(100)))

	requestDuration.Observe(time.Since(start).Seconds())
}
