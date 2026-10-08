---
title: "Configuration Schema (krewire.yaml)"
description: "Canonical reference for krewire.yaml: metadata, build targets, server port bindings, SSG settings, worker queues, and environment overlays."
date: "2026-10-08"
---

# Configuration Schema (`krewire.yaml`)

Krewire operates under the **One Config** rule: all settings across all eight workloads are declared in a single file—**`krewire.yaml`**.

There is no separate `ssg.yaml` or tool-specific configuration file. Configuration is typed, validated against schema rules (`packages/validation`), and loaded using `packages/config`.

---

## Canonical Schema Overview

```yaml
# 1. Project Identification (Required)
project:
  name: "my-service"            # Unique project name (DNS-safe identifier)
  kind: "service"               # Workload kind: app, cli, site, book, worker, service, infra, runtime
  version: "0.1.0"              # Semantic version (SemVer)
  author: "Krewire Team"        # Author or organization name

# 2. Source Content Input (For site and book workloads)
input: "content"                # Root directory for Markdown manuscripts or docs

# 3. Build & Compilation Pipeline
build:
  output: ".krewire/build"       # Output directory for compiled assets or binaries
  base: "/"                     # Base URL path for static asset resolution
  tags: ["netgo"]               # Custom Go build tags
  target: "binary"              # Build target override: binary, site, book, wasm, infra

# 4. Local Development Server
dev:
  port: "8080"                  # HTTP port for local development server
  host: "127.0.0.1"             # Host binding address
  watch:                        # Custom directories to monitor for file changes
    - "pages"
    - "components"
    - "layouts"
    - "assets"

# 5. Workload-Specific Settings (Optional depending on kind)

# Background Worker Settings (kind: worker)
worker:
  concurrency: 10               # Maximum concurrent workers
  queue: "memory"               # Queue connection URL (memory, redis://, nats://)
  dead_letter_queue: "failed"   # Queue for failed messages

# Microservice Settings (kind: service)
service:
  port: "8080"
  discovery:
    provider: "kubernetes"      # consul, etcd, kubernetes, static
    service_name: "my-service"
  rate_limit:
    requests_per_second: 500
    burst: 100

# Cloud Infrastructure Settings (kind: infra)
infra:
  provider: "aws"               # aws, gcp, kubernetes
  region: "ap-southeast-1"
  state:
    backend: "s3"
    bucket: "krewire-infra-state"
```

---

## Schema Sections Reference

### `project`

| Key | Type | Required | Description |
| :--- | :--- | :---: | :--- |
| `name` | String | Yes | Name of the project or service. |
| `kind` | String | Yes | Workload kind (`app`, `cli`, `site`, `book`, `worker`, `service`, `infra`, `runtime`). |
| `version` | String | Yes | Current semantic release version string (`vX.Y.Z`). |
| `author` | String | No | Author, team, or maintainer name. |

### `build`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `output` | String | `.krewire/build` | Directory path where build artifacts are written. |
| `base` | String | `/` | Base URL prefix for static links and assets. |
| `tags` | Array | `[]` | Go build tags passed to the compiler. |
| `target` | String | auto | Explicit target compiler dispatch override. |

### `dev`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `port` | String / Int | `8080` | Local port for the dev server (can be overridden by `KIW_PORT`). |
| `host` | String | `127.0.0.1` | Network interface address to bind to. |
| `watch` | Array | auto | Additional filesystem paths to observe for live reload. |

---

## Environment Overlays & Dotenv (`.env`)

Krewire automatically loads `.env` files from the project root during local development:

```bash
# .env
KIW_PORT=3000
DATABASE_URL=postgres://user:pass@localhost:5432/db
```

Values declared in environment variables take precedence over settings in `krewire.yaml`.
