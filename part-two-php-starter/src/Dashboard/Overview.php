<?php

namespace App\Dashboard;

use App\Catalog\Service;

use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\DashboardCursorSync;
use Grafana\Foundation\Dashboard\DashboardLinkBuilder;
use Grafana\Foundation\Dashboard\TextBoxVariableBuilder;
use Grafana\Foundation\Dashboard\VariableRefresh;

class Overview
{
    public static function forService(Service $service): SDKDashboard\Dashboard
    {
        $builder = (new DashboardBuilder(title: $service->name.' service overview'))
            ->uid($service->name.'-overview')
            ->tags([$service->name, 'generated'])
            ->readonly()
            ->time('now-30m', 'now')
            ->tooltip(DashboardCursorSync::crosshair())
            ->refresh('10s')
            ->link(
                (new DashboardLinkBuilder('GitHub Repository'))
                    ->url($service->repositoryUrl)
                    ->targetBlank(true)
            )
            ->withVariable(
                (new TextBoxVariableBuilder('logs_filter'))
                    ->label('Logs filter')
            )
            ->withVariable(self::logLevelsVariable($service))
        ;

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

        return $builder->build();
    }

    private static function logLevelsVariable(Service $service): SDKDashboard\QueryVariableBuilder
    {
        return (new SDKDashboard\QueryVariableBuilder('logs_level'))
            ->label('Logs level')
            ->datasource(Common::lokiDatasourceRef())
            ->query([
				'label'  => 'level',
				'stream' => "{service=\"$service->name\"}",
				'type'   =>   1,
            ])
            ->refresh(VariableRefresh::onTimeRangeChanged())
            ->includeAll(true)
            ->allValue('.*')
        ;
    }
}