# Krewire.com — Developer & Contributor Documentation

Welcome to the internal documentation for **`krewire.com`**, the official website and documentation portal for the Krewire digital SDLC ecosystem.

This repository is **100% built with Krewire** dogfooding the `site` workload and `.kiw` DSL. It demonstrates real-world static site generation, documentation manuscript rendering, multi-language internationalization, and scoped styling without external frontend frameworks.

---

## Repository Overview

- **Workload Kind**: `site` (declared in `krewire.yaml`).
- **Build Engine**: `kiw build` driven by `pkg/kiw`, `pkg/book`, and `pkg/web/ssg`.
- **Runtime Model**: Pure static files generated into `.krewire/build/`, served locally via `kiw serve` or in production via lightweight Nginx containers.
- **Go Module Requirement**: None (`go.mod` is not needed for `site` workloads).

---

## Directory Layout

```
krewire.com/
├── krewire.yaml           # Devtool configuration (kind, i18n, book mount, nav)
├── pages/                 # File-based routes (index.kiw → /)
├── layouts/               # SSG layouts (Landing.kiw, Docs.kiw)
├── components/            # Reusable .kiw components (SearchModal.kiw, TabbedCode.kiw)
├── content/               # Markdown docs collections
│   ├── en/docs/           # English documentation manuscripts
│   └── id/docs/           # Indonesian documentation manuscripts
├── lang/                  # i18n translation dictionaries (en.json, id.json)
├── public/                # Static assets (images, fonts, client JS, CSS)
├── docker/                # Production web server configs (nginx.conf)
├── docs/                  # Contributor and developer guides (this folder)
├── Dockerfile             # Multi-stage production container build
└── AGENTS.md              # AI agent guidelines for this repository
```

---

## Documentation Navigation

- [Getting Started](./getting-started.md) — Local setup, previewing with `kiw serve`, and building with `kiw build`.
- [Architecture](./architecture.md) — SSG compilation pipeline, `.kiw` DSL integration, and production hosting model.
- [Contributing](./contributing.md) — Contribution workflow, content writing rules, theme conventions, and quality gates.
