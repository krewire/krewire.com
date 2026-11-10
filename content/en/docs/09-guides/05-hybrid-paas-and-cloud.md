---
title: "Hybrid PaaS & Managed Cloud"
description: "Operating Krewire Cloud: self-hosted single-binary PaaS engine on private VPS vs managed cloud SaaS platform."
date: "2026-10-24"
---

# Hybrid PaaS & Managed Cloud (Krewire Cloud)

Krewire Cloud is designed as both a self-hosted hybrid PaaS engine for private VPS hosting and a managed multi-tenant cloud platform.

> **Live Portal:** Krewire Cloud is available at [cloud.krewire.com](https://cloud.krewire.com). The `kiw deploy` commands in section 1 are part of the open-source core and work independently on your own infrastructure. See [Core & Platform Services](/project/core-and-platform-services).

## 1. Self-Hosted Hybrid PaaS Engine

Deploy your private Go-native workloads on any Linux VPS (Ubuntu, Debian, Alpine) using standard tooling:

```bash
# Package production binaries with embedded assets
kiw deploy --target binary

# Generate systemd supervisor units for zero-downtime service management
kiw deploy --target systemd
```

- **Extreme RAM Efficiency**: Go workloads run in 15–30MB RAM. A single $5/month VPS can host dozens of live applications.
- **Zero-Downtime Rolling Deploys**: Supervised process pool spawns new instances in isolated execution sandboxes, validates health checks, and dynamically updates the reverse proxy.
- **Automatic Let's Encrypt TLS**: Automated ACME certificates for wildcard subdomains and custom domains.

## 2. Managed Multi-Tenant Cloud Platform

The managed platform serves teams that prefer automated cloud deployments with zero VPS management:

- `kiw deploy` + push-to-deploy Git webhook: bundles and stages projects locally, then continuous deployment rolls them out to production.
- Real-time build streaming and centralized log aggregation.
- Local Indonesian billing via QRIS and Bank Virtual Accounts.
- Full data sovereignty compliance under Indonesia's Personal Data Protection Law (UU PDP).
