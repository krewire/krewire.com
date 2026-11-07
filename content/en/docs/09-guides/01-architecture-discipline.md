---
title: "Architecture & Engineering Discipline"
description: "Core architectural invariants: the 6-tier Pyramid, strict downward dependencies, peer decoupling via Bridge/Gateway, Single Source of Truth (SSOT), and Single Responsibility Principle."
date: "2026-11-07"
---

# Architecture & Engineering Discipline

The Krewire codebase is governed by non-negotiable architectural invariants defined in `ADR-6CAB8C` and `ADR-E8P9D`.

## 1. The Architectural Pyramid (Strict Downward Dependencies)

The ecosystem is structured strictly as an architectural **Pyramid**:

```text
▲
│                       [ Tier 5: Applications & Portals ]
│               krewire.com · dash (client workspace) · starters · templates
│                                        │
│               [ Tier 4: Enterprise Platforms & Workload Engines ]
│            cloud (PaaS)  ◄───[GATEWAY / BRIDGE]───►  panel (Server Admin)
│                 │                                         │
│                 └───►  auth (SSO / OIDC) ◄────────────────┘
│                                        │
│                     [ Tier 3: Developer Driver & CLI ]
│                                tools/kiw (kiw)
│                                        │
│                 [ Tier 2: Modular Domain Libraries (pkg/*) ]
│                 web · ui · kiw-dsl · orm · cache · resilience · tui
│                                        │
│                    [ Tier 1: Core Foundation & Kernel ]
│               pkg/model · pkg/lifecycle · pkg/errs · pkg/version
│                                        │
│                 [ Tier 0: Go Standard Library & OS Primitives ]
▼
```

- **Strict Downward-Only**: Dependencies flow exclusively downward (`Tier N → Tier < N`). Upward dependencies (`Tier N → Tier > N`) are strictly prohibited and enforced by CI AST checks (`layering_test.go`).
- **Core Kernel Isolation**: `pkg/model` (Tier 1) depends only on Go standard library packages and `pkg/errs`. It never imports higher-level packages or developer tooling.
- **CLI Encapsulation**: No library or platform package may **ever** import `kiw` (Tier 3). Generated user projects depend on Krewire libraries, never on `kiw`.

## 2. Peer Decoupling & The Bridge/Gateway Pattern

Platforms and libraries situated at the same tier are **peers** and must maintain lateral independence:

- **Zero Peer Coupling**: Sibling platforms at Tier 4 (`cloud`, `panel`, `auth`, `dash`, `krew`) must never import each other's source code or private internals.
- **Bridge & Gateway Mediation**: Any communication between peer platforms or across boundaries must traverse a standard Bridge or Gateway:
  - **API Gateway & Reverse Proxy**: Routes HTTP REST and gRPC traffic with route authorization and rate limiting.
  - **Federated Identity Bridge**: Standard OpenID Connect (OIDC), OAuth2 Bearer Tokens, and public JWKS validation (zero shared database credentials).
  - **Event / Message Bus**: Asynchronous Pub/Sub messaging for decoupled event broadcasting.

## 3. Single Source of Truth (SSOT)

A value that the ecosystem must agree upon is declared **once** in the package that owns it:

- **Banned:** Copying literal tables or constant lists into multiple files.
- **Banned:** Re-deriving ordering logic that has an existing canonical resolver.
- **Banned:** Duplicating version strings in markdown when they are declared in code.
- **Required:** One owner per fact. Export the resolver function rather than reimplementing it in consumer packages.

## 4. Pure Go (`CGO_ENABLED=0`) Invariant

Every library and workload binary must cross-compile cleanly to Linux, macOS (`darwin/arm64`), and Windows (`windows/amd64`) without C compiler toolchains.
