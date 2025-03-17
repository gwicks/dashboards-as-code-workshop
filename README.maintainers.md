# Notes

## Starting the stack

```console
docker compose -f docker-compose.yaml up --watch --build
```

`--watch` will watch for changes in `./dummy/` and `./service-catalog`, and rebuild+restart the services when needed.

The Grafana instance is accessible at `http://localhost:3003` (credentials: `admin` / `admin`)

## Part two

Reference implementations for "Part two" live in `./part-two-[language]/`

Each folder must contain a readme describes how to deal with dependencies and run the code.
