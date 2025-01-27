package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Service struct {
	Name    string
	Version string
	HTTP    []HTTPEndpoint
	GRPC    []GRPCEndpoint
}

type HTTPEndpoint struct {
	Method string
	Code   int
	Path   string
	// Better with NormFloat64() * desiredStdDev + desiredMean ?
	Count   [2]int // [min, max]
	Latency [2]int // [min, max], in ms
}

type GRPCEndpoint struct {
	Service string
	Method  string
	Code    string // https://github.com/grpc/grpc-go/blob/4103cfc52a951673d441f8b2c02eee96e31f1897/codes/code_string.go#L31
	// Better with NormFloat64() * desiredStdDev + desiredMean ?
	Count   [2]int // [min, max]
	Latency [2]int // [min, max], in ms
}

var services = []Service{
	{
		Name:    "users",
		Version: "v1.2.3",
	},
	{
		Name:    "payments",
		Version: "v2.0.1",
	},
	{
		Name:    "orders",
		Version: "v1.24.42",
		HTTP: []HTTPEndpoint{
			{
				Method:  http.MethodGet,
				Code:    http.StatusOK,
				Path:    "/api/orders/{order_ref}",
				Count:   [2]int{4, 9},
				Latency: [2]int{100, 350},
			},
			{
				Method:  http.MethodPost,
				Code:    http.StatusOK,
				Path:    "/api/orders",
				Count:   [2]int{1, 3},
				Latency: [2]int{350, 650},
			},
		},
	},
	{
		Name:    "products",
		Version: "v3.6.9",
		HTTP: []HTTPEndpoint{
			{
				Method:  http.MethodGet,
				Code:    http.StatusOK,
				Path:    "/api/products",
				Count:   [2]int{1, 20},
				Latency: [2]int{50, 1_000},
			},
			{
				Method:  http.MethodGet,
				Code:    http.StatusOK,
				Path:    "/api/products/{product_id}",
				Count:   [2]int{0, 10},
				Latency: [2]int{30, 900},
			},
			{
				Method:  http.MethodGet,
				Code:    http.StatusOK,
				Path:    "/api/products/{product_id}/reviews",
				Count:   [2]int{0, 8},
				Latency: [2]int{90, 1_400},
			},
			{
				Method:  http.MethodPost,
				Code:    http.StatusNotFound,
				Path:    "/api/products/{product_id}",
				Count:   [2]int{0, 2},
				Latency: [2]int{5, 10},
			},
		},
		GRPC: []GRPCEndpoint{
			{
				Service: "grpc.ProductService",
				Method:  "updateProductDetails",
				Code:    "OK",
				Count:   [2]int{1, 3},
				Latency: [2]int{15, 55},
			},
			{
				Service: "grpc.ProductService",
				Method:  "updateProductStop",
				Code:    "OK",
				Count:   [2]int{1, 10},
				Latency: [2]int{5, 25},
			},
			{
				Service: "grpc.ProductService",
				Method:  "updateProductStop",
				Code:    "InvalidArgument",
				Count:   [2]int{0, 8},
				Latency: [2]int{5, 25},
			},
		},
	},
}

var (
	serviceInfos = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "app_infos",
		Help: "Basic informations about a service",
	}, []string{"service", "version"})

	httpRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "HTTP requests.",
	}, []string{"service", "code", "method", "path"})
	httpRequestsDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_requests_duration_seconds",
		Help:    "HTTP requests durations in seconds.",
		Buckets: prometheus.DefBuckets,
	}, []string{"service", "code", "method", "path"})

	grpcRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "grpc_server_handled_total",
		Help: "gRPC requests.",
	}, []string{"service", "grpc_service", "grpc_method", "grpc_code"})
	grpcRequestsDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "grpc_server_handling_seconds",
		Help:    "Histogram of response latency (seconds) of gRPC requests.",
		Buckets: prometheus.DefBuckets,
	}, []string{"service", "grpc_service", "grpc_method", "grpc_code"})
)

func emitFakeMetrics() {
	for {
		for _, service := range services {
			emitFakeMetricsForService(service)
		}

		time.Sleep(5 * time.Second)
	}
}

func emitFakeMetricsForService(service Service) {
	serviceInfos.With(prometheus.Labels{"service": service.Name, "version": service.Version}).Set(1)

	emitFakeHTTPMetricsForService(service)
	emitFakeGRPCMetricsForService(service)
}

func emitFakeGRPCMetricsForService(service Service) {
	for _, grpcCall := range service.GRPC {
		labels := prometheus.Labels{
			"service":      service.Name,
			"grpc_service": grpcCall.Service,
			"grpc_method":  grpcCall.Method,
			"grpc_code":    grpcCall.Code,
		}

		requestsCount := rand.N(grpcCall.Count[1]) + grpcCall.Count[0]
		grpcRequests.With(labels).Add(float64(requestsCount))

		for i := 0; i < requestsCount; i++ {
			duration := rand.N(grpcCall.Latency[1]) + grpcCall.Latency[0]
			grpcRequestsDuration.With(labels).Observe(float64(duration) / 1000)
		}
	}
}

func emitFakeHTTPMetricsForService(service Service) {
	for _, httpCall := range service.HTTP {
		labels := prometheus.Labels{
			"service": service.Name,
			"code":    strconv.Itoa(httpCall.Code),
			"method":  httpCall.Method,
			"path":    httpCall.Path,
		}

		requestsCount := rand.N(httpCall.Count[1]) + httpCall.Count[0]
		httpRequests.With(labels).Add(float64(requestsCount))

		for i := 0; i < requestsCount; i++ {
			duration := rand.N(httpCall.Latency[1]) + httpCall.Latency[0]
			httpRequestsDuration.With(labels).Observe(float64(duration) / 1000)
		}
	}
}

func main() {
	httpPort := "8080"
	if port := os.Getenv("HTTP_PORT"); port != "" {
		httpPort = port
	}

	go emitFakeMetrics()

	http.Handle("/metrics", promhttp.Handler())

	log.Print(fmt.Sprintf("Listening on :%s...", httpPort))
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
