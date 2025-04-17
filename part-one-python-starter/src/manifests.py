from grafana_foundation_sdk.models import dashboard, resource


class Manifest:
    @classmethod
    def dashboard(cls, folder_uid: str, dash: dashboard.Dashboard) -> resource.Manifest:
        return resource.Manifest(
            api_version="dashboard.grafana.app/v1alpha1",
            kind="Dashboard",
            metadata=resource.Metadata(
                annotations={"grafana.app/folder": folder_uid},
                name=dash.uid,
            ),
            spec=dash,
        )

