from grafana_foundation_sdk.builders import common, dashboard, heatmap, logs, loki, prometheus, stat, text, timeseries
from grafana_foundation_sdk.models.common import (
    BigValueGraphMode,
    BigValueColorMode,
    BigValueJustifyMode,
    BigValueTextMode,
    VizOrientation,
    LogsSortOrder,
    GraphGradientMode,
    LegendDisplayMode,
    LegendPlacement,
    TooltipDisplayMode,
    SortOrder,
    GraphThresholdsStyleMode,
    HeatmapCellLayout,
    AxisPlacement,
    ScaleDistribution,
)
from grafana_foundation_sdk.models.dashboard import (
    DataSourceRef,
    FieldColorModeId,
    ThresholdsMode,
    Threshold,
)
from grafana_foundation_sdk.models.heatmap import (
    HeatmapColorMode,
    HeatmapColorScale,
)
from grafana_foundation_sdk.models.prometheus import PromQueryFormat
from grafana_foundation_sdk.models.text import TextMode

# This file contains a series of utility functions to simplify the creation
# of panels while providing a consistent "look and feel".


def loki_datasource_ref() -> DataSourceRef:
    """
    Returns a reference to the Loki datasource used by the
    workshop stack.
    """
    return DataSourceRef(type_val="loki", uid="loki")

def prometheus_datasource_ref() -> DataSourceRef:
    """
    Returns a reference to the Prometheus datasource used by the
    workshop stack.
    """
    return DataSourceRef(type_val="prometheus", uid="prometheus")

def stat_panel() -> stat.Panel:
    """
    Creates a pre-configured stat panel.
    """
    return (
        stat.Panel()
        .color_scheme(dashboard.FieldColor().mode(FieldColorModeId.THRESHOLDS))
        .graph_mode(BigValueGraphMode.AREA)
        .color_mode(BigValueColorMode.VALUE)
        .justify_mode(BigValueJustifyMode.AUTO)
        .text_mode(BigValueTextMode.AUTO)
        .orientation(VizOrientation.AUTO)
        .thresholds(
            dashboard.ThresholdsConfig()
            .mode(ThresholdsMode.ABSOLUTE)
            .steps([
                Threshold(color="green"),
                Threshold(color="red", value=80),
            ])
        )
    )

def text_panel(content: str) -> text.Panel:
    """
    Creates a pre-configured text panel.
    """
    return (
        text.Panel()
        .mode(TextMode.MARKDOWN)
        .content(content)
    )

def timeseries_panel() -> timeseries.Panel:
    """
    Creates a pre-configured timeseries panel.
    """
    return (
        timeseries.Panel()
        .line_width(1)
        .point_size(5)
        .fill_opacity(20)
        .gradient_mode(GraphGradientMode.OPACITY)
        .legend(
            common.VizLegendOptions()
            .display_mode(LegendDisplayMode.LIST)
            .placement(LegendPlacement.BOTTOM)
            .show_legend(True)
        )
        .tooltip(
            common.VizTooltipOptions()
            .mode(TooltipDisplayMode.SINGLE)
            .sort(SortOrder.NONE)
        )
        .color_scheme(
            dashboard.FieldColor()
            .mode(FieldColorModeId.PALETTE_CLASSIC)
        )
        .thresholds_style(
            common.GraphThresholdsStyleConfig()
            .mode(GraphThresholdsStyleMode.OFF)
        )
    )

def heatmap_panel() -> timeseries.Panel:
    """
    Creates a pre-configured heatmap panel.
    """
    return (
        heatmap.Panel()
        .color(
            heatmap.HeatmapColorOptions()
            .mode(HeatmapColorMode.SCHEME)
            .scheme("RdYlBu")
            .fill("dark-orange")
            .scale(HeatmapColorScale.EXPONENTIAL)
            .exponent(0.5)
            .steps(64)
            .reverse(False)
        )
        .filter_values(heatmap.FilterValueRange().le(1e-09))
        .rows_frame(heatmap.RowsHeatmapOptions().layout(HeatmapCellLayout.AUTO))
        .y_axis(
            heatmap.YAxisConfig()
            .unit("s")
            .reverse(False)
            .axis_placement(AxisPlacement.LEFT)
        )
        .show_legend()
        .mode(TooltipDisplayMode.SINGLE)
        .hide_y_histogram()
        .show_color_scale(True)
        .scale_distribution(common.ScaleDistributionConfig().type_val(ScaleDistribution.LINEAR))
        .hide_from(
            common.HideSeriesConfig()
            .tooltip(False)
            .legend(False)
            .viz(False)
        )
    )


def log_panel() -> logs.Panel:
    """
    Creates a pre-configured logs panel.
    """
    return (
        logs.Panel()
        .datasource(loki_datasource_ref())
        .show_time(True)
        .sort_order(LogsSortOrder.DESCENDING)
        .enable_log_details(True)
    )

def loki_query(expression: str) -> loki.Dataquery:
    """
    Creates a Loki query pre-configured for range vectors.
    """
    return (
        loki.Dataquery()
        .expr(expression)
        .query_type("range")
        .legend_format("__auto")
    )

def prometheus_query(expression: str) -> prometheus.Dataquery:
    """
    Creates a Prometheus query pre-configured for range vectors.
    """
    return (
        prometheus.Dataquery()
        .expr(expression)
        .range_val()
        .format_val(PromQueryFormat.TIME_SERIES)
        .legend_format("__auto")
    )

def instant_prometheus_query(expression: str) -> prometheus.Dataquery:
    """
    Creates a Prometheus query pre-configured for instant
    vectors and table data formatting.
    """
    return (
        prometheus.Dataquery()
        .expr(expression)
        .instant()
        .format_val(PromQueryFormat.TABLE)
        .legend_format("__auto")
    )
