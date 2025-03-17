# Prerequisites

For the purposes of this workshop, an entire stack is provided with:

 * Grafana
 * Alloy
 * Loki
 * Prometheus

To run it, you will need [Docker](https://docs.docker.com/engine/) and [Docker Compose](https://docs.docker.com/compose/).

## Start the stack

```shell
docker compose -f docker-compose.yaml up --build
```

> [!NOTE]
> The Grafana instance is accessible at `http://localhost:3003` (credentials: `admin` / `admin`)

## Install Grizzly

[Grizzly](https://grafana.github.io/grizzly/) is a CLI tool used to manage Grafana resources.

Make sure it is installed:

<details>
    <summary><b>For macOS (with Homebrew)</b></summary>

```shell
brew install grizzly
```
</details>

<details>
    <summary><b>For macOS (Apple silicon)</b></summary>

```shell
sudo curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-darwin-arm64"
sudo chmod +x /usr/local/bin/grr
```
</details>

<details>
    <summary><b>For macOS (Intel chips)</b></summary>

```shell
sudo curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-darwin-amd64"
sudo chmod +x /usr/local/bin/grr
```
</details>

<details>
    <summary><b>For Linux</b></summary>

```shell
sudo curl -fSL -o "/usr/local/bin/grr" "https://github.com/grafana/grizzly/releases/download/v0.7.1/grr-linux-amd64"
sudo chmod +x /usr/local/bin/grr
```
</details>

## Configure Grizzly

With Grizzly installed, configure it to connect to the lab's Grafana stack:

```shell
grr config set grafana.url http://localhost:3003
grr config set grafana.user admin
grr config set grafana.token admin
```

Check that Grizzly can indeed talk to the stack:

```shell
grr config check
```

## Next steps

[Start the lab!](./part-one.md)
