# Hands-on lab: Grafana as code

## Installing dependencies

```shell
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

## Running the code

```shell
python main.py
```

It will generate a single dashboard, with a hardcoded service configuration.
This mode is meant for development, to be used alongside Grizzly:

```shell
grr serve --only-spec --kind Dashboard -w -S 'python main.py' .
```

## Where should I start?

The [`main.py`](./main.py) file is the entrypoint both for the development and
deployment *modes*.

The [`./src/dashboard.py`](./src/dashboard.py) file defines a `dashboard_for_service()`
function that will be called to generate a dashboard for a given service.

The [`./src/common.py`](./src/common.py) file contains a few utilities related
to panel creations with sensible defaults and configuration.

> [!TIP]
> It is highly recommended that every panel created for your dashboard use one
> of these utility functions.

## Deploying the dashboards

```shell
python main.py --deploy
```

This will call the service catalog and deploy a dashboard for each service it
describes.
