---
title: "Open Core & Licensing"
description: "How Krewire's open-core boundary works today: MIT-licensed public repositories, the tier classification in pkg/model, the design rules that protect the free core, and what is not yet defined."
date: "2026-11-09"
---

# Open Core & Licensing

Krewire is an open-source ecosystem with an **open-core** design: the public core is MIT-licensed, and optional platform services are developed separately. This guide describes the licensing boundary **as it exists today** and separates it from what is not yet defined.

---

## 1. What is licensed today

| Layer | Repositories | License |
| :--- | :--- | :--- |
| Open-source core | `krewire`, `krew`, `ui`, `krewire.com`, `krewire.github.io` | MIT (public) |
| Platform services | `dash`, `cloud`, `auth`, `forge` | Platform services with dedicated portals; see [Core & Platform Services](/project/core-and-platform-services) |

See [Openness & Self-Hosting](/project/openness-and-self-hosting) for links to each `LICENSE` file.

---

## 2. The tier classification (`pkg/model`)

`pkg/model` defines four ordered tiers — `free` → `pro` → `team` → `enterprise` — and `kiw tier [--json] [--check <tier>]` reports a project's tier posture and can assert that it stays within a tier (for example `kiw tier --check free`).

```text
┌──────────────────────────────┐
│  4. enterprise               │
├──────────────────────────────┤
│  3. team                     │
├──────────────────────────────┤
│  2. pro                      │
├──────────────────────────────┤
│  1. free  ← the entire core  │
└──────────────────────────────┘
```

Current state (verifiable in [`version.go`](https://github.com/krewire/krewire/blob/main/version.go)): the core monorepo declares `Tier = model.TierFree` — *"All code in krewire/krewire is 100% open source; all packages are currently free-tier by default."*

| Tier | Status |
| :--- | :--- |
| **`free`** | The whole public core. MIT-licensed. |
| **`pro`** | Defined as a classification only. No packages, license, or price are published. |
| **`team`** | Defined as a classification only. No packages, license, or price are published. |
| **`enterprise`** | Defined as a classification only. No packages, license, contract terms, or price are published. |

Pricing, commercial license models, and support terms for higher tiers are **not yet defined**. Any future definition will be published, with an owner and date, on [Core & Platform Services](/project/core-and-platform-services).

---

## 3. Design rules that protect the free core

These rules are Krewire's stated design commitments for the open-core boundary:

### Rule 1: Zero Runtime Lock-in
A project referencing only `free` packages must **build, deploy, and run with zero license keys or cloud call-homes**. There is no trial expiration or hidden limit on the free tier. The network calls that the core CLI does make are listed in [Openness & Self-Hosting](/project/openness-and-self-hosting).

### Rule 2: Additive Wrapper Boundaries
Any future paid module may wrap, harden, or optimize a `free` module, but must **never be a prerequisite for a free module**.

### Rule 3: Strict Hierarchy Subsets
Each tier is a superset of the tier below it: `pro` includes all of `free`, `team` includes all of `pro` and `free`, and `enterprise` includes all lower tiers.
