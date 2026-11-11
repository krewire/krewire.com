---
title: "Command Catalog"
description: "Complete command catalog for kiw: syntax, flags, workload support, and examples for every command in the Krewire developer CLI."
date: "2026-10-08"
---

# Command Catalog

The **`kiw`** CLI provides an end-to-end command set spanning the entire software development lifecycle—from project initialization to production deployment.

---

## 1. Project Initialization & Scaffolding

### `kiw new <name>`

Scaffold a new project in a new directory with a pre-configured template.

```bash
kiw new my-app [--site | --app | --cli | --book | --worker | --service | --infra]
```

- `--site`: Static site generator with `.kiw` Single-File Components.
- `--app`: Fullstack web monolith with `net/http` and embedded assets.
- `--cli`: Terminal utility with Elm-style Model-Update-View architecture.
- `--book`: Documentation book with Markdown manuscripts and `pkg/book`.
- `--worker`: Asynchronous job worker daemon.
- `--service`: Cloud-native microservice.
- `--infra`: Infrastructure as Code in Go.

### `kiw init`

Equip an existing project directory with Krewire workload structures:

```bash
kiw init [--site | --book | --cli | --worker | --service | --infra]
```

- `--template <git-url>`: Clone an external starter template repository into an empty directory.

### `kiw make <kind> <name>`

Rapidly scaffold application artifacts with idiomatic conventions:

```bash
kiw make component <Name>       # Scaffold components/<Name>.kiw
kiw make component <Name.Part>  # Scaffold compound component components/<Name>/<Part>.kiw
kiw make page <name>            # Scaffold pages/<name>.kiw
kiw make handler <Name>         # Scaffold internal/handlers/<name>.go
kiw make migration <name>       # Scaffold migrations/YYYYMMDDHHMMSS_<name>.go
kiw make model <Name>           # Scaffold internal/models/<name>.go (with orm.BaseModel)
kiw make seeder <Name>          # Scaffold seeds/<name>_seeder.go
kiw make factory <Name>         # Scaffold factories/<name>_factory.go
kiw make repository <Name>      # Scaffold internal/repositories/<name>_repository.go
kiw make middleware <Name>      # Scaffold internal/middleware/<name>.go
kiw make service <Name>         # Scaffold internal/services/<name>.go
kiw make job <Name>             # Scaffold internal/jobs/<name>_job.go
kiw make chapter <name>         # Scaffold book chapter manuscript
kiw make doc <name>             # Scaffold docs article
```

Supports colon syntax (e.g. `kiw make:component Card` or `kiw make component:Card`) and explicit flags (`--kind`, `--name`). Supports compound dotted component names (`kiw make component Card.Header` automatically scaffolds `components/Card/Header.kiw` for seamless `<Card.Header />` compound invocation).

### `kiw migrate <up [n] | down [n] | redo | status | create <name>>`

Manage database schema migrations with transactional safety and advisory locks:

```bash
kiw migrate create create_users   # Generate a new Go schema migration
kiw migrate up                    # Apply all pending migrations
kiw migrate up 1                  # Apply the next single migration
kiw migrate down                  # Roll back the last migration batch
kiw migrate redo                  # Roll back and re-apply atomically
kiw migrate status                # Display migration status table
```

Flags:
- `--dsn <dsn>`: Override database connection string.
- `--driver <driver>`: Database driver (`sqlite`, `postgres`, `mysql`).
- `--dir <path>`: Migrations directory (default: `migrations`).
- `--lock-timeout <duration>`: Lock acquisition timeout (default: `30s`).

### `kiw seed [--class <name>] [--dsn <dsn>] [--driver <driver>] [--dir <dir>]`
Alias: `kiw db:seed`

Seed the database with test data and initial system state via Go seeders and model factories:

```bash
kiw seed                       # Execute DatabaseSeeder
kiw seed --class UserSeeder    # Execute a specific seeder class
kiw db:seed                    # Laravel-compatible alias
```

---

## 2. Development & Execution

### `kiw dev`

Start the local development server with automatic file watching and live reload:

```bash
kiw dev [--port 8080]
```

- Watches `pages/`, `layouts/`, `components/`, `public/`, and Go source files.
- Automatically re-renders templates or restarts Go server processes in `<10ms`.

### `kiw run [task|args]`

Run a Go entry point or a custom task defined in `krewire.yaml`:

```bash
kiw run [task-name] [-- args...]
```

Examples:
```bash
kiw run                # Runs the default Go entry point
kiw run build-css      # Executes the 'build-css' task from krewire.yaml
kiw run -- --verbose   # Forwards '--verbose' flag to the application
```

### `kiw serve`

Start a local HTTP server to preview static output (`site` or `book` workloads) or listen on a web port:

```bash
kiw serve [--port 3000] [--dir .krewire/build]
```

### `kiw dash` (alias: `kiw dashboard`)

Launch an interactive local developer dashboard and real-time observability telemetry server, or manage AI-native virtual developer workspaces and isolated execution sandboxes:

```bash
kiw dash [--port 4000] [--host 127.0.0.1]
kiw dash start --sandbox=go-1.27
kiw dash list
kiw dash status
kiw dash console [--port 9000]
kiw dash stop [id]
```

- **Interactive UI (`/`)**: Dark-themed web dashboard inspecting detected pages, handlers, configured microservices, and live memory metrics.
- **Real-Time Telemetry (`/api/events`)**: Server-Sent Events (SSE) stream pushing goroutine count, memory allocation, and GC telemetry every 2 seconds.
- **Virtual Workspace Sandboxes**:
  - `start [--sandbox=<image>]`: Mount and attach to an isolated developer workspace sandbox.
  - `list`: List running developer workspaces and sandboxes.
  - `status`: Inspect Dash server connectivity, quota, and health.
  - `console`: Open or print Web IDE console URL.
  - `stop [id]`: Stop a running sandbox.
- **REST & MCP Endpoints**:
  - `GET /api/info`: Project shape, Go toolchain version, services, and route table.
  - `GET /api/metrics`: Live memory allocations and goroutine counter.
  - `GET /api/health`: Health status probe for local containers and automation.
  - `GET /mcp/sse`: Native Model Context Protocol (MCP) server for AI coding agents.

### `kiw mcp`

Start a native Model Context Protocol (MCP) server for AI coding agents (Antigravity, Cursor, Claude Desktop, OpenCode):

```bash
kiw mcp --stdio                      # Standard stdio transport for AI IDEs
kiw mcp --port 4000 --host 127.0.0.1 # Standalone HTTP/SSE transport
```

- **MCP Resources**:
  - `krewire://project/info`: Project metadata, detected kind, Go version, and configured services.
  - `krewire://project/routes`: All discovered `.kiw` pages, Markdown articles, and Go HTTP handlers.
  - `krewire://runtime/metrics`: Live memory allocation, system memory, goroutines, and GC stats.
  - `krewire://runtime/health`: Continuous health status probe.
- **MCP Tools**:
  - `kiw_inspect_project`: Inspect project configuration, architecture shape, and environment.
  - `kiw_list_routes`: Extract the full routing tree.
  - `kiw_make_scaffold`: Generate components, pages, handlers, migrations, models, middleware, and services.
  - `kiw_get_metrics`: Real-time telemetry probe for AI performance audits.
### `kiw worker`

Run asynchronous background workers and queue processors:

```bash
kiw worker [--queue memory|file|kv] [--concurrency <N>] [--dlq list|clear]
```

- `--queue`: Queue storage backend (`memory`, `file`, or `kv`).
- `--concurrency`: Number of concurrent job workers (default derived from `krewire.yaml` or CPU count).
- `--dlq list`: Inspect dead-lettered failed jobs.
- `--dlq clear`: Flush and purge all dead-lettered jobs.

### `kiw gateway`

Run a high-performance API gateway and reverse proxy for microservices:

```bash
kiw gateway [--port 8080] [--host 0.0.0.0] [--routes "/api/orders=http://localhost:8081,/api/users=http://localhost:8082"]
```

- Dispatches routes based on `gateway:` configuration in `krewire.yaml` or explicit `--routes` flag.
- Built-in circuit breaking, distributed rate limiting, CORS, distributed tracing headers, and round-robin load balancing.

---

## 3. Building & Deployment

### `kiw build`

Compile the current project into release-ready artifacts according to its workload kind:

```bash
kiw build [--target site|book|binary|wasm|infra] [--output <path>]
```

- **`site`**: Generates optimized static HTML and scoped CSS into `.krewire/build`.
- **`book`**: Compiles Markdown manuscripts into multi-chapter documentation portal.
- **`app`/`cli`/`service`/`worker`**: Compiles a single Go binary with embedded assets.
- **`wasm`**: Compiles a `pkg/runtime` client bundle into a `.wasm` module plus JS glue.
- **`infra`**: Compiles and verifies the cloud infrastructure topology plan.

### `kiw deploy`

Validate the project, run quality gates, stage distribution artifacts into `.krewire/dist/`, and publish to target environments:

```bash
kiw deploy [--target binary|systemd|gh-pages|infra] [--branch <name>] [--remote <name>] [--dry-run]
```

- `--dry-run`: Preview deployment steps without publishing or mutating cloud state.
- `--target binary`: Build and stage standalone executable into distribution directory.
- `--target systemd`: Generate and configure Linux systemd unit service files (`/etc/systemd/system/<app>.service`) with custom user, working directory, and auto-restart policies.
- `--target gh-pages`: Commit and push the static output to GitHub Pages (`gh-pages` branch).
- `--target infra`: Provision declared cloud resources via provider adapters.

### `kiw bundle`

Bundle all eligible CSS and JS assets into `assets/krewire.css` and `assets/krewire.js`:

```bash
kiw bundle [--minify] [--output <dir>]
```

- `--minify` / `-m`: enable CSS and JS minification.
- `--output <dir>` / `-o`: target output directory (default: derived from `krewire.yaml` or `.krewire/build`).

---

## 4. Code Quality & Testing

### `kiw test`

Execute unit and integration tests across all Go packages:

```bash
kiw test [--filter <pattern>] [--watch] [--coverage] [--verbose] [--json]
```

Spawns `go test ./...` with standardized colorized output and test summary reporting.

- `--filter <pattern>` (alias `--run`): run only tests matching a name or file pattern.
- `--watch` / `-w`: watch files and rerun tests on change.
- `--coverage`: run with coverage (`go test -cover`); `--coverprofile <file>` writes the profile.
- `--verbose` / `-v`, `--json`, `--reporter verbose|dot|json`: output control.
- `--update` / `-u`: update snapshots/golden files; `--list`: list tests instead of running.
- `--count <n>`: pass through `go test -count` (e.g. `1` to disable the test cache).

### `kiw vet`

Run static code analysis and optional security audits:

```bash
kiw vet [--sec] [--vuln]
```

Runs Go's official `go vet ./...` analyzer, catching potential defects, dead code, and invalid struct tags.

- `--sec`: security posture checks — secrets scan, config hygiene, and `.gitignore` coverage.
- `--vuln`: vulnerability audit via `govulncheck` and checksum verification.

### `kiw fmt`

Inspect or format source code and `.kiw` component files:

```bash
kiw fmt [--write]
```

- Without `--write`: Prints unformatted files and exits with code `1` if formatting is required (useful for CI gates).
- With `--write`: Formats all Go and component files in-place according to canonical conventions.

---

## 5. Ecosystem & Management

### `kiw info`

Display detected workload kind, active configuration, and Go environment diagnostics:

```bash
kiw info
```

*Example Output:*
```text
Project:      storefront (v0.1.0)
Kind:         app (Fullstack Monolith)
Config:       /home/reasvyn/workspace/storefront/krewire.yaml
Output:       bin/server
Go Version:   go1.27.1 (linux/amd64)
```

### `kiw version`

Print the version of the `kiw` CLI and underlying Krewire packages:

```bash
kiw version
```

### `kiw compat`

Validate version compatibility across all module declarations in the workspace:

```bash
kiw compat
```

### `kiw tier`

Inspect the project's open-core tier and validate open-core boundary invariants:

```bash
kiw tier [--json] [--check <tier>]
```

- `--json`: output tier posture in structured JSON.
- `--check <tier>`: assert the project does not exceed the specified tier (e.g. `--check free`).

### `kiw doctor`

Diagnose the developer environment and project readiness:

```bash
kiw doctor [--json]
```

Checks toolchain prerequisites, workspace layout, and the prerequisites a project must meet before it can build. `--json` emits a machine-readable report for CI.

### `kiw ws <command>`

Manage multi-package Go workspaces (`go.work`):

```bash
kiw ws info         # Show workspace members and active modules
kiw ws list         # List paths of all packages
kiw ws sync         # Synchronize go.work and go.mod dependencies
kiw ws exec <cmd>   # Execute a command across every workspace module
```

### `kiw boost install`

Equip any project repository with the **Krewire Boost** AI agent system:

```bash
kiw boost install [path]
```

Installs the autonomous agent definitions, custom slash commands, and autonomous SDLC workflows into `.agents/` and `AGENTS.md`.

### `kiw add <package[@version]>`

Add a Go module or an opt-in ecosystem plugin (such as Tailwind CSS):

```bash
kiw add twcss@latest
kiw add tailwindcss
kiw add github.com/krewire/krewire@v0.2.0
```

- **Zero JS by Default**: Krewire defaults to zero JavaScript and standalone CSS tokens.
- **Opt-In Technology**: When adding `twcss`, `kiw` automatically initializes `tailwind.config.js` and enables the Tailwind compiler runner (standalone CLI binary in `bin/tailwindcss`, Bun, PNPM, or npm/npx fallback).

### `kiw remove <package>`

Uninstall a package or deactivate an opt-in plugin:

```bash
kiw remove twcss
```

### `kiw update [packages...] [flags]`

Update project dependencies to their latest compatible versions. When run without arguments, updates all direct and indirect dependencies in `go.mod` and tidies the module files:

```bash
kiw update                      # Update all dependencies in project
kiw update --dry-run            # Preview available updates without modifying files
kiw update krewire ui           # Update specific ecosystem packages to @latest
kiw update github.com/foo/bar   # Update specific external dependency
```

- `--dry-run`: Previews available dependency updates without touching disk.
- `--tidy`: Runs `go mod tidy` automatically after updating (default `true`).
- Automatically updates `package.json` dependencies via `npm update` if frontend plugins are configured.

### `kiw self-update [flags]`

Check for and install the latest stable release of the `kiw` devtool binary:

```bash
kiw self-update          # Update kiw binary to latest stable release
kiw self-update --check  # Check if an update is available without downloading
kiw self-update --force  # Force reinstall the latest version
kiw self-update --dry-run # Simulate update process
```

- `--check`: Inspects GitHub Releases for a newer stable version.

> **Note (2026-11-09):** no GitHub Release has been published for `krewire/krewire` yet, so `kiw self-update` currently relies on the Go toolchain fallback (Go 1.27.1+).

- `--force`: Reinstalls or updates even if already on the latest version.
- `--version <ver>`: Installs a specific target version.
- Performs atomic binary replacement on Linux, macOS, and Windows with automatic Go toolchain fallback.

### `kiw forge <subcommand>`

Declarative web, form, and admin panel builder for Go and `.kiw` (Krewire Forge).

```bash
kiw forge panel admin                                                        # Scaffold admin dashboard layout
kiw forge resource product --model=Product                                   # Scaffold full CRUD resource
kiw forge form contact --fields="name:string,email:string,message:textarea"  # Generate declarative form
kiw forge studio --port=8089                                                 # Launch interactive visual builder studio
```

### `kiw ui <subcommand>`

Manage Pop-Brutalist `.kiw` design system components and templates:

```bash
kiw ui list               # List all available components in Krewire UI
kiw ui add <name>         # Install component directly into ./components/
kiw ui diff <name>        # Compare local component with upstream template
```

### `kiw pkg <init|pack|add|list|export>`

Manage standalone, independently distributable `.kiw` component packages and UI kits ([ADR-K7P9L](https://github.com/krewire/internal/blob/main/docs/adr/ADR-K7P9L-pluggable-reusable-kiw-dsl-components.md)). Enables modular packaging, tar.gz (`.kpkg`) archiving with SHA-256 integrity verification, local installation into `./.krewire/pkg/`, and automated style injection with zero npm dependencies:

```bash
kiw pkg init @acme/ui              # Scaffolds standalone package with kiw.pkg.json
kiw pkg pack ./my-kit              # Packs package into a distributable .kpkg archive
kiw pkg add ./dist/ui-1.0.0.kpkg   # Installs package archive into ./.krewire/pkg/
kiw pkg add ./local-kit            # Installs from local directory into ./.krewire/pkg/
kiw pkg list                       # Lists installed component packages in ./.krewire/pkg
kiw pkg list --json                # Machine-readable JSON inventory of installed packages
kiw pkg export Card                # Validates and exports self-contained component
```

- **Manifest (`kiw.pkg.json`)**: Declares package metadata, exported component paths, and package stylesheets (`styles`).
- **Namespace & Aliasing**: Components from installed packages can be invoked using their package namespace (e.g. `<acme:Button />`) or imported directly via `<script> import Button from '@acme/ui'; </script>`.
- **Integrity & Security**: `kiw pkg pack` and `kiw pkg add` enforce SHA-256 checksums and `sec.SafePath` bounds to prevent ZipSlip and directory traversal vulnerabilities.
