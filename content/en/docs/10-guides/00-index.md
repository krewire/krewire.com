---
title: "Guides & Best Practices"
description: "Engineering guides, open-core licensing, agentic development, hybrid cloud, and production deployment operations."
date: "2026-11-12"
---

# Guides & Best Practices

This chapter collects practical engineering guides, operational practices, and integration recipes for building and running Krewire workloads in production. For the architectural invariants themselves, continue to [Architecture Deep-Dive](/architecture).

Whether you are contributing to core packages or authoring mission-critical applications on top of Krewire, these guides ensure your codebase remains maintainable, secure, and sustainable over long horizons.

> **Note:** The architectural invariants (layered pyramid, ecosystem DAG, scope levels) moved to the dedicated [Architecture Deep-Dive](/architecture) chapter, so that this chapter stays focused on practical guides.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">10.1</span>
      <h3 class="chapter-card-title">
        <a href="/guides/agentic-development">Agentic Development <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pair with AI coding agents using native Model Context Protocol (MCP), Unified AI SDK structured outputs, and Boost presets.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">10.2</span>
      <h3 class="chapter-card-title">
        <a href="/guides/deployment-and-operations">Deployment &amp; Operations <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Production deployment playbooks: multi-stage Docker builds, Nginx reverse proxy configuration, systemd daemon management, and zero-downtime rollouts.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">10.3</span>
      <h3 class="chapter-card-title">
        <a href="/guides/hybrid-paas-and-cloud">Hybrid PaaS &amp; Managed Cloud <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Running Krewire Cloud on private VPS infrastructure vs deploying to managed multi-tenant cloud platforms.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">10.4</span>
      <h3 class="chapter-card-title">
        <a href="/guides/open-core-and-licensing">Open Core &amp; Licensing <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Understand the four-rung open core ladder (free → pro → team → enterprise), strict module decoupling, and zero runtime lock-in guarantees.
    </p>
  </div>

</div>
