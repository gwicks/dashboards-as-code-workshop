package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

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

	// TODO:
	// - "Version" panel. Height: 4, Span: 4
	// - "service description" panel. Height: 4, Span: 4
	// - "Logs volume" panel. Height: 4, Span: 16

	// gRPC row, if relevant
	// TODO: define a "gRPC" row with the following panels:
	// - "gRPC Requests" panel. Height: 8
	// - "gRPC Requests latencies" panel. Height: 8
	// - "GRPC Logs" panel. Height: 8, Span: 24

	// HTTP row, if relevant
	// TODO: define an "HTTP" row with the following panels:
	// - "HTTP Requests" panel. Height: 8
	// - "HTTP Requests latencies" panel. Height: 8
	// - "HTTP Logs" panel. Height: 8, Span: 24

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
