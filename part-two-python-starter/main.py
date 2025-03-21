import argparse, sys

from grafana_foundation_sdk.cog.encoder import JSONEncoder

from src.catalog import Config as CatalogConfig, Client as Catalog, Service
from src.grafana import Config as GrafanaConfig, Client as Grafana
from src.dashboard import dashboard_for_service

def print_development_dashboard():
    service = Service(
        name='products',
        description='A service related to products',
        has_http=True,
        has_grpc=True,
        repository_url='http://github.com/org/products-service',
    )

    dashboard = dashboard_for_service(service)
    print(JSONEncoder(sort_keys=True, indent=2).encode(dashboard.build()))

def fetch_services_and_deploy():
    catalog = Catalog(CatalogConfig.from_env())
    grafana = Grafana(GrafanaConfig.from_env())
    services = catalog.services()

    for service in services:
        dashboard = dashboard_for_service(service)
        folder_uid = grafana.find_or_create_folder(service.name)

        grafana.persist_dashboard(folder_uid, dashboard)
    
    print(f"{len(services)} dashboards deployed")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(prog='part-two')
    parser.add_argument('--deploy', action='store_true', help='Fetch the list of services from the catalog and deploy a dashboard for each entry')

    args = parser.parse_args()

    if not args.deploy:
        print_development_dashboard()
        sys.exit(0)

    fetch_services_and_deploy()
