import json, typing
from grafana_foundation_sdk.models import dashboard
from grafana_foundation_sdk.cog.encoder import JSONEncoder


class Manifest:
    api_version: str
    kind: str
    metadata: dict[str, str]
    spec: typing.Any

    def __init__(self, api_version: str, kind: str = "", metadata: dict[str, str] = None, spec: typing.Any = None):
        self.api_version = api_version
        self.kind = kind
        self.metadata = {} if metadata is None else metadata
        self.spec = spec

    @classmethod
    def dashboard(cls, folder_uid: str, dash: dashboard.Dashboard) -> typing.Self:
        dash_json = JSONEncoder(sort_keys=True, indent=2).encode(dash)
        data = json.loads(dash_json)
        return cls(
            api_version="dashboard.grafana.app/v1alpha1",
            kind="Dashboard",
            metadata={
                "annotations": {
                    "grafana.app/folder": folder_uid,
                },
                "name": dash.uid,
            },
            spec=data,
        )

    def as_data(self) -> dict[str, object]:
        return {
            "apiVersion": self.api_version,
            "kind": self.kind,
            "metadata": self.metadata,
            "spec": self.spec,
        }
