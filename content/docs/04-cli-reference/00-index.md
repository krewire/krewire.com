---
title: "CLI Reference"
description: "Comprehensive guide to the kiw CLI: architecture, global flags, environment variables, exit codes, and everyday developer workflows."
date: "2026-10-08"
---

# CLI Reference

The **`kiw`** CLI is the unified entry point for the entire Krewire ecosystem. It drives all eight project kinds (`app`, `cli`, `site`, `book`, `worker`, `service`, `infra`, `kernel`) behind one binary and one configuration file (**`krewire.yaml`**).

`kiw` dogfoods Krewire's own `packages/tui` and `packages/term` libraries, ensuring zero external dependencies, near-instant startup times (<2ms), and consistent POSIX exit code discipline.

---

## What is in this Chapter?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.1</span>
      <h3 class="chapter-card-title">
        <a href="/docs/cli-reference/commands">Command Catalog <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Detailed documentation of every <code>kiw</code> command: flags, usage syntax, supported workloads, and execution behavior.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.2</span>
      <h3 class="chapter-card-title">
        <a href="/docs/cli-reference/configuration">Configuration Schema (krewire.yaml) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Comprehensive specification of the <code>krewire.yaml</code> schema: project metadata, build targets, dev server ports, and workload-specific blocks.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.3</span>
      <h3 class="chapter-card-title">
        <a href="/docs/cli-reference/scripting">Task Automation &amp; Scripting <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Define custom pipeline tasks, test scripts, and build chains directly in <code>krewire.yaml</code> with cross-platform argument forwarding.
    </p>
  </div>

</div>

---

## Global Options & Environment Variables

Every `kiw` command respects the following global flags and environment variable overrides:

| Flag / Variable | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `--help`, `-h` | Boolean | `false` | Display command usage and flag descriptions. |
| `--verbose`, `-v` | Boolean | `false` | Enable verbose debug output and diagnostics. |
| `KIW_PORT` | Integer | `8080` | Override the development or preview server HTTP port. |
| `KIW_ENV` | String | `development` | Active runtime environment (`development`, `staging`, `production`). |
| `NO_COLOR` | Flag | `false` | Disable ANSI color sequences in accordance with the `NO_COLOR` standard. |

---

## POSIX Exit Code Contract

`kiw` adheres to strict exit code standards across all commands (`packages/kern/errs`):

| Exit Code | Constant | Meaning |
| :---: | :--- | :--- |
| **`0`** | `ExitCodeSuccess` | Command completed successfully with zero errors. |
| **`1`** | `ExitCodeFailure` | General operational or execution failure (build failed, tests failed, runtime error). |
| **`2`** | `ExitCodeUsage` | Command-line syntax error (invalid flag, missing required argument, unknown command). |
