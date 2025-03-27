<?php

namespace App\Dashboard;

use Grafana\Foundation\Dashboard\DataSourceRef;
use Grafana\Foundation\Logs;
use Grafana\Foundation\Loki;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Prometheus\PromQueryFormat;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Text;
use Grafana\Foundation\Timeseries;

class Common
{

    public static function statPanel(): Stat\PanelBuilder
    {
        return (new Stat\PanelBuilder());
    }

    public static function logPanel(): Logs\PanelBuilder
    {
        return (new Logs\PanelBuilder())
            // TODO: configure default options for logs panels
        ;
    }

    public static function textPanel(string $content): Text\PanelBuilder
    {
        return (new Text\PanelBuilder())
	        // TODO: configure default options for text panels
        ;
    }

    public static function timeseriesPanel(): Timeseries\PanelBuilder
    {
        return (new Timeseries\PanelBuilder())
            // TODO: configure default options for timeseries panels
        ;
    }

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