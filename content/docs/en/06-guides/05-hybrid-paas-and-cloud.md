---
title: "Hybrid PaaS & Managed Cloud"
description: "Operating Krewire Cloud: self-hosted single-binary PaaS engine on private VPS vs managed cloud SaaS platform."
date: "2026-10-24"
---

# Hybrid PaaS & Managed Cloud (`cloud`)

Krewire Cloud adopts **Model C (Hybrid Open-Core PaaS & Managed Cloud SaaS)** per [ADR-C9L2P](file:///home/reasvyn/workspace/krewire/internal/docs/adr/ADR-C9L2P-hybrid-paas-platform-architecture.md).

## 1. Self-Hosted Open-Core PaaS Engine

Run your private Go-native platform on any Linux VPS (Ubuntu, Debian, Alpine) with a single command:

```bash
cd cloud
go run ./cmd/cloud-server
```

- **Extreme RAM Efficiency**: Go workloads run in 15–30MB RAM. A single $5/month VPS can host dozens of live applications.
- **Zero-Downtime Rolling Deploys**: Supervised process pool binds new instances to dynamic ports, validates health checks, and dynamically updates the reverse proxy.
- **Automatic Let's Encrypt TLS**: Automated ACME certificates for wildcard subdomains and custom domains.

## 2. Managed Multi-Tenant Cloud Platform (`cloud.krewire.com`)

For teams that prefer instantaneous serverless deployments without VPS management:

- `kiw deploy --cloud`: Directly bundles and stages projects into production.
- Real-time build streaming and centralized log aggregation.
- Local Indonesian billing via QRIS and Bank Virtual Accounts.
- Full data sovereignty compliance under Indonesia's Personal Data Protection Law (UU PDP).
