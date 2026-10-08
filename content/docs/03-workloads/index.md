---
title: "Workloads"
description: "Comprehensive guide to Krewire's 8 official workload kinds: app, cli, site, book, worker, service, infra, and runtime. Learn architecture, execution models, and progressive scaling."
date: "2026-10-08"
---

# Workloads

Krewire provides a unified developer experience across eight first-class workload kinds. Rather than stitching together disjointed libraries and runtimes, every workload is defined in a single configuration file (**`krewire.yaml`**) and orchestrated through one developer tool (**`kiw`**).

Workload kind selection informs the compiler, dev server, and deployment pipeline while preserving standard Go ergonomics and strict downward-only package dependencies.

---

## What is in this Chapter?

This chapter provides comprehensive architectural documentation, configuration guides, and code examples for each of the eight supported workloads:

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.1</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/app">Fullstack Monolith (app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Single-binary web platforms with Go <code>net/http</code>, embedded static assets, SSR templates, session management, and integrated dependency injection.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.2</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/cli">Terminal CLI &amp; TUI (cli) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Interactive terminal utilities and command trees powered by the lightweight <code>packages/tui</code> toolkit and POSIX-compliant flag parser.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.3</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/site">Static Site Generator (site) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      High-speed marketing sites and landing pages with file-based routing, compile-time scoped <code>.kiw</code> components, and zero JavaScript overhead.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.4</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/book">Documentation Book (book) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-chapter documentation portals and manuals compiled from Markdown manuscripts using the zero-dependency <code>mdbind</code> engine.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.5</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/worker">Background Worker (worker) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Asynchronous queue processing, scheduled cron tasks, retry policies, backoff timers, and dead-letter queues via <code>packages/cloud/worker</code>.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.6</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/service">Microservices (service) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Cloud-native microservices with service discovery, API gateway routing, rate limiting, circuit breaking, and OpenTelemetry distributed tracing.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.7</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/infra">Cloud Infrastructure (infra) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Declarative Infrastructure as Code (IaC) written in Go structs, supporting deterministic execution plans and multi-cloud provisioning.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.8</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/runtime">WebAssembly Runtime (runtime) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Client-side Go WebAssembly runtimes with virtual DOM diffing, partial island hydration, and bi-directional JavaScript interop.
    </p>
  </div>

</div>

---

## Workload Comparison Matrix

| Workload Kind | Entry Point | Primary Packages | Output Artifact | Deployment Target |
| :--- | :--- | :--- | :--- | :--- |
| **`app`** | `cmd/server/main.go` or `main.go` | `web`, `app`, `ui`, `sec` | Single Go Binary | VPS, Docker, VM |
| **`cli`** | `cmd/<name>/main.go` | `tui`, `term`, `kern/errs` | Standalone CLI Binary | User Terminal, PATH |
| **`site`** | `pages/index.kiw` | `kiw`, `web/ssg`, `ui` | Static HTML / CSS / JS | GitHub Pages, S3, Nginx |
| **`book`** | `content/` or `manuscript/` | `mdbind`, `markdown` | Static Book Site | Documentation CDN |
| **`worker`** | `cmd/worker/main.go` | `cloud/worker`, `resilience` | Long-running Daemon | Docker, K8s, Systemd |
| **`service`** | `cmd/service/main.go` | `cloud/service`, `sec` | RPC/HTTP Service | Kubernetes, ECS, Bare Metal |
| **`infra`** | `infra/main.go` | `cloud/infra` | Deployment Plan / State | AWS, GCP, Azure, K8s |
| **`runtime`** | `main.go` (`js/wasm`) | `runtime` | `.wasm` + loader script | Browser / Web Client |

---

## Declaring Workloads in `krewire.yaml`

To set the workload for your project, configure the `project.kind` field:

```yaml
project:
  name: "storefront"
  kind: app # app, cli, site, book, worker, service, infra, runtime
  version: "0.1.0"
  author: "Acme Corp"
```

When `project.kind` is omitted, `kiw` automatically infers the kind using shape markers:
- Directories containing `pages/` or an `ssg:` key are detected as **`site`**.
- Directories containing `content/` or `book.yaml` are detected as **`book`**.
- Directories containing a `worker:` key are detected as **`worker`**.
- Directories containing a `service:` key are detected as **`service`**.
- Directories containing an `infra:` key are detected as **`infra`**.
- Repositories containing `main.go` or `cmd/` without specialized keys default to **`app`**.

> [!TIP]
> Always declare `project.kind` explicitly in `krewire.yaml` to ensure deterministic CLI behavior across CI/CD and developer workstations.
