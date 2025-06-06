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
	// * "Version" panel
	//   - type: `stat`
	//   - query: `app_infos{service="[service_name]"}`
	//	   - instant: true
	//   - datasource: Prometheus datasource ref (see Common.prometheusDatasourceRef())
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
	//   - datasource: Prometheus datasource ref (see Common.prometheusDatasourceRef())
	//   - query legend format: `{{ grpc_method }} – {{ grpc_code }}`
	//   - unit: requests per second (reqps)
	//   - height: 8
	//   - span: 12
	// * "gRPC Requests latencies" panel
	//   - type: `heatmap`
	//   - query: `sum(increase(grpc_server_handling_seconds_bucket{service="[service_name]"}[$__rate_interval])) by (le)`
	//   - query format: `heatmap`
	//   - datasource: Prometheus datasource ref (see Common.prometheusDatasourceRef())
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
	//   - datasource: Prometheus datasource ref (see Common.prometheusDatasourceRef())
	//   - query legend format: `{{code}} - {{ method }} {{ path }}`
	//   - unit: requests per second (reqps)
	//   - height: 8
	//   - span: 12
	// * "HTTP Requests latencies" panel
	//   - type: `heatmap`
	//   - query: `sum(increase(http_requests_duration_seconds_bucket{service="[service_name]"}[$__rate_interval])) by (le)`
	//   - query format: `heatmap`
	//   - datasource: Prometheus datasource ref (see Common.prometheusDatasourceRef())
	//   - height: 8
	//   - span: 12
	// * "HTTP Logs" panel
	//   - type: `logs`
	//   - query: `{service="[service_name]", source="http", level=~"$logs_level"} |~ "$logs_filter"`
	//   - height: 8
	//   - span: 24

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
