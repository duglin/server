# xRegistry Documentation

This is the full documentation set for the xRegistry server (`xrserver`) and
CLI (`xr`). Docs are organized by audience and follow a simple rule:
**start simple, link to "advanced" sections/pages only when you need them.**

If you're new here, start with [Getting Started](getting-started.md).

## For Users (the `xr` CLI)

| Doc | What's in it |
| --- | --- |
| [Getting Started](getting-started.md) | Run a server, make your first API calls, in under 5 minutes |
| [Concepts](concepts.md) | Registries, Groups, Resources, Versions, XIDs — the vocabulary you need |
| [`xr` CLI Reference](cli/xr.md) | Every command and flag, man-page style |
| [Configuration Reference](configuration.md) | Config files and environment variables for `xr` and `xrserver` |
| [UI Guide](ui.md) | Using and customizing the built-in web explorer |
| [Conformance Testing](conformance.md) | Using `xr conform` to check a server against the spec |

## For Admins (running `xrserver`)

| Doc | What's in it |
| --- | --- |
| [Installation & Admin Guide](installation.md) | Docker, standalone binaries, Kubernetes, MySQL, TLS/auth |
| [`xrserver` CLI Reference](cli/xrserver.md) | Every command and flag, man-page style |
| [Configuration Reference](configuration.md) | Config files, environment variables, `--set` overrides |

## For Developers (contributing to this repo)

| Doc | What's in it |
| --- | --- |
| [Developer Guide](developers.md) | Building, testing, coding conventions, PR process |
| [Design Notes](design.md) | Implementation-specific patterns, gotchas, and DB conventions |

## Community

| Doc | What's in it |
| --- | --- |
| [Community](community.md) | Mailing list, Slack, security reporting, code of conduct |

---

Something missing or unclear? Docs issues are just as welcome as code issues —
please [open one](https://github.com/xregistry/server/issues).
