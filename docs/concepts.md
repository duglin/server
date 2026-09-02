# Concepts

A terse reference for the vocabulary and structure used throughout xRegistry.
See the [xRegistry specification](https://github.com/xregistry/spec) for the
full, authoritative definitions — this page just orients you.

## Entity hierarchy

Every Registry has the same shape, defined by its **model**:

```
Registry                          (the root; one per Registry)
└── Group collection              (e.g. "dirs")
    └── Group                     (e.g. "d1")
        └── Resource collection   (e.g. "files")
            └── Resource          (e.g. "f1")
                └── Version collection ("versions")
                │   └── Version   (e.g. "1", "2", ...)
                └── "meta"        (versioning policy/state for the Resource)
```

- **Group** and **Resource** type names (e.g. `dirs`/`files`) are defined by
  the model — they aren't fixed. A Registry can define any number of Group
  types, each with any number of Resource types.
- Creating a child implicitly creates its parents if they don't exist yet
  (e.g. `PUT`-ing a Version auto-creates the Group and Resource above it).
- A Resource with no Version specified auto-creates Version `"1"`.

## Multiple Registries, one server

A single `xrserver` can host many Registries at once:

- `/` — the **default Registry** (configurable, see
  [Configuration](configuration.md)).
- `/xregs/NAME` — any other Registry hosted on this server, by name.
- `/xreg` — alias for whichever Registry is the default, useful when
  `--rootapp=xreg` is set (see [`xrserver`](cli/xrserver.md)).

## XIDs

An **XID** is the absolute, canonical path to any entity, always starting
with `/` and using the collection's plural name, e.g.:

```
/dirs/d1/files/f1/versions/2
```

XIDs are stable identifiers independent of the HTTP `self` URL's host/scheme,
and are what cross-entity references (e.g. `xref` attributes) use.

## Documents vs metadata (`$details`)

A Resource/Version with `hasdocument=true` (the default) returns its
**document content** (e.g. the actual file bytes) when you `GET` it
directly. To read/write its xRegistry **metadata** instead (`name`,
`description`, `epoch`, ...), append `$details` to the URL:

```
GET /dirs/d1/files/f1          -> the file's contents
GET /dirs/d1/files/f1$details  -> the file's xRegistry metadata
```

`$details` is only valid on Resources and Versions — not on Groups or the
Registry root.

## The model

The **model** describes which Group/Resource types a Registry supports, and
which attributes they allow. Manage it with `xr model` (see
[`xr` reference](cli/xr.md)) or by `PUT`-ing a `model`/`modelsource` document
directly.

## Conformance / format checking

xRegistry ships format checkers (JSON Schema, Avro, Protobuf, XML Schema,
JSON Structure, ...) that can validate Resource documents against a schema
referenced by the model. See [Design Notes](design.md) for how formats are
registered internally.

## Where to go next

- [`xr` CLI Reference](cli/xr.md)
- [Installation & Admin Guide](installation.md)
- [UI Guide](ui.md)
