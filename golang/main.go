package main

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v11"
)

type config struct {
	CatalogEndpoint string `env:"CATALOG_ENDPOINT" envDefault:"http://localhost:8082/api/services"`
	GrafanaHost     string `env:"GRAFANA_HOST" envDefault:"localhost:3003"`
	// TODO: use it
	GrafanaToken string `env:"GRAFANA_TOKEN"`
}

func main() {
	cfg, err := env.ParseAs[config]()
	if err != nil {
		log.Fatal(err)
	}

	deploy := false
	flag.BoolVar(&deploy, "deploy", false, "Fetch the list of services from the catalog and deploy a dashboard for each entry")
	flag.Parse()

	// By default, let's assume we're in "development mode" and print a service's
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
