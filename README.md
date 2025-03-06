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

### Install Grizzly

Make sure Grizzly is [installed](https://grafana.github.io/grizzly/installation/), and [configured for the Grafana instance exposed by this stack](https://grafana.github.io/grizzly/configuration/).

**For MacOS (with Homebrew):**

```
brew install grizzly
```

**For MacOS (Apple silicon):**

```
curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-darwin-arm64"
chmod +x /usr/local/bin/grr
```

**For MacOS (Intel chips):**

```
curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-darwin-amd64"
chmod +x /usr/local/bin/grr
```

**For Linux:**

```
curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-linux-amd64"
chmod +x /usr/local/bin/grr
```

## Deploy the dashboards

```console
cd ./golang
go run *.go -deploy
```

This will call the service catalog and deploy a dashboard for each service it describes.
