---
title: "Documentation — Comprehensive Learning Path"
description: "Follow a staged learning path from your first Krewire site to production workloads and platform architecture."
date: "2026-11-12"
---

# Krewire Documentation

> **Mission:** Krewire's mission is to make it nearly free — in financial cost, operational burden, and cognitive effort — for anyone to build, own, and operate digital products that are secure, sustainable, and scalable. Read more in [Why Krewire](/overview/why-krewire).

**Krewire** is a set of independent Go packages driven by one CLI (**`kiw`**) and one config file (**`krewire.yaml`**). You use it to build websites, documentation, web applications, CLIs, and — when you need them — workers, microservices, gateways, and infrastructure.

You do **not** need to read about architecture to use it. Follow the stages below in order, or jump directly to the path that matches your goal.

---

## Learning Path at a Glance

| Stage | Goal | Read |
| :--- | :--- | :--- |
| **0. Orient** | Understand what Krewire is and whether it fits. | [1. Overview](/overview) |
| **1. First success** | Install `kiw`, scaffold a project, and build something runnable. | [2. Getting Started](/getting-started) |
| **2. Build products** | Choose the right workload and learn its lifecycle and commands. | [3. Workloads](/workloads) |
| **3. Learn the foundations** | Understand the domain model, lifecycle, config, validation, and resilience. | [4. Core & Foundation Libraries](/core-libraries) |
| **4. Build interfaces** | Add HTTP routing, application containers, WASM, and a design system. | [5. Web & UI Framework](/web-ui) · [6. Krewire UI](/ui) |
| **5. Add data and services** | Persist data, cache, store files, send mail, and integrate payments. | [7. Data, Storage & Services](/data-services) |
| **6. Secure and scale** | Apply security controls, identity, cloud, and AI platform capabilities. | [8. Security & Ecosystem Platforms](/security-platforms) |
| **7. Operate and reference** | Look up commands, configuration, and production playbooks. | [9. CLI Reference](/cli-reference) · [10. Guides & Best Practices](/guides) |
| **8. Go deep** | Learn the architectural invariants that keep the ecosystem composable. | [11. Architecture Deep-Dive](/architecture) |

---

## New Here? Start With the Beginner Track

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/quickstart">Quickstart — 5 minutes <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Install <code>kiw</code>, scaffold a site, edit a page, add a component, run the dev server, and build. Zero prior knowledge of Krewire required.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/fullstack-app-tutorial">Fullstack app — 15 minutes <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      A task manager with SQLite, <code>.kiw</code> pages, Krewire UI components, validated CSRF-protected forms, and a single binary at the end.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3</span>
      <h3 class="chapter-card-title">
        <a href="/architecture">Architecture Deep-Dive <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      For platform engineers and contributors: the layered pyramid, the multi-repo DAG, scope levels, and how data flows between build time and request time. Optional for everyone else.
    </p>
  </div>

</div>

---

## Choose a Path by Goal

### "I want to build a static website or documentation" — `site`, `book`

1. [Quickstart](/getting-started/quickstart) — your first site in five minutes.
2. [The `.kiw` DSL](/getting-started/dsl) — pages, layouts, components, scoped CSS.
3. [Site workload](/workloads/site) or [Book workload](/workloads/book) — content collections, Markdown chapters, i18n.
4. [Krewire UI components](/ui/components) — 66 ready-made components.
5. [Deployment & Operations](/guides/deployment-and-operations) — ship the static build.

### "I want to build a web application with a backend API and database" — `app`

1. [Quickstart](/getting-started/quickstart) — learn the `kiw` loop on a small site.
2. [Build a Fullstack App in 15 Minutes](/getting-started/fullstack-app-tutorial) — SQLite, forms, CSRF, single binary.
3. [App workload](/workloads/app) and [Application container (`pkg/app`)](/web-ui/app).
4. [Web engine & routing (`pkg/web`)](/web-ui/web) and [Database & ORM (`pkg/orm`)](/data-services/database-and-orm).
5. [Security controls](/security-platforms/security) and [Authentication](/security-platforms/auth).
6. [Data flow: build time vs request time](/architecture/data-flow) — when to render statically and when per request.

### "I want to build an enterprise platform with microservices" — `service`, `gateway`, `worker`

1. [Why Krewire?](/overview/why-krewire) and the [workload matrix](/overview/krewire-workloads).
2. [Service](/workloads/service), [Gateway](/workloads/gateway), and [Worker](/workloads/worker) workloads.
3. [Distributed systems (`pkg/cloud`)](/security-platforms/cloud-distributed) and [Resilience](/core-libraries/resilience).
4. [Architecture Deep-Dive](/architecture) — pyramid, ecosystem DAG, scope levels.
5. [Deployment & Operations](/guides/deployment-and-operations) and [Hybrid PaaS & Cloud](/guides/hybrid-paas-and-cloud).

---

## Recommended Reading Order by Role

| Role | Read in this order |
| :--- | :--- |
| **Frontend developer** | [Quickstart](/getting-started/quickstart) → [`.kiw` DSL](/getting-started/dsl) → [UI components](/ui/components) → [UI design system (`pkg/ui`)](/web-ui/ui) → [Web runtime & WASM](/web-ui/web-runtime) → [Site workload](/workloads/site) |
| **Backend developer** | [Quickstart](/getting-started/quickstart) → [Fullstack app tutorial](/getting-started/fullstack-app-tutorial) → [`pkg/web`](/web-ui/web) → [`pkg/app`](/web-ui/app) → [Database & ORM](/data-services/database-and-orm) → [Validation](/core-libraries/validation) → [Data flow](/architecture/data-flow) |
| **DevOps engineer** | [Installation](/getting-started/installation) → [Configuration](/getting-started/configuration) → [CLI commands](/cli-reference/commands) → [Scripting & CI/CD](/cli-reference/scripting) → [Deployment & Operations](/guides/deployment-and-operations) → [Hybrid PaaS & Cloud](/guides/hybrid-paas-and-cloud) |
| **Platform engineer** | [Overview](/overview) → [Workload matrix](/overview/krewire-workloads) → [Architecture Deep-Dive](/architecture) → [Model & primitives](/core-libraries/model) → [Lifecycle kernel](/core-libraries/lifecycle) → [Open core & licensing](/guides/open-core-and-licensing) |

---

## All Chapters

| Chapter | What you will find |
| :--- | :--- |
| [1. Overview](/overview) | Why Krewire exists, its philosophy, and the eight workload kinds. |
| [2. Getting Started](/getting-started) | **Beginner track**: quickstart, fullstack tutorial, installation, configuration, `.kiw` DSL. |
| [3. Workloads](/workloads) | One guide per workload: `app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`. |
| [4. Core & Foundation Libraries](/core-libraries) | Model, lifecycle, config, validation, resilience, testing. |
| [5. Web & UI Framework](/web-ui) | Router, application container, WASM runtime, and TUI. |
| [6. Krewire UI](/ui) | The design system and all 66 components, with live previews. |
| [7. Data, Storage & Services](/data-services) | ORM, object storage, filesystem, cache, mail, payments. |
| [8. Security & Ecosystem Platforms](/security-platforms) | Security controls, auth, identity, cloud, AI agents. |
| [9. CLI Reference](/cli-reference) | Every `kiw` command and `krewire.yaml` field. |
| [10. Guides & Best Practices](/guides) | Deployment, agentic development, hybrid cloud, open-core licensing. |
| [11. Architecture Deep-Dive](/architecture) | Platform engineering: pyramid, ecosystem DAG, scope levels, data flow, and the staged learning path for teams. |
