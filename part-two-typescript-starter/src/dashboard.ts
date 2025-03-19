import {
    DashboardBuilder,
    DashboardCursorSync,
    DashboardLinkBuilder,
    QueryVariableBuilder,
    TextBoxVariableBuilder,
    VariableRefresh,
} from '@grafana/grafana-foundation-sdk/dashboard';
import { Service } from './catalog';
import { lokiDatasourceRef } from './common';

export const dashboardForService = (service: Service): DashboardBuilder => {
    const builder = new DashboardBuilder(`${service.name} service overview`)
        .uid(`${service.name}-overview`)
        .tags([service.name, 'generated'])
        .readonly()
        .tooltip(DashboardCursorSync.Crosshair)
        .refresh('10s')
        .time({ from: 'now-30m', to: 'now' })
        .link(
            new DashboardLinkBuilder('GitHub Repository')
                .url(service.github)
                .targetBlank(true)
        )
        .withVariable(new TextBoxVariableBuilder('logs_filter').label('Logs filter'))
        .withVariable(logLevelsVariable(service))
    ;

    // TODO:
    // - "Version" panel. Height: 4, Span: 4
    // - "service description" panel. Height: 4, Span: 4
    // - "Logs volume" panel. Height: 4, Span: 16

	// gRPC row, if relevant
    if (service.has_grpc) {
        // TODO: define a "gRPC" row with the following panels:
        // - "gRPC Requests" panel. Height: 8
        // - "gRPC Requests latencies" panel. Height: 8
        // - "GRPC Logs" panel. Height: 8, Span: 24
    }

	// HTTP row, if relevant
    if (service.has_grpc) {
        // TODO: define an "HTTP" row with the following panels:
        // - "HTTP Requests" panel. Height: 8
        // - "HTTP Requests latencies" panel. Height: 8
        // - "HTTP Logs" panel. Height: 8, Span: 24
    }

    return builder;
};

const logLevelsVariable = (service: Service): QueryVariableBuilder => {
    return new QueryVariableBuilder('logs_level')
        .label('Logs level')
        .datasource(lokiDatasourceRef())
        .query({
            type: 1,
            label: 'level',
            stream: `{service="${service.name}"}`,
        })
        .refresh(VariableRefresh.OnTimeRangeChanged)
        .includeAll(true)
        .allValue('.*');
};
