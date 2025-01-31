package initializers

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Default metrics for the application
	
	// Counter for the number of requests made to GET /posts
	GetPostsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "get_posts_counter",
			Help: "Number of requests made to GET /posts",
		},
	)

	// Histogram for the time taken to process a request to GET /posts
	GetPostsDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "get_posts_duration",
			Help:    "Time taken to process a request to GET /posts in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)

	// Histogram for the time taken to process a request
	RequestMetrics = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:   "http_request_duration_seconds",
			Help:  "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status"},
	)

	// Request Counter according to http method
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests made.",
		},
		[]string{"method", "path", "status"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(GetPostsCounter)
	prometheus.MustRegister(GetPostsDuration)
	prometheus.MustRegister(RequestMetrics)
	prometheus.MustRegister(RequestCounter)
}
