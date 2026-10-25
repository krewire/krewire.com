---
title: "Packages & Libraries"
description: "Reference catalog for Krewire's modular domain libraries: strict downward dependencies, SSOT invariant, package catalog, and integration patterns."
date: "2026-10-24"
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
        <a href="/packages/model-and-lifecycle">Model, Lifecycle &amp; Errors (pkg/model, pkg/lifecycle, pkg/errs) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      The foundational control plane: workload kinds, lifecycle supervision, error taxonomy, environment resolution, and version contracts.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.2</span>
      <h3 class="chapter-card-title">
        <a href="/packages/web-and-app">Web &amp; Application (pkg/web, pkg/app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      HTTP routing engine, middleware pipelines, RFC 7807 problem details, dependency injection container, and server lifecycle runners.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.3</span>
      <h3 class="chapter-card-title">
        <a href="/packages/security-and-auth">Security &amp; Auth (pkg/sec, pkg/auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Hardened security defaults, OWASP Top 10 middleware, CSRF tokens, SSRF protection, PII masking, JWT verification, and basic authentication.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.4</span>
      <h3 class="chapter-card-title">
        <a href="/packages/ui-and-tui">User Interfaces (pkg/ui, pkg/tui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Reusable Single-File Components, design tokens, light/dark themes, and interactive terminal user interface (TUI) toolkits.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.5</span>
      <h3 class="chapter-card-title">
        <a href="/packages/storage-and-fs">Storage &amp; Filesystem (pkg/storage, pkg/fs, pkg/file) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Backend-independent object storage (S3, local disk, memory), atomic file replacements, and mockable filesystem boundaries.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.6</span>
      <h3 class="chapter-card-title">
        <a href="/packages/resilience-and-testing">Resilience &amp; Testing (pkg/resilience, pkg/testing) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Exponential backoff retries, circuit breakers, dead-letter dispatch, and the testing harness with HTTP and UI assertion utilities.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.7</span>
      <h3 class="chapter-card-title">
        <a href="/packages/cloud-and-distributed">Cloud &amp; Distributed (pkg/cloud) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Asynchronous worker queues, cron schedulers, microservice registry and mesh, and Infrastructure as Code (IaC) schema.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.8</span>
      <h3 class="chapter-card-title">
        <a href="/packages/config-and-validation">Config &amp; Validation (pkg/config, pkg/validation) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Hierarchical config loaders with dotenv and CLI overrides, paired with composable, zero-reflection input validation rules.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.9</span>
      <h3 class="chapter-card-title">
        <a href="/packages/ai-and-agents">AI &amp; Agent Orchestration (pkg/ai) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-provider LLM integrations, OpenAI/Anthropic/Gemini adapters, streaming completions, and prompt template engines.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.10</span>
      <h3 class="chapter-card-title">
        <a href="/packages/database-and-orm">Database &amp; Multi-Driver ORM (pkg/orm) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Fluent query builder, model entity mapping, and multi-driver database engine supporting SQLite, MySQL, PostgreSQL, MongoDB, and Elasticsearch.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.11</span>
      <h3 class="chapter-card-title">
        <a href="/packages/caching">Multi-Driver Cache (pkg/cache, pkg/redis) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      In-memory and Redis distributed caching with TTL, expiration sweeps, and connection pooling.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.12</span>
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
      <span class="chapter-card-num">4.13</span>
      <h3 class="chapter-card-title">
        <a href="/packages/commerce-and-payments">Commerce &amp; Local Payments (pkg/commerce) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Indonesia-first payment integration with Midtrans supporting dynamic QRIS, Virtual Accounts (BCA, Mandiri, BRI, BNI), and credit cards.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.14</span>
      <h3 class="chapter-card-title">
        <a href="/packages/identity-and-auth">Enterprise Identity &amp; Auth (krewire/auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Multi-tenant user identity, PBKDF2 password hashing, RFC 6238 TOTP 2FA, enterprise SSO (OIDC/SAML2), hierarchical RBAC, and responsive web auth UI.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.15</span>
      <h3 class="chapter-card-title">
        <a href="/packages/cloud-paas">Hybrid PaaS Platform (krewire/cloud) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Go-native self-hosted PaaS engine and managed cloud SaaS, featuring hermetic builders, supervised process pools, and dynamic reverse proxy edge routing.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.16</span>
      <h3 class="chapter-card-title">
        <a href="/packages/web-runtime">Web Runtime &amp; WASM (pkg/runtime) <span class="arrow">→</span></a>
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
