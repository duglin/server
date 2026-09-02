# xrserver (xRegistry Server)

`xrserver` boots and manages the xRegistry API server and its backing MySQL
database. See [Installation & Admin Guide](../installation.md) for
deployment options (Docker, Kubernetes, external MySQL, TLS/auth).

## `xrserver` Command Summary

<!-- XRSERVER HELP START -->
```yaml
xrserver [command]
  # Global flags:
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
      --dontcreate          Don't create DB/reg if missing
  -?, --help                Help for commands
      --help-all            Help for all commands
  -p, --port int            HTTP Listen port (8080*)
      --recreatedb          Recreate the DB
      --recreatereg         Recreate registry
  -r, --registry string     Default Registry name
      --rootapp string      Root application (ui,xreg) (default "ui")
      --samples             Load sample registries
      --set stringArray     Override configFile property: --set NAME[:VALUE]
      --ui-dir string       Serve new UI from this directory (dev mode)
  -v, --verbose             Be chatty
      --verify              Verify loading and exit
      --version             Print command version string

xrserver db [command]
  # Manage mysql databases
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver db create NAME
  # Create a new mysql DB
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -f, --force               Delete existing DB first
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver db delete NAME
  # Delete a mysql DB
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -f, --force               Ignore DB missing error
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver db get NAME
  # Get details about a mysql DB
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver db list
  # List the databases
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -?, --help                Help for commands
  -o, --output string       Output format: json, table*
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver help [command]
  # Help about any command

xrserver registry [command]
  # Manage xRegistries
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver registry create ID...
  # Create one or more xRegistry
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -f, --force               Ignore existing registry
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver registry delete ID...
  # Delete one or more registries
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -f, --force               Ignore missing registry
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver registry get ID
  # Get details about a registry
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver registry list
  # List the registries
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
  -?, --help                Help for commands
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --version             Print command version string

xrserver run
  # Run server (the default command)
      --config string       Config file ($HOME/.xrserver)
      --db string           DB name (registry*)
      --dbhost string       DB host address (127.0.0.1*)
      --dbpassword string   DB password (password*)
      --dbport int          DB host port (3306*)
      --dbuser string       DB user (root*)
      --dontcreate          Don't create DB/reg if missing
  -?, --help                Help for commands
  -p, --port int            HTTP Listen port (8080*)
      --recreatedb          Recreate the DB
      --recreatereg         Recreate registry
  -r, --registry string     Default Registry name(xRegistry*)
      --rootapp string      Root application (ui,xreg) (default "ui")
      --samples             Load sample registries
      --set stringArray     Override configFile property: --set NAME[:VALUE]
  -v, --verbose             Be chatty
      --verify              Verify loading and exit
      --version             Print command version string
```
<!-- XRSERVER HELP END -->

## Example Commands
```yaml
# Start server on port 8080 and load sample data ('run' is optional)
xrserver --samples
xrserver run --samples

# Drop & recreate the database, then run the server ('run' is optional)
xrserver --recreatedb
xrserver run --recreatedb

# Create a new registry named "myregistry"
xrserver registry create myregistry

# List all registries
xrserver registry list
```

## `xrserver` Environment Variables & Config File

Covered in the [Configuration Reference](../configuration.md), including
the full list of `DB*`/`XR_*` environment variables and the `.xrserver`
config file format (`rootapp`, `path.*`, `ui.*`, `db.*`, ...).
