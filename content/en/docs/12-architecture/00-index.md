---
title: "Architecture Deep-Dive"
description: "Platform engineering track: the layered dependency pyramid, the multi-repo ecosystem DAG, the five scope levels, data flow between build time and request time, and a staged learning path for teams."
date: "2026-11-09"
---

# Architecture Deep-Dive (Platform Engineering)

<nav class="track-crumbs" aria-label="Architecture Deep-Dive track">
<span class="track-crumbs-label">Deep-Dive track</span>
<ol>
<li aria-current="step"><span>Overview</span></li>
<li><a href="/architecture/layered-pyramid">1 · Pyramid</a></li>
<li><a href="/architecture/ecosystem-dag">2 · Ecosystem DAG</a></li>
<li><a href="/architecture/scope-levels">3 · Scope levels</a></li>
<li><a href="/architecture/data-flow">4 · Data flow</a></li>
</ol>
</nav>

This chapter is for **platform engineers, architects, and contributors** to the Krewire repositories. It explains the rules that keep dozens of independent packages and several product repositories composable: which code may import which, how repositories depend on each other, how systems are scoped, and where data enters a page.

> **Tip:** **Building an app or a site?** You can skip this chapter entirely. The [beginner track](/getting-started) and the workload guides are written so that none of the concepts here are required. Come back when you design a multi-service platform, contribute to the core, or need to reason about build-time versus request-time rendering.

---

## Pages in This Track

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.1</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/layered-pyramid">Layered Pyramid &amp; Invariants <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Layers 0–5, strict downward-only imports, peer decoupling through bridges and gateways, Single Source of Truth, and the pure-Go rule.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.2</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/ecosystem-dag">Ecosystem DAG <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      How the repositories (<code>krewire</code>, <code>ui</code>, <code>krew</code>, <code>cloud</code>, <code>auth</code>, <code>dash</code>, <code>panel</code>, the sites) depend on each other at compile time and integrate at runtime.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.3</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/scope-levels">Scope Levels <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Workspace → Module → Domain → Service → Unit: the five levels codified in <code>pkg/model</code> and how specs and tests target them.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.4</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/data-flow">Data Flow: SSG vs Runtime <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      How data reaches a <code>.kiw</code> view at build time (<code>site</code>, <code>book</code>) and at request time (<code>app</code>), with Go handler recipes and client-side fetch patterns.
    </p>
  </div>

</div>

---

## Staged Learning Path for Teams

Engineering teams adopting Krewire across several products can follow this six-stage path. It covers the whole documentation, so treat it as a curriculum rather than a prerequisite.

```text
[Stage 1: Core Mental Model] ───► [Stage 2: Developer Onboarding]
             │                                   │
             ▼                                   ▼
[Stage 3: Workload Mastery]   ───► [Stage 4: Domain Packages Deep-Dive]
             │                                   │
             ▼                                   ▼
[Stage 5: Daily Dev & CLI]    ───► [Stage 6: Cloud, Operations & Scale]
```

### Stage 1: Core Mental Model & Architecture
1. [**Why Krewire?**](/overview/why-krewire) — Go resource efficiency, single binaries, no npm/bundler toolchain, and data sovereignty.
2. [**Core Philosophy**](/overview/philosophy) — The Three Pillars (*Secure, Sustainable, Scalable*) and the stdlib-first convention.
3. [**Workload Matrix**](/overview/krewire-workloads) — The eight project kinds and how they compose without rewrites.
4. [**Layered Pyramid**](/architecture/layered-pyramid), [**Ecosystem DAG**](/architecture/ecosystem-dag), and [**Scope Levels**](/architecture/scope-levels) — the rules in this chapter.

### Stage 2: Developer Onboarding
1. [**Beginner track**](/getting-started) — [Quickstart](/getting-started/quickstart) and [Fullstack app tutorial](/getting-started/fullstack-app-tutorial).
2. [**Installation**](/getting-started/installation), [**Configuration**](/getting-started/configuration), and [**Directory Structure**](/getting-started/directory-structure).
3. [**The `.kiw` DSL**](/getting-started/dsl) and [**Agentic Development**](/getting-started/agentic-development).

### Stage 3: Workload Specialization
- **Content:** [Book](/workloads/book), [Site](/workloads/site).
- **Terminal tools:** [CLI & TUIs](/workloads/cli).
- **Web:** [App](/workloads/app), [Gateway](/workloads/gateway).
- **Distributed systems:** [Worker](/workloads/worker), [Service](/workloads/service), [Infra](/workloads/infra).

### Stage 4: Domain Libraries
- **Foundations:** [Model](/core-libraries/model), [Lifecycle](/core-libraries/lifecycle), [Config](/core-libraries/config), [Validation](/core-libraries/validation), [Resilience](/core-libraries/resilience), [Testing](/core-libraries/testing).
- **Web & UI:** [Web engine](/web-ui/web), [Application container](/web-ui/app), [Web runtime & WASM](/web-ui/web-runtime), [UI](/web-ui/ui), [UI components](/ui/components), [TUI](/web-ui/tui).
- **Data:** [Database & ORM](/data-services/database-and-orm), [Storage](/data-services/storage), [Filesystem](/data-services/filesystem), [Caching](/data-services/caching), [Mail](/data-services/mail), [Commerce & payments](/data-services/commerce-and-payments).
- **Security & platforms:** [Security](/security-platforms/security), [Auth](/security-platforms/auth), [Identity](/security-platforms/identity-and-auth), [Distributed systems](/security-platforms/cloud-distributed), [Hybrid PaaS](/security-platforms/cloud-paas), [AI orchestration](/security-platforms/ai-and-agents), [Autonomous agents](/security-platforms/autonomous-agent), [Virtual workspace](/security-platforms/virtual-workspace).

### Stage 5: Daily Workflows & CLI
1. [**CLI Command Reference**](/cli-reference/commands)
2. [**Configuration Schema Reference**](/cli-reference/configuration)
3. [**Custom Scripting & Automation**](/cli-reference/scripting)

### Stage 6: Operations, Cloud & the Open-Core Boundary
1. [**Deployment & Operations**](/guides/deployment-and-operations)
2. [**Core & Platform Services**](/project/core-and-platform-services) — what is open source and the current availability of Krewire Cloud, Auth, Dash, and Forge.
3. [**Open Core & Licensing**](/guides/open-core-and-licensing)

---

## Sources

The rules on these pages summarize the specifications in the core repository, which remain authoritative:

- `krewire/docs/foundation/architecture.md` — monorepo layers and the six layering invariants.
- `krewire/docs/foundation/ecosystem-dag.md` — the ecosystem DAG and compile-time vs runtime boundaries.
- `krewire/pkg/model/scope.go` — the scope hierarchy as code.
- `krewire/layering_test.go` — the tests that enforce the layering.
