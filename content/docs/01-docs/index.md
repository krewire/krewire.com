---
title: "Overview"
description: "Welcome to Krewire — Modular Go Libraries for Every Workload. Learn the core architecture, philosophy, workloads, and developer workflows."
date: "2026-09-29"
---

# Overview

Welcome to the **Krewire Documentation**.

**Krewire** is an open-source, end-to-end **digital SDLC ecosystem** delivered as **Modular Libraries**: independent, single-purpose Go packages you compose behind one CLI (**`kiw`**) and one unified configuration (**`krewire.yaml`**) to build, test, and deploy **eight distinct workloads**:

1. **Fullstack Web Monoliths (`app`)**
2. **Terminal Interfaces & TUIs (`cli`)**
3. **High-Speed Static Sites (`site`)**
4. **Structured Documentation Books (`book`)**
5. **Background Workers & Job Queues (`worker`)**
6. **Distributed Microservices (`service`)**
7. **Cloud Infrastructure as Code (`infra`)**
8. **Client-Side WebAssembly Runtimes (`runtime`)**

---

## What is in this Chapter?

This **Overview** chapter introduces you to the core philosophy, technical architecture, and developer workflow of Krewire:

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.1</span>
      <h3 class="chapter-card-title">
        <a href="/docs/krewire-framework">Krewire Framework <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Explore the modular monolith architecture, the package ecosystem (<code>framework</code>, <code>libs</code>, <code>kiw</code>, <code>mdbind</code>), and the core web engine.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.2</span>
      <h3 class="chapter-card-title">
        <a href="/docs/why-krewire">Why Krewire? <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Understand how Krewire solves modern toolchain fragmentation, replaces multi-language stacks, and enables progressive system growth without rewrites.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.3</span>
      <h3 class="chapter-card-title">
        <a href="/docs/krewire-workloads">Krewire Workloads <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Deep dive into the 8 official workload kinds (<code>app</code>, <code>cli</code>, <code>site</code>, <code>book</code>, <code>worker</code>, <code>service</code>, <code>infra</code>, <code>runtime</code>).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.4</span>
      <h3 class="chapter-card-title">
        <a href="/docs/upgrade-guide">Upgrade Guide <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Learn the SemVer zero-breakage promise, migration steps for <code>krewire.yaml</code> decoupling, and how to update CLI & framework dependencies.
    </p>
  </div>

</div>

---

## Three Pillars

Everything Krewire ships is held to three pillars — **secure**, **sustainable**, and **scalable** — the filter a capability must pass before it lands:

| Pillar | In one line |
| :--- | :--- |
| **Secure** | Security is the default: a stdlib-first dependency graph, OWASP/CWE-aligned controls, and secrets referenced — never stored. |
| **Sustainable** | Near-zero cost, engineered: single static binaries, opt-in batteries, and no license fees — ever. |
| **Scalable** | Growth without rewrite: the progressive pipeline from static site to services and infra, from a $5 VPS to hyperscale. |

## Core Tenets

Krewire is engineered around four guiding tenets:

1. **Go as the Architectural Foundation**  
   Standard Go (`net/http`, `html/template`, `embed`, `context`, `slog`, `flag`) powers every layer. Single-binary deployments with embedded static assets ensure near-instant startups, negligible RAM usage, and zero runtime dependencies.
2. **Modular at Every Scope (SRP & SoC)**  
   Engineered using strict Single Responsibility and High Cohesion principles (`KWF-5ZHQV`). Start with a simple modular monolith and cleanly extract services or background workers as your traffic scales — without restructuring your business logic.
3. **Zero JavaScript Fatigue**  
   Generate lightning-fast static sites, documentation portals, and fullstack applications with `.kiw` component files without requiring Node.js, `npm`, webpack, or fragile `node_modules` dependency trees.
4. **Unified Developer Experience**  
   Manage your entire software lifecycle — creation, local development, asset compilation, testing, and deployment — through **one single CLI binary (`kiw`)** and a single devtool configuration file (**`krewire.yaml`**).

---

## Documentation Roadmap

Explore the comprehensive manual for the Krewire ecosystem:

| Chapter | Title | Focus |
| :--- | :--- | :--- |
| **1.** | [**Overview**](/docs) | Ecosystem architecture, core tenets, three pillars, and workloads overview. |
| **2.** | [**Getting Started**](/getting-started) | Prerequisites, 5-minute quickstart, project anatomy, agentic SDLC, and `.kiw` DSL. |
| **3.** | [**Workloads in Depth**](/workloads) | Detailed architecture and guides for all 8 workloads (`app`, `cli`, `site`, `book`, `worker`, `service`, `infra`, `runtime`). |
| **4.** | [**CLI Reference**](/cli-reference) | Complete `kiw` command catalog, `krewire.yaml` schema specification, and task automation. |
| **5.** | [**Packages & Libraries**](/packages) | Domain libraries catalog: `kern`, `web`, `app`, `sec`, `auth`, `ui`, `tui`, `storage`, `resilience`, and `testing`. |
| **6.** | [**Architecture & Guides**](/guides) | Spec-driven engineering methodology, open core licensing rules, and production deployment playbooks. |

---

## Next Steps

To begin building with Krewire, explore the subchapters in order:

- Proceed to [**1.1 Krewire Framework →**](/docs/krewire-framework)
- Jump to [**2. Getting Started →**](/getting-started)
- Or dive into [**3. Workloads in Depth →**](/workloads)
