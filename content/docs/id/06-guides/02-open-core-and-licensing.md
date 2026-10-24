---
title: "Open Core & Licensing"
description: "The four-rung open core ladder: free, pro, team, and enterprise. Architecture rules, commercial boundaries, and zero runtime lock-in."
date: "2026-10-08"
---

# Open Core & Licensing

Krewire is an open-source community based in Indonesia, operating under a transparent **Open Core** business model.

Our licensing framework guarantees complete freedom for builders while offering advanced operational, compliance, and enterprise features through commercial tiers.

---

## 1. The Four-Rung Tier Ladder

The ecosystem organizes modules into four distinct rungs:

```text
┌──────────────────────────────────────────────┐
│  4. Enterprise (B2B, air-gapped, custom SLA)  │
├──────────────────────────────────────────────┤
│  3. Team (Multi-tenant governance, audits)   │
├──────────────────────────────────────────────┤
│  2. Pro (Advanced telemetry, high throughput)│
├──────────────────────────────────────────────┤
│  1. Free (Core libraries, CLI, full SDLC)    │
└──────────────────────────────────────────────┘
```

| Tier | License Model | Pricing | Target Audience |
| :--- | :--- | :--- | :--- |
| **`free`** | Permissive Open Source (MIT) | $0 Forever | Individual developers, startups, open-source projects. |
| **`pro`** | Commercial Source-Available | Published Self-Serve | Growing teams needing hardened drivers and metrics. |
| **`team`** | Commercial License | Published Organization | Mid-market companies needing SSO, RBAC, and multi-tenant policies. |
| **`enterprise`** | Negotiated B2B Contract | Custom Quote | Air-gapped deployments, custom compliance, and dedicated 24/7 SLAs. |

---

## 2. Inviolable Architectural Rules

As documented in `internal/docs/open-core.md`, three rules protect the integrity of the open-source core:

### Rule 1: Zero Runtime Lock-in
A project referencing only `free` packages must **build, deploy, and run with zero license keys or cloud call-homes**. There is never a trial expiration or hidden limit on the free tier.

### Rule 2: Additive Wrapper Boundaries
Paid modules may wrap, harden, or optimize a `free` module, but a paid module must **never be a prerequisite for a free module**.

### Rule 3: Strict Hierarchy Subsets
Each rung is a strict superset of the rung below it:
- `pro` includes all of `free`.
- `team` includes all of `pro` and `free`.
- `enterprise` is negotiated per organization and is never a published fixed price.
