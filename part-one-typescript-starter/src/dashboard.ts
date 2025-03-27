import {
    DashboardBuilder,
    DashboardCursorSync,
} from '@grafana/grafana-foundation-sdk/dashboard';
import * as logs from '@grafana/grafana-foundation-sdk/logs';
import * as stat from '@grafana/grafana-foundation-sdk/stat';
import * as text from '@grafana/grafana-foundation-sdk/text';
import * as timeseries from '@grafana/grafana-foundation-sdk/timeseries';
import {
    logPanel,
    statPanel,
    textPanel,
    timeseriesPanel
} from './common';

export const exampleDashboard = (): DashboardBuilder => {
    const builder = new DashboardBuilder(`Test dashboard`)
        .uid(`test-dashboard`)
        .tags(['test', 'generated'])
        .readonly()
        .tooltip(DashboardCursorSync.Crosshair)
        .refresh('10s')
        .time({ from: 'now-30m', to: 'now' })
    ;

    builder
        .withPanel(prometheusVersionStat())
        .withPanel(descriptionText())
        .withPanel(unfilteredLogs())
        .withPanel(prometheusGoroutinesTimeseries())
    ;

    return builder;
};

export const prometheusVersionStat = (): stat.PanelBuilder => {
    return statPanel()
	    // TODO: configure the panel
    ;
};

export const descriptionText = (): text.PanelBuilder => {
    return textPanel(``)
        // TODO: configure the panel
    ;
};

export const unfilteredLogs = (): logs.PanelBuilder => {
    return logPanel()
        // TODO: configure the panel
    ;
};

export const prometheusGoroutinesTimeseries = (): timeseries.PanelBuilder => {
    return timeseriesPanel()
        // TODO: configure the panel
    ;
};
