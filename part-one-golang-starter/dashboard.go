package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/logs"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/text"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func testDashboard() *dashboard.DashboardBuilder {
	builder := dashboard.NewDashboardBuilder("Test dashboard").
		Uid("test-dashboard").
		Tags([]string{"test", "generated"}).
		Time("now-30m", "now").
		Tooltip(dashboard.DashboardCursorSyncCrosshair).
		Refresh("10s")

	builder.
		WithPanel(prometheusVersionStat()).
		WithPanel(descriptionText()).
		WithPanel(unfilteredLogs()).
		WithPanel(prometheusGoroutinesTimeseries())

	return builder
}

func prometheusVersionStat() *stat.PanelBuilder {
	reduceOptions := common.NewReduceDataOptionsBuilder().Calcs([]string{"last"}).Fields("/^version$/")
	return statPanel().Title("Prometheus version").Transparent(true).ReduceOptions(reduceOptions).Datasource(prometheusDatasourceRef()).WithTarget(instantPrometheusQuery("prometheus_build_info{}"))
	// TODO: configure the panel
	//
	//  * `title`: `Prometheus version`
	//  * `transparent` set to `true`
	//  * `reduce` options
	//    * `calcs` set to `["last"]`
	//    * `fields` set to `/^version$/`
	//  * Instant Prometheus query: `prometheus_build_info{}` (see instantPrometheusQuery())
	//  * `datasource`: Prometheus datasource ref (see prometheusDatasourceRef())
	//
	// See: https://grafana.github.io/grafana-foundation-sdk/v11.6.x+cog-v0.0.x/go/Reference/stat/builder-PanelBuilder/
}

func descriptionText() *text.PanelBuilder {
	return textPanel("Text panels are supported too! Even with *markdown* text :)").Transparent(true)
	// TODO: configure the panel
	//
	//  * `content`: `Text panels are supported too! Even with *markdown* text :)`
	//  * `transparent` set to `true`
	//
	// See: https://grafana.github.io/grafana-foundation-sdk/v11.6.x+cog-v0.0.x/go/Reference/text/builder-PanelBuilder/
}

func unfilteredLogs() *logs.PanelBuilder {
	return logPanel().Title("Logs").Datasource(lokiDatasourceRef()).WithTarget(lokiQuery("{job=\"app_logs\"}"))
	// TODO: configure the panel
	//
	//  * `title`: `Logs`
	//  * Loki query: `{job="app_logs"}` (see lokiQuery())
	//  * `datasource`: loki datasource ref (see lokiDatasourceRef())
	//
	// See: https://grafana.github.io/grafana-foundation-sdk/v11.6.x+cog-v0.0.x/go/Reference/logs/builder-PanelBuilder/
}

func prometheusGoroutinesTimeseries() *timeseries.PanelBuilder {
	return timeseriesPanel().Title("Prometheus goroutines").Datasource(prometheusDatasourceRef()).WithTarget(prometheusQuery("go_goroutines{job=\"prometheus\"}"))
	// TODO: configure the panel
	//
	//  * `title`: `Prometheus goroutines`
	//  * Prometheus query: `go_goroutines{job="prometheus"}` (see prometheusQuery())
	//  * `datasource`: prometheus datasource ref (see prometheusDatasourceRef())
	//
	// See: https://grafana.github.io/grafana-foundation-sdk/v11.6.x+cog-v0.0.x/go/Reference/timeseries/builder-PanelBuilder/
}
