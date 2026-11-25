---
title: "Layered Pyramid & Invariants"
description: "Core architectural invariants: the Layer 0–5 pyramid, strict downward dependencies, peer decoupling via Bridge/Gateway, Single Source of Truth (SSOT), and the pure-Go rule."
date: "2026-11-09"
---

# Layered Pyramid & Invariants

<nav class="track-crumbs" aria-label="Architecture Deep-Dive track">
<span class="track-crumbs-label">Deep-Dive track</span>
<ol>
<li><a href="/architecture">Overview</a></li>
<li aria-current="step"><span>1 · Pyramid</span></li>
<li><a href="/architecture/ecosystem-dag">2 · Ecosystem DAG</a></li>
<li><a href="/architecture/scope-levels">3 · Scope levels</a></li>
<li><a href="/architecture/data-flow">4 · Data flow</a></li>
</ol>
</nav>

The Krewire codebase is governed by non-negotiable architectural invariants across all packages and workloads. They are enforced by tests, not by convention.

## 1. The Ecosystem Architectural Battery

The entire Krewire ecosystem flows through a unified architectural battery progression:

```text
Go Stdlib ──► Krewire Libraries ──► Krewire UI | Krewire Auth | Krewire Krew ──► Krewire Cloud ──► Krewire Dash
                                                                                                  (Unified Workspace)
```

Every stage in this progression is driven and composed through the `krew` developer CLI and configured declaratively via `krewire.yaml`.

## 2. The Ecosystem Pyramid (Strict Downward Dependencies)

Seen from the whole ecosystem, every repository sits on one layer of a pyramid (source: `krewire/docs/foundation/ecosystem-dag.md`):

```text
Layer 5: Unified Studio & Apex Portals     Dash (virtual workspace) · krewire.com · krewire.github.io
   ↓
Layer 4: Platform Engines & Cloud          Cloud (PaaS) ∥ Auth (IAM / SSO) ∥ Krew (AI agent)
   ↓
Layer 3: Universal Developer Driver & CLI (krew driven by krewire.yaml)
   ↓
Layer 2: UI Foundation & Domain Libraries  ui (.kiw components) · hub · krewire modules (web, orm, sec, …)
   ↓
Layer 1: Core Foundation & Kernel          model · errs · version · lifecycle
   ↓
Layer 0: Go Standard Library & OS Primitives
```

- **Strict downward-only:** a layer may import only layers below it (`Layer N → Layer < N`). Upward imports are forbidden and caught by `layering_test.go`.
- **Kernel isolation:** `model` depends only on the standard library and `errs`.
- **CLI encapsulation:** no library or platform package may import `krew`. Generated user projects depend on Krewire libraries, never on the CLI.

## 3. Inside the Core Repository

Zooming into `github.com/krewire/krewire`, the same idea is applied as a ranked ladder of packages (source: `krewire/docs/foundation/architecture.md`):

| Layer | Contents | May import |
| :--- | :--- | :--- |
| 5 | Scaffolding blueprints (`templates/*`) | Layers 0–4 |
| 4 | Developer tooling (`krew`) | Layers 0–3 |
| 3 | Domain libraries (`web`, `sec`, `orm`, `ui`, …) | Layers 0–2 and sibling domain packages, acyclically |
| 2 | Execution & supervision (`lifecycle`, `env`, `log`) | Layers 0–1 |
| 1 | Declarative core (`errs`, `model`, `version`) | Standard library (`model`, `version` may import `errs`) |
| 0 | Go standard library | — |

Five invariants keep this ladder honest:

| # | Invariant | Enforced by |
| :--- | :--- | :--- |
| 1 | **Downward-only directionality** — every import edge goes to a lower layer, or stays acyclic inside Layer 3. | `TestFoundationLayering`, `TestUIDoesNotImportWeb` in `layering_test.go` |
| 2 | **Base floor** — `errs` imports only the standard library; `model` and `version` add only `errs`. | `errs/dependency_test.go`, `model/dependency_test.go` |
| 3 | **No upward names** — foundation packages contain no string literals naming higher packages. | `TestFoundationNamesNoHigherPackage` |
| 4 | **Hermetic core** — `model` has no third-party or other Krewire imports. | `TestModel_StdlibBoundary` |
| 5 | **Tooling isolation** — no foundational domain module or `templates/` imports `krew`. | `TestPackagesNeverImportKiw` |

## 4. Peer Decoupling & the Bridge/Gateway Pattern

Components on the same layer are **peers** and stay laterally independent:

- **Zero peer coupling:** parallel products (`dash`, `cloud`, `auth`, `krew`) never import each other's Go packages or private source trees.
- **Bridge & gateway mediation:** communication between peers, or with the horizontal capabilities, goes through standard interfaces:
  - **API gateway / reverse proxy** for HTTP and gRPC traffic with per-route authorization and rate limiting.
  - **Federated identity** — OpenID Connect, OAuth 2.0 bearer tokens, and public JWKS validation, with no shared database credentials.
  - **Event bus / SSE streams** for asynchronous notifications.
- **No mandatory coupling:** hosted `cloud` and `auth` are never prerequisites; `dash` runs self-hosted without them.

The [Ecosystem DAG](/architecture/ecosystem-dag) page lists the module edges that actually exist between the repositories today.

## 5. Single Source of Truth (SSOT)

A value the ecosystem must agree on is declared **once**, in the package that owns it:

- **Banned:** copying literal tables or constant lists into multiple files.
- **Banned:** re-deriving ordering logic that already has a canonical resolver.
- **Banned:** duplicating version strings in Markdown when they are declared in code.
- **Required:** one owner per fact; export the resolver instead of re-implementing it in consumers.

## 6. Pure Go (`CGO_ENABLED=0`)

Every library and workload binary must cross-compile to Linux, macOS (`darwin/arm64`), and Windows (`windows/amd64`) without a C toolchain. CI checks this with `CGO_ENABLED=0 go build ./...`. When an application needs a database driver, choose a pure-Go one — the [fullstack tutorial](/getting-started/fullstack-app-tutorial) uses `modernc.org/sqlite` for exactly this reason.

---

**Next:** [**11.2 Ecosystem DAG →**](/architecture/ecosystem-dag)
