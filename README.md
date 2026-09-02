[![CI](https://github.com/xregistry/server/actions/workflows/master-ci.yaml/badge.svg)](https://github.com/xregistry/server/actions/workflows/ci.yaml)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B162%2Fgithub.com%2Fxregistry%2Fserver.svg?type=small)](https://app.fossa.com/projects/custom%2B162%2Fgithub.com%2Fxregistry%2Fserver?ref=badge_small)

<img src="https://github.com/cncf/artwork/raw/main/projects/xregistry/horizontal/color/xregistry-horizontal-color.svg" alt="xRegistry"></img><br>

# xRegistry Implementation

A full implementation of the [xRegistry specification](https://xregistry.io) —
a metadata registry for discovering, validating, and versioning resources in
distributed systems. Includes:

- **`xrserver`** — the HTTP API server (backed by MySQL) with a built-in
  web UI for browsing registries.
- **`xr`** — a CLI for managing and querying registries from your terminal
  or scripts.

> **Note:** This project is a work-in-progress. See [`todo`](todo) for
> planned work, or file an [issue](https://github.com/xregistry/server/issues)
> if you hit a problem.

## Try it in 30 seconds

```bash
docker run -ti -p 8080:8080 ghcr.io/xregistry/xrserver-all --samples
```

Then open [`http://localhost:8080/ui`](http://localhost:8080/ui) in a
browser, or query it directly:

```bash
curl localhost:8080
```

👉 Or skip local setup entirely: [**try the hosted demo**](http://xregistry.soaphub.org?ui).

## Documentation

Full docs live in [`docs/`](docs/README.md). Highlights:

| I want to... | Read |
| --- | --- |
| Get up and running fast | [Getting Started](docs/getting-started.md) |
| Understand core concepts (registries, groups, resources, XIDs) | [Concepts](docs/concepts.md) |
| Use the `xr` CLI | [`xr` reference](docs/cli/xr.md) |
| Run/deploy the server (Docker, Kubernetes, MySQL, config) | [Installation & Admin Guide](docs/installation.md) |
| Use the web UI | [UI Guide](docs/ui.md) |
| Check a server's spec compliance | [Conformance Testing](docs/conformance.md) |
| Contribute code | [Developer Guide](docs/developers.md) |
| Understand internal design decisions | [Design Notes](docs/design.md) |
| Get help / talk to the community | [Community](docs/community.md) |

See the full [documentation index](docs/README.md) for everything else.
