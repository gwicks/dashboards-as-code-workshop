package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/heatmap"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func httpRequestsTimeseries(service Service) *timeseries.PanelBuilder {
	return timeseriesPanel().
		Title("HTTP Requests").
		Unit("reqps").
		WithTarget(prometheusQuery(fmt.Sprintf("rate(http_requests_total{service=\"%s\"}[$__rate_interval])", service.Name)).
			LegendFormat("{{code}} - {{ method }} {{ path }}"),
		).
		Datasource(prometheusDatasourceRef())
}

func httpLatenciesHeatmap(service Service) *heatmap.PanelBuilder {
	return heatmapPanel().
		Title("HTTP Requests latencies").
		WithTarget(prometheusQuery(fmt.Sprintf("sum(increase(http_requests_duration_seconds_bucket{service=\"%s\"}[$__rate_interval])) by (le)", service.Name)).
			Format(prometheus.PromQueryFormatHeatmap),
		).
		Datasource(prometheusDatasourceRef())
}
