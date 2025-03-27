from grafana_foundation_sdk.builders import dashboard, logs, stat, text
from grafana_foundation_sdk.models.dashboard import DashboardCursorSync

from .common import (
    timeseries_panel,
    log_panel,
    stat_panel,
    text_panel,
)

def example_dashboard() -> dashboard.Dashboard:
    builder = (
        dashboard.Dashboard("Test dashboard")
        .uid("test-dashboard")
        .tags(["test", "generated"])
        .readonly()
        .time("now-30m", "now")
        .tooltip(DashboardCursorSync.CROSSHAIR)
        .refresh("10s")
    )

    builder.with_panel(prometheus_version_stat())
    builder.with_panel(description_text())
    builder.with_panel(unfiltered_logs())
    builder.with_panel(prometheus_goroutines_timeseries())

    return builder

def prometheus_version_stat() -> stat.Panel:
    return (
        stat_panel()
	    # TODO: configure the panel
    )

def description_text() -> text.Panel:
    return (
        text_panel("")
	    # TODO: configure the panel
    )

def unfiltered_logs() -> logs.Panel:
    return (
        log_panel()
	    # TODO: configure the panel
    )

def prometheus_goroutines_timeseries() -> logs.Panel:
    return (
        timeseries_panel()
	    # TODO: configure the panel
    )