---
title: "Core & Platform Services"
description: "The boundary between Krewire's open-source core and the optional platform services (Workspace/Dash, Cloud, Auth, Panel): delivery model, maturity, dependencies, availability, and pricing status."
date: "2026-11-09"
---

# Core & Platform Services

Krewire has two layers:

1. **The open-source core** — the public, MIT-licensed repositories (`krewire`, `krew`, `ui`) that contain the libraries, the `kiw` CLI, and the tooling for all eight workloads.
2. **Optional platform services** — Krewire Workspace (Dash), Cloud, Auth, and Panel, developed in **private** repositories.

The core is complete on its own. Platform services are meant to add operational convenience on top of it; they are never a prerequisite for building, running, or deploying software with the core.

> **Status as of 2026-11-09:** none of the four platform services is publicly available — there is no public sign-up, download, or hosted endpoint. The subdomains `dash.krewire.com`, `cloud.krewire.com`, `auth.krewire.com`, and `panel.krewire.com` currently serve this website. No prices, plans, or commercial licenses are published.

---

## At a glance

| Component | Source | Delivery model | Availability today | Maturity |
| :--- | :--- | :--- | :--- | :--- |
| **Core libraries & `kiw` CLI** (`krewire`) | Public, MIT | Self-hosted (you run it) | Available — `go install` or source | Per workload, see [Maturity](/project/maturity-and-roadmap) |
| **Krew** AI agent (`krew`) | Public, MIT | Self-hosted (embed, CLI, HTTP/SSE, MCP) | Available — source | Experimental |
| **Krewire UI** (`ui`) | Public, MIT | Copy-and-own components | Available — source | Experimental |
| **Krewire Workspace** (Dash) | Private | Intended: hosted service | Not publicly available | Experimental |
| **Krewire Cloud** | Private | Intended: hybrid — managed SaaS and self-hosted engine | Not publicly available | Experimental |
| **Krewire Auth** | Private | Intended: hosted identity service and standalone server | Not publicly available | Experimental |
| **Krewire Panel** | Private | Intended: self-hosted single binary | Not publicly available | Experimental |

Maturity labels use the shared definitions in [Maturity, Roadmap & Support](/project/maturity-and-roadmap). "Intended" delivery models come from each private repository's README; they describe design intent, not a service you can use today.

---

## Is the core paywalled?

**No.** This is a verifiable licensing fact, not a promise about the future:

- Every public repository is MIT-licensed (see [Openness & Self-Hosting](/project/openness-and-self-hosting)).
- The core declares its own tier in code as free: [`version.go`](https://github.com/krewire/krewire/blob/main/version.go) sets `Tier = model.TierFree` with the comment *"All code in krewire/krewire is 100% open source; all packages are currently free-tier by default."*
- `pkg/model` defines a tier ladder (`free`, `pro`, `team`, `enterprise`) and `kiw tier --check <tier>` can assert that a project stays within a tier. **No `pro`, `team`, or `enterprise` packages, licenses, or prices exist publicly today.** The ladder is a classification mechanism, not a price list.
- The core's stated design commitment ([Philosophy](/overview/philosophy)) is that no paid license is ever a prerequisite to build, deploy, or operate a complete system with it.

---

## What hosted services would add

The value of a hosted service is **operating** software, not access to it. When a platform service becomes available, the operational value it is designed to provide beyond the underlying software is:

- running, upgrading, and monitoring the service so you do not have to;
- managed multi-tenant isolation, backups, and TLS;
- integration across services (for example, single sign-on from Auth into Workspace and Cloud).

None of this is offered today. Until it is, the equivalent path is to self-host your own output with the core (see [Deployment & Operations](/guides/deployment-and-operations)).

---

## Product details

Descriptions below summarise each private repository's README. Functionality cannot be inspected publicly while the source is private.

### Krewire Workspace (Dash)

- **What it is:** an AI-native virtual workspace server — sandboxed workspaces, a browser PTY terminal, real-time presence, and an embedded `krew` agent, shipped as one daemon (`krewire-dash-server`).
- **Depends on:** `krewire` (core), `krew`, and `auth` (minimum versions declared in its `version.go`).
- **Delivery model:** intended as a hosted service at `dash.krewire.com`. Not live.
- **Version:** v0.2.0 (git tag in the private repository).
- **Docs:** [AI-Native Virtual Workspace](/security-platforms/virtual-workspace).

### Krewire Cloud

- **What it is:** a hybrid PaaS engine and multi-tenant control plane — hermetic Go builds, a process supervisor, a host-based edge gateway, and quota/metering (`pkg/billing`) designed for future SaaS billing.
- **Depends on:** `krewire` (core) and `auth`.
- **Delivery model:** intended as hybrid — a managed multi-tenant service and a self-hosted engine (`krewire-cloud-server`). Not publicly available in either form.
- **Version:** v0.2.0 (git tag in the private repository).
- **Docs:** [Hybrid PaaS Platform](/security-platforms/cloud-paas).

### Krewire Auth

- **What it is:** an identity server — PBKDF2 password hashing, TOTP MFA, OIDC/SAML 2.0 SSO, hierarchical RBAC, and audit logging, served by `krewire-auth-server`.
- **Depends on:** `krewire` (core).
- **Delivery model:** intended as a hosted identity service (`auth.krewire.com`) and a standalone server. Not publicly available.
- **Version:** v0.2.0 (git tag in the private repository).
- **Docs:** [Enterprise Identity & Auth Platform](/security-platforms/identity-and-auth). For independent use, the public `pkg/auth` and `pkg/sec` primitives need no server.

### Krewire Panel

- **What it is:** a self-hosted server administration panel — host telemetry, systemd services, Nginx virtual hosts, ACME TLS, databases, firewall, cron, and Docker — as one binary (`krewire-panel-server`).
- **Depends on:** `krewire` (core) and `auth`.
- **Delivery model:** intended as self-hosted software you install on your own server. Not publicly distributed.
- **Version:** `version.go` declares 0.2.0; no release tag exists yet.
- **Docs:** no dedicated documentation page yet.

---

## Adopting platform services independently and incrementally

When platform services become available, each is designed to be adopted on its own:

- You can use the core without any platform service — today this is the only option.
- Platform services are separate repositories and binaries; using one does not require moving your code into it.
- Shared identity (Auth) is a declared dependency of Workspace, Cloud, and Panel; that dependency will be documented here before any of them is offered.

---

## Pricing and availability: ownership and update process

- **Current pricing:** none published.
- **Owner:** Krewire Team (maintainers of [`krewire/krewire.com`](https://github.com/krewire/krewire.com)).
- **Update process:** any availability or pricing information will be added to this page by pull request, dated, and reviewed as part of the tagged site release process described in [Maturity, Roadmap & Support](/project/maturity-and-roadmap). Information not on this page should not be treated as current.
