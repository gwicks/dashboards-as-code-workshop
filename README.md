# Notes

## Start the stack

```console
docker compose -f docker-compose.yaml up --watch
```

`--watch` will watch for changes in `./dummy/` and `./service-catalog`, and rebuild+restart the services when needed.

The Grafana instance is accessible at `http://localhost:3003` (credentials: `admin` / `admin`)

## Develop the dashboards

The dashboard's code lives in `./golang/`.

It can be run with:

```console
cd ./golang
go run *.go
```

It will generate a single dashboard, with a hardcoded service configuration.
This mode is meant for development, to be used alongside grizzly:

```console
cd ./golang
grr serve -w -S 'go run *.go' .
```

Make sure Grizzly is [installed](https://grafana.github.io/grizzly/installation/), and [configured for the Grafana instance exposed by this stack](https://grafana.github.io/grizzly/configuration/).

## Deploy the dashboards

```console
cd ./golang
go run *.go -deploy
```

This will call the service catalog and deploy a dashboard for each service it describes.
