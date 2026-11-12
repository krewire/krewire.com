---
title: "Maturity, Roadmap & Support"
description: "Consistent maturity labels for Krewire workloads and products with linked evidence, a roadmap that separates shipped, in-progress, proposed, and deferred work, realistic maintenance and support expectations, and where to report bugs and security issues."
date: "2026-11-09"
---

# Maturity, Roadmap & Support

This page sets realistic expectations about what is ready, what is experimental, what is planned, and what kind of support exists.

> **Project-wide status:** Krewire is **pre-1.0**. The core monorepo tag is `v0.2.0`, and minor releases may contain breaking changes as documented in the [Upgrade Guide](/overview/upgrade-guide).

---

## Maturity labels

One set of labels is used for workloads **and** products:

| Label | Definition |
| :--- | :--- |
| **Stable** | Covered by automated tests (`go test ./...`), running in a production workload (for example krewire.com or the `kiw` CLI itself), pure Go (`CGO_ENABLED=0`), with documentation and starter scaffolding. |
| **Usable (Beta)** | Core pipelines and drivers are functional and covered by unit tests; suitable for single-node, self-hosted, or early production use; known limitations exist. |
| **Experimental (Alpha)** | Core interfaces and prototypes exist; broader automated and production verification is still underway. A component without public production evidence is not rated above this label. |
| **Planned** | Described as intent (roadmap, README) but not implemented in a tagged version. |

**Important:** "Stable" describes implementation readiness. It is **not** an API-compatibility guarantee — until `v1.0.0`, minor versions may still change public APIs.

---

## Workload maturity

| Workload | Label | Evidence | Notes from verification on this site |
| :--- | :--- | :--- | :--- |
| `site` | Stable | This website is built with `kiw build` (krewire.com). | Verified. |
| `book` | Stable | These docs (English and Indonesian) are rendered by `pkg/book`. | Verified. |
| `cli` | Stable | The `kiw` CLI is itself a `cli` workload (tools/kiw). | Verified. |
| `app` | Stable | Upstream matrix cites SSR + WebAssembly compilation and tests. | The upstream evidence does not cite a public production deployment. |
| `gateway` | Stable | Upstream matrix cites the `api.krewire.com` gateway. | Upstream gateway endpoints for private services are experimental. |
| `worker` | Usable (Beta) | Queue processing and graceful shutdown tested in `tools/kiw`. | Distributed broker adapters are listed as planned upstream. |
| `service` | Usable (Beta) | REST runtime with tracing, circuit breakers, rate limiting. | gRPC generation needs the external `protoc` toolchain. |
| `infra` | Usable (Beta) | Resource planning (`kiw build --plan`) and state reconciliation. | Provider drivers target local Docker and Linux systemd. |

## Product maturity

| Component | Label | Evidence |
| :--- | :--- | :--- |
| `kiw` CLI & core libraries | Per workload (above) | Public source, tag `v0.2.0`. |
| Krew (`krew`) | Experimental | Public source, tag `v0.2.0`; no public production deployment cited. |
| Krewire UI (`ui`) | Experimental | Public source, no tag yet; used by this website's components. |
| Krewire Dash, Cloud, Auth | Experimental | Private source, tag `v0.2.0`; not publicly available. |
| Krewire Forge | Experimental | Private source; `version.go` declares 0.2.0, no tag yet; not publicly available. |

Delivery models and availability are on [Core & Platform Services](/project/core-and-platform-services).

---

## Roadmap

Status definitions:

- **Shipped** — in a tagged version, or merged to `main` (marked "unreleased" until tagged).
- **In progress** — an open issue or pull request exists in a public repository.
- **Proposed** — published as intent in a repository README or document; not scheduled.
- **Deferred** — explicitly postponed.

### Shipped

- `v0.2.0` of the core: `kiw` CLI with scaffolding for all eight workloads, `pkg/*` libraries, the `site`/`book` pipeline that builds this website, and the release/compatibility tooling (`kiw release`, `kiw compat`).
- Advanced multi-driver database, migrations (`schema.Blueprint`), seeders, model factories, and generic repository in `pkg/orm`.
- Model C: Hybrid Go + `.kiw` component architecture, typed Go component API (`pkg/ui/node`, `pkg/ui/widget`), bi-directional Go ↔ `.kiw` interop, and package-based distribution via Go modules (with optional `kiw ui vendor`/`eject` escape hatch).

### In progress

- Standalone `.kiw` component packaging (`.kpkg`) refinements and registry/property-decoder polish.

### Proposed

- **Distributed worker coordination:** Redis and RabbitMQ broker adapters for `worker`.
- **Hierarchical memory & agent mesh:** Advanced autonomous agent features in `krew`.
- **MicroVM isolation:** Sandboxed workspace execution in Krewire Dash.

Years in the proposed items are the targets stated in those READMEs, not commitments made by this page.

### Deferred

No item is currently recorded as deferred.

### Known gaps (not yet on a roadmap)

- No GitHub Release with pre-built `kiw` binaries; the installer falls back to `go install` (Go 1.27.1+).
- No `SECURITY.md`, code of conduct, `CHANGELOG`, or public CI test results for the core repositories.

---

## Maintenance and support expectations

- **Who maintains it:** the Krewire Team, on a best-effort basis.
- **No guaranteed support:** there is no SLA, no guaranteed response time, and no paid support offering.
- **No long-term support (LTS) release** exists yet. Only the latest version on `main` is maintained; fixes are not backported to older tags.
- **Compatibility:** pre-1.0, minor releases may break APIs; breaking changes are expected to come with migration notes in the [Upgrade Guide](/overview/upgrade-guide). Module compatibility is checked with `kiw compat`.
- **Security fixes:** handled as they are reported; no formal disclosure timeline or supported-versions policy is defined yet.

---

## Where to report

| What | Where |
| :--- | :--- |
| Bugs and feature requests | GitHub Issues in the relevant public repository: `krewire`, `krew`, `ui`, `krewire.com` |
| Questions and ideas | GitHub Discussions on `krewire/krewire.com` |
| Release history | Release tags of `krewire/krewire` |
| Contribution standards | [Contributing & Community](/project/contributing) |

### Reporting security issues

Please **do not** open a public issue for a suspected vulnerability. Until a dedicated security policy exists, email [hello@krewire.com](mailto:hello@krewire.com) with the subject line `SECURITY` and include the affected repository, version or commit, and reproduction steps.
