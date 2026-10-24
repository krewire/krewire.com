# Krewire — krewire/krewire.com

Source for the unified Krewire site (landing + docs), **100% built with Krewire** (`site` workload, file-based `.kiw`).

**Positioning:** this site is the public face of Krewire — an end-to-end **digital SDLC ecosystem** in Go, from specification to production operations, held to three pillars: **secure**, **sustainable**, and **scalable**. Canonical pillar definitions live in `internal/docs/project-vision.md`.

- **Stack:** `krewire.yaml` (kind `site`, `base: "/"`) + `pages/*.kiw` + `layouts/*.kiw` + `content/<locale>/docs/*.md` → `kiw build` → `.krewire/build/` (no `go.mod` needed per `KWF-DF3PL`).
- **Design:** Inspired by `laravel.com` — sparse hero with code snippet, 8 workload cards, ecosystem strip, docs sidebar. Theme toggle (`auto`/`light`/`dark`) via `localStorage`, scoped CSS (`data-kiw-*`), theme tokens (`--primary` `#39D353`).
- **Version:** `v0.1.0` — single source `krewire.yaml` `project.version`; injected as `.Version` into every page (badges/footer), never hardcoded in content.
- **Spec:** tracked as GitHub Issues in `krewire/krewire.com` labelled `spec` (see [`AGENTS.md`](AGENTS.md)).

## Quick start

```bash
# from this repo (no go.mod needed)
kiw build        # → .krewire/build/
kiw serve        # preview at http://localhost:8080
```

The output has two layers:

- `pages/index.kiw` → `.krewire/build/index.html` (landing page, `Landing` layout).
- `content/<locale>/docs/*.md` → `.krewire/build/docs/**` for the default locale plus `.krewire/build/<locale>/docs/**` per locale (rendered by `pkg/book`).

## Deploy

- **Source:** `krewire/krewire.com` `main` (this repo)
- **Production:** the Docker image copies `.krewire/build` into nginx (`docker/nginx.conf`); domain configuration is managed separately.

The `main` branch is the source of truth. The production deployment is independent of `krewire/krewire.github.io`.

## Structure

```
krewire.yaml           # devtool config: kind, build output, dev server, i18n, nav
pages/                 # file-based routes (index.kiw → /)
layouts/               # ssg page layouts (Landing for the home page)
lang/<locale>/         # translation dictionaries consumed by public/assets/krewire.js
content/<locale>/docs/ # Markdown documentation collections (en, id) rendered by pkg/book
public/                # Static assets (favicon, images, client JS/CSS)
docker/                # nginx configuration for the production image
```

## Ecosystem

Krewire consists of 8 repositories (`krewire`, `krew`, `cloud`, `auth`, `workspace`, `internal`, `krewire.com`, `krewire.github.io`). This site dogfoods `pkg/web/ssg` + `pkg/kiw` DSL (`<markdown>`, `dict` helper, optional frontmatter) and validates the `site` path end-to-end.
