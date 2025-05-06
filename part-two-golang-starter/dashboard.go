package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/text"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func verisonPanelForSvc(serviceName string) *stat.PanelBuilder {
	promQuery := fmt.Sprintf("app_infos{service=\"%s\"}", serviceName)
	versionPanelReduceOpts := common.NewReduceDataOptionsBuilder().Calcs([]string{"last"}).Fields("/^version$/").Values(false)
	return stat.NewPanelBuilder().Title("Version").
		Datasource(prometheusDatasourceRef()).
		WithTarget(instantPrometheusQuery(promQuery)).
		Transparent(true).ReduceOptions(versionPanelReduceOpts)
}

func serviceDescPanel(mdText string) *text.PanelBuilder {
	return text.NewPanelBuilder().Content(mdText).Mode(text.TextModeMarkdown).Transparent(true)
}

func logsVolumePanelForSvc(serviceName string) *timeseries.PanelBuilder {
	// * "Logs volume" panel. Height: 4, Span: 16
	//   - type: `timeseries`
	//   - query: `sum by (level) (count_over_time({service="[service_name]", level=~"$logs_level"} |~ "$logs_filter" [$__auto]))`
	//     - legend format: `{{level}}`
	//   - stacking mode: normal
	//   - `legend` options:
	//     - displayMode: `list`
	//     - placement: `bottom`
	//     - showLegend: `true`
	//   - draw style: bars
	//   - override by name:
	//     - name: "INFO"
	//     - value: `color = {"mode": "fixed", "fixedColor": "green"}`
	//   - override by name:
	//     - name: "ERROR"
	//     - value: `color = {"mode": "fixed", "fixedColor": "red"}`
	//   - height: 4
	//   - span: 16
	lokiQueryStr := fmt.Sprintf("sum by (level) (count_over_time({service=\"%s\", level=~\"$logs_level\"} |~ \"$logs_filter\" [$__auto]))", serviceName)
	return timeseries.NewPanelBuilder().
		Title("Logs volume").
		Datasource(lokiDatasourceRef()).
		WithTarget(lokiQueryWithLegendFmt(lokiQueryStr, "{{level}}")).
		Stacking(common.NewStackingConfigBuilder().Mode(common.StackingModeNormal)).
		DrawStyle(common.GraphDrawStyleBars).
		Legend(common.NewVizLegendOptionsBuilder().
			Placement(common.LegendPlacementRight).
			ShowLegend(true).
			DisplayMode(common.LegendDisplayModeList),
		).
		OverrideByName("INFO", []dashboard.DynamicConfigValue{
			{
				Id:    "color",
				Value: map[string]any{"mode": "fixed", "fixedColor": "green"},
			},
		}).
		OverrideByName("ERROR", []dashboard.DynamicConfigValue{
			{
				Id:    "color",
				Value: map[string]any{"mode": "fixed", "fixedColor": "red"},
			},
		})
}

func dashboardForService(service Service) *dashboard.DashboardBuilder {
	builder := dashboard.NewDashboardBuilder(fmt.Sprintf("%s service overview", service.Name)).
		Uid(fmt.Sprintf("%s-overview", service.Name)).
		Tags([]string{service.Name, "generated"}).
		Readonly().
		Time("now-30m", "now").
		Tooltip(dashboard.DashboardCursorSyncCrosshair).
		Refresh("10s").
		Link(dashboard.NewDashboardLinkBuilder("GitHub Repository").
			Type(dashboard.DashboardLinkTypeLink).
			Url(service.RepositoryURL).
			TargetBlank(true),
		).
		WithVariable(dashboard.NewTextBoxVariableBuilder("logs_filter").
			Label("Logs filter"),
		).
		WithVariable(logLevelsVariable(service))

	builder.
		WithPanel(verisonPanelForSvc(service.Name).Height(4).Span(4)).
		WithPanel(serviceDescPanel(service.Description).Height(4).Span(4)).
		WithPanel(logsVolumePanelForSvc(service.Name).Height(4).Span(16))

	if service.HasGRPC {
		builder.WithRow(dashboard.NewRowBuilder("gRPC")).
			WithPanel(grpcRequestsTimeseries(service).Height(8)).
			WithPanel(grpcLatenciesHeatmap(service).Height(8)).
			WithPanel(grpcLogsPanel(service).Height(8).Span(24))
	}

	if service.HasHTTP {
		builder.WithRow(dashboard.NewRowBuilder("HTTP")).
			WithPanel(httpRequestsForSvc(service).Height(8).Span(12)).
			WithPanel(httpRequestLatenciesForSvc(service).Height(8).Span(12)).
			WithPanel(httpLogsForSvc(service).Height(8).Span(24))
	}
	// TODO:
	// * "Version" panel
	//   - type: `stat`
	//   - query: `app_infos{service="[service_name]"}`
	//	   - instant: true
	//   - datasource: Prometheus datasource ref (see prometheusDatasourceRef())
	//   - transparent: true
	//   - reduce options:
	//     - values: false
	//     - calcs: ["last"]
	//     - fields: "/^version$/"
	//   - height: 4
	//   - span: 4
	// * "service description" panel
	//   - type: `text`
	//   - transparent: true
	//   - height: 4
	//   - span: 4
	// * "Logs volume" panel. Height: 4, Span: 16
	//   - type: `timeseries`
	//   - query: `sum by (level) (count_over_time({service="[service_name]", level=~"$logs_level"} |~ "$logs_filter" [$__auto]))`
	//     - legend format: `{{level}}`
	//   - stacking mode: normal
	//   - `legend` options:
	//     - displayMode: `list`
	//     - placement: `bottom`
	//     - showLegend: `true`
	//   - draw style: bars
	//   - override by name:
	//     - name: "INFO"
	//     - value: `color = {"mode": "fixed", "fixedColor": "green"}`
	//   - override by name:
	//     - name: "ERROR"
	//     - value: `color = {"mode": "fixed", "fixedColor": "red"}`
	//   - height: 4
	//   - span: 16

	// gRPC row, if relevant
	// TODO: define a "gRPC" row with the following panels:
	// * "gRPC Requests" panel
	//   - type: `timeseries`
	//   - query: `rate(grpc_server_handled_total{service="[service_name]"}[$__rate_interval])`
	//   - datasource: Prometheus datasource ref (see prometheusDatasourceRef())
	//   - query legend format: `{{ grpc_method }} – {{ grpc_code }}`
	//   - unit: requests per second (reqps)
	//   - height: 8
	//   - span: 12
	// * "gRPC Requests latencies" panel
	//   - type: `heatmap`
	//   - query: `sum(increase(grpc_server_handling_seconds_bucket{service="[service_name]"}[$__rate_interval])) by (le)`
	//   - query format: `heatmap`
	//   - datasource: Prometheus datasource ref (see prometheusDatasourceRef())
	//   - height: 8
	//   - span: 12
	// * "GRPC Logs" panel
	//   - type: `logs`
	//   - query: `{service="[service_name]", source="grpc", level=~"$logs_level"} |~ "$logs_filter"`
	//   - height: 8
	//   - span: 24

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

	return builder
}

func logLevelsVariable(service Service) *dashboard.QueryVariableBuilder {
	return dashboard.NewQueryVariableBuilder("logs_level").
		Label("Logs level").
		Datasource(lokiDatasourceRef()).
		Query(dashboard.StringOrMap{
			Map: map[string]any{
				"label":  "level",
				"stream": fmt.Sprintf(`{service="%s"}`, service.Name),
				"type":   1,
			},
		}).
		Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
		IncludeAll(true).
		AllValue(".*")
}
