---
title: "Packages & Libraries"
description: "Reference catalog for Krewire's modular domain libraries: strict downward dependencies, SSOT invariant, package catalog, and integration patterns."
date: "2026-10-08"
---

# Packages & Libraries

Krewire is delivered as **Modular Libraries**: independent, single-purpose Go packages that you compose, rather than a monolithic framework you inherit.

Every package in the `krewire/packages` catalog adheres to strict architectural boundaries:
- **Strict Downward-Only Dependencies:** Domain primitives and rules reside in `packages/kern/model`. Higher-level libraries never inject upward dependencies.
- **Single Source of Truth (SSOT):** Facts, constants, and validation rules are owned by exactly one package and imported elsewhere.
- **Standard Library Alignment:** Zero unnecessary third-party dependencies. Where the Go standard library provides a capability (`net/http`, `context`, `slog`, `embed`), Krewire builds upon it rather than replacing it.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.1</span>
      <h3 class="chapter-card-title">
        <a href="/docs/packages/kern">Kernel &amp; Model (kern) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      The foundational control plane: workload kinds, lifecycle supervision, error taxonomy, environment detection, and version contracts.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.2</span>
      <h3 class="chapter-card-title">
        <a href="/docs/packages/web-and-app">Web &amp; Application (web, app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      HTTP routing engine, middleware pipelines, RFC 7807 problem details, dependency injection container, and server lifecycle runners.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.3</span>
      <h3 class="chapter-card-title">
        <a href="/docs/packages/security-and-auth">Security &amp; Auth (sec, auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Hardened security defaults, OWASP Top 10 middleware, CSRF tokens, SSRF protection, PII masking, JWT verification, and basic authentication.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.4</span>
      <h3 class="chapter-card-title">
        <a href="/docs/packages/ui-and-tui">User Interfaces (ui, tui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Reusable Single-File Components, design tokens, light/dark themes, and interactive terminal user interface (TUI) toolkits.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.5</span>
      <h3 class="chapter-card-title">
        <a href="/docs/packages/storage-and-fs">Storage &amp; Filesystem (storage, fs, file) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Backend-independent object storage (S3, local, memory), atomic file replacements, and mockable filesystem boundaries.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.6</span>
      <h3 class="chapter-card-title">
        <a href="/docs/packages/resilience-and-testing">Resilience &amp; Testing (resilience, testing) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Exponential backoff retries, circuit breakers, dead-letter dispatch, and the spec-driven test harness with traceability citations.
    </p>
  </div>

</div>

---

## Ecosystem Dependency Hierarchy

As documented in `internal/docs/architecture.md`, dependencies flow strictly downward:

```text
templates/boost (AI agent starters)
       │
       ▼
tools/kiw (Developer CLI)
       │
       ▼
packages/* (Domain libraries: web, sec, auth, ui, tui, storage...)
       │
       ▼
packages/kern (Kernel model, lifecycle, errors, environment, version)
```

No library or package may ever import `kiw`. Domain rules belong in `packages/kern/model`, and runtime mechanics belong in `packages/kern/lifecycle`.
