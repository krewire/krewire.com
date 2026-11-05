---
title: "Web & UI Framework"
description: "Reference catalog for Krewire's web and user interface libraries: HTTP radix routing, application container, WebAssembly runtime, .kiw UI components, and TUI."
date: "2026-10-25"
---

# Web & UI Framework

Krewire provides a complete, cohesive web and presentation layer designed for sub-millisecond latency, minimal memory footprints (~15MB RAM), and pure Go execution.

The framework supports both server-rendered and interactive experiences:
- **Radix Tree Routing:** Zero-allocation routing engine with parameter extraction and middleware pipelines.
- **Unified Monolith Architecture:** Application container with automatic dependency injection and lifecycle supervision.
- **Client-Side Go WASM:** Run pure Go in the browser with virtual DOM diffing and SSR hydration.
- **Component & Terminal Toolkits:** Reusable `.kiw` design tokens and Elm-architecture terminal applications.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.1</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/web">Web Engine & HTTP Routing (pkg/web) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Zero-allocation radix tree HTTP router, middleware pipeline, RFC 7807 problem details, parameter extraction, and render helpers.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.2</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/app">Application Container & Monolith (pkg/app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Dependency injection container, application lifecycle management, and graceful process supervision for web monoliths.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.3</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/web-runtime">Web Runtime & WebAssembly (pkg/runtime) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Client-side Go WebAssembly compilation, reactive signals, virtual DOM reconciler, and server-side rendering (SSR) hydration.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.4</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/ui">Web UI Components & Tokens (pkg/ui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Reusable Single-File Components (.kiw), semantic CSS design tokens, typography scales, and accessible component primitives.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.5</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/tui">Terminal UI Toolkit (pkg/tui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Terminal user interface framework built on Elm/Bubble Tea architecture, featuring ANSI styling, viewport scrolling, and interactive prompts.
    </p>
  </div>

</div>
