---
title: "Core & Foundation Libraries"
description: "Reference catalog for Krewire's core libraries: workload models, lifecycle supervision, configuration, validation, resilience, and testing harness."
date: "2026-10-25"
---

# Core & Foundation Libraries

The foundational tier of Krewire establishes the architectural guarantees, domain rules, and lifecycle supervision required by all workloads.

Packages in this layer adhere to strict design principles:
- **Strict Downward-Only Dependencies:** Domain primitives and rules reside in `pkg/model` and `pkg/errs`. Higher packages never inject upward dependencies.
- **Single Source of Truth (SSOT):** Facts, constants, and validation rules are owned by exactly one package and imported elsewhere.
- **Zero Third-Party Bloat:** Where Go standard library packages suffice, Krewire builds upon them without introducing external dependency chains.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.1</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/model">Workload Model & Domain Primitives (pkg/model, pkg/errs) <span class="arrow">→</span></a>
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
        <a href="/core-libraries/lifecycle">Lifecycle Supervision & Kernel (pkg/lifecycle) <span class="arrow">→</span></a>
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
        <a href="/core-libraries/config">Hierarchical Configuration (pkg/config) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Deterministic configuration loading with precedence ordering across CLI flags, environment variables, and layered .env files.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.4</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/validation">Composable Input Validation (pkg/validation) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Zero-reflection, composable data validation rules for API payloads, forms, and business logic schemas.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.5</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/resilience">Resilience Patterns & Circuit Breakers (pkg/resilience) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Circuit breakers, jittered exponential retry backoff, deadline budgets, and fault isolation primitives for distributed operations.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.6</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/testing">Testing Harness & Assertions (pkg/testing) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Fast assertion utilities, HTTP response recorders, ephemeral server fixtures, and golden test helpers.
    </p>
  </div>

</div>
