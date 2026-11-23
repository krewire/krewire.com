---
title: "Krewire Workloads"
description: "In-depth guide to Krewire's 8 official workload kinds: app, cli, site, book, worker, service, gateway, and infra."
date: "2026-09-29"
---

# Krewire Workloads

Krewire unifies eight distinct architectural workload kinds under one command-line interface (**`krew`**) and one declarative devtool configuration (**`krewire.yaml`**).

Krewire handles 8 progressive workloads within a unified project architecture without rigid workload boundaries or mode locking. Declare capabilities and configuration directly in `krewire.yaml`:

```yaml
project:
  name: "my-service"
  version: "0.1.0"
```

---

## The 8 Workloads Matrix

| Workload | Kind Name | Primary Use Case | CLI Execution | Key Packages |
| :--- | :--- | :--- | :--- | :--- |
| **Fullstack Monolith** | `app` | Web applications, SSR, JSON APIs, sessions, database access | `krew run`, `krew dev` | `pkg/web`, `pkg/app`, `pkg/ui` |
| **Terminal CLI & TUI** | `cli` | Command-line utilities, interactive developer TUIs | `krew run`, `krew new my-cli --cli` | `pkg/tui` |
| **Static Site (SSG)** | `site` | Landing pages, marketing websites, blogs, portfolios | `krew build`, `krew dev` | `pkg/web/ssg`, `pkg/kiw` |
| **Technical Book** | `book` | Multi-chapter documentation, books, software manuals | `krew build --target book`, `krew dev` | `pkg/book` |
| **Background Worker** | `worker` | Asynchronous job queues, cron schedules, retries, DLQ | `krew worker` | `pkg/cloud/worker` |
| **Microservice** | `service` | Distributed RPC/HTTP APIs, service registry, gateways | `krew run`, `krew dash` | `pkg/cloud/service` |
| **Cloud Infra (IaC)** | `infra` | Provider-agnostic cloud resource provisioning (AWS/K8s) | `krew deploy --target infra` | `pkg/cloud/infra` |
| **API Gateway** | `gateway` | Declarative reverse proxy, per-route middleware, service-discovery balancing | `krew gateway`, `krew run` | `pkg/web/gateway` |

---

## 1. Fullstack Monolith (`app`)

The `app` workload produces a single, self-contained static Go binary that embeds static assets, templates, database migrations, and HTTP routing:

### Characteristics

- **Zero External Dependencies:** Built on Go's `net/http` and `embed.FS`. No external web servers (e.g. Apache, Nginx) required for local development or basic deployment.
- **Integrated Template Engine:** Server-Side Rendering (SSR) via `html/template` or `.kiw` component compilation.
- **REST & JSON APIs:** Expressive routing with groups, parameter binding, and RFC 7807 Problem Details.
- **Modular Monolith Ready:** Internal domain boundaries allow future extraction of services without code rewrites.

### Scaffolding & Running

```bash
krew new my-app --app
cd my-app
krew run
```

---

## 2. Terminal CLI & TUI (`cli`)

The `cli` workload is tailored for building command-line utilities, developer tools, and rich interactive terminal user interfaces:

### Characteristics

- **Terminal Model:** Powered by `pkg/tui` with an Elm/Bubble Tea-inspired Model-Update-View architecture.
- **Flag Parsing & Commands:** POSIX-compliant flag parser with automatic `--help` generation and exit-code discipline (`pkg/errs`).
- **Structured Logging:** Integrated with Go's `log/slog` for structured, colorized terminal output.

### Scaffolding & Running

```bash
krew new my-cli --cli
cd my-cli
krew run -- --verbose
```

---

## 3. Static Site Generator (`site`)

The `site` workload provides an ultra-fast static site generator powered by the `.kiw` component DSL:

### Characteristics

- **File-Based Routing:** `pages/index.kiw` serves `/`; `pages/about.kiw` serves `/about`.
- **Zero Client-Side JS:** Emits pure semantic HTML and scoped CSS by default.
- **Tailwind CSS Integration:** Automatically detects `tailwind.config.js` and compiles minified utility classes.
- **Incremental Builds:** Fast rebuilds during development (`krew dev`) with millisecond turnaround.

### Scaffolding & Running

```bash
krew new my-site --site
cd my-site
krew dev # Starts dev server at localhost:8080 with live reload
```

---

## 4. Technical Book (`book`)

The `book` workload transforms Markdown manuscripts into responsive, reader-friendly documentation portals via the integrated `pkg/book` engine:

### Characteristics

- **Automated Chapter Numbering:** Numbered directory and file prefixes (`01-overview.md`) map to `1. Overview`, `1.1 Subchapter`, etc.
- **Responsive Navigation:** Interactive collapsible sidebar, chapter filter search input, and mobile drawer toggler.
- **Dark/Light Theming:** Persistent theme switcher syncing with OS preferences via `localStorage`.
- **Hybrid Site Coexistence:** A book can be mounted cleanly under `/docs/` while an SSG landing page owns `/` in the same output directory.

### Scaffolding & Running

```bash
krew new my-docs --book
cd my-docs
krew dev
```

---

## 5. Background Worker (`worker`)

The `worker` workload runs background asynchronous job processors, distributed task queues, and scheduled cron jobs:

### Characteristics

- **Robust Retries & Backoff:** Configurable exponential backoff, jitter, and maximum retry attempts.
- **Dead-Letter Queues (DLQ):** Failed jobs are safely quarantined for inspection and replay.
- **Graceful Shutdown:** Respects OS signals (`SIGINT`, `SIGTERM`), draining active tasks before exit.

### Running

```bash
krew worker --queue memory
```

---

## 6. Microservice (`service`)

The `service` workload equips high-throughput Go microservices with cloud-native distributed service patterns:

### Characteristics

- **Service Registry & Discovery:** First-class integrations with Consul, etcd, NATS, and Kubernetes DNS.
- **API Gateway Routing:** Path rewriting, weighted canary traffic shifting, and token bucket rate limiting.
- **Resilience Primitives:** Circuit breakers, request timeouts, and concurrency bulkheads.
- **Distributed Tracing:** Native OpenTelemetry (OTel) instrumentation with W3C `traceparent` propagation.

### Running

```bash
krew run
```

---

## 7. Cloud Infrastructure as Code (`infra`)

The `infra` workload allows developers to declare cloud architecture directly in Go:

### Characteristics

- **Declarative Go Structs:** Define VPCs, databases, object storage, and compute clusters using typed Go declarations instead of HCL or YAML.
- **Deterministic Planning:** `krew deploy --target infra --plan` generates an execution graph with zero side-effects before applying.
- **Provider Support:** Multi-cloud drivers covering AWS (ECS, Lambda, RDS, S3, CloudFront) and Kubernetes (Deployments, Services, Ingress).

### Running

```bash
krew deploy --target infra --plan
krew deploy --target infra
```

---

## 8. API Gateway (`gateway`)

The `gateway` workload is a standalone reverse proxy that fronts one or more upstream services with a declarative route table:

### Characteristics

- **Declarative Routing:** Exact and prefix routes, per-route method matching, and either a fixed `target_url` or a registry-resolved `service`.
- **Per-Route Middleware:** CORS, tracing, timeout, rate limiting, and circuit breaking are registered by name and composed per route.
- **Service Discovery:** Round-robin balancing over registered endpoints; an unavailable upstream returns `502` with a structured Problem JSON body.
- **One Route Table, Two Runners:** `krew run` executes the project binary; `krew gateway` runs the identical `gateway:` block without a rebuild.

### Running

```bash
krew gateway --port 8080 --routes "/api=user-service"
```

> [!NOTE]
> Client-side WebAssembly is a **library**, not a workload kind. `pkg/runtime`
> compiles Go to WASM inside an `app` or `site` project via
> `krew build --target wasm` — see [Web Runtime & WASM](/web-ui/web-runtime).

---

## Summary

By selecting a workload kind in `krewire.yaml`, developers access tailored ergonomics while retaining the universal benefits of the Go runtime and standard library.

Proceed to [**1.4 Upgrade Guide →**](/overview/upgrade-guide) to understand versioning, deprecations, and configuration migrations.
