package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/heatmap"
	"github.com/grafana/grafana-foundation-sdk/go/logs"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

// HTTP row, if relevant
// TODO: define an "HTTP" row with the following panels:
// * "HTTP Requests" panel
//   - type: `timeseries`
//   - query: `rate(http_requests_total{service="[service_name]"}[$__rate_interval])`
//   - datasource: Prometheus datasource ref (see prometheusDatasourceRef())
//   - query legend format: `{{code}} - {{ method }} {{ path }}`
//   - unit: requests per second (reqps)
//   - height: 8
//   - span: 12
// * "HTTP Requests latencies" panel
//   - type: `heatmap`
//   - query: `sum(increase(http_requests_duration_seconds_bucket{service="[service_name]"}[$__rate_interval])) by (le)`
//   - query format: `heatmap`
//   - datasource: Prometheus datasource ref (see prometheusDatasourceRef())
//   - height: 8
//   - span: 12
// * "HTTP Logs" panel
//   - type: `logs`
//   - query: `{service="[service_name]", source="http", level=~"$logs_level"} |~ "$logs_filter"`
//   - height: 8
//   - span: 24

func httpRequestsForSvc(service Service) *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().Title("HTTP Requests").
		Datasource(prometheusDatasourceRef()).
		Unit(units.RequestsPerSecond).
		WithTarget(prometheusQuery(fmt.Sprintf("rate(http_requests_total{service=\"%s\"}[$__rate_interval])", service.Name)).
			LegendFormat("{{code}} - {{ method }} {{ path }}"))

}
func httpRequestLatenciesForSvc(service Service) *heatmap.PanelBuilder {
	return heatmap.NewPanelBuilder().
		Title("HTTP Requests latencies").
		WithTarget(prometheusQuery(fmt.Sprintf("sum(increase(http_requests_duration_seconds_bucket{service=\"%s\"}[$__rate_interval])) by (le)", service.Name)).
			Format(prometheus.PromQueryFormatHeatmap),
		).
		Datasource(prometheusDatasourceRef())
}

func httpLogsForSvc(service Service) *logs.PanelBuilder {
	logsQueryStr := fmt.Sprintf(`{service="%s", source="http", level=~"$logs_level"} |~ "$logs_filter"`, service.Name)
	return logs.NewPanelBuilder().
		Datasource(lokiDatasourceRef()).
		ShowTime(true).
		SortOrder(common.LogsSortOrderDescending).
		EnableLogDetails(true).Title("HTTP Logs").WithTarget(lokiQuery(logsQueryStr))
}
