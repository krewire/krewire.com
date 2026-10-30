# Getting Started with Krewire.com

This guide walks you through setting up and running **`krewire.com`** locally.

---

## Prerequisites

1. **`kiw` CLI**: The unified Krewire ecosystem CLI.
   - Built from the `krewire` core monorepo:
     ```bash
     cd /home/reasvyn/workspace/krewire/krewire
     go build -o ../bin/kiw ./tools/kiw/cmd/kiw
     ```
   - Ensure `bin/kiw` is accessible in your `PATH` or invoke it directly via `../bin/kiw`.
2. **Docker** *(optional)*: For testing the production Nginx container build.
3. **No Go Toolchain Required Locally**: As a pure `site` workload, `krewire.com` does not need a local `go.mod`.

---

## Local Development Workflow

### 1. Build Static Output

Generate the static website into `.krewire/build/`:

```bash
cd /home/reasvyn/workspace/krewire/krewire.com
kiw build
```

The build pipeline performs:
- Route generation from `pages/*.kiw` using layouts in `layouts/`.
- Markdown compilation of `content/<locale>/docs/**/*.md` into `/docs/...` and `/<locale>/docs/...` via `pkg/book`.
- Translation dictionary aggregation from `lang/`.
- Asset copying and bundling from `public/`.

### 2. Live Preview Server

Start the local development preview server:

```bash
kiw serve
```

By default, the server listens at `http://localhost:8080`. You can configure the port in `krewire.yaml` under `dev.port`.

---

## Authoring Pages and Components

### `.kiw` Template Syntax

Pages and layouts use Krewire's `.kiw` format:
- **YAML Frontmatter**: Metadata declared at the top between `---` delimiters.
- **Scoped Styling**: `<style>` tags are automatically scoped with unique `data-kiw-*` attributes.
- **Components**: Reusable components in `components/` can be invoked using custom tags (e.g. `<SearchModal />` or `<TabbedCode />`).
- **Markdown Blocks**: Embed markdown directly inside template files using `<markdown>...</markdown>`.

Example page template:

```html
---
title: "Custom Feature"
layout: "Landing"
---
<section class="feature-hero">
  <h1>{{ .Title }}</h1>
  <markdown>
    Explore our modular architecture with **zero overhead**.
  </markdown>
</section>

<style>
.feature-hero h1 {
  color: var(--primary);
}
</style>
```

---

## Editing Documentation Content

The official documentation lives under `content/`:
- **English**: `content/en/docs/`
- **Indonesian**: `content/id/docs/`

Files are organized with numbered prefixes for ordering (e.g., `01-overview`, `02-getting-started`).
When adding or updating docs:
1. Mirror changes across both `en` and `id` trees when applicable.
2. Maintain frontmatter metadata (`title`, `description`, etc.).
3. Rebuild with `kiw build` to verify table-of-contents generation.

---

## Previewing Production Container

To build and run the production Docker image locally:

```bash
docker build -t krewire-site .
docker run -p 8080:80 krewire-site
```

Visit `http://localhost:8080` to verify Nginx routing, asset headers, and fallback behaviors.
