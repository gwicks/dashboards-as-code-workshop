package lab;

import com.grafana.foundation.dashboard.*;
import lab.catalog.Service;

import java.util.List;
import java.util.Map;

import static lab.Common.*;

public class Overview {
    public static DashboardBuilder forService(Service service) {
        DashboardBuilder builder = new DashboardBuilder(service.name+" service overview").
                uid(service.name+"-overview").
                tags(List.of("generated", service.name)).
                readonly().
                time(new DashboardDashboardTimeBuilder().
                        from("now-30m").
                        to("now")
                ).
                tooltip(DashboardCursorSync.CROSSHAIR).
                refresh("10s").
                link(new DashboardLinkBuilder("GitHub Repository").
                        url(service.repositoryUrl).
                        targetBlank(true)
                ).
                withVariable(new TextBoxVariableBuilder("logs_filter").
                        label("Logs filter")
                ).
                withVariable(logsLevelsVariable(service));

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

        return builder;
    }

    private static QueryVariableBuilder logsLevelsVariable(Service service) {
        return new QueryVariableBuilder("logs_level").
                label("Logs level").
                datasource(lokiDatasourceRef()).
                query(StringOrMap.createMap(Map.of(
                        "label", "level",
                        "stream", "{service=\""+service.name+"\"}",
                        "type", 1
                ))).
                refresh(VariableRefresh.ON_TIME_RANGE_CHANGED).
                includeAll(true).
                allValue(".*");
    }
}
