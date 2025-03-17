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
grr serve -w -S 'go run *.go' .
```

## Deploying the dashboards

```shell
go run *.go -deploy
```

This will call the service catalog and deploy a dashboard for each service it describes.
