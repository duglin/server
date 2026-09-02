# Getting Started

Requires [Docker](https://www.docker.com/). This starts the xRegistry server
with an embedded MySQL database and loads sample data:

```bash
docker run -ti -p 8080:8080 ghcr.io/xregistry/xrserver-all --samples
```

> `--samples` preloads a few sample Registries so you have something to
> explore. Watch the `Loading: /...` lines in the output for their URL
> paths. Leave it off to start with an empty Registry.

The API is now up at `http://localhost:8080`, reachable with any HTTP client:

```bash
$ curl localhost:8080
{
  "specversion": "1.0-rc4",
  "registryid": "xRegistry",
  "self": "http://localhost:8080/",
  "xid": "/",
  "epoch": 1,
  "createdat": "2025-05-20T16:06:00.652061965Z",
  "modifiedat": "2025-05-20T16:06:00.652061965Z"
}
```

Or, more conveniently, with the [`xr` CLI](cli/xr.md):

```bash
$ xr get
{
  "specversion": "1.0-rc4",
  "registryid": "xRegistry",
  ...
}
```

That's the top-level metadata of the **default Registry**, which starts out
empty. Point at one of the sample Registries instead by adding
`/xregs/NAME` to the server address (see [Concepts](concepts.md) for what
that path means):

```bash
$ export XR_SERVER=localhost:8080/xregs/DocStore   # or: xr -s ...
$ xr get
{
  "registryid": "DocStore",
  "self": "http://localhost:8080/xregs/DocStore/",
  "name": "DocStore Registry",
  "documentsurl": "http://localhost:8080/xregs/DocStore/documents",
  "documentscount": 2,
  ...
}
```

## Browse it visually

Point a browser at the built-in web UI:

```
http://localhost:8080/ui
```

See the [UI Guide](ui.md) for a tour.

## Next steps

- [Concepts](concepts.md) — understand Registries, Groups, Resources, XIDs
- [`xr` CLI Reference](cli/xr.md) — create, update, delete, query data
- [Installation & Admin Guide](installation.md) — production setup options
- Try the [`samples/doc-store`](../samples/doc-store) script for a scripted
  walkthrough with sample data
