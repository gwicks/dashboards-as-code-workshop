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

The [`./src/Dashboard/Playground.php`](./src/Dashboard/Playground.php) file defines a `Playground::create()`
static method that will be called to generate the dashboard.

The [`./src/Dashboard/Common.php`](./src/Dashboard/Common.php) file is where "base functions" for each panel type should be defined.

> [!TIP]
> It is highly recommended that every panel created for your dashboard use one
> of these utility functions.

## Deploying the dashboards

```shell
php index.php --manifests
```

This will generate a YAML manifest for the test dashboard.
The manifest is written under `./manifests/` by default and can be deployed
from the CLI:

```shell
grr apply ./manifests
```
