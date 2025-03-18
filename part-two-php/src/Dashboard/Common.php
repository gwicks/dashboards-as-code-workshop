<?php

namespace App\Dashboard;

use Grafana\Foundation\Common\AxisPlacement;
use Grafana\Foundation\Common\BigValueColorMode;
use Grafana\Foundation\Common\BigValueGraphMode;
use Grafana\Foundation\Common\BigValueJustifyMode;
use Grafana\Foundation\Common\BigValueTextMode;
use Grafana\Foundation\Common\GraphGradientMode;
use Grafana\Foundation\Common\GraphThresholdsStyleConfigBuilder;
use Grafana\Foundation\Common\GraphThresholdsStyleMode;
use Grafana\Foundation\Common\HeatmapCellLayout;
use Grafana\Foundation\Common\HideSeriesConfigBuilder;
use Grafana\Foundation\Common\LegendDisplayMode;
use Grafana\Foundation\Common\LegendPlacement;
use Grafana\Foundation\Common\LogsSortOrder;
use Grafana\Foundation\Common\ScaleDistribution;
use Grafana\Foundation\Common\ScaleDistributionConfigBuilder;
use Grafana\Foundation\Common\SortOrder;
use Grafana\Foundation\Common\TooltipDisplayMode;
use Grafana\Foundation\Common\VisibilityMode;
use Grafana\Foundation\Common\VizLegendOptionsBuilder;
use Grafana\Foundation\Common\VizOrientation;
use Grafana\Foundation\Common\VizTooltipOptionsBuilder;
use Grafana\Foundation\Dashboard\DataSourceRef;
use Grafana\Foundation\Dashboard\FieldColorBuilder;
use Grafana\Foundation\Dashboard\FieldColorModeId;
use Grafana\Foundation\Dashboard\ThresholdsConfigBuilder;
use Grafana\Foundation\Dashboard\ThresholdsMode;
use Grafana\Foundation\Heatmap;
use Grafana\Foundation\Heatmap\CellValuesBuilder;
use Grafana\Foundation\Heatmap\FilterValueRangeBuilder;
use Grafana\Foundation\Heatmap\HeatmapColorMode;
use Grafana\Foundation\Heatmap\HeatmapColorOptionsBuilder;
use Grafana\Foundation\Heatmap\HeatmapColorScale;
use Grafana\Foundation\Heatmap\RowsHeatmapOptionsBuilder;
use Grafana\Foundation\Heatmap\YAxisConfigBuilder;
use Grafana\Foundation\Logs;
use Grafana\Foundation\Loki;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Prometheus\PromQueryFormat;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Text;
use Grafana\Foundation\Text\TextMode;
use Grafana\Foundation\Timeseries;

class Common
{
    public static function lokiDatasourceRef(): DataSourceRef
    {
        return new DataSourceRef(
            type: 'loki',
            uid: 'loki',
        );
    }

    public static function prometheusDatasourceRef(): DataSourceRef
    {
        return new DataSourceRef(
            type: 'prometheus',
            uid: 'prometheus',
        );
    }

    public static function statPanel(): Stat\PanelBuilder
    {
        return (new Stat\PanelBuilder())
            ->colorScheme((new FieldColorBuilder())->mode(FieldColorModeId::thresholds()))
            ->graphMode(BigValueGraphMode::area())
            ->colorMode(BigValueColorMode::value())
            ->justifyMode(BigValueJustifyMode::auto())
            ->textMode(BigValueTextMode::auto())
            ->orientation(VizOrientation::auto())
            ->thresholds(
                (new ThresholdsConfigBuilder())
                    ->mode(ThresholdsMode::absolute())
                    ->steps([
                        ['color' => 'green'],
                        ['value' => 80, 'color' => 'red'],
                    ])
            )
        ;
    }

    public static function heatmapPanel(): Heatmap\PanelBuilder
    {
        return (new Heatmap\PanelBuilder())
            ->color(
                (new HeatmapColorOptionsBuilder())
                ->mode(HeatmapColorMode::scheme())
                ->scheme('RdYlBu')
                ->fill('dark-orange')
                ->scale(HeatmapColorScale::exponential())
                ->exponent(0.5)
                ->steps(64)
                ->reverse(false)
            )
            ->filterValues((new FilterValueRangeBuilder())->le(1e-09))
            ->rowsFrame((new RowsHeatmapOptionsBuilder())->layout(HeatmapCellLayout::auto()))
            ->showValue(VisibilityMode::auto())
            ->cellValues(new CellValuesBuilder())
            ->yAxis(
                (new YAxisConfigBuilder())
                    ->unit('s')
                    ->reverse(false)
                    ->axisPlacement(AxisPlacement::left())
            )
            ->showLegend()
            ->mode(TooltipDisplayMode::single())
            ->hideYHistogram()
            ->showColorScale(true)
            ->scaleDistribution((new ScaleDistributionConfigBuilder())->type(ScaleDistribution::linear()))
            ->hideFrom(
                (new HideSeriesConfigBuilder())
                    ->tooltip(false)
                    ->legend(false)
                    ->viz(false)
            )
        ;
    }

    public static function logPanel(): Logs\PanelBuilder
    {
        return (new Logs\PanelBuilder())
            ->datasource(self::lokiDatasourceRef())
            ->showTime(true)
            ->sortOrder(LogsSortOrder::descending())
            ->enableLogDetails(true)
        ;
    }

    public static function textPanel(string $content): Text\PanelBuilder
    {
        return (new Text\PanelBuilder())
            ->mode(TextMode::markdown())
            ->content($content)
        ;
    }

    public static function timeseriesPanel(): Timeseries\PanelBuilder
    {
        return (new Timeseries\PanelBuilder())
            ->lineWidth(1)
            ->pointSize(5)
            ->fillOpacity(20)
            ->gradientMode(GraphGradientMode::opacity())
            ->legend(
                (new VizLegendOptionsBuilder())
                    ->displayMode(LegendDisplayMode::list())
                    ->placement(LegendPlacement::bottom())
                    ->showLegend(true)
            )
            ->tooltip(
                (new VizTooltipOptionsBuilder())
                    ->mode(TooltipDisplayMode::single())
                    ->sort(SortOrder::none())
            )
            ->colorScheme((new FieldColorBuilder())->mode(FieldColorModeId::paletteClassic()))
            ->thresholdsStyle((new GraphThresholdsStyleConfigBuilder())->mode(GraphThresholdsStyleMode::off()))
        ;
    }

    public static function instantPrometheusQuery(string $expression): Prometheus\DataqueryBuilder
    {
        return (new Prometheus\DataqueryBuilder())
            ->expr($expression)
            ->instant()
            ->format(PromQueryFormat::table())
            ->legendFormat('__auto')
        ;
    }

    public static function prometheusQuery(string $expression): Prometheus\DataqueryBuilder
    {
        return (new Prometheus\DataqueryBuilder())
            ->expr($expression)
            ->range()
            ->format(PromQueryFormat::timeSeries())
            ->legendFormat('__auto')
        ;
    }

    public static function lokiQuery(string $expression): Loki\DataqueryBuilder
    {
        return (new Loki\DataqueryBuilder())
            ->expr($expression)
            ->QueryType('range')
            ->legendFormat('__auto')
        ;
    }
}