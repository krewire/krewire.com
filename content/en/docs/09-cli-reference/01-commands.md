---
title: "Command Catalog"
description: "Complete command catalog for krew: syntax, flags, workload support, and examples for every command in the Krewire developer CLI toolchain."
date: "2026-11-24"
---

# Command Catalog

The **`krew`** CLI provides an end-to-end command set spanning the entire software development lifecycle—from project initialization and AI-assisted development to quality gates and production deployment.

*(Note: The shorthand alias `kiw` is automatically installed as a symlink pointing to `krew` for terminal convenience).*

---

## 1. Core Workflow

### `krew chat`

Start an interactive conversational AI coding session in the terminal:

```bash
krew chat [options]
```

- `--model <model>`: Model override (OpenAI, Anthropic, Gemini, Ollama).
- `--persona <persona>`: Agent persona (default: `coding`). Options: `coding`, `devops`, `worker`, `general`.
- `--resume <session-id>`: Resume an existing saved conversation session ID.
- `--max-steps <int>`: Maximum reasoning steps per prompt (default: `25`).
- `--workspace <dir>`: Target workspace directory (default: `.`).

### `krew agent <prompt>`

Execute an autonomous task or prompt headless to completion:

```bash
krew agent "Add unit tests for auth middleware" [options]
```

- `--brief <path>`: Path to a Markdown Task Brief file or GitHub issue reference.
- `--sdlc`: Execute governed SDLC workflow phases (`Understand -> Plan -> Implement -> Verify -> Review -> Report`).
- `--persona <persona>`: Agent persona specialization (default: `coding`).
- `--model <model>`: Model override.
- `--max-steps <int>`: Bounded reasoning steps (default: `25`).
- `--workspace <dir>`: Workspace directory (default: `.`).

### `krew new <name>`

Scaffold a new project in a new directory with a pre-configured template:

```bash
krew new my-app [--site | --app | --cli | --book | --worker | --service | --infra]
```

- `--site`: Static site generator with `.kiw` Single-File Components.
- `--app`: Fullstack web monolith with `net/http` and embedded assets.
- `--cli`: Terminal utility with Elm-style Model-Update-View architecture.
- `--book`: Documentation book with Markdown manuscripts and `book`.
- `--worker`: Asynchronous background job worker daemon.
- `--service`: Cloud-native microservice.
- `--infra`: Infrastructure as Code in Go.

### `krew dev`

Start the local development server with automatic file watching and live reload:

```bash
krew dev [--port 8080]
```

- Watches `pages/`, `layouts/`, `components/`, `public/`, and Go source files.
- Automatically re-renders templates or restarts Go server processes in `<10ms`.

### `krew build`

Compile and bundle the project for production:

```bash
krew build [--release]
```

- `site`: Emits compiled static HTML, CSS, and JS into `.krewire/build/`.
- `book`: Compiles Markdown manuscripts and navigation into static pages.
- `app`, `cli`, `worker`, `service`: Produces standalone static CGO-free binaries (`CGO_ENABLED=0`) with embedded assets.

### `krew test`

Run the automated test suite of the current project:

```bash
krew test [flags]
```

Runs standard Go unit tests (`go test -v ./...`) across all packages with coverage calculation.

### `krew verify`

Run reproducible quality gates and produce an evidence-backed task verification record:

```bash
krew verify [--json] [--ledger <path>]
```

- `--json`: Output verification record in machine-readable JSON format.
- `--ledger <path>`: Path to tamper-evident ledger JSON file to append evidence block.
- `--task <id>`: Task identifier for evidence tracking (default: `active-task`).
- `--verify-ledger`: Verify cryptographic integrity of existing ledger file without running gates.

### `krew add <package[@version]>`

Add a package, plugin, or copy Krewire UI components into `./components`:

```bash
krew add ui/button ui:minimal      # Install components into ./components/
krew add twcss                     # Enable Tailwind CSS compiler runner
krew add github.com/krewire/ui     # Add Go UI component package
```

### `krew dash` (alias: `krew dashboard`)

Launch an interactive local developer dashboard and real-time observability telemetry server, or manage AI-native virtual developer workspaces and isolated execution sandboxes:

```bash
krew dash [--port 4000] [--host 127.0.0.1]
krew dash start --sandbox=go-1.27
krew dash list
krew dash status
krew dash console [--port 9000]
krew dash stop [id]
```

- **Interactive UI (`/`)**: Dark-themed web dashboard inspecting detected pages, handlers, configured microservices, and live memory metrics.
- **Real-Time Telemetry (`/api/events`)**: Server-Sent Events (SSE) stream pushing goroutine count, memory allocation, and GC telemetry every 2 seconds.
- **Virtual Workspace Sandboxes**: Mount and attach to isolated developer execution sandboxes.

### `krew make <kind> <name>`

Rapidly scaffold application artifacts with idiomatic conventions:

```bash
krew make component <Name>       # Scaffold components/<Name>.kiw
krew make component <Name.Part>  # Scaffold compound component components/<Name>/<Part>.kiw
krew make page <name>            # Scaffold pages/<name>.kiw
krew make handler <Name>         # Scaffold internal/handlers/<name>.go
krew make migration <name>       # Scaffold migrations/YYYYMMDDHHMMSS_<name>.go
krew make model <Name>           # Scaffold internal/models/<name>.go (with orm.BaseModel)
krew make seeder <Name>          # Scaffold seeds/<name>_seeder.go
krew make factory <Name>         # Scaffold factories/<name>_factory.go
krew make repository <Name>      # Scaffold internal/repositories/<name>_repository.go
krew make middleware <Name>      # Scaffold internal/middleware/<name>.go
krew make service <Name>         # Scaffold internal/services/<name>.go
krew make job <Name>             # Scaffold internal/jobs/<name>_job.go
krew make chapter <name>         # Scaffold book chapter manuscript
krew make doc <name>             # Scaffold docs article
```

Supports colon syntax (e.g. `krew make:component Card` or `krew make component:Card`) and explicit flags (`--kind`, `--name`). Supports compound dotted component names (`krew make component Card.Header`).

### `krew pkg <init|pack|add|list|export>`

Manage standalone, independently distributable `.kiw` component packages and UI kits (ADR-K7P9L):

```bash
krew pkg init @acme/ui              # Scaffolds standalone package with kiw.pkg.json
krew pkg pack ./my-kit              # Packs package into a distributable .kpkg archive
krew pkg add ./dist/ui-1.0.0.kpkg   # Installs package archive into ./.krewire/
krew pkg add ./local-kit            # Installs from local directory into ./.krewire/
krew pkg list                       # Lists installed component packages in ./.krewire/pkg
krew pkg list --json                # Machine-readable JSON inventory of installed packages
krew pkg export Card                # Validates and exports self-contained component
```

---

## 2. Project Management

### `krew init`

Equip an existing project directory with Krewire workload structures:

```bash
krew init [--site | --book | --cli | --worker | --service | --infra]
```

- `--template <git-url>`: Clone an external starter template repository into an empty directory.

### `krew remove <package>`

Uninstall a package or deactivate an opt-in plugin:

```bash
krew remove twcss
```

### `krew update [packages...] [flags]`

Update project dependencies to their latest compatible versions:

```bash
krew update                      # Update all dependencies in project
krew update --dry-run            # Preview available updates without modifying files
krew update krewire ui           # Update specific ecosystem packages to @latest
krew update github.com/foo/bar   # Update specific external dependency
```

### `krew migrate <up [n] | down [n] | redo | status | create <name>>`

Manage database schema migrations with transactional safety and advisory locks:

```bash
krew migrate create create_users   # Generate a new Go schema migration
krew migrate up                    # Apply all pending migrations
krew migrate up 1                  # Apply the next single migration
krew migrate down                  # Roll back the last migration batch
krew migrate redo                  # Roll back and re-apply atomically
krew migrate status                # Display migration status table
```

Flags:
- `--dsn <dsn>`: Override database connection string.
- `--driver <driver>`: Database driver (`sqlite`, `postgres`, `mysql`).
- `--dir <path>`: Migrations directory (default: `migrations`).
- `--lock-timeout <duration>`: Lock acquisition timeout (default: `30s`).

### `krew seed [--class <name>] [--dsn <dsn>] [--driver <driver>] [--dir <dir>]`

Seed the database with test data and initial system state via Go seeders and model factories:

```bash
krew seed                       # Execute DatabaseSeeder
krew seed --class UserSeeder    # Execute a specific seeder class
```

### `krew ui <subcommand>`

Manage Pop-Brutalist `.kiw` design system components and templates:

```bash
krew ui list               # List all available components in Krewire UI
krew ui add <name>         # Install component directly into ./components/
krew ui diff <name>        # Compare local component with upstream template
```

---

## 3. Inspection & Diagnostics

### `krew version`

Print the version of the `krew` CLI and underlying Krewire packages:

```bash
krew version [--json]
```

### `krew info`

Print comprehensive diagnostic information about the current workspace:

```bash
krew info [--json]
```

### `krew compat`

Validate cross-module compatibility against the ecosystem dependency graph:

```bash
krew compat [--json]
```

### `krew context`

Inspect, validate, and lint authoritative project context for human engineers and autonomous AI agents:

```bash
krew context [--lint] [--json] [--scope <level>]
```

- `--lint`: Run context integrity checks and return error if missing or conflicting.
- `--json`: Output complete project context in structured JSON.
- `--scope <level>`: Filter context rules by architectural scope (`workspace`, `module`, `domain`, `service`, `unit`).

### `krew bench`

Run performance benchmarks and evaluate baselines:

```bash
krew bench [--cli | --core | --workflow] [--iterations 50] [--compare] [--json]
```

- `--cli`: Benchmark CLI startup latency and command dispatch.
- `--core`: Benchmark Core primitive performance (parsers, renderers, models).
- `--workflow`: Benchmark end-to-end SDLC workflows.
- `--compare`: Compare results against baseline report and report regressions.
- `--save <path>`: Save benchmark report artifact to disk.

### `krew tier`

Inspect and validate open-core tier licensing posture:

```bash
krew tier [--check free|pro|enterprise]
```

### `krew config`

Validate `krewire.yaml` against its JSON Schema or inspect the schema:

```bash
krew config validate [--strict] [--json]
```

- `--strict`: Treat warnings (unknown or deprecated keys) as failures.
- `--json`: Print validation issues as JSON.

### `krew doctor`

Diagnose developer environment, toolchain prerequisites, and workspace integrity:

```bash
krew doctor
```

### `krew self-update [flags]`

Check for and install the latest stable release of the `krew` binary:

```bash
krew self-update          # Update krew binary to latest stable release
krew self-update --check  # Check if an update is available without downloading
krew self-update --force  # Force reinstall the latest version
```

---

## 4. Build & Preview

### `krew bundle`

Bundle and optimize static CSS and JS assets:

```bash
krew bundle [--minify]
```

### `krew serve`

Preview production build output without watching:

```bash
krew serve [--port 3000] [--dir .krewire/build]
```

---

## 5. Development & Runtime

### `krew run [task|args]`

Run a Go entry point, file, or a custom task defined in `krewire.yaml`:

```bash
krew run [task-name] [-- args...]
```

### `krew vet`

Run Go static analysis, security posture audit, and vulnerability checks:

```bash
krew vet [--sec] [--vuln]
```

### `krew fmt`

Check formatting and format source files according to Go and `.kiw` standards:

```bash
krew fmt [--write]
```

### `krew worker`

Start background worker daemons and process job queues:

```bash
krew worker [--queues default,emails,reports]
```

### `krew gateway`

Run edge API gateway reverse proxy for microservices:

```bash
krew gateway [--port 8080]
```

### `krew mcp`

Start native Model Context Protocol (MCP) server for AI coding agents:

```bash
krew mcp --stdio                      # Standard stdio transport for AI IDEs
krew mcp --port 4000 --host 127.0.0.1 # Standalone HTTP/SSE transport
```

---

## 6. Ship & Operations

### `krew deploy`

Stage deployable artifacts in `.krewire/dist/`:

```bash
krew deploy [--dry-run]
```

### `krew ws <command>`

Manage multi-repository and multi-project workspaces:

```bash
krew ws list
krew ws sync
krew ws run <command>
```

### `krew boost`

Manage the embedded Krewire Boost autonomous AI agent framework:

```bash
krew boost install [path]
krew boost status
```

### `krew release`

Bump versions and stage a release across ecosystem modules:

```bash
krew release framework --bump minor --apply
```

- `--all`: Release every Krewire module (workspace maintainer mode).
- `--bump <level>`: Version component to increment (`patch`, `minor`, `major`).
- `--apply`: Write changes to disk (default is dry-run).
- `--qualify`: Run release qualification checks against compatibility matrix.

---

## 7. Autonomous Agent & Headless

### `krew daemon`

Launch background HTTP/REST and Server-Sent Events (SSE) server for remote agent hosting:

```bash
krew daemon [--port 8080] [--host 0.0.0.0] [--token <secret>]
```

### `krew sessions`

List past conversation sessions with message and token counts:

```bash
krew sessions [--dir <path>]
```

### `krew export <session-id>`

Export a conversation session transcript into clean Markdown:

```bash
krew export <session-id> [--format markdown]
```
