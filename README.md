# Hands-on lab: Grafana as code

## Install Grizzly

Make sure [Grizzly](https://grafana.github.io/grizzly/) is installed:

<details>
    <summary><b>For macOS (with Homebrew)</b></summary>

```shell
brew install grizzly
```
</details>

<details>
    <summary><b>For macOS (Apple silicon)</b></summary>

```shell
curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-darwin-arm64"
chmod +x /usr/local/bin/grr
```
</details>

<details>
    <summary><b>For macOS (Intel chips)</b></summary>

```shell
curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-darwin-amd64"
chmod +x /usr/local/bin/grr
```
</details>

<details>
    <summary><b>For Linux</b></summary>

```shell
curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-linux-amd64"
chmod +x /usr/local/bin/grr
```
</details>

## Configure Grizzly

With Grizzly installed, lets configure it to connect to our local Grafana stack:

```shell
grr config set grafana.url http://localhost:3003
grr config set grafana.user admin
grr config set grafana.token admin
```

## Start the stack

```shell
docker compose -f docker-compose.yaml up --watch --build
```

`--watch` will watch for changes in `./dummy/` and `./service-catalog`, and rebuild+restart the services when needed.

The Grafana instance is accessible at `http://localhost:3003` (credentials: `admin` / `admin`)

## Verify Grizzly's configuration

Now that our Grafana stack is up and running, let's check that Grizzly can indeed talk to it:

```shell
grr config check
```

## Develop the dashboards

The dashboard's code lives in `./golang-starter/`.

It can be run with:

```shell
cd ./golang-starter
go mod tidy
go mod vendor
go run *.go
```

It will generate a single dashboard, with a hardcoded service configuration.
This mode is meant for development, to be used alongside Grizzly:

```shell
cd ./golang
grr serve -w -S 'go run *.go' .
```

## Write code and iterate!

TBD

## Deploy the dashboards

```shell
cd ./golang
go run *.go -deploy
```

This will call the service catalog and deploy a dashboard for each service it describes.
