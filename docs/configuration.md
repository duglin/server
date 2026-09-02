# Configuration Reference

Both `xr` and `xrserver` are configured the same way, in this priority
order (highest wins): **command-line flags** > **environment variables** >
**config file** > **built-in defaults**.

## Config files

Format (one setting per line):

```
# comment
name: value
```

A config file is looked for, in order: the path given by `--config`, then
`./FILENAME` (current directory), then `$HOME/FILENAME`. If none exist,
built-in defaults are used silently.

### `xr` — `.xr`

| Name | Meaning |
| --- | --- |
| `server.url` | Default xRegistry server URL (same as `-s`/`XR_SERVER`) |
| `header.KEY` | Add HTTP header `KEY` to every request `xr` makes, e.g. for auth tokens |

Example `~/.xr`:

```
server.url: registry.example.com
header.Authorization: Bearer abc123
```

### `xrserver` — `.xrserver`

| Name | Default | Meaning |
| --- | --- | --- |
| `http.port` | `8080` | HTTP listen port |
| `db.name` | `registry` | MySQL database name |
| `db.host` | `127.0.0.1` | MySQL host |
| `db.port` | `3306` | MySQL port |
| `db.user` | `root` | MySQL user |
| `db.password` | `password` | MySQL password |
| `rootapp` | `ui` | What's served at `/` — `ui` or `xreg` (see below) |
| `ui.dir` | _(built-in)_ | Serve the UI from this directory instead of the embedded copy (dev mode) |
| `ui.xrui.json` | _(none)_ | Path to a custom `xrui.json` UI config file (see [UI Guide](ui.md)) |
| `path.ui` | `ui` | URL segment the UI is served under, e.g. `/ui` |
| `path.defaultreg` | `xreg` | URL segment used as an alias for the default Registry |
| `path.regcollection` | `xregs` | URL segment prefix for named Registries, e.g. `/xregs/NAME` |

Any of these can also be set with `xrserver --set NAME:VALUE` at the
command line, without editing the file.

`rootapp` controls what lives at the server's `/` path:
- `ui` (default) — the web UI is served at `/`; the API is offset to
  `/<path.defaultreg>` (e.g. `/xreg`).
- `xreg` — the API is served directly at `/`; the UI moves to `/<path.ui>`.

## Environment variables

### `xr`

| Env Var | Meaning |
| --- | --- |
| `XR_SERVER` | Same as `-s`/`server.url` — the xRegistry server to talk to |
| `XR_VERBOSE` | Chattiness level (see `-v`) |
| `XR_SHOWLOGS` | `true` to show server-side logs even on success (used by `xr conform`) |

### `xrserver`

| Env Var | Maps to | Meaning |
| --- | --- | --- |
| `XR_PORT` | `http.port` | HTTP listen port (`8080*`) |
| `DBHOST` | `db.host` | MySQL host (`127.0.0.1*`) |
| `DBPORT` | `db.port` | MySQL port (`3306*`) |
| `DBUSER` | `db.user` | MySQL user (`root*`) |
| `DBPASSWORD` | `db.password` | MySQL password (`password*`) |
| `DBNAME` | `db.name` | MySQL database name (`registry*`) |
| `XR_VERBOSE` | — | Chattiness: `0`=none, `1`=start-up info, `2`=HTTP requests* , `3+`=debug |
| `XR_MODEL_PATH` | — | Search path for sample model files (dev/test use) |
| `XR_LOAD_LARGE` | — | If set, loads a very large sample Registry (perf testing) |

`*` marks the default value.

## See also

- [`xr` CLI Reference](cli/xr.md)
- [`xrserver` CLI Reference](cli/xrserver.md)
- [Installation & Admin Guide](installation.md)
