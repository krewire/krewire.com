# Contributing to Krewire.com

Thank you for your interest in contributing to **`krewire.com`**, the official website and documentation portal for the Krewire digital SDLC ecosystem!

This document outlines the development workflow, quality gates, content authoring guidelines, and commit standards for contributors.

---

## 1. Principles & Language Standards

- **English as Single Source of Truth:**
  - All repository files, components, English documentation, and commit messages must be written in **English**.
  - Indonesian documentation lives under `content/id/docs/` and maintains strict 1:1 structural and contextual parity with `content/en/docs/`.
- **Zero-Bloat Dogfooding:**
  - The site is 100% built with Krewire (`site` workload) and compiled using `kiw build`.
  - Do not introduce Node.js/npm dependencies, external CSS frameworks, or heavy client-side JavaScript.
  - Scoped styling belongs in `.kiw` `<style>` blocks; shared tokens live in `public/assets/krewire.css`.

---

## 2. Local Setup & Tooling

### Prerequisites

1. **`kiw` CLI**:
   ```bash
   curl -fsSL https://krewire.com/scripts/install.sh | sh
   ```
   Or install from the core Go repository:
   ```bash
   go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
   ```
2. **Python 3**: Used for running verification and utility scripts in `scripts/`.
3. **Docker** *(optional)*: For testing the production container build.

### Development Commands

```bash
# Build the entire static site into .krewire/build/
kiw build

# Start the local development server (default: http://localhost:8080)
kiw serve
```

---

## 3. Documentation Authoring & 1:1 Parity

Documentation manuscripts are organized into numbered chapters under `content/`:

- **English**: `content/en/docs/`
- **Indonesian**: `content/id/docs/`

### Authoring Rules

1. **1:1 File & Path Parity:** Every new or renamed document in `content/en/docs/` must have a corresponding counterpart in `content/id/docs/` with the exact same relative file path and numbered prefix.
2. **YAML Frontmatter:** Every documentation file must include frontmatter with `title:` and `description:`.
3. **Code Blocks:** Always declare explicit language identifiers on fenced code blocks (e.g., ````go`, ````yaml`, ````bash`, ````html`).
4. **Internal Links:** Link to other docs chapters using absolute paths (e.g., `[Overview](/overview)` or `[Button](/ui/components/button)`). For links to non-book routes such as the UI playground, use `https://krewire.com/ui/playground`.
5. **No Hardcoded Versions:** Do not hardcode release versions into documentation prose unless specifically referring to historical migration guides. Dynamic pages receive `.Version` from `krewire.yaml`.

---

## 4. Reusable Components & UI Playground

- Reusable `.kiw` components reside in `components/`.
- Interactive UI component previews are rendered in `ui/playground` and embedded in individual component documentation pages.
- When adding or modifying core UI components in `components/` or documentation previews, regenerate static preview artifacts:
  ```bash
  python3 scripts/gen-ui-previews.py
  ```

---

## 5. Verification & Quality Gates

Before opening a pull request or pushing commits, ensure all checks pass:

```bash
# 1. Check 1:1 bilingual documentation parity across all files
python3 scripts/check-i18n-parity.py

# 2. Check translation keys synchronization
python3 scripts/sync-i18n.py --check

# 3. Clean static site generation
kiw build
```

`kiw build` must exit code 0 and compile all static assets and documentation into `.krewire/build/` without errors or broken routes.

---

## 6. Git Workflow & Commit Discipline

### Branching Strategy

- **`main`**: The primary branch representing the stable documentation state. All feature and documentation pull requests target `main`.
- **`hotfix`**: Dedicated branch for urgent production fixes bypassing standard release gating. Changes merged to `hotfix` trigger emergency deployment via [`.github/workflows/hotfix.yaml`](.github/workflows/hotfix.yaml).

### Conventional Commits

We strictly follow Conventional Commits formatting:

- `feat(site|ui|docs): description`
- `fix(site|ui|docs): description`
- `docs(content): description`
- `style(ui): description`
- `refactor(site): description`
- `chore(deps|scripts): description`

### Atomic Commits

Keep commits focused and atomic (recommended ≤~300 LOC per commit). If implementing a tracked specification, reference the issue number and requirement ID (e.g., `feat(site): implement KRW-LAND-001 (ref #12, req KRW-LAND-001)`).

---

## 7. Additional Resources

- [Internal Documentation Hub](docs/index.md)
- [Getting Started Guide](docs/getting-started.md)
- [Architecture Deep Dive](docs/architecture.md)
- [Contributing Details](docs/contributing.md)
- [AI Agent Rules (AGENTS.md)](AGENTS.md)
