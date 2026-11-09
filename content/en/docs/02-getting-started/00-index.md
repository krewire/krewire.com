---
title: "Getting Started"
description: "The beginner track: install kiw, build your first site in five minutes, then a fullstack app with SQLite and forms. No architecture background required."
date: "2026-11-09"
---

# Getting Started

<nav class="track-crumbs" aria-label="Beginner track">
<span class="track-crumbs-label">Beginner track</span>
<ol>
<li><a href="/">Choose a path</a></li>
<li aria-current="step"><span>Getting Started</span></li>
<li><a href="/getting-started/quickstart">1 · Quickstart</a></li>
<li><a href="/getting-started/fullstack-app-tutorial">2 · Fullstack app</a></li>
<li><a href="/getting-started/dsl">3 · The .kiw DSL</a></li>
</ol>
</nav>

This chapter is the **beginner track**. It assumes you know a little Go and a terminal — nothing else. You do not need to know about architecture layers, dependency graphs, scope levels, or ADRs to build with Krewire; those live in the separate [Architecture Deep-Dive](/architecture) for platform engineers and contributors.

---

## 1. The Beginner Track

Follow these in order. Every command in them was run against the current `kiw` CLI.

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.1</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/quickstart">Quickstart (5 minutes) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Install <code>kiw</code>, scaffold a static site, edit a page, add your own component, run the dev server, and produce a production build.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.2</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/fullstack-app-tutorial">Build a Fullstack App (15 minutes) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      A task manager with <code>kiw new --app</code>, SQLite through <code>pkg/orm</code>, a <code>.kiw</code> page with Krewire UI components, validated CSRF-protected forms, and a single-binary build.
    </p>
  </div>

</div>

When you finish the track, keep the reference pages of this chapter at hand:

| # | Page | Read it when you want to… |
| :--- | :--- | :--- |
| 2.3 | [Installation](/getting-started/installation) | install on another OS, pin a version, or set up shell completion. |
| 2.4 | [Configuration](/getting-started/configuration) | understand every field of `krewire.yaml`. |
| 2.5 | [Directory Structure](/getting-started/directory-structure) | know where files go for each workload. |
| 2.6 | [Agentic Development](/getting-started/agentic-development) | pair with AI coding agents on a Krewire project. |
| 2.7 | [The `.kiw` DSL](/getting-started/dsl) | learn the full component syntax: props, slots, scoped CSS, scripts. |

---

## 2. Prerequisites

| Requirement | Version | Check with |
| :--- | :--- | :--- |
| **Go** | 1.27.1 or newer | `go version` |
| **Git** | any recent version | `git --version` |
| **OS** | Linux, macOS, or Windows (PowerShell or WSL2) | — |

Krewire needs **no Node.js, no npm, and no bundler**. If Go is missing or older than 1.27.1, install it from [go.dev/dl](https://go.dev/dl) first.

---

## 3. Which Workload Should I Start With?

Every Krewire project has one *kind*. Pick the row that matches what you are building — you can grow into the others later without moving your components.

| I want to build… | Kind | Scaffold | Learn more |
| :--- | :--- | :--- | :--- |
| A website, landing page, or blog | `site` | `kiw new my-site --site` | [Site workload](/workloads/site) |
| Documentation or a book | `book` | `kiw new my-docs --book` | [Book workload](/workloads/book) |
| A web app with a database and forms | `app` | `kiw new my-app --app` | [App workload](/workloads/app) |
| A command-line tool | `cli` | `kiw new my-tool --cli` | [CLI workload](/workloads/cli) |
| A background job processor | `worker` | `kiw new my-worker --worker` | [Worker workload](/workloads/worker) |
| An HTTP microservice | `service` | `kiw new my-svc --service` | [Service workload](/workloads/service) |
| A reverse proxy in front of services | `gateway` | `kiw new my-gw --gateway` | [Gateway workload](/workloads/gateway) |
| Cloud infrastructure definitions | `infra` | `kiw new my-infra --infra` | [Infra workload](/workloads/infra) |

Not sure? Start with `--site` (Quickstart) or `--app` (Fullstack tutorial).

---

## 4. Everyday Commands

| Command | What it does |
| :--- | :--- |
| `kiw new <name> --site` | Create a new project (`--app`, `--book`, `--cli`, … for other kinds). |
| `kiw dev` | Run the project locally with rebuild on save (`http://localhost:8080`). |
| `kiw build` | Production build: static files in `.krewire/build/` for sites and books, a binary in `.krewire/build/bin/` for apps. |
| `kiw serve --addr :3000` | Preview the static build over HTTP. |
| `kiw make component Card` | Generate a starter file (`component`, `page`, `model`, `migration`, `handler`, …). |
| `kiw test` | Run the project's Go tests. |
| `kiw doctor` | Diagnose your toolchain and workspace. |
| `kiw help <command>` | Show the options of any command. |

The complete list is in the [CLI Command Reference](/cli-reference/commands).

---

## 5. Common Beginner Pitfalls

**`kiw: command not found`** — the install directory is not on your `PATH`. Add `export PATH="$HOME/.local/bin:$PATH"` (shell installer) or `export PATH="$(go env GOPATH)/bin:$PATH"` (`go install`) to `~/.zshrc` or `~/.bashrc`, then open a new terminal.

**`address already in use`** — another program owns port 8080. Change `dev.port` in `krewire.yaml`, or preview a build on another address with `kiw serve --addr :3000`.

**`requires go >= 1.27.1`** — your Go is older than the one Krewire targets. Install a current Go from [go.dev/dl](https://go.dev/dl); distribution packages are often several versions behind.

---

## Next Steps

- Start the track: [**2.1 Quickstart →**](/getting-started/quickstart)
- Prefer to read the big picture first? See [**1. Overview →**](/overview).
