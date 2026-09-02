# UI Guide

`xrserver` ships with a built-in single-page web app for browsing any
xRegistry server — local or remote.

## Accessing it

By default the UI is available at `/ui` on any running server:

```
http://localhost:8080/ui
```

If the server is run with `--rootapp=xreg`, the UI moves to whatever
`path.ui` is configured to (still `/ui` by default) — see
[Configuration](configuration.md).

## Grid vs. List view

The Home screen offers two layouts for browsing a Registry's contents:

- **Grid** — card-based, good for visual/browse-y exploration.
- **List** — dense, table-based, good for scanning many entities at once.

Toggle between them from the view's layout control; your choice is
remembered per-browser (via `localStorage`), not sent to the server.

## Viewing raw JSON vs. rendered data

Every entity view has a toggle to switch between the rendered/table view
and the raw JSON returned by the server — useful when you need to see
exactly what the API returns (e.g. while debugging a model or a client).

## Adding and finding other servers

The UI isn't tied to the server it's served from — you can point it at any
number of xRegistry servers via the **Config** menu, and switch between
them. Known servers, along with any custom labels or hidden/reordered
entries, are stored in your browser's `localStorage`, not on the server.

> **Tip:** Because server lists live in your browser, you can bookmark
> [`hub.xregistry.io`](https://hub.xregistry.io) (a public, empty UI
> instance) and use it as your own personal, multi-server xRegistry
> browser — add the servers you care about, hide/delete the sample
> `hub` entry, and your bookmark becomes a customized dashboard.

## Customizing the look and feel: `xrui.json`

Server admins can customize the UI's branding via an `xrui.json` file
(point `xrserver` at a custom one with `ui.xrui.json` — see
[Configuration](configuration.md)). Supported fields (only `//`-style
comments allowed in the file itself):

| Field | Meaning |
| --- | --- |
| `servers` | Array of xRegistry server URLs to pre-populate |
| `title` | HTML for the page title/header |
| `summary` | HTML shown under the title (ignored if `headerHTML` is set) |
| `headerHTML` | Path to a file whose contents replace the top of the page entirely |
| `footer` | HTML shown at the bottom of the page (ignored if `footerHTML` is set) |
| `footerHTML` | Path to a file whose contents replace the bottom of the page entirely |
| `footerAlign` | `left`, `center`, or `right` (default) — alignment of `footer` |

Substitutions available in `title`/`summary`/`footer` text:
- `$COMMIT` — first 12 characters of the UI's git commit hash
- `[text](url)` — rendered as a Markdown link

See [`registry/ui/xrui.json`](../registry/ui/xrui.json) for a working
example.

## See also

- [Installation & Admin Guide](installation.md)
- [Configuration Reference](configuration.md)
