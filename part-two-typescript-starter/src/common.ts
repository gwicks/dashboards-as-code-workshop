import {
    DataSourceRef,
    FieldColorBuilder,
    FieldColorModeId,
    ThresholdsConfigBuilder,
    ThresholdsMode,
} from '@grafana/grafana-foundation-sdk/dashboard';
import * as common from '@grafana/grafana-foundation-sdk/common';
import * as heatmap from '@grafana/grafana-foundation-sdk/heatmap';
import * as logs from '@grafana/grafana-foundation-sdk/logs';
import * as loki from '@grafana/grafana-foundation-sdk/loki';
import * as prometheus from '@grafana/grafana-foundation-sdk/prometheus';
import * as stat from '@grafana/grafana-foundation-sdk/stat';
import * as text from '@grafana/grafana-foundation-sdk/text';
import * as timeseries from '@grafana/grafana-foundation-sdk/timeseries';

// This file contains a series of utility functions to simplify the creation
// of panels while providing a consistent "look and feel".

// Creates a pre-configured stat panel.
export const statPanel = (): stat.PanelBuilder => {
    return new stat.PanelBuilder()
        .colorScheme(new FieldColorBuilder().mode(FieldColorModeId.Thresholds))
        .graphMode(common.BigValueGraphMode.Area)
        .colorMode(common.BigValueColorMode.Value)
        .justifyMode(common.BigValueJustifyMode.Auto)
        .textMode(common.BigValueTextMode.Auto)
        .orientation(common.VizOrientation.Auto)
        .thresholds(
            new ThresholdsConfigBuilder()
                .mode(ThresholdsMode.Absolute)
                .steps([
                    {value: null, color: 'green'},
                    {value: 80, color: 'red'},
                ])
        )
    ;
};

// Creates a text panel pre-configured for markdown content.
export const textPanel = (content: string): text.PanelBuilder => {
    return new text.PanelBuilder()
        .mode(text.TextMode.Markdown)
        .content(content)
    ;
};

// Creates a pre-configured timeseries panel.
export const timeseriesPanel = (): timeseries.PanelBuilder => {
    return new timeseries.PanelBuilder()
        .lineWidth(1)
        .pointSize(5)
        .fillOpacity(20)
        .gradientMode(common.GraphGradientMode.Opacity)
        .legend(
            new common.VizLegendOptionsBuilder()
                .displayMode(common.LegendDisplayMode.List)
                .placement(common.LegendPlacement.Bottom)
                .showLegend(true)
        )
        .tooltip(
            new common.VizTooltipOptionsBuilder()
                .mode(common.TooltipDisplayMode.Single)
                .sort(common.SortOrder.None)
        )
        .colorScheme(new FieldColorBuilder().mode(FieldColorModeId.PaletteClassic))
        .thresholdsStyle(
            new common.GraphThresholdsStyleConfigBuilder()
                .mode(common.GraphThresholdsStyleMode.Off)
        )
    ;
};

// Creates a pre-configured heatmap panel.
export const heatmapPanel = (): heatmap.PanelBuilder => {
    return new heatmap.PanelBuilder()
        .color(
            new heatmap.HeatmapColorOptionsBuilder()
                .mode(heatmap.HeatmapColorMode.Scheme)
                .scheme('RdYlBu')
                .fill('dark-orange')
                .scale(heatmap.HeatmapColorScale.Exponential)
                .exponent(0.5)
                .steps(64)
                .reverse(false)
        )
        .filterValues(new heatmap.FilterValueRangeBuilder().le(1e-09))
        .rowsFrame(new heatmap.RowsHeatmapOptionsBuilder().layout(common.HeatmapCellLayout.Auto))
        .showValue(common.VisibilityMode.Auto)
        .cellValues(new heatmap.CellValuesBuilder())
        .yAxis(
            new heatmap.YAxisConfigBuilder()
                .unit('s')
                .reverse(false)
                .axisPlacement(common.AxisPlacement.Left)
        )
        .showLegend()
        .mode(common.TooltipDisplayMode.Single)
        .hideYHistogram()
        .showColorScale(true)
        .scaleDistribution(
            new common.ScaleDistributionConfigBuilder()
                .type(common.ScaleDistribution.Linear)
        )
        .hideFrom(
            new common.HideSeriesConfigBuilder()
                .tooltip(false)
                .legend(false)
                .viz(false)
        )
    ;
};

// Creates a pre-configured logs panel.
export const logPanel = (): logs.PanelBuilder => {
    return new logs.PanelBuilder()
        .datasource(lokiDatasourceRef())
        .showTime(true)
        .sortOrder(common.LogsSortOrder.Descending)
        .enableLogDetails(true)
    ;
};

// Creates a Prometheus query pre-configured for range vectors.
export const prometheusQuery = (expression: string): prometheus.DataqueryBuilder => {
	return new prometheus.DataqueryBuilder()
        .expr(expression)
        .range()
        .format(prometheus.PromQueryFormat.TimeSeries)
        .legendFormat('__auto')
    ;
};

// Creates a Prometheus query pre-configured for instant vectors and table data
// formatting.
export const instantPrometheusQuery = (expression: string): prometheus.DataqueryBuilder => {
	return new prometheus.DataqueryBuilder()
        .expr(expression)
        .instant()
        .format(prometheus.PromQueryFormat.Table)
        .legendFormat('__auto')
    ;
};

// Creates a Loki query pre-configured for range vectors.
export const lokiQuery = (expression: string): loki.DataqueryBuilder => {
	return new loki.DataqueryBuilder()
        .expr(expression)
        .queryType('range')
        .legendFormat('__auto')
    ;
};

// Returns a reference to the Prometheus datasource used by the workshop stack.
export const prometheusDatasourceRef = (): DataSourceRef => {
	return {
		type: 'prometheus',
		uid:  'prometheus',
	};
};

// Returns a reference to the Loki datasource used by the workshop stack.
export const lokiDatasourceRef = (): DataSourceRef => {
	return {
		type: 'loki',
		uid:  'loki',
	};
};
