---
title: "Configuration Schema (krewire.yaml)"
description: "Canonical reference for krewire.yaml: project identity, build and content inputs, dev server, book/SSG/i18n blocks, and every workload-specific section."
date: "2026-10-25"
---

# Configuration Schema (`krewire.yaml`)

Krewire operates under the **One Config** rule: all settings across all eight workloads are declared in a single file—**`krewire.yaml`**.

There is no separate `ssg.yaml` or tool-specific configuration file. Configuration is typed, validated against schema rules (`pkg/validation`), and loaded using `pkg/config` (schema owner: `tools/kiw/internal/config`).

---

## Canonical Schema Overview

```yaml
# 1. Project Identity
project:
  name: "my-service"            # Unique project name (DNS-safe identifier)
  kind: "service"               # app, cli, site, book, worker, service, gateway, infra
  version: "v0.1.0"             # Product version, injected as .Version into pages

title: "My Service"             # Site/book title
author: "Krewire Team"          # Site/book author
env: "local"                    # local, production, testing (KIW_ENV overrides)

# 2. Content Input
input: "content"                # Content directory for site and book workloads

# 3. Build Pipeline
build:
  output: ".krewire/build"      # Output directory (default .krewire/build)
  base: "/"                     # URL base the site is served under
  include: ["**/*.md"]          # Content globs to include
  exclude: ["**/README.md"]     # Content globs to drop after include
  depends_on: []                # Shell commands run before the build

# 4. Local Development Server
dev:
  port: 8080                    # Dev server port (integer; --addr overrides)

# 5. Workload-Specific Settings (only the block matching project.kind applies)

# kind: worker
worker:
  concurrency: 10               # Maximum concurrent workers
  queues: ["default", "email"]  # Queue names this worker consumes

# kind: service
service:
  name: "user-service"
  port: 8080

# kind: gateway
gateway:
  host: ""
  port: 8080
  routes:
    - path: /api
      prefix: true
      strip_prefix: true
      target_url: http://localhost:8081
      timeout: 10s
    - path: /users
      service: user-service
      middleware: [cors]
      rate_limit: 100

# kind: infra
infra:
  provider: "aws"               # aws, gcp, kubernetes
  region: "ap-southeast-1"
  state: "s3://krewire-infra-state"
```

---

## Schema Sections Reference

### `project`

| Key | Type | Required | Description |
| :--- | :--- | :---: | :--- |
| `name` | String | Yes | Name of the project or service. |
| `kind` | String | Yes | Workload kind (`app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`). |
| `version` | String | No | Semantic release version (`vX.Y.Z`); folded into the effective version used by every pipeline. |

A pre-scaffold bootstrap project is **not** a kind: `kiw info` reports `none`
with the marker `pre-init` until `kiw init` equips a variant.

### Top-level metadata

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `title` | String | module name | Site or book title. |
| `author` | String | — | Author shown in book pages and metadata. |
| `version` | String | `project.version` | Product version injected as `.Version`. |
| `env` | String | `local` | Target environment (`local`, `production`, `testing`). |
| `debug` | Bool | `false` | Enable debug mode (`KIW_DEBUG` overrides). |
| `input` | String | `content` | Content directory for site and book workloads. |
| `nav` | Array | `[]` | Navbar links (`text` + `url`). |
| `footer` | String | — | Footer text. |
| `scripts` | Map | `{}` | Named shell tasks run with `kiw run <task>`. |

### `build`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `output` | String | `.krewire/build` | Directory where build artifacts are written. |
| `base` | String | `/` | Base URL prefix for static links and assets. |
| `include` | Array | `["**/*.md"]` | Content globs to include; an empty list disables the filter. |
| `exclude` | Array | README notes | Content globs removed after inclusion. |
| `depends_on` | Array | `[]` | Shell commands run in order before the build. |
| `outputs` | Map | `{}` | Runtime name → output directory, for polyglot projects. |

The top-level `output:` and `base:` keys are legacy aliases; the nested
`build.*` value wins when both are set, and the `--output` / `--base` flags win
over both. Declare the value once — do not repeat it at both levels.

### `dev`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `port` | Int | `8080` | Local dev server port; the `--addr` flag wins over it. |

### `book`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `mount` | String | — | Sub-path the book occupies in output and URL space (e.g. `/docs/`). |
| `toc` | Bool | auto | Force the generated root table-of-contents page on or off. |

### `i18n`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `default_locale` | String | `en` | Locale served at the site root. |
| `fallback_locale` | String | `en` | Locale used when a key is missing. |
| `base_path` | String | `lang` | Directory holding `<locale>/site.json` dictionaries. |
| `locales` | Array | `[]` | Locales to build (`content/<locale>/docs/` → `.krewire/build/<locale>/docs/`). |

### `ssg`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `description` | String | — | Injected into every page's template data. |
| `layouts` | Array | `[]` | Named layouts (`name`, `body`, `style`, or a `ui` shell). |
| `components` | Array | `[]` | Named components (`name`, `body`, `style`). |
| `pages` | Array | `[]` | Output pages (`path`, `title`, `layout`, `root`, `data`). |
| `assets` | Map | `{}` | Output path to file contents. |
| `auto_assets` | Block | on | CSS/JS injection; see [Configuration](/getting-started/configuration). |

A project with a `pages/` directory builds as `site` without this block. The
`ssg:` key is required only for a fully declarative site that has no page files
on disk.

### `wasm`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `entry` | String | `./wasm` | Main package built with `GOOS=js GOARCH=wasm`. |
| `name` | String | `runtime` | Output module name. |

Client-side WebAssembly is a **library target**, not a workload kind: build it
with `kiw build --target wasm` inside an `app` or `site` project.

### `worker`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `concurrency` | Int | `1` | Maximum concurrent workers. |
| `queues` | Array | `[]` | Queue names this worker consumes. |

### `service`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `name` | String | `project.name` | Service name used for registration and discovery. |
| `port` | Int | `8080` | Port the service binds. |

`services` accepts a list of the same shape when one project declares several
services.

### `gateway`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `host` | String | `""` | Bind address (empty means all interfaces). |
| `port` | Int | `8080` | Port the gateway binds (`GATEWAY_ADDR` overrides both). |
| `routes` | Array | `[]` | Route table; see below. |

Route keys:

| Key | Type | Description |
| :--- | :--- | :--- |
| `path` | String | Exact path, or prefix when `prefix: true`. |
| `method` | String | Optional HTTP method filter. |
| `service` | String | Upstream resolved through the service registry. |
| `target_url` | String | Fixed upstream URL (alternative to `service`). |
| `prefix` | Bool | Treat `path` as a prefix match. |
| `strip_prefix` | Bool | Remove `path` from the upstream request path. |
| `middleware` | Array | Names registered on the gateway via `WithMiddleware`. |
| `rate_limit` | Int | Requests per second for this route (enforced). |
| `timeout` | String | Per-route deadline, e.g. `10s` (enforced; exceeded → `504`). |

### `infra`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `provider` | String | `aws` | Target provider (`aws`, `gcp`, `kubernetes`). |
| `region` | String | — | Provider region. |
| `state` | String | — | State backend location. |
| `resources` | Array | `[]` | Declared resources applied by `kiw deploy --target infra`. |

### `runtime`

| Key | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `namespace` | String | — | Prefix for Krewire environment variables (e.g. `go` → `GO_*`). |
| `env_file` | String | `.env` | Path to the environment file loaded at startup. |
| `port` | Int | — | Default runtime port; the first port consulted when resolving a listen address. |

---

## Environment Overlays & Dotenv (`.env`)

`kiw` loads the environment file from the project root (path from
`runtime.env_file`, default `.env`) without overwriting variables that are
already set in the process environment:

```bash
# .env
DATABASE_URL=postgres://user:pass@localhost:5432/db
LOG_LEVEL=debug
```

`KIW_ENV` and `KIW_DEBUG` are read directly from the environment and take
precedence over the `env:` and `debug:` keys in `krewire.yaml`.

The dev server port is **not** an environment variable. The listen address
resolves in this order:

1. `--addr` flag (`kiw serve --addr :3000`).
2. `runtime.port`, then `dev.port`, then `server.port`, then the first
   `services[].port`.
3. `:8080`.
