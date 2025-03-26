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

It will generate a single dashboard and print its representation to stdout.
This mode is meant for development, to be used alongside Grizzly:

```shell
grr serve --only-spec --kind Dashboard  -w -S 'go run *.go' .
```

## Where should I start?

The [`main.go`](./main.go) file is the entrypoint both for the development and
deployment *modes*.

The [`dashboard.go`](./dashboard.go) file defines a `testDashboard()`
function that will be called to generate the dashboard.

The [`common.go`](./common.go) file contains a few utility functions related
to panel creations with sensible defaults and configuration.

> [!TIP]
> It is highly recommended that every panel created for your dashboard use one
> of these utility functions.

## Deploying the dashboards

```shell
go run *.go -manifests
```

This will generate a YAML manifest for the test dashboard.
The manifest is written under `./manifests/` by default and can be deployed
from the CLI:

```shell
grr apply ./manifests
```
