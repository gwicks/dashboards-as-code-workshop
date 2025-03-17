# Notes

## Starting the stack

```console
docker compose -f docker-compose.yaml up --watch --build
```

`--watch` will watch for changes in `./dummy/` and `./service-catalog`, and rebuild+restart the services when needed.

The Grafana instance is accessible at `http://localhost:3003` (credentials: `admin` / `admin`)

## Service catalog

The fake service catalog is implemented as a small Go application serving a static JSON file (the catalog)
over HTTP.

Code: `./service-catalog`

Catalog endpoint: `http://localhost:8082/api/services`

## Fake services

Fake services are implemented by a single Go application which randomly emits metrics and logs.

Code: `./dummy`

The metrics emitted for each service can somewhat be configured. See the [`./dummy/services.go`](./dummy/services.go) file
