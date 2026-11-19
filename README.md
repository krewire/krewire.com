# Krewire — krewire/krewire.com

Source repository for the unified Krewire public website and official documentation portal, **100% built with Krewire** (`site` workload dogfooding the `.kiw` DSL and `pkg/book` manuscript compiler).

[![Version](https://img.shields.io/badge/version-v0.2.3-39D353.svg)](krewire.yaml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Workload: site](https://img.shields.io/badge/workload-site-brightgreen.svg)](krewire.yaml)

---

## Mission & Purpose

- **Mission:** Make it nearly free — in financial cost, operational burden, and cognitive effort — for anyone to build, own, and operate digital products that are secure, sustainable, and scalable (see [Why Krewire](content/en/docs/01-overview/01-why-krewire.md)).
- **Positioning:** This site is the public face of Krewire — an end-to-end **digital SDLC ecosystem** in Go, spanning specification to production operations across eight native workloads (`app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`).
- **Zero-Dependency Dogfooding:** Demonstrates real-world Static Site Generation (SSG), multi-language book compilation, scoped styling, and client-side modal search with zero Node.js/npm dependencies and no `go.mod` requirement in pure site mode.

---

## Key Features & Architecture

- **Unified Build Engine:** Built via `kiw build`, powered by `pkg/kiw` (template engine), `pkg/book` (documentation manuscript engine), and `pkg/web/ssg` (static site generator).
- **Single Source of Truth Versioning:** `krewire.yaml` (`project.version: "v0.2.3"`) automatically propagates as `.Version` to all templates, component footers, and asset cache-busting queries (`?v=0.2.3`).
- **Bilingual Documentation:** Full 1:1 documentation parity between English (`content/en/docs/`) and Indonesian (`content/id/docs/`), verified by automated parity tooling.
- **Krewire UI & Playground:** 66 single-file components with zero runtime dependencies. Core components feature live light/dark interactive previews at `/ui/playground` and inside documentation pages.
- **Pop-Brutalism Design System:** Scoped CSS (`data-kiw-*`), theme tokens (`--primary` `#39D353`, `--border` `2.5px solid`), accessible theme toggling (`auto`/`light`/`dark`) via `localStorage`.

---

## Quick Start

### Prerequisites

Install the unified `kiw` CLI tool:

```bash
# Via official installer script
curl -fsSL https://krewire.com/scripts/install.sh | sh

# Or install with Go toolchain
go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
```

### Local Development

No `go.mod` is required to develop and build `krewire.com`:

```bash
# Build the static site into .krewire/build/
kiw build

# Start the local development server at http://localhost:8080
kiw serve
```

The output is emitted to `.krewire/build/`:

- `pages/index.kiw` → `.krewire/build/index.html` (Landing page)
- `content/en/docs/**/*.md` → `.krewire/build/docs/**` (English documentation)
- `content/id/docs/**/*.md` → `.krewire/build/id/docs/**` (Indonesian documentation)
- `ui/playground.html` & `ui/preview/*.html` → `.krewire/build/ui/**` (UI component playground and interactive previews)

---

## Project Structure

```text
krewire.com/
├── krewire.yaml           # Devtool configuration (kind, version, i18n, book mount, nav)
├── pages/                 # File-based routes (index.kiw → /)
├── layouts/               # SSG layouts (Landing.kiw, Docs.kiw)
├── components/            # Reusable .kiw components (Navbar, Footer, SearchModal, etc.)
├── content/               # Bilingual documentation manuscripts
│   ├── en/docs/           # English documentation (134 chapters/subchapters)
│   └── id/docs/           # Indonesian documentation (134 chapters/subchapters)
├── lang/                  # i18n translation dictionaries (en.json, id.json)
├── public/                # Static assets (images, icons, fonts, CSS, client JS)
├── scripts/               # Automation & quality assurance tooling
│   ├── check-i18n-parity.py   # Validates 1:1 file parity between en and id docs
│   ├── sync-i18n.py           # Validates and syncs translation keys
│   ├── gen-ui-previews.py     # Generates static HTML previews for UI components
│   └── archive-docs-version.sh# Archives documentation snapshots
├── examples/              # Tutorial companion code
│   └── todo-app/          # "Build a Fullstack App in 15 Minutes" companion app
├── ui/                    # UI component playground and variant previews
├── docker/                # Hardened Nginx server configuration (nginx.conf)
├── Dockerfile             # Multi-stage production container image
├── docs/                  # Contributor and internal developer guides
├── CONTRIBUTING.md        # Contributor workflow and contribution standards
└── AGENTS.md              # Instructions and rules for AI developer agents
```

---

## Quality Gates & Verification

Before committing changes or cutting a release, verify all quality gates:

```bash
# 1. Verify 1:1 bilingual documentation parity across all markdown files
python3 scripts/check-i18n-parity.py

# 2. Check translation dictionary sync
python3 scripts/sync-i18n.py --check

# 3. Clean static build (must exit 0 and produce valid HTML)
kiw build
```

---

## Deployment Pathways

- **Standard Production Release**: Triggered exclusively by pushing stable SemVer Git tags (`vX.Y.Z` or `vX.Y.Z-lts*`) on `main`. Workflow: [`.github/workflows/deploy.yaml`](.github/workflows/deploy.yaml).
- **Emergency Hotfix Deployment**: Bypasses standard release gating and Git tags. Pushed directly to the dedicated `hotfix` branch. Workflow: [`.github/workflows/hotfix.yaml`](.github/workflows/hotfix.yaml). Automatically falls back to the previous stable release if container health checks fail.

---

## Developer Documentation

- [Getting Started Guide](docs/getting-started.md) — Local environment setup, previewing, and asset management.
- [Architecture Deep Dive](docs/architecture.md) — SSG compilation pipeline, `.kiw` engine, and container runtime.
- [Contributing Guidelines](docs/contributing.md) — Pull request workflow, bilingual authoring, and styling rules.
- [Agent Guidelines](AGENTS.md) — Single Source of Truth, commit discipline, and automated progress workflows.

---

## License

This repository is licensed under the [MIT License](LICENSE).
