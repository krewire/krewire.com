# Architecture of Krewire.com

This document describes the architectural layout, compilation pipeline, subsystem design, and deployment strategy of the **`krewire.com`** website.

---

## Architectural Role

`krewire.com` serves as the official public website and documentation portal for the Krewire digital SDLC ecosystem.
It is built as an end-to-end reference implementation for the **`site`** workload kind, validating that high-performance websites and documentation can be produced with zero Node.js/npm dependencies and standard Go toolchain primitives.

```text
                      +-----------------------------+
                      |   Source Files & Content    |
                      |  - pages/*.kiw              |
                      |  - layouts/*.kiw            |
                      |  - components/*.kiw         |
                      |  - content/{en,id}/docs/*.md|
                      |  - public/ & lang/          |
                      +--------------+--------------+
                                     |
                                     v
                      +-----------------------------+
                      |          krew build          |
                      |  - pkg/kiw (template parser)|
                      |  - pkg/book (manuscripts)   |
                      |  - pkg/web/ssg (generator)  |
                      |  - auto_assets (bundling)   |
                      +--------------+--------------+
                                     |
                                     v
                      +-----------------------------+
                      |     .krewire/build/         |
                      |  Static HTML, CSS, JS, Img  |
                      +--------------+--------------+
                                     |
                      +--------------+--------------+
                      |                             |
                      v                             v
           [Local Preview Server]        [Production Docker Image]
             krew serve (:8080)              Nginx Alpine (:80)
```

---

## Key Subsystems

### 1. The `.kiw` Template Engine (`pkg/kiw`)
- Parses YAML frontmatter, scoped `<style>` blocks, inline `<markdown>`, and custom tags.
- Injects standard page context into layouts and templates:
  - `.Title`, `.Description`: Page metadata.
  - `.Version`: Propagated from `krewire.yaml` (SSOT: `v0.2.3`).
  - `.Nav`: Navigation links configured in `krewire.yaml`.

### 2. Documentation Book Subsystem (`pkg/book`)
- Configured via `book:` block in `krewire.yaml` (`mount: "/docs"`, `toc: true`).
- Parses markdown manuscripts from `content/en/docs/` and `content/id/docs/`.
- Automatically extracts headings to generate hierarchical tables of contents with independent subchapter toggle states.
- Generates HTML pages rendered through `layouts/Docs.kiw`.
- Mounts default English documentation under `/docs/` and localized Indonesian documentation under `/id/docs/`.

### 3. Asset Injection & Bundling
- Controlled by `auto_assets.bundle: true` in `krewire.yaml`.
- All CSS and JS files in `public/assets/` are automatically bundled and injected into `<head>`:
  - Scoped CSS rules prevent style leaks across components (`data-kiw-*`).
  - Cache-busting queries (`?v=<version>`) are automatically appended based on `project.version`.

### 4. Client Search Subsystem
- Documentation topics are indexed during build into search JSON manifests.
- Handled at runtime by `components/SearchModal.kiw` and client scripts in `public/assets/krewire.js`.
- Fast, client-side modal search with keyboard shortcuts (`Cmd/Ctrl + K`) supporting both English and Indonesian documents.

### 5. UI Component Previews & Playground
- Contains 66 single-file `.kiw` components in the Krewire UI design system.
- Interactive live previews for core components (`Button`, `Badge`, `Input`, `Card`, `Modal`, `Tabs`, `Alert`) are statically generated via `scripts/gen-ui-previews.py`.
- Previews feature live light/dark theme toggles and one-click copyable `.kiw` component and usage snippets.
- Accessible at `/ui/playground` and embedded directly in component documentation pages.

### 6. Internationalization (i18n) & Parity Pipeline
- Configured in `krewire.yaml` (`locales: ["en", "id"]`, `default_locale: "en"`).
- **Separation of Concerns**:
  - `content/{locale}/docs/`: Long-form markdown manuscripts compiled by `pkg/book`.
  - `lang/{locale}/site.json`: UI and shell localized strings for the client translation engine.
- Strict 1:1 bilingual documentation parity is enforced:
  - `python3 scripts/check-i18n-parity.py` validates that all 134 manuscripts in English have identical Indonesian counterparts.
  - `python3 scripts/sync-i18n.py --check` ensures all translation keys match across dictionaries.

---

## Directory Ownership & Workload Boundaries

The official portal balances static site generation with production runtime daemons:

1. **Framework-Reserved Discovery Roots**:
   - `pages/`: Evaluated by `pkg/web/ssg` for URL routing.
   - `layouts/`: Evaluated by `pkg/kiw` for master layout wrappers.
   - `krewire.yaml`: Workload specification.
2. **Conventional Content Roots**:
   - `components/`: Project components. Subdirectories are recursively scanned and registered both by basename and namespaced identifier.
   - `content/`: Markdown manuscripts parsed exclusively by the book engine.
   - `lang/`: JSON translation dictionaries.
   - `public/`: Copied verbatim to build output root (`.krewire/build/`).
3. **Platform Extensions**:
   - `cmd/krewire-mcp-server/`: Go daemon serving Model Context Protocol endpoints.
   - `internal/mcp/`: Private packages for the MCP server.
   - `examples/todo-app/`: Isolated example with independent `go.mod`.
   - `.krewire/build/`: Transient compilation artifact output; excluded from source control.

---

## Dual Production Deployment Pathways

The repository employs two distinct deployment pathways:

1. **Standard Production Release (SemVer)**:
   - Triggered exclusively by pushing stable SemVer Git tags (`vX.Y.Z` or `vX.Y.Z-lts*`) on `main`.
   - Workflow: [`.github/workflows/deploy.yaml`](../.github/workflows/deploy.yaml).
   - Generates production static build, verifies health checks, and deploys container.

2. **Emergency Hotfix Deployment (Fast-Track Bypass)**:
   - Bypasses tag requirements and standard multi-stage approval.
   - Pushed directly to the dedicated `hotfix` branch (`git checkout hotfix && git merge main --ff-only && git push origin hotfix`).
   - Workflow: [`.github/workflows/hotfix.yaml`](../.github/workflows/hotfix.yaml).
   - **Automated Rollback Guarantee**: If container health checks or live smoke tests fail, the workflow immediately rolls back to the last known stable production image (`krewire-app:stable`).

---

## Nginx Production Runtime

The container build defined in `Dockerfile` and `docker/nginx.conf` features:

- Multi-stage build based on pinned `nginx:alpine` image.
- Removal of default Nginx configuration to prevent configuration leaks.
- Static Gzip compression for text, HTML, CSS, JavaScript, and JSON.
- Clean URL rewrites (e.g. `/docs` → `/docs/index.html`).
- Cache-Control headers for version-hashed assets.
- Hardened security headers: `X-Frame-Options`, `X-Content-Type-Options`, `Referrer-Policy`, and `Content-Security-Policy`.
- Container health check:
  ```dockerfile
  HEALTHCHECK --interval=5s --timeout=3s --start-period=2s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://127.0.0.1/ || exit 1
  ```
