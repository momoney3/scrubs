// Package tools use to setup prometheus
package tools

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// func main() {
// 	http.Handle("/metrics", promhttp.Handler())
// 	http.ListenAndServe(":8080", nil)p
// }

var OpsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "myapp",
	Help: "the total number of processed events",
})

var CurrentUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "myapp_current_users",
	Help: "Current number of active users",
})

var RequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "myapp_request_duration_seconds",
	Help:    "Histogram of response time for handler",
	Buckets: prometheus.DefBuckets,
})

func init() {
	prometheus.MustRegister(OpsProcessed)
	prometheus.MustRegister(CurrentUsers)
	prometheus.MustRegister(RequestDuration)
}

func Process() {
	start := time.Now()

	OpsProcessed.Inc()

	CurrentUsers.Set(float64(rand.Intn(100)))

	RequestDuration.Observe(time.Since(start).Seconds())
}
