# Hands-on lab: Grafana as code

## Installing dependencies

```shell
composer install
```

## Running the code

```shell
php index.php
```

It will generate a single dashboard, with a hardcoded service configuration.
This mode is meant for development, to be used alongside Grizzly:

```shell
grr serve --only-spec --kind Dashboard -w -S 'php index.php' .
```

## Where should I start?

The [`index.php`](./index.php) file is the entrypoint both for the development and
deployment *modes*.

The [`./src/Dashboard/Overview.php`](./src/Dashboard/Overview.php) file defines a `Overview::forService()`
static method that will be called to generate a dashboard for a given service.

The [`./src/Dashboard/Common.php`](./src/Dashboard/Common.php) file contains a few utilities related
to panel creations with sensible defaults and configuration.

> [!TIP]
> It is highly recommended that every panel created for your dashboard use one
> of these utility functions.

## Deploying the dashboards

```shell
php index.php --deploy
```

This will call the service catalog and deploy a dashboard for each service it
describes.
