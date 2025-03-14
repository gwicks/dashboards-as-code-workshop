package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func dashboardForService(service Service) *dashboard.DashboardBuilder {
	builder := dashboard.NewDashboardBuilder(fmt.Sprintf("%s service overview", service.Name)).
		Uid(fmt.Sprintf("%s-overview", service.Name)).
		Tags([]string{service.Name, "generated"}).
		Readonly().
		Time("now-30m", "now").
		Tooltip(dashboard.DashboardCursorSyncCrosshair).
		Refresh("10s").
		Link(dashboard.NewDashboardLinkBuilder("GitHub Repository").
			Url(service.RepositoryURL).
			TargetBlank(true),
		)

	// Overview
	builder.
		WithPanel(versionStat(service).Height(4).Span(4)).
		WithPanel(descriptionText(service).Height(4).Span(4)).
		WithPanel(logsVolumeTimeseries(service).Height(4).Span(16))

	// gRPC row, if relevant
	if service.HasGRPC {
		builder.WithRow(dashboard.NewRowBuilder("gRPC")).
			WithPanel(grpcRequestsTimeseries(service).Height(8)).
			WithPanel(grpcLatenciesHeatmap(service).Height(8)).
			WithPanel(grpcLogsPanel(service).Height(8).Span(24))
	}

	// HTTP row, if relevant
	if service.HasHTTP {
		builder.
			WithRow(dashboard.NewRowBuilder("HTTP")).
			WithPanel(httpRequestsTimeseries(service).Height(8)).
			WithPanel(httpLatenciesHeatmap(service).Height(8)).
			WithPanel(httpLogsPanel(service).Height(8).Span(24))

	}

	return builder
}

func fetchServicesAndDeploy(cfg config) error {
	services, err := fetchServices(cfg)
	if err != nil {
		return err
	}

	client := grafanaClient(cfg)

	for _, service := range services {
		serviceDashboard, err := dashboardForService(service).Build()
		if err != nil {
			return err
		}

		folderUid, err := findOrCreateFolder(client, service.Name)
		if err != nil {
			return err
		}

		_, err = client.Dashboards.PostDashboard(&models.SaveDashboardCommand{
			FolderUID: folderUid,
			Dashboard: serviceDashboard,
			Overwrite: true,
		})
		if err != nil {
			return fmt.Errorf("failed posting dashboard for service '%s': %w", service.Name, err)
		}
	}

	return nil
}

func printDevelopmentDashboard(service Service) {
	serviceDashboard, err := dashboardForService(service).Build()
	if err != nil {
		panic(err)
	}
	dashboardJson, err := json.MarshalIndent(serviceDashboard, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dashboardJson))
}
