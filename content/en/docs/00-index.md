---
title: "Documentation & Learning Path"
description: "Welcome to Krewire — Modular Go Libraries for Every Workload. Comprehensive learning path from first principles to production cloud operations."
date: "2026-10-24"
---

# Documentation & Comprehensive Learning Path

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

## Recommended Learning Path

To master Krewire efficiently—whether you are an individual developer or an enterprise engineering team—follow this 6-stage progressive learning path:

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
Understand why Krewire adopts modular libraries instead of monolithic frameworks, the Three Pillars (*Secure, Sustainable, Scalable*), and the Single Source of Truth (SSOT) invariant.

1. [**Why Krewire?**](/docs/overview/why-krewire) — Extreme Go resource efficiency (~15MB RAM), sub-millisecond cold starts, zero npm/bundler fatigue, and data sovereignty.
2. [**Core Philosophy**](/docs/overview/philosophy) — The Three Pillars (*Secure, Sustainable, Scalable*), stdlib-first convention, and the strict downward-only dependency pyramid.
3. [**Workload Matrix**](/docs/overview/krewire-workloads) — Understanding the 8 project kinds and how they compose progressively without rewrites.

### Stage 2: Developer Onboarding & First Steps
Bootstrap your local Go toolchain, master the single CLI `kiw`, explore file-based conventions, and learn the `.kiw` component DSL.

1. [**Installation**](/docs/getting-started/installation) — Setting up Go 1.27.1+ and installing the `kiw` CLI.
2. [**Unified Configuration**](/docs/getting-started/configuration) — Managing your entire project via `krewire.yaml`.
3. [**Directory Structure**](/docs/getting-started/directory-structure) — Standard codebase layout (`cmd/`, `pkg/`, `pages/`, `.krewire/`).
4. [**The `.kiw` DSL**](/docs/getting-started/dsl) — Single-File Components with YAML frontmatter, scoped `<style>`, `<markdown>`, and JSX-like component sugar.
5. [**Agentic Development**](/docs/getting-started/agentic-development) — Pairing with AI agents using the Boost starter and `AGENTS.md` context discipline.

### Stage 3: Workload Specialization
Build real-world applications by diving into the specific workload kind tailored to your requirements:

1. **Content & Documentation:**
   - [**Documentation Books (`book`)**](/docs/workloads/book) — Ultra-fast manuscript-to-site compilation with `pkg/book`.
   - [**Static Sites (`site`)**](/docs/workloads/site) — Fast, SEO-optimized static sites built with pure Go.
2. **Terminal Tools:**
   - [**Interactive CLI & TUIs (`cli`)**](/docs/workloads/cli) — Building rich command-line interfaces with `pkg/tui`.
3. **Web & Fullstack Services:**
   - [**Fullstack & API Monoliths (`app`)**](/docs/workloads/app) — High-throughput web applications with `pkg/web` and `pkg/app`.
   - [**Client WebAssembly Runtimes (`runtime`)**](/docs/workloads/runtime) — Interactive browser UIs with Go WASM and VDOM progressive hydration.
4. **Distributed Systems & Cloud:**
   - [**Background Workers (`worker`)**](/docs/workloads/worker) — Async jobs, queues, cron schedulers, and dead-letter queues.
   - [**Microservices (`service`)**](/docs/workloads/service) — Resilient services with dynamic service discovery and tracing.
   - [**Infrastructure as Code (`infra`)**](/docs/workloads/infra) — Declarative cloud provisioning with plans and state locks.

### Stage 4: Domain Packages Deep-Dive
Leverage Krewire's modular, independent batteries to power core application concerns:

- [**Kernel & Lifecycle (`pkg/model`, `pkg/lifecycle`, `pkg/errs`)**](/docs/packages/kern) — Taxonomy, domain rules, and process supervisors.
- [**Web & HTTP Routing (`pkg/web`, `pkg/app`)**](/docs/packages/web-and-app) — HTTP multiplexers, middleware pipelines, and DI containers.
- [**Security & Identity (`pkg/sec`, `pkg/auth`)**](/docs/packages/security-and-auth) — Hardened OWASP middleware, PBKDF2 password hashing, TOTP 2FA, SSO (OIDC/SAML2), and RBAC.
- [**User Interfaces (`pkg/ui`, `pkg/tui`)**](/docs/packages/ui-and-tui) — Reusable components, design tokens, and themes.
- [**Storage & Caching (`pkg/storage`, `pkg/cache`, `pkg/redis`)**](/docs/packages/storage-and-fs) — Multi-driver object storage (memory, disk, S3) and Redis caching.
- [**Data, Mail & Commerce (`pkg/orm`, `pkg/mail`, `pkg/commerce`)**](/docs/packages) — Multi-driver ORM (SQLite, MySQL, Postgres, MongoDB, Elasticsearch), email dispatch (SMTP, Resend, Sendgrid), and Indonesian payments (Midtrans QRIS/VA).
- [**Resilience & Quality Assurance (`pkg/resilience`, `pkg/testing`)**](/docs/packages/resilience-and-testing) — Circuit breakers, retry backoff, and test harnesses.
- [**AI Orchestration (`pkg/ai`)**](/docs/packages/ai-and-agents) — Multi-provider LLM integrations and prompt engines.

### Stage 5: Daily Developer Workflows & CLI Tooling
Accelerate developer productivity with automated test suites, formatting, project introspection, and custom build tasks:

1. [**CLI Command Reference**](/docs/cli-reference/commands) — Comprehensive reference for `kiw new`, `init`, `dev`, `build`, `run`, `test`, `compat`, and `release`.
2. [**Configuration Schema Reference**](/docs/cli-reference/configuration) — Full field specification for `krewire.yaml`.
3. [**Custom Scripting & Automation**](/docs/cli-reference/scripting) — Defining pre-build tasks and automated hooks.

### Stage 6: Enterprise Operations, Cloud & Monetization
Scale workloads to production, deploy to self-hosted or managed PaaS, and leverage the Open Core business model:

1. [**Deployment & Operations Guide**](/docs/guides/deployment-and-operations) — Single-binary containerization, zero-downtime rolling deploys, and automated ACME SSL/TLS reverse proxies.
2. **Krewire Cloud (`cloud`)** — Deploying to the Go-native hybrid PaaS engine or managed SaaS platform (`kiw deploy --cloud`).
3. [**Open Core Model & Licensing**](/docs/guides/open-core-and-licensing) — The four-tier ladder (*free*, *pro*, *team*, *enterprise*) and data sovereignty under UU PDP.
