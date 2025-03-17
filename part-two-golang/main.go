package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

type config struct {
	CatalogEndpoint string `env:"CATALOG_ENDPOINT" envDefault:"http://localhost:8082/api/services"`
	GrafanaHost     string `env:"GRAFANA_HOST" envDefault:"localhost:3003"`
	GrafanaUser     string `env:"GRAFANA_USER" envDefault:"admin"`
	GrafanaPassword string `env:"GRAFANA_PASSWORD" envDefault:"admin"`
}

func main() {
	cfg, err := env.ParseAs[config]()
	if err != nil {
		log.Fatal(err)
	}

	deploy := false
	flag.BoolVar(&deploy, "deploy", false, "Fetch the list of services from the catalog and deploy a dashboard for each entry")
	flag.Parse()

	// By default, assume we're in "development mode" and print a single
	// dashboard to stdout.
	if !deploy {
		service := Service{
			Name:    "product",
			HasHTTP: true,
			HasGRPC: true,
		}

		printDevelopmentDashboard(service)

		return
	}

	// Otherwise, fetch the list services from the catalog and deploy a
	// dashboard for each of them
	if err := fetchServicesAndDeploy(cfg); err != nil {
		log.Fatal(err)
	}
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

		err = persistDashboard(client, folderUid, serviceDashboard)
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
