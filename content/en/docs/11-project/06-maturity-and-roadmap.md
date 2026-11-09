---
title: "Maturity, Roadmap & Support"
description: "Consistent maturity labels for Krewire workloads and products with linked evidence, a roadmap that separates shipped, in-progress, proposed, and deferred work, realistic maintenance and support expectations, and where to report bugs and security issues."
date: "2026-11-09"
---

# Maturity, Roadmap & Support

This page sets realistic expectations about what is ready, what is experimental, what is planned, and what kind of support exists. Status was last verified on **2026-11-09** against the repositories.

> **Project-wide status:** Krewire is **pre-1.0**. The core monorepo's latest tag is `v0.2.0` (2026-11-05), and no GitHub Release with downloadable binaries has been published. Under the project's [release and compatibility specification](https://github.com/krewire/krewire/blob/main/docs/foundation/release-and-compatibility.md), `0.y.0` releases may contain breaking changes, documented in the [Upgrade Guide](/overview/upgrade-guide).

---

## Maturity labels

One set of labels is used for workloads **and** products. The first three definitions are taken from the core repository's [workload maturity matrix](https://github.com/krewire/krewire/blob/main/docs/workloads/maturity-matrix.md).

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
| `site` | Stable | This website is built with `kiw build` ([repository](https://github.com/krewire/krewire.com)). | Verified. |
| `book` | Stable | These docs (English and Indonesian) are rendered by `pkg/book`. | Verified. |
| `cli` | Stable | The `kiw` CLI is itself a `cli` workload ([source](https://github.com/krewire/krewire/tree/main/tools/kiw)). | Verified. |
| `app` | Stable | Upstream matrix cites SSR + WebAssembly compilation and tests. | The upstream evidence does not cite a public production deployment. |
| `gateway` | Stable | Upstream matrix cites the `api.krewire.com` gateway. | The gateway responds publicly, but its `/auth`, `/cloud`, and `/dash` upstreams returned HTTP 502 on 2026-11-09. |
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
| Panel | Experimental | Private source; `version.go` declares 0.2.0, no tag yet; not publicly available. |

Delivery models and availability are on [Core & Platform Services](/project/core-and-platform-services).

---

## Roadmap

Status definitions:

- **Shipped** — in a tagged version, or merged to `main` (marked "unreleased" until tagged).
- **In progress** — an open issue or pull request exists in a public repository.
- **Proposed** — published as intent in a repository README or document; not scheduled.
- **Deferred** — explicitly postponed.

### Shipped

- `v0.2.0` of the core (2026-11-05): `kiw` CLI with scaffolding for all eight workloads, `pkg/*` libraries, the `site`/`book` pipeline that builds this website, and the release/compatibility tooling (`kiw release`, `kiw compat`).
- Unreleased on `main` of `krewire/krewire` after `v0.2.0`: `kiw self-update`/`kiw update`, a release specification with workload integration tests and interoperability docs, and nested subchapters in `pkg/book` ([commits](https://github.com/krewire/krewire/compare/v0.2.0...main)).

### In progress

| Item | Tracking |
| :--- | :--- |
| Website: mission, audiences, openness, core/services boundary, community journey, impact report, maturity & roadmap (this chapter) | [`krewire/krewire.com` issues #15–#21](https://github.com/krewire/krewire.com/issues) |
| UI system contract and integration boundary with `.kiw` | [`krewire/ui` #1](https://github.com/krewire/ui/issues/1) |
| Community portal: path from discovery to first project; sponsorship and support information | [`krewire/krewire.github.io` #8, #9](https://github.com/krewire/krewire.github.io/issues) |

### Proposed

- **Core implementation gaps** (from the [maturity matrix](https://github.com/krewire/krewire/blob/main/docs/workloads/maturity-matrix.md)): zero-configuration build success for all eight starters (P0); distributed worker coordination (P1); standard Prometheus endpoints for `service` and `gateway` (P1).
- **Krew** ([README roadmap](https://github.com/krewire/krew#readme)): hierarchical memory, OpenTelemetry step tracing, and agent identity (stated for 2027); agent-to-agent mesh and `v1.0.0` LTS (stated for 2028).
- **Krewire Dash** (private README roadmap): MicroVM isolation and snapshots (stated for 2027); collaborative editing and `v1.0.0` LTS (stated for 2028).

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
| Bugs and feature requests | GitHub Issues in the relevant public repository: [`krewire`](https://github.com/krewire/krewire/issues), [`krew`](https://github.com/krewire/krew/issues), [`ui`](https://github.com/krewire/ui/issues), [`krewire.com`](https://github.com/krewire/krewire.com/issues) |
| Questions and ideas | [GitHub Discussions on `krewire/krewire.com`](https://github.com/krewire/krewire.com/discussions) |
| Release history | [Tags of `krewire/krewire`](https://github.com/krewire/krewire/tags) — written release notes are not yet published |
| Contribution standards | [Contributing & Community](/project/contributing) |

### Reporting security issues

Please **do not** open a public issue for a suspected vulnerability. Until a dedicated security policy exists, email [hello@krewire.com](mailto:hello@krewire.com) with the subject line `SECURITY` and include the affected repository, version or commit, and reproduction steps.

As of 2026-11-09, no repository has a `SECURITY.md`, and GitHub private vulnerability reporting is not enabled on the public repositories. This section will be updated when either changes.

---

## How this page is kept current

- **Owner:** Krewire Team (maintainers of [`krewire/krewire.com`](https://github.com/krewire/krewire.com)).
- **When:** reviewed as part of every tagged release of this website (see the "Release checklist" in this repository's `AGENTS.md`), and whenever a workload's maturity, a product's availability, or a roadmap item changes.
- **How:** re-verify tags (`git tag`), repository visibility, the upstream maturity matrix, open issues, and live endpoints; update the tables and the "last verified" date at the top of this page.
