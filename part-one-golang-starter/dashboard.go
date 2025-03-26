package main

import (
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
	return statPanel()
	// TODO: configure the panel
}

func descriptionText() *text.PanelBuilder {
	return textPanel("")
	// TODO: configure the panel
}

func unfilteredLogs() *logs.PanelBuilder {
	return logPanel()
	// TODO: configure the panel
}

func prometheusGoroutinesTimeseries() *timeseries.PanelBuilder {
	return timeseriesPanel()
	// TODO: configure the panel
}
