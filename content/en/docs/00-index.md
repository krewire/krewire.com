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
8. **API Gateways (`gateway`)**

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

1. [**Why Krewire?**](/overview/why-krewire) — Extreme Go resource efficiency (~15MB RAM), sub-millisecond cold starts, zero npm/bundler fatigue, and data sovereignty.
2. [**Core Philosophy**](/overview/philosophy) — The Three Pillars (*Secure, Sustainable, Scalable*), stdlib-first convention, and the strict downward-only dependency pyramid.
3. [**Workload Matrix**](/overview/krewire-workloads) — Understanding the 8 project kinds and how they compose progressively without rewrites.

### Stage 2: Developer Onboarding & First Steps
Bootstrap your local Go toolchain, master the single CLI `kiw`, explore file-based conventions, and learn the `.kiw` component DSL.

1. [**Installation**](/getting-started/installation) — Setting up Go 1.27.1+ and installing the `kiw` CLI.
2. [**Unified Configuration**](/getting-started/configuration) — Managing your entire project via `krewire.yaml`.
3. [**Directory Structure**](/getting-started/directory-structure) — Standard codebase layout (`cmd/`, `pkg/`, `pages/`, `.krewire/`).
4. [**The `.kiw` DSL**](/getting-started/dsl) — Single-File Components with YAML frontmatter, scoped `<style>`, `<markdown>`, and JSX-like component sugar.
5. [**Agentic Development**](/getting-started/agentic-development) — Pairing with AI agents using the Boost starter and `AGENTS.md` context discipline.

### Stage 3: Workload Specialization
Build real-world applications by diving into the specific workload kind tailored to your requirements:

1. **Content & Documentation:**
   - [**Documentation Books (`book`)**](/workloads/book) — Ultra-fast manuscript-to-site compilation with `pkg/book`.
   - [**Static Sites (`site`)**](/workloads/site) — Fast, SEO-optimized static sites built with pure Go.
2. **Terminal Tools:**
   - [**Interactive CLI & TUIs (`cli`)**](/workloads/cli) — Building rich command-line interfaces with `pkg/tui`.
3. **Web & Fullstack Services:**
   - [**Fullstack & API Monoliths (`app`)**](/workloads/app) — High-throughput web applications with `pkg/web` and `pkg/app`.
   - [**API Gateways (`gateway`)**](/workloads/gateway) — Declarative reverse proxies with per-route middleware, rate limiting, and service discovery.
4. **Distributed Systems & Cloud:**
   - [**Background Workers (`worker`)**](/workloads/worker) — Async jobs, queues, cron schedulers, and dead-letter queues.
   - [**Microservices (`service`)**](/workloads/service) — Resilient services with dynamic service discovery and tracing.
   - [**Infrastructure as Code (`infra`)**](/workloads/infra) — Declarative cloud provisioning with plans and state locks.

### Stage 4: Domain Packages Deep-Dive
Leverage Krewire's modular, independent batteries to power core application concerns:

- [**Workload Model & Domain Primitives (`pkg/model`, `pkg/errs`)**](/packages/model) — Workload kinds, project metadata, error taxonomy, and version contracts.
- [**Lifecycle Supervision & Kernel (`pkg/lifecycle`)**](/packages/lifecycle) — Module registry, dependency ordering, and graceful shutdown.
- [**Web Engine & HTTP Routing (`pkg/web`)**](/packages/web) — High-performance radix tree router, middleware pipeline, and render helpers.
- [**Application Container & Monolith (`pkg/app`)**](/packages/app) — Dependency injection container, application lifecycle management, and signal draining.
- [**Security Controls & Hardening (`pkg/sec`)**](/packages/security) — Hardened OWASP defense headers, CSRF tokens, SSRF protection, and PII redaction.
- [**Authentication Primitives (`pkg/auth`)**](/packages/auth) — Cryptographic JWT tokens, claims verification, and constant-time HTTP Basic Auth.
- [**Web UI Components & Tokens (`pkg/ui`)**](/packages/ui) — Reusable Single-File Components (.kiw), semantic CSS design tokens, and themes.
- [**Terminal UI Toolkit (`pkg/tui`)**](/packages/tui) — Terminal user interfaces, Elm/Bubble Tea architecture, and ANSI styling.
- [**Multi-Driver Object Storage (`pkg/storage`)**](/packages/storage) — Multi-driver object storage (memory, disk, S3 / MinIO).
- [**Virtual Filesystem & Atomic Files (`pkg/fs`, `pkg/file`)**](/packages/filesystem) — Testable virtual filesystem boundaries and crash-safe atomic file writes.
- [**Resilience Patterns & Circuit Breakers (`pkg/resilience`)**](/packages/resilience) — Circuit breakers, jittered retry backoff, and failure isolation.
- [**Testing Utilities & Assertions (`pkg/testing`)**](/packages/testing) — Test assertion harness, HTTP helpers, and mock fixtures.
- [**Hierarchical Configuration (`pkg/config`)**](/packages/config) — Deterministic configuration loading from CLI flags, env vars, and .env files.
- [**Composable Input Validation (`pkg/validation`)**](/packages/validation) — Zero-reflection composable data validation for forms and APIs.
- [**Multi-Driver Cache (`pkg/cache`, `pkg/redis`)**](/packages/caching) — In-memory and Redis distributed caching with TTL and invalidation sweeps.
- [**Transactional Mailer (`pkg/mail`)**](/packages/mail) — Multi-driver transactional email dispatch (SMTP, Resend, SendGrid, Postmark).
- [**Database & Multi-Driver ORM (`pkg/orm`)**](/packages/database-and-orm) — Multi-driver ORM (SQLite, MySQL, Postgres, MongoDB, Elasticsearch).
- [**Commerce & Indonesian Payments (`pkg/commerce`)**](/packages/commerce-and-payments) — Midtrans QRIS, Virtual Accounts (BCA, Mandiri, BRI, BNI), and credit cards.
- [**Enterprise Identity & Auth Platform (`krewire/auth`)**](/packages/identity-and-auth) — Multi-tenant identity, PBKDF2, TOTP 2FA, SSO (OIDC/SAML2), and RBAC.
- [**Hybrid PaaS Platform (`krewire/cloud`)**](/packages/cloud-paas) — Go-native self-hosted PaaS engine and managed cloud control plane.
- [**Web Runtime & WASM (`pkg/runtime`)**](/packages/web-runtime) — Client-side Go WebAssembly, virtual DOM diffing, and SSR hydration.
- [**AI Orchestration (`pkg/ai`)**](/packages/ai-and-agents) — Multi-provider LLM integrations and prompt engines.

### Stage 5: Daily Developer Workflows & CLI Tooling
Accelerate developer productivity with automated test suites, formatting, project introspection, and custom build tasks:

1. [**CLI Command Reference**](/cli-reference/commands) — Comprehensive reference for `kiw new`, `init`, `dev`, `build`, `run`, `test`, `compat`, and `release`.
2. [**Configuration Schema Reference**](/cli-reference/configuration) — Full field specification for `krewire.yaml`.
3. [**Custom Scripting & Automation**](/cli-reference/scripting) — Defining pre-build tasks and automated hooks.

### Stage 6: Enterprise Operations, Cloud & Monetization
Scale workloads to production, deploy to self-hosted or managed PaaS, and leverage the Open Core business model:

1. [**Deployment & Operations Guide**](/guides/deployment-and-operations) — Single-binary containerization, zero-downtime rolling deploys, and automated ACME SSL/TLS reverse proxies.
2. **Krewire Cloud (`cloud`)** — Deploying to the Go-native hybrid PaaS engine or managed SaaS platform (`kiw deploy --target infra` with `pkg/cloud/infra`).
3. [**Open Core Model & Licensing**](/guides/open-core-and-licensing) — The four-tier ladder (*free*, *pro*, *team*, *enterprise*) and data sovereignty under UU PDP.
