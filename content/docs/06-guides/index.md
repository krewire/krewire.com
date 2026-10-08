---
title: "Architecture & Guides"
description: "Engineering guides, architectural discipline, spec-driven development, open-core licensing model, and production deployment operations."
date: "2026-10-08"
---

# Architecture & Guides

This chapter collects deep-dive engineering guides, architectural invariants, and operational practices that govern software development across the Krewire ecosystem.

Whether you are contributing to core packages or authoring mission-critical applications on top of Krewire, these guides ensure your codebase remains maintainable, secure, and sustainable over long horizons.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">6.1</span>
      <h3 class="chapter-card-title">
        <a href="/docs/guides/spec-driven-development">Spec-Driven Development <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Learn the spec-first engineering methodology: requirement numbering (KWL/KWN/KWM), traceability citations in Go tests, and automated quality gates.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">6.2</span>
      <h3 class="chapter-card-title">
        <a href="/docs/guides/open-core-and-licensing">Open Core &amp; Licensing <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Understand the four-rung open core ladder (free → pro → team → enterprise), strict module decoupling, and zero runtime lock-in guarantees.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">6.3</span>
      <h3 class="chapter-card-title">
        <a href="/docs/guides/deployment-and-operations">Deployment &amp; Operations <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Production deployment playbooks: multi-stage Docker builds, Nginx reverse proxy configuration, systemd daemon management, and zero-downtime rollouts.
    </p>
  </div>

</div>
