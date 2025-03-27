from grafana_foundation_sdk.builders import dashboard
from grafana_foundation_sdk.models.dashboard import (
    DashboardCursorSync,
    VariableRefresh,
)

from .catalog import Service
from .common import loki_datasource_ref

def dashboard_for_service(service: Service) -> dashboard.Dashboard:
    builder = (
        dashboard.Dashboard(f"{service.name} service overview")
        .uid(f"{service.name}-overview")
        .tags([service.name, "generated"])
        .readonly()
        .time("now-30m", "now")
        .tooltip(DashboardCursorSync.CROSSHAIR)
        .refresh("10s")
        .link(
            dashboard.DashboardLink("GitHub Repository")
            .url(service.repository_url)
            .target_blank(True)
        )
        .with_variable(
            dashboard.TextBoxVariable("logs_filter")
            .label("Logs filter")
        )
        .with_variable(log_levels_variable(service))
    )

    # TODO:
    # - "Version" panel. Height: 4, Span: 4
    # - "service description" panel. Height: 4, Span: 4
    # - "Logs volume" panel. Height: 4, Span: 16

    # gRPC row, if relevant
    # TODO: define a "gRPC" row with the following panels:
    # - "gRPC Requests" panel. Height: 8
    # - "gRPC Requests latencies" panel. Height: 8
    # - "GRPC Logs" panel. Height: 8, Span: 24

    # HTTP row, if relevant
    # TODO: define an "HTTP" row with the following panels:
    # - "HTTP Requests" panel. Height: 8
    # - "HTTP Requests latencies" panel. Height: 8
    # - "HTTP Logs" panel. Height: 8, Span: 24

    return builder

def log_levels_variable(service: Service) -> dashboard.QueryVariable:
    return (
        dashboard.QueryVariable("logs_level")
        .label("Logs level")
        .datasource(loki_datasource_ref())
        .query({
            'label': 'level',
            'stream': '{service="%s"}'%service.name,
            'type': 1,
        })
        .refresh(VariableRefresh.ON_TIME_RANGE_CHANGED)
        .include_all(True)
        .all_value(".*")
    )