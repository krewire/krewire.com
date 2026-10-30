# Architecture of Krewire.com

This document describes the architectural layout, compilation pipeline, and deployment strategy of the **`krewire.com`** website.

---

## Architectural Role

`krewire.com` serves as the public face and official documentation portal of the Krewire ecosystem.
It is built as an end-to-end reference implementation for the **`site`** workload kind, validating that static sites and docs can be produced without external Node.js/JavaScript toolchains.

```
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
                      |          kiw build          |
                      |  - pkg/kiw (template parser)|
                      |  - pkg/book (manuscripts)   |
                      |  - pkg/web/ssg (generator)  |
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
             kiw serve (:8080)              Nginx Alpine (:80)
```

---

## Key Subsystems

### 1. The `.kiw` Template Engine
- Evaluated by `pkg/kiw` in the core repository.
- Processes YAML frontmatter, scoped `<style>` blocks, inline `<markdown>`, and embedded components.
- Injects standard page context:
  - `.Title`, `.Description`: Page metadata.
  - `.Version`: Propagated from `krewire.yaml` (SSOT).
  - `.Nav`: Navigation links configured in `krewire.yaml`.

### 2. Documentation Book Subsystem (`pkg/book`)
- Configured via `book:` block in `krewire.yaml` (`mount: "/docs"`, `toc: true`).
- Parses markdown files in `content/en/docs/` and `content/id/docs/`.
- Extracts headings to build hierarchical tables of contents.
- Wraps output using `layouts/Docs.kiw`.

### 3. Client Search Subsystem
- Documentation topics are indexed during build into search JSON manifests.
- Handled at runtime by `components/SearchModal.kiw` and client scripts in `public/assets/`.
- Fast, client-side modal search with keyboard shortcuts (`Cmd/Ctrl + K`).

### 4. Internationalization (i18n)
- Declared under `i18n:` in `krewire.yaml` (`locales: ["en", "id"]`).
- Default English routes mount at `/` and `/docs/...`.
- Indonesian routes mount under `/id/` and `/id/docs/...`.
- Translation dictionaries live in `lang/en/` and `lang/id/`.

---

## Production Deployment Model

1. **Static Build**: `kiw build` generates optimized static files in `.krewire/build/`.
2. **Container Image**: A multi-stage Docker build copies `.krewire/build/` into a hardened `nginx:alpine` image.
3. **Nginx Runtime**:
   - Gzip compression for static text/HTML/CSS/JS.
   - Cache-Control headers for hashed assets.
   - Clean URLs (rewriting `/docs` to `/docs/index.html`).
   - Standard security headers (`X-Frame-Options`, `X-Content-Type-Options`, `Referrer-Policy`).
