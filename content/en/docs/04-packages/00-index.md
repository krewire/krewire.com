---
title: "Packages & Libraries"
description: "Reference catalog for Krewire's modular domain libraries: strict downward dependencies, SSOT invariant, package catalog, and integration patterns."
date: "2026-10-25"
---

# Packages & Libraries

Krewire is delivered as **Modular Libraries**: independent, single-purpose Go packages that you compose, rather than a monolithic framework you inherit.

Every package in the `krewire/pkg` catalog adheres to strict architectural boundaries:
- **Strict Downward-Only Dependencies:** Domain primitives and rules reside in `pkg/model` and `pkg/errs`. Higher-level libraries never inject upward dependencies.
- **Single Source of Truth (SSOT):** Facts, constants, and validation rules are owned by exactly one package and imported elsewhere.
- **Standard Library Alignment:** Zero unnecessary third-party dependencies. Where the Go standard library provides a capability (`net/http`, `context`, `log/slog`, `embed`), Krewire builds upon it rather than replacing it.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.1</span>
      <h3 class="chapter-card-title">
        <a href="/packages/model">Workload Model & Domain Primitives (pkg/model, pkg/errs) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      The foundational domain model: eight workload kinds, project metadata, scope hierarchy, error taxonomy, environment detection, and version contracts.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.2</span>
      <h3 class="chapter-card-title">
        <a href="/packages/lifecycle">Lifecycle Supervision & Kernel (pkg/lifecycle) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Imperative control plane of Krewire: kernel supervisor, module registry, dependency ordering, and graceful process shutdown.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.3</span>
      <h3 class="chapter-card-title">
        <a href="/packages/web">Web Engine & HTTP Routing (pkg/web) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Zero-allocation radix tree HTTP router, middleware pipeline, RFC 7807 problem details, parameter extraction, and render helpers.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.4</span>
      <h3 class="chapter-card-title">
        <a href="/packages/app">Application Container & Monolith (pkg/app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Dependency injection container, application lifecycle management, and graceful process supervision for web monoliths.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.5</span>
      <h3 class="chapter-card-title">
        <a href="/packages/security">Security Controls & Hardening (pkg/sec) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Hardened security defaults, OWASP Top 10 middleware, CSRF tokens, SSRF protection, and PII masking.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.6</span>
      <h3 class="chapter-card-title">
        <a href="/packages/auth">Authentication Primitives (pkg/auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Core cryptographic authentication primitives, HMAC-SHA256 JWT tokens, and constant-time HTTP Basic Auth.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.7</span>
      <h3 class="chapter-card-title">
        <a href="/packages/ui">Web UI Components & Tokens (pkg/ui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Reusable Single-File Components (.kiw), semantic CSS design tokens, and light/dark theme switching engine.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.8</span>
      <h3 class="chapter-card-title">
        <a href="/packages/tui">Terminal UI Toolkit (pkg/tui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Interactive terminal user interfaces, Elm/Bubble Tea architecture, ANSI color styling, tables, and progress bars.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.9</span>
      <h3 class="chapter-card-title">
        <a href="/packages/storage">Multi-Driver Object Storage (pkg/storage) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Uniform object storage abstraction with swappable drivers for memory, local disk, and Amazon S3 / MinIO.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.10</span>
      <h3 class="chapter-card-title">
        <a href="/packages/filesystem">Virtual Filesystem & Atomic Files (pkg/fs, pkg/file) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Testable virtual filesystem boundaries, atomic crash-safe file writes, and path safety.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.11</span>
      <h3 class="chapter-card-title">
        <a href="/packages/resilience">Resilience Patterns & Circuit Breakers (pkg/resilience) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Fault-tolerant network communication with jittered retry policies, timeout propagation, and circuit breakers.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.12</span>
      <h3 class="chapter-card-title">
        <a href="/packages/testing">Testing Utilities & Assertions (pkg/testing) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Testing assertion harness, HTTP test helpers, UI component assertions, and mock fixtures.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.13</span>
      <h3 class="chapter-card-title">
        <a href="/packages/cloud-distributed">Cloud & Distributed Systems (pkg/cloud) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Asynchronous worker queues, cron schedulers, microservice registry and mesh, and Infrastructure as Code schema.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.14</span>
      <h3 class="chapter-card-title">
        <a href="/packages/config">Hierarchical Configuration (pkg/config) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Deterministic configuration loading from CLI flags, environment variables, .env files, and structs.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.15</span>
      <h3 class="chapter-card-title">
        <a href="/packages/validation">Composable Input Validation (pkg/validation) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Composable data validation without reflection overhead, providing structured errors for APIs and forms.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.16</span>
      <h3 class="chapter-card-title">
        <a href="/packages/ai-and-agents">AI & Agent Orchestration (pkg/ai) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-provider LLM integrations, OpenAI/Anthropic/Gemini adapters, streaming completions, and prompt templates.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.17</span>
      <h3 class="chapter-card-title">
        <a href="/packages/database-and-orm">Database & Multi-Driver ORM (pkg/orm) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Fluent query builder, model entity mapping, and multi-driver database engine supporting SQLite, MySQL, PostgreSQL, MongoDB, and Elasticsearch.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.18</span>
      <h3 class="chapter-card-title">
        <a href="/packages/caching">Multi-Driver Cache (pkg/cache, pkg/redis) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      High-performance in-memory and Redis distributed caching with TTL, expiration sweeps, and connection pooling.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.19</span>
      <h3 class="chapter-card-title">
        <a href="/packages/mail">Transactional Mailer (pkg/mail) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-driver transactional email dispatch supporting SMTP, Resend, SendGrid, Postmark, and in-memory testing.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.20</span>
      <h3 class="chapter-card-title">
        <a href="/packages/commerce-and-payments">Commerce & Local Payments (pkg/commerce) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Indonesia-first payment integration with Midtrans supporting dynamic QRIS, Virtual Accounts (BCA, Mandiri, BRI, BNI), and credit cards.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.21</span>
      <h3 class="chapter-card-title">
        <a href="/packages/identity-and-auth">Enterprise Identity & Auth (krewire/auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-tenant user identity, PBKDF2 password hashing, RFC 6238 TOTP 2FA, enterprise SSO (OIDC/SAML2), and hierarchical RBAC.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.22</span>
      <h3 class="chapter-card-title">
        <a href="/packages/cloud-paas">Hybrid PaaS Platform (krewire/cloud) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Go-native self-hosted PaaS engine and managed cloud SaaS, featuring hermetic builders, supervised process pools, and reverse proxy edge routing.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.23</span>
      <h3 class="chapter-card-title">
        <a href="/packages/web-runtime">Web Runtime & WASM (pkg/runtime) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Client-side Go WebAssembly: virtual DOM diffing, generic hooks, the component registry, SSR hydration markers, and a reproducible WASM build pipeline.
    </p>
  </div>

</div>

---

## Ecosystem Dependency Hierarchy

As documented in `internal/docs/foundation/architecture.md`, dependencies flow strictly downward:

```text
templates/boost (AI agent starters)
       │
       ▼
tools/kiw (Developer CLI)
       │
       ▼
pkg/* (Domain libraries: web, sec, auth, ui, tui, storage, orm, cache, mail...)
       │
       ▼
pkg/model, pkg/lifecycle, pkg/errs (domain model, supervision, errors, environment, version)
```

No library or package may ever import `kiw`. Domain rules belong in `pkg/model`, and runtime mechanics belong in `pkg/lifecycle`.
