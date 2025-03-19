from grafana_foundation_sdk.cog.encoder import JSONEncoder

from src.catalog import Service
from src.dashboard import dashboard_for_service

def printDevelopmentDashboard():
    service = Service(
        name='products',
        description='A service related to products',
        has_http=True,
        has_grpc=True,
        repository_url='http://github.com/org/products-service',
    )

    dashboard = dashboard_for_service(service)
    print(JSONEncoder(sort_keys=True, indent=2).encode(dashboard.build()))

if __name__ == '__main__':
    printDevelopmentDashboard()
