# Krewire.com — Developer & Contributor Documentation

Welcome to the internal engineering documentation for **`krewire.com`**, the official public website and documentation portal for the Krewire digital SDLC ecosystem.

This repository is **100% built with Krewire**, dogfooding the native `site` workload and `.kiw` DSL. It demonstrates real-world static site generation, documentation manuscript rendering, multi-language internationalization, and scoped styling without external frontend frameworks or Node.js dependencies.

---

## Repository Overview

- **Workload Kind:** `site` (declared in `krewire.yaml`).
- **Current Version:** `v0.2.3` (single source of truth in `krewire.yaml`).
- **Build Engine:** `kiw build` driven by `pkg/kiw`, `pkg/book`, and `pkg/web/ssg`.
- **Runtime Model:** Pure static assets emitted into `.krewire/build/`, served locally via `kiw serve` or in production via lightweight hardened Nginx containers.
- **Go Module Requirement:** None (`go.mod` is not needed for `site` workloads).
- **Internationalization:** Bilingual 1:1 documentation parity between English (`en`) and Indonesian (`id`).

---

## Directory Layout

```text
krewire.com/
├── krewire.yaml           # Devtool configuration (kind, version, i18n, book mount, nav)
├── pages/                 # File-based routes (index.kiw → /)
├── layouts/               # SSG layouts (Landing.kiw, Docs.kiw)
├── components/            # Reusable .kiw components (Navbar, Footer, SearchModal, etc.)
├── content/               # Markdown documentation collections
│   ├── en/docs/           # English documentation manuscripts (134 documents)
│   └── id/docs/           # Indonesian documentation manuscripts (134 documents)
├── lang/                  # i18n translation dictionaries (en.json, id.json)
├── public/                # Static assets (images, logos, fonts, CSS, client JS)
├── scripts/               # Quality assurance and automation scripts
│   ├── check-i18n-parity.py   # Validates 1:1 file parity between en and id docs
│   ├── sync-i18n.py           # Validates and syncs translation dictionaries
│   ├── gen-ui-previews.py     # Generates interactive UI previews for components
│   └── archive-docs-version.sh# Historical version archiving utility
├── examples/              # Companion tutorial projects
│   └── todo-app/          # Fullstack app tutorial companion code
├── ui/                    # UI component playground and variant previews
├── docker/                # Production web server configs (nginx.conf)
├── Dockerfile             # Multi-stage production container build
├── docs/                  # Contributor and developer guides (this folder)
├── CONTRIBUTING.md        # Public contribution guide
└── AGENTS.md              # AI agent guidelines and quality gates
```

---

## Documentation Navigation

- [**Getting Started**](./getting-started.md) — Local environment setup, `kiw` installation, live preview with `kiw serve`, and building with `kiw build`.
- [**Architecture**](./architecture.md) — SSG compilation pipeline, `.kiw` DSL integration, component playground, and production hosting model.
- [**Contributing**](./contributing.md) — Contribution workflow, bilingual authoring standards, theme conventions, and quality gates.
- [**Agent Guidelines**](../AGENTS.md) — Progress tracking, Git discipline, and automated quality gates.
