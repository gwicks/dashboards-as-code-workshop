import { DataSourceRef } from '@grafana/grafana-foundation-sdk/dashboard';
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
    return new stat.PanelBuilder();
};

// Creates a text panel pre-configured for markdown content.
export const textPanel = (_content: string): text.PanelBuilder => {
    return new text.PanelBuilder()
        // TODO: configure default options for text panels
    ;
};

// Creates a pre-configured timeseries panel.
export const timeseriesPanel = (): timeseries.PanelBuilder => {
    return new timeseries.PanelBuilder()
        // TODO: configure default options for timeseries panels
    ;
};

// Creates a pre-configured logs panel.
export const logPanel = (): logs.PanelBuilder => {
    return new logs.PanelBuilder()
        // TODO: configure default options for logs panels
    ;
};

// Creates a pre-configured heatmap panel.
export const heatmapPanel = (): heatmap.PanelBuilder => {
    return new heatmap.PanelBuilder()
        .color(
            new heatmap.HeatmapColorOptionsBuilder()
                .mode(heatmap.HeatmapColorMode.Scheme)
                .scheme('RdYlBu')
                .scale(heatmap.HeatmapColorScale.Exponential)
                .steps(64)
        )
        .filterValues(new heatmap.FilterValueRangeBuilder().le(1e-09))
        .yAxis(new heatmap.YAxisConfigBuilder().unit('s'))
        .mode(common.TooltipDisplayMode.Single)
        .scaleDistribution(
            new common.ScaleDistributionConfigBuilder()
                .type(common.ScaleDistribution.Linear)
        )
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
