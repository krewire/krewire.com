---
title: "Hybrid PaaS Platform"
description: "Architecture, runner, process supervisor, WebAuth console UI, Git push webhooks, and dynamic reverse proxy edge router in Krewire Cloud."
date: "2026-10-31"
---

# Hybrid PaaS Platform (Krewire Cloud)

Krewire Cloud is designed as both a hybrid PaaS engine and a multi-tenant managed platform.

> **Live Portal:** Krewire Cloud is available at [cloud.krewire.com](https://cloud.krewire.com) featuring hybrid PaaS orchestration, hermetic builds, process supervision, and dynamic edge routing. See [Core & Platform Services](/project/core-and-platform-services) for architectural details.

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
- Dynamic process pool supervisor allocating isolated execution sandboxes.
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

## Platform Access & Deployment

Krewire Cloud is accessible via [cloud.krewire.com](https://cloud.krewire.com). In addition to cloud hosting, you can deploy any Krewire project to your own infrastructure today using the open-source `kiw` CLI:

### Deploying Your Application

You can deploy any Krewire project directly using the `kiw` CLI:

```bash
# Stage production binary and deployable artifacts
kiw deploy --target binary

# Generate production systemd service definition
kiw deploy --target systemd

# For static and documentation sites (GitHub Pages)
kiw deploy --target gh-pages
```

### Production Docker Container

To package any Krewire application into an ultra-lightweight (<15MB) standalone container for deployment to any VPS, Docker host, or cloud Kubernetes cluster:

```dockerfile
# Multi-stage build for Krewire workloads
FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/server .

FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]
```

### Self-Hosted and Managed Editions

No self-hosted or managed edition of Krewire Cloud is offered publicly, and no enterprise subscription, pricing, or licensing terms are published. Any future availability will be announced on [Core & Platform Services](/project/core-and-platform-services).
