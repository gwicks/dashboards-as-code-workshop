package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/heatmap"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func grpcRequestsTimeseries(service Service) *timeseries.PanelBuilder {
	return timeseriesPanel().
		Title("gRPC Requests").
		Unit("reqps").
		WithTarget(prometheusQuery(fmt.Sprintf("rate(grpc_server_handled_total{service=\"%s\"}[$__rate_interval])", service.Name)).
			LegendFormat("{{ grpc_method }} – {{ grpc_code }}"),
		).
		Datasource(prometheusDatasourceRef())
}

func grpcLatenciesHeatmap(service Service) *heatmap.PanelBuilder {
	return heatmapPanel().
		Title("gRPC Requests latencies").
		WithTarget(prometheusQuery(fmt.Sprintf("sum(increase(grpc_server_handling_seconds_bucket{service=\"%s\"}[$__rate_interval])) by (le)", service.Name)).
			Format(prometheus.PromQueryFormatHeatmap),
		).
		Datasource(prometheusDatasourceRef())
}
