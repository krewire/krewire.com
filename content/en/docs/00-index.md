---
title: "Documentation & Learning Path"
description: "Welcome to Krewire — Modular Go Libraries for Every Workload. Comprehensive learning path from first principles to production cloud operations."
date: "2026-10-24"
---

# Documentation & Comprehensive Learning Path

Welcome to the **Krewire Documentation**.

> **Mission:** Krewire's mission is to make it nearly free — in financial cost, operational burden, and cognitive effort — for anyone to build, own, and operate digital products that are secure, sustainable, and scalable. Read more in [About Krewire](/project).

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

### Stage 4: Modular Domain Libraries Deep-Dive
Leverage Krewire's modular, independent batteries grouped into four domain tiers:

1. **Core & Foundations:**
   - [**Workload Model & Domain Primitives (`pkg/model`, `pkg/errs`)**](/core-libraries/model) — Workload kinds, project metadata, error taxonomy, and version contracts.
   - [**Lifecycle Supervision & Kernel (`pkg/lifecycle`)**](/core-libraries/lifecycle) — Module registry, dependency ordering, and graceful shutdown.
   - [**Hierarchical Configuration (`pkg/config`)**](/core-libraries/config) — Deterministic configuration loading from CLI flags, env vars, and .env files.
   - [**Composable Input Validation (`pkg/validation`)**](/core-libraries/validation) — Zero-reflection composable data validation for forms and APIs.
   - [**Resilience Patterns & Circuit Breakers (`pkg/resilience`)**](/core-libraries/resilience) — Circuit breakers, jittered retry backoff, and failure isolation.
   - [**Testing Utilities & Assertions (`pkg/testing`)**](/core-libraries/testing) — Test assertion harness, HTTP helpers, and mock fixtures.

2. **Web & UI Framework:**
   - [**Web Engine & HTTP Routing (`pkg/web`)**](/web-ui/web) — High-performance radix tree router, middleware pipeline, and render helpers.
   - [**Application Container & Monolith (`pkg/app`)**](/web-ui/app) — Dependency injection container, application lifecycle management, and signal draining.
   - [**Web Runtime & WASM (`pkg/runtime`)**](/web-ui/web-runtime) — Client-side Go WebAssembly, virtual DOM diffing, and SSR hydration.
   - [**Web UI Components & Tokens (`pkg/ui`)**](/web-ui/ui) — Reusable Single-File Components (.kiw), semantic CSS design tokens, and themes.
   - [**Krewire UI Components Catalog (66 Components)**](/ui/components) — Complete component-by-component documentation, props, slots, and interactive examples.
   - [**Terminal UI Toolkit (`pkg/tui`)**](/web-ui/tui) — Terminal user interfaces, Elm/Bubble Tea architecture, and ANSI styling.

3. **Data, Storage & Services:**
   - [**Database & Multi-Driver ORM (`pkg/orm`)**](/data-services/database-and-orm) — Multi-driver ORM (SQLite, MySQL, Postgres, MongoDB, Elasticsearch).
   - [**Multi-Driver Object Storage (`pkg/storage`)**](/data-services/storage) — Multi-driver object storage (memory, disk, S3 / MinIO).
   - [**Virtual Filesystem & Atomic Files (`pkg/fs`, `pkg/file`)**](/data-services/filesystem) — Testable virtual filesystem boundaries and crash-safe atomic file writes.
   - [**Multi-Driver Cache (`pkg/cache`, `pkg/redis`)**](/data-services/caching) — In-memory and Redis distributed caching with TTL and invalidation sweeps.
   - [**Transactional Mailer (`pkg/mail`)**](/data-services/mail) — Multi-driver transactional email dispatch (SMTP, Resend, SendGrid, Postmark).
   - [**Commerce & Indonesian Payments (`pkg/commerce`)**](/data-services/commerce-and-payments) — Midtrans QRIS, Virtual Accounts (BCA, Mandiri, BRI, BNI), and credit cards.

4. **Security & Ecosystem Platforms:**
   - [**Security Controls & Hardening (`pkg/sec`)**](/security-platforms/security) — Hardened OWASP defense headers, CSRF tokens, SSRF protection, and PII redaction.
   - [**Authentication Primitives (`pkg/auth`)**](/security-platforms/auth) — Cryptographic JWT tokens, claims verification, and constant-time HTTP Basic Auth.
   - [**Enterprise Identity & Auth Platform (Krewire Auth)**](/security-platforms/identity-and-auth) — Multi-tenant identity, PBKDF2, TOTP 2FA, SSO (OIDC/SAML2), and RBAC.
   - [**Cloud & Distributed Systems (`pkg/cloud`)**](/security-platforms/cloud-distributed) — Cluster discovery, distributed locks, health checks, and telemetry.
   - [**Hybrid PaaS Platform (Krewire Cloud)**](/security-platforms/cloud-paas) — Go-native self-hosted PaaS engine and managed cloud control plane.
   - [**AI Orchestration & LLMs (`pkg/ai`)**](/security-platforms/ai-and-agents) — Multi-provider LLM integrations and prompt engines.
   - [**Autonomous AI Agent Engine (`krew`)**](/security-platforms/autonomous-agent) — Standalone and embedded autonomous agent runtime.
   - [**AI-Native Virtual Workspace (Krewire Dash)**](/security-platforms/virtual-workspace) — Multi-agent virtual workspace and autonomous integration engine.

### Stage 5: Daily Developer Workflows & CLI Tooling
Accelerate developer productivity with automated test suites, formatting, project introspection, and custom build tasks:

1. [**CLI Command Reference**](/cli-reference/commands) — Comprehensive reference for `kiw new`, `init`, `dev`, `build`, `run`, `test`, `compat`, and `release`.
2. [**Configuration Schema Reference**](/cli-reference/configuration) — Full field specification for `krewire.yaml`.
3. [**Custom Scripting & Automation**](/cli-reference/scripting) — Defining pre-build tasks and automated hooks.

### Stage 6: Operations, Cloud & the Open-Core Boundary
Scale workloads to production on infrastructure you control, and understand the open-core boundary:

1. [**Deployment & Operations Guide**](/guides/deployment-and-operations) — Single-binary containerization, zero-downtime rolling deploys, and automated ACME SSL/TLS reverse proxies.
2. [**Core & Platform Services**](/project/core-and-platform-services) — What is open source, what is an optional platform service, and the current availability of Krewire Cloud, Auth, Workspace, and Panel (none is publicly available yet).
3. [**Open Core Model & Licensing**](/guides/open-core-and-licensing) — MIT-licensed core, the tier classification in `pkg/model`, and the design rules that keep the free core free.
