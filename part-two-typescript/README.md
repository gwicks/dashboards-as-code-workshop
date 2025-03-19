# Hands-on lab: Grafana as code

## Installing dependencies

```shell
yarn install
```

## Running the code

```shell
yarn dev
```

It will generate a single dashboard, with a hardcoded service configuration.
This mode is meant for development, to be used alongside Grizzly:

```shell
grr serve --only-spec --kind Dashboard -w -S 'yarn -s dev' .
```

## Where should I start?

The [`./src/index.ts`](./src/index.ts) file is the entrypoint both for the development and
deployment *modes*.

The [`./src/dashboard.ts`](./src/dashboard.ts) file defines a `dashboardForService()`
function that will be called to generate a dashboard for a given service.

The [`./src/common.ts`](./src/common.ts) file contains a few utilities related
to panel creations with sensible defaults and configuration.

> [!TIP]
> It is highly recommended that every panel created for your dashboard use one
> of these utility functions.

## Deploying the dashboards

```shell
TBD
```

This will call the service catalog and deploy a dashboard for each service it
describes.
