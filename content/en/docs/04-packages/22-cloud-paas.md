---
title: "Hybrid PaaS Platform"
description: "Architecture, runner, process supervisor, WebAuth console UI, Git push webhooks, and dynamic reverse proxy edge router in krewire/cloud."
date: "2026-10-31"
---

# Hybrid PaaS Platform (`github.com/krewire/cloud`)

`krewire/cloud` powers both the self-hosted open-core PaaS engine and the multi-tenant managed SaaS platform (`cloud.krewire.com`), fulfilling `ADR-C9L2P`.

---

## Subsystems & Architecture

### 1. `pkg/controlplane`
- Manages tenant accounts, projects, deployment records, and environment variable secrets.
- Enforces plan limits and quotas (`free`, `pro`, `team`, `enterprise`).
- Indonesian Personal Data Protection (UU PDP) compliant tenant data isolation and auditing.

### 2. `pkg/runner`
- Executes hermetic Go compilations (`CGO_ENABLED=0`, `-trimpath`, `-s -w`).
- Produces compact, standalone binaries running in ~15MB RAM with zero external runtime dependencies.

### 3. `pkg/supervisor`
- Dynamic process pool supervisor allocating isolated localhost ports.
- Automated health check probes before traffic cutover.
- Graceful zero-downtime rolling deploys.

### 4. `pkg/gateway`
- Dynamic reverse proxy edge router.
- Automated Let's Encrypt ACME TLS certificate provisioning for wildcard subdomains and custom domains.
- Host-based HTTP/WebSocket edge routing and connection pooling.

### 5. `pkg/api` & Web Console
- **WebAuth Console UI:** Integrated login, SSO, and TOTP 2FA web console UI (`/webauth/login`, `/webauth/sso`, `/webauth/mfa`).
- **Continuous Deployment Webhooks:** Automated Git push webhook listener (`/api/v1/webhooks/git`) triggering immediate build and deploy pipelines.
- **Log Streaming:** Real-time build and container execution log streaming delivered over Server-Sent Events (SSE) at `/api/v1/apps/{id}/logs`.

---

## Daemon & Deployment (`cmd/krewire-cloud-server`)

The single-binary daemon runs on any Linux VPS (Ubuntu, Debian, Alpine):

```bash
cd cloud
go run ./cmd/krewire-cloud-server
```

Or deploy as a container using the production multi-stage Dockerfile:

```bash
docker build -t krewire-cloud-server ./cloud
```
