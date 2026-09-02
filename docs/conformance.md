# Conformance Testing

`xr conform` checks whether a running server implements the xRegistry
specification correctly — useful both for this server and for validating
other/third-party xRegistry implementations.

## Basic usage

```bash
xr conform localhost:8080
```

Runs the full test suite against the given server(s) and prints a
pass/fail summary. With no arguments, it targets whatever server `xr` is
already configured to talk to (`-s`, `XR_SERVER`, or config file — see
[Configuration](configuration.md)).

You can test multiple servers in one run:

```bash
xr conform localhost:8080 other-host:8080
```

## Useful flags

| Flag | Effect |
| --- | --- |
| `-l`, `--logs` | Show server-side logs even for tests that pass (default: only on failure) |
| `--failfast` | Stop at the first failing test instead of running the whole suite |
| `-d`, `--depth` | How much detail to print for each result (default `2`) |
| `--nowrap` | Don't wrap long output lines |

See the full [`xr` CLI Reference](cli/xr.md) for every flag.

## Exit code

`xr conform` exits non-zero if any test fails — safe to use directly in CI:

```bash
xr conform localhost:8080 || exit 1
```

## See also

- [Concepts](concepts.md) — background on the model this checks against
- [Developer Guide](developers.md) — how conformance tests fit into `make test`
