# Hands-on lab: Grafana as code

## Installing dependencies

```shell
go mod tidy
go mod vendor
```

## Running the code

```shell
go run *.go
```

It will generate a single dashboard, with a hardcoded service configuration.
This mode is meant for development, to be used alongside Grizzly:

```shell
grr serve --only-spec --kind Dashboard  -w -S 'go run *.go' .
```

## Where should I start?

The [`main.go`](./main.go) file is the entrypoint both for the development and
deployment *modes*.

The [`dashboard.go`](./dashboard.go) file defines a `dashboardForService()`
function that will be called to generate a dashboard for a given service in
both cases.

The [`common.go`](./common.go) file contains a few utility functions related
to panel creations with sensible defaults and configuration.

> [!TIP]
> It is highly recommended that every panel created for your dashboard use one
> of these utility functions.

## Deploying the dashboards

```shell
go run *.go -manifests
```

This will call the service catalog and generate a dashboard manifest for each
service it describes.
These manifests are written under `./manifests/` by default and can be deployed
from the CLI:

```shell
grr apply ./manifests
```
