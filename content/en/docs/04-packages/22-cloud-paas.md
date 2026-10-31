---
title: "Hybrid PaaS Platform"
description: "Architecture, runner, process supervisor, and dynamic reverse proxy edge router in krewire/cloud."
date: "2026-10-24"
---

# Hybrid PaaS Platform (`github.com/krewire/cloud`)

`krewire/cloud` powers both the self-hosted open-core PaaS engine and the multi-tenant managed SaaS platform (`cloud.krewire.com`), fulfilling `ADR-C9L2P`.

## Subsystems

1. **`pkg/controlplane`**:
   - Manages tenant accounts, projects, deployment records, and environment variable secrets.
   - Enforces plan limits and quotas (`free`, `pro`, `team`, `enterprise`).
2. **`pkg/runner`**:
   - Executes hermetic Go compilations (`CGO_ENABLED=0`, `-trimpath`, `-s -w`).
   - Produces compact, standalone binaries running in ~15MB RAM.
3. **`pkg/supervisor`**:
   - Dynamic process pool supervisor allocating isolated localhost ports.
   - Performs automated health check probes before routing traffic.
   - Graceful zero-downtime rolling deploys.
4. **`pkg/gateway`**:
   - Dynamic reverse proxy edge router.
   - Resolves subdomains (`<app-slug>.krewire.dev`) and custom domains.
   - Auto-ACME TLS certificate provisioning hook.
5. **`cmd/cloud-server`**:
   - Single-binary daemon deployable on any Linux VPS with one command:
     ```bash
     go run ./cmd/cloud-server
     ```
