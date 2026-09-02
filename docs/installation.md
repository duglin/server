# Installation & Admin Guide

## Getting the bits

| Artifact | Where |
| --- | --- |
| Container images | [`github.com/orgs/xregistry/packages`](https://github.com/orgs/xregistry/packages) |
| Standalone executables (`xr`, `xrserver`) | [GitHub `dev` release](https://github.com/xregistry/server/releases/tag/dev) |

Three container images are published:

- `xrserver` — the API server only; you provide a MySQL database.
- `xrserver-all` — the API server **plus** an embedded MySQL database, for
  quick starts/demos. Not recommended for production (no data durability
  guarantees beyond the container's lifetime unless you mount a volume).
- `xr` — the CLI, also included inside both server images.

## Quickest path (single container, embedded DB)

```bash
docker run -ti -p 8080:8080 ghcr.io/xregistry/xrserver-all --samples
```

See [Getting Started](getting-started.md) for what to do next.

## Standalone binaries

Download `xrserver`/`xr` from the [`dev` release](https://github.com/xregistry/server/releases/tag/dev),
then run `xrserver` against a MySQL instance you manage yourself (see
below) and/or run `xr` against any xRegistry server.

## Using an external MySQL database

Run `xrserver` (the plain image, not `-all`) and point it at your MySQL
instance via flags, environment variables, or a config file — see the
[Configuration Reference](configuration.md) for every knob
(`db.host`, `db.port`, `db.user`, `db.password`, `db.name`, and their
`DBHOST`/`DBPORT`/`DBUSER`/`DBPASSWORD`/`DBNAME` env var equivalents):

```bash
docker run -ti -p 8080:8080 \
  -e DBHOST=mysql.example.com -e DBUSER=xr -e DBPASSWORD=secret \
  ghcr.io/xregistry/xrserver
```

By default `xrserver` creates the database and its schema automatically
on first run. Use `--dontcreate` to disable that (e.g. once you've
provisioned the schema yourself) and `--recreatedb`/`--recreatereg` when
you explicitly want a clean slate — see the
[`xrserver` CLI Reference](cli/xrserver.md).

<details>
<summary>Persisting MySQL data with Docker volumes</summary>

If you're running MySQL in its own container (rather than `xrserver-all`'s
embedded instance, or a managed cloud MySQL), mount a volume so data
survives container restarts:

```bash
docker run -d --name xr-mysql \
  -e MYSQL_ROOT_PASSWORD=password \
  -v xr-mysql-data:/var/lib/mysql \
  -p 3306:3306 mysql:8
```

Then point `xrserver` at it as shown above.
</details>

## Kubernetes

Minimal example manifests are provided under
[`misc/deploy.yaml`](../misc/deploy.yaml) (the `xrserver` Pod/Service) and
[`misc/mysql.yaml`](../misc/mysql.yaml) (a MySQL Pod/Service for it to use).
They're intentionally bare-bones — a starting point, not a production
recipe.

```bash
kubectl apply -f misc/mysql.yaml
kubectl apply -f misc/deploy.yaml
```

<details>
<summary>Making this production-ready</summary>

The sample manifests are single Pods with no persistence, secrets
management, or ingress/TLS configured. For production you'll likely want
to add, at minimum:

- A `PersistentVolumeClaim` for MySQL's data directory, and/or a managed
  cloud MySQL instance instead of running your own.
- A `Secret` for `DBPASSWORD` rather than a plaintext env var.
- A `Deployment` (not a bare `Pod`) for `xrserver` so it's restarted/
  rescheduled automatically.
- An `Ingress` (or equivalent) in front of `xrserver` for TLS — see
  below.
</details>

## TLS and authentication

`xrserver` does not terminate TLS or perform authentication itself. Put a
reverse proxy (nginx, Envoy, your cloud provider's load balancer/ingress,
etc.) in front of it to add:

- TLS termination
- Authentication/authorization (e.g. OAuth2 proxy, mTLS, basic auth)
- Rate limiting

Once you have an auth layer, the `xr` CLI can send whatever credentials
it requires via custom headers — see `header.KEY` in the
[Configuration Reference](configuration.md).

## Configuring the web UI

See the [UI Guide](ui.md) for `xrui.json` (branding) and how the UI locates
its API server.

## Next steps

- [Getting Started](getting-started.md)
- [Configuration Reference](configuration.md)
- The [`samples/doc-store`](../samples/doc-store) script for a scripted
  setup with sample data
