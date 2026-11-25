# Getting Started with Krewire.com

This guide walks you through setting up, running, and contributing to **`krewire.com`** locally.

---

## Prerequisites

1. **`krew` CLI**: The unified developer toolchain and CLI driver.
   - Install using the official installer:
     ```bash
     curl -fsSL https://krewire.com/install.sh | bash
     ```

   - Or install via Go:
     ```bash
     go install github.com/krewire/krew/cmd/krew@latest
     ```

   - Or build locally from the `krew` repository:
     ```bash
     cd /home/krewire/workspace/krewire/krew
     go build -o ../bin/krew ./cmd/krew
     ```

   - Verify installation:
     ```bash
     krew version
     ```
2. **Python 3**: Required to execute verification and utility scripts in `scripts/`.
3. **Docker** *(optional)*: For testing the production Nginx container build.
4. **No Go Toolchain Required Locally**: As a pure `site` workload, `krewire.com` does not need a local `go.mod`.

---

## Local Development Workflow

### 1. Build Static Output

Generate the complete static website into `.krewire/build/`:

```bash
krew build
```

The build pipeline performs:

- Route generation from `pages/*.kiw` using layouts in `layouts/`.
- Markdown compilation of `content/en/docs/**/*.md` into `/docs/...` via `book`.
- Markdown compilation of `content/id/docs/**/*.md` into `/id/docs/...` via `book`.
- Scoped CSS compilation for all `.kiw` component styles.
- Automatic asset copying and bundling from `public/` into `.krewire/build/assets/`.

### 2. Live Preview Server

Start the local development preview server:

```bash
krew serve
```

By default, the server listens at `http://localhost:8080`. You can configure the port in `krewire.yaml` under `dev.port`.

---

## Authoring Pages and Components

### `.kiw` Template Syntax

Pages and layouts use Krewire's `.kiw` single-file component format:

- **YAML Frontmatter**: Metadata declared at the top between `---` delimiters.
- **Scoped Styling**: `<style>` tags are automatically scoped with unique `data-kiw-*` attributes.
- **Components**: Reusable components in `components/` can be invoked using custom tags (e.g. `<SearchModal />`, `<Navbar />`, `<Footer />`).
- **Markdown Blocks**: Embed markdown directly inside template files using `<markdown>...</markdown>`.

Example page template (`pages/example.kiw`):

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

Files are organized with numbered prefixes for hierarchical ordering (e.g., `01-overview`, `02-getting-started`, `06-ui`).

When adding or updating docs:

1. **Maintain 1:1 Parity**: Ensure every document in `content/en/docs/` has an identical counterpart in `content/id/docs/`.
2. **Preserve Frontmatter**: Maintain frontmatter metadata (`title`, `description`).
3. **Verify Parity**:
   ```bash
   python3 scripts/check-i18n-parity.py
   ```
4. **Rebuild & Verify**:
   ```bash
   krew build
   ```

---

## UI Component Previews

Krewire UI components have live interactive previews in `ui/playground` and inside documentation pages:

- Previews are statically generated using `scripts/gen-ui-previews.py`.
- If you update core UI components in `components/` or add new preview variants, re-run:
  ```bash
  python3 scripts/gen-ui-previews.py
  ```

---

## Previewing Production Container

To build and run the production Docker image locally:

```bash
# Build the production container
docker build -t krewire-site .

# Run the container locally
docker run -d -p 8080:80 --name krewire-site-test krewire-site
```

Visit `http://localhost:8080` to verify Nginx routing, clean URLs, security headers, and asset caching behavior.
