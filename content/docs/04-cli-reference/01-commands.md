---
title: "Command Catalog"
description: "Complete command catalog for kiw: syntax, flags, workload support, and examples for every command in the Krewire developer CLI."
date: "2026-10-08"
---

# Command Catalog

The **`kiw`** CLI provides an end-to-end command set spanning the entire software development lifecycle—from kernel initialization to production deployment.

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
- `--book`: Documentation book with Markdown manuscripts and `mdbind`.
- `--worker`: Asynchronous job worker daemon.
- `--service`: Cloud-native microservice.
- `--infra`: Infrastructure as Code in Go.

### `kiw init`

Equip an existing project directory with Krewire workload structures:

```bash
kiw init [--site | --book | --cli | --worker | --service | --infra]
```

- `--template <git-url>`: Clone an external starter template repository into an empty directory.

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
- **`runtime`**: Compiles client-side Go WebAssembly into `public/app.wasm`.
- **`infra`**: Compiles and verifies the cloud infrastructure topology plan.

### `kiw deploy`

Validate the project, run quality gates, stage distribution artifacts into `.krewire/dist/`, and publish to target environments:

```bash
kiw deploy [--target binary|gh-pages|infra] [--branch <name>] [--remote <name>] [--dry-run]
```

- `--dry-run`: Preview deployment steps without publishing or mutating cloud state.
- `--target gh-pages`: Commit and push the static output to GitHub Pages (`gh-pages` branch).
- `--target infra`: Provision declared cloud resources via provider adapters.

---

## 4. Code Quality & Testing

### `kiw test`

Execute unit and integration tests across all Go packages:

```bash
kiw test [--verbose] [--race] [--cover]
```

Spawns `go test ./...` with standardized colorized output and test summary reporting.

### `kiw vet`

Run static code analysis:

```bash
kiw vet
```

Runs Go's official `go vet ./...` analyzer, catching potential defects, dead code, and invalid struct tags.

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

Print the version of the `kiw` CLI and underlying Krewire framework:

```bash
kiw version
```

### `kiw compat`

Validate version compatibility across all module declarations in the workspace:

```bash
kiw compat
```

### `kiw ws <command>`

Manage multi-package Go workspaces (`go.work`):

```bash
kiw ws info         # Show workspace members and active modules
kiw ws list         # List paths of all packages
kiw ws sync         # Synchronize go.work and go.mod dependencies
kiw ws exec <cmd>   # Execute a command across every workspace module
```

### `kiw boost install`

Equip any project repository with the **Krewire Boost** AI agent guild:

```bash
kiw boost install [path]
```

Installs the autonomous agent guild definitions, custom slash commands, and spec-driven SDLC workflows into `.agents/` and `AGENTS.md`.
