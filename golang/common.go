package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/heatmap"
	"github.com/grafana/grafana-foundation-sdk/go/logs"
	"github.com/grafana/grafana-foundation-sdk/go/loki"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/text"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func statPanel() *stat.PanelBuilder {
	return stat.NewPanelBuilder().
		ColorScheme(dashboard.NewFieldColorBuilder().Mode(dashboard.FieldColorModeIdThresholds)).
		GraphMode(common.BigValueGraphModeArea).
		ColorMode(common.BigValueColorModeValue).
		JustifyMode(common.BigValueJustifyModeAuto).
		TextMode(common.BigValueTextModeAuto).
		Orientation(common.VizOrientationAuto).
		Thresholds(dashboard.NewThresholdsConfigBuilder().
			Mode("absolute").
			Steps([]dashboard.Threshold{
				{Color: "green"},
				{Value: cog.ToPtr[float64](80), Color: "red"},
			}),
		)
}

func textPanel() *text.PanelBuilder {
	return text.NewPanelBuilder().
		Mode(text.TextModeMarkdown)
}

func timeseriesPanel() *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().
		LineWidth(1).
		PointSize(5).
		FillOpacity(20).
		GradientMode(common.GraphGradientModeOpacity).
		Legend(common.NewVizLegendOptionsBuilder().
			DisplayMode(common.LegendDisplayModeList).
			Placement(common.LegendPlacementBottom).
			ShowLegend(true),
		).
		Tooltip(common.NewVizTooltipOptionsBuilder().
			Mode(common.TooltipDisplayModeSingle).
			Sort(common.SortOrderNone),
		).
		ColorScheme(dashboard.NewFieldColorBuilder().Mode(dashboard.FieldColorModeIdPaletteClassic)).
		ThresholdsStyle(common.NewGraphThresholdsStyleConfigBuilder().Mode(common.GraphThresholdsStyleModeOff))
}

func heatmapPanel() *heatmap.PanelBuilder {
	return heatmap.NewPanelBuilder().
		Color(heatmap.NewHeatmapColorOptionsBuilder().
			Mode("scheme").
			Scheme("RdYlBu").
			Fill("dark-orange").
			Scale("exponential").
			Exponent(0.5).
			Steps(0x40).
			Reverse(false),
		).
		FilterValues(heatmap.NewFilterValueRangeBuilder().Le(1e-09)).
		RowsFrame(heatmap.NewRowsHeatmapOptionsBuilder().Layout("auto")).
		ShowValue("").
		CellValues(heatmap.NewCellValuesBuilder()).
		YAxis(heatmap.NewYAxisConfigBuilder().
			Unit("s").
			Reverse(false).
			AxisPlacement("left"),
		).
		ShowLegend().
		Mode("single").
		HideYHistogram().
		ShowColorScale(true).
		ScaleDistribution(common.NewScaleDistributionConfigBuilder().Type("linear")).
		HideFrom(common.NewHideSeriesConfigBuilder().
			Tooltip(false).
			Legend(false).
			Viz(false),
		)
}

func logPanel() *logs.PanelBuilder {
	return logs.NewPanelBuilder().
		Datasource(lokiDatasourceRef()).
		ShowTime(true).
		SortOrder(common.LogsSortOrderDescending).
		EnableLogDetails(true)
}

func prometheusQuery(expression string) *prometheus.DataqueryBuilder {
	return prometheus.NewDataqueryBuilder().
		Expr(expression).
		Range().
		Format(prometheus.PromQueryFormatTimeSeries).
		LegendFormat("__auto")
}

func instantPrometheusQuery(expression string) *prometheus.DataqueryBuilder {
	return prometheus.NewDataqueryBuilder().
		Expr(expression).
		Instant().
		Format(prometheus.PromQueryFormatTable).
		LegendFormat("__auto")
}

func lokiQuery(expression string) *loki.DataqueryBuilder {
	return loki.NewDataqueryBuilder().
		Expr(expression).
		QueryType("range").
		LegendFormat("__auto")
}

func prometheusDatasourceRef() dashboard.DataSourceRef {
	return dashboard.DataSourceRef{
		Type: cog.ToPtr[string]("prometheus"),
		Uid:  cog.ToPtr[string]("prometheus"),
	}
}

func lokiDatasourceRef() dashboard.DataSourceRef {
	return dashboard.DataSourceRef{
		Type: cog.ToPtr[string]("loki"),
		Uid:  cog.ToPtr[string]("loki"),
	}
}
