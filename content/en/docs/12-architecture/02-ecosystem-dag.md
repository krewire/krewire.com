---
title: "Ecosystem DAG"
description: "The multi-repository dependency graph of the Krewire ecosystem: compile-time Go module edges, runtime integration protocols, product parallelism, and the checks that keep the graph acyclic."
date: "2026-11-09"
---

# Ecosystem DAG

<nav class="track-crumbs" aria-label="Architecture Deep-Dive track">
<span class="track-crumbs-label">Deep-Dive track</span>
<ol>
<li><a href="/architecture">Overview</a></li>
<li><a href="/architecture/layered-pyramid">1 · Pyramid</a></li>
<li aria-current="step"><span>2 · Ecosystem DAG</span></li>
<li><a href="/architecture/scope-levels">3 · Scope levels</a></li>
<li><a href="/architecture/data-flow">4 · Data flow</a></li>
</ol>
</nav>

Krewire is not one repository. The core libraries and CLI, the UI catalog, the agent engine, the platform services, and the websites live in separate repositories that depend on each other as a **directed acyclic graph (DAG)**. This page shows that graph and the rules that keep it acyclic.

## 1. The Ecosystem Architectural Battery

All components, platforms, and workloads advance along a progressive architectural battery, driven entirely by the `kiw` devtool CLI and declarative `krewire.yaml` schema:

```text
Go Stdlib ──► Krewire Libraries ──► Krewire UI | Krewire Auth | Krewire Krew ──► Krewire Forge | Krewire Cloud ──► Krewire Dash
                                                                                                               (Unified Workspace)
```

## 2. Compile-Time Edges (Go Module Imports)

These edges are Go module dependencies: an arrow means "imports packages from". They were read from each repository's `go.mod` and import statements on 9 November 2026.

```text
                 krewire.com    krewire.github.io        (Layer 5 — built by kiw, no Go module)
                        │              │
                        └──────┬───────┘   uses kiw + kui at build time
                               ▼
     dash ──────┐        cloud ─────┐                    (Layer 4 — products & platform services)
      │  │      │         │  │      │
      │  └──► auth ◄──────┘  │      │        forge
      │         │            │      │          │
      └──► krew ◄────────────┘      │          │
            │   │                   │          │
            ▼   ▼                   ▼          ▼
        ┌──────────────── github.com/krewire/krewire ────────────────┐  (Layers 1–3 — core + kiw)
        │  tools/kiw · pkg/* · pkg/model · pkg/errs · pkg/lifecycle   │
        └─────────────────────────────┬──────────────────────────────┘
                                      ▼
                             github.com/krewire/ui  (kui)              (Layer 2 — no dependencies)
                                      ▼
                             Go standard library                       (Layer 0)
```

| Repository | Module | Imports from the ecosystem |
| :--- | :--- | :--- |
| `ui` (kui) | `github.com/krewire/ui` | nothing — standard library only (`TestModuleHasNoExternalDependencies`) |
| `krewire` (core + `kiw`) | `github.com/krewire/krewire` | `ui` (`pkg/ui/components` re-exports the kui `embed.FS`) |
| `krew` (agent engine) | `github.com/krewire/krew` | `krewire` |
| `auth` | `github.com/krewire/auth` | `krewire` |
| `panel` (Forge) | `github.com/krewire/panel` | `krewire` |
| `cloud` | `github.com/krewire/cloud` | `krewire`, `auth`, `krew` |
| `dash` | `github.com/krewire/dash` | `krewire`, `auth`, `krew` |
| `krewire.com`, `krewire.github.io` | — | none in Go; built with `kiw build`, which renders kui components |

Two properties matter:

- **The graph is acyclic.** `kui` imports nothing, the core imports only `kui`, and every product imports downward. No product imports another product: `dash` and `panel` (Forge) share no Go code.
- **The CLI is a leaf for consumers.** Products import `pkg/*` packages; none imports `tools/kiw`.

## 3. Runtime Integrations (Decoupled Protocols)

Compile-time imports are kept to shared libraries. Products talk to platform capabilities over standard protocols instead:

- **Identity:** OpenID Connect / OAuth 2.0 with public JWKS verification against `auth`.
- **HTTP APIs:** REST behind an edge reverse proxy or a [gateway workload](/workloads/gateway).
- **Events:** an asynchronous event bus or Server-Sent Events streams.

Hosted `cloud` and `auth` are optional. `forge` and `dash` run completely self-hosted; see [Core & Platform Services](/project/core-and-platform-services) for what is publicly available today.

## 4. Product Parallelism

| Product | Responsibility | Depends on the other? |
| :--- | :--- | :--- |
| **Krewire Forge** | Declarative web, form, resource, and admin panel builder for Go and `.kiw`. | No |
| **Krewire Dash** | AI-native virtual developer workspace, execution sandboxes, and collaboration console. | No |

Neither depends on the other at compile time or deployment time, so each can be released, forked, or replaced on its own.

## 5. How the Graph Is Kept Honest

| Check | Where |
| :--- | :--- |
| Downward-only imports and foundation ranking | `krewire/layering_test.go` |
| kui has no module dependencies | `ui/components/contract_test.go` |
| Pure-Go portability | `CGO_ENABLED=0 go build ./...` in CI |
| Formatting and vet | `gofmt -l .`, `go vet ./...` |

When you add a repository or a new import between repositories, place it on a layer first, then make sure the new edge points downward. If two peers need to cooperate, add a protocol (an API, a token, an event) — not an import.

---

**Previous:** [12.1 Layered Pyramid](/architecture/layered-pyramid) · **Next:** [**12.3 Scope Levels →**](/architecture/scope-levels)
