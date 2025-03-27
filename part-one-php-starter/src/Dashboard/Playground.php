<?php

namespace App\Dashboard;

use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\DashboardCursorSync;
use Grafana\Foundation\Logs;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Text;
use Grafana\Foundation\Timeseries;

class Playground
{
    public static function create(): SDKDashboard\Dashboard
    {
        $builder = (new DashboardBuilder(title: 'Test dashboard'))
            ->uid('test-dashboard')
            ->tags(['test', 'generated'])
            ->readonly()
            ->time('now-30m', 'now')
            ->tooltip(DashboardCursorSync::crosshair())
            ->refresh('10s')
        ;

        $builder
            ->withPanel(self::prometheusVersionStat())
            ->withPanel(self::descriptionText())
            ->withPanel(self::unfilteredLogs())
            ->withPanel(self::prometheusGoroutinesTimeseries())
        ;

        return $builder->build();
    }

    private static function prometheusVersionStat(): Stat\PanelBuilder
    {
        return Common::statPanel()
	        // TODO: configure the panel
        ;
    }

    private static function descriptionText(): Text\PanelBuilder
    {
        return Common::textPanel('')
            // TODO: configure the panel
        ;
    }

    private static function unfilteredLogs(): Logs\PanelBuilder
    {
        return Common::logPanel()
            // TODO: configure the panel
        ;
    }

    private static function prometheusGoroutinesTimeseries(): Timeseries\PanelBuilder
    {
        return Common::timeseriesPanel()
            // TODO: configure the panel
        ;
    }
}