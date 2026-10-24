---
title: "Documentation Book (book)"
description: "Compile Markdown manuscripts into responsive, book-shaped documentation portals with mdbind, automated chapter ordering, search, and themes."
date: "2026-10-08"
---

# Documentation Book (`book`)

The **`book`** workload compiles folders of Markdown manuscripts into structured, responsive, book-shaped documentation portals.

Powered by **`mdbind`** (`github.com/krewire/mdbind`), the book engine operates with zero external dependencies and compiles hundreds of documentation pages into static HTML in less than **15 milliseconds**.

---

## Architectural Characteristics

1. **Pure Go Standard Library:** Zero npm scripts, no Python runtime, and no Node.js dependencies.
2. **Automated Chapter Ordering:** Files and directories with numeric prefixes (`01-overview`, `02-guide`) are automatically converted into numbered chapters and subchapters with clean URL slugs.
3. **Responsive Navigation & Search:** Built-in collapsible sidebar, quick chapter search filter, keyboard shortcut navigation (`/` to search), and breadcrumbs.
4. **Instant Live Authoring:** `mdbind serve` provides a local HTTP server that re-renders Markdown files dynamically on browser refresh.
5. **Dark & Light Theming:** Built-in theme switcher with automatic OS preference detection and `localStorage` persistence.

---

## Scaffolding a `book` Project

```bash
kiw new my-book --book
cd my-book
```

Directory layout:

```text
my-book/
├── krewire.yaml          # Project configuration
└── content/              # Documentation manuscript files
    ├── 01-introduction/
    │   ├── index.md      # Chapter 1 overview
    │   ├── 01-welcome.md # Subchapter 1.1
    │   └── 02-setup.md   # Subchapter 1.2
    ├── 02-architecture/
    │   ├── index.md      # Chapter 2 overview
    │   └── 01-layers.md  # Subchapter 2.1
    └── 03-api/
        └── index.md      # Chapter 3 overview
```

---

## Chapter Ordering Rules

`mdbind` uses deterministic naming rules to generate human-friendly navigation:

| Directory or File Name | Number | Slug | Display Title |
| :--- | :--- | :--- | :--- |
| `01-introduction` | 1 | `introduction` | `1. Introduction` |
| `01-welcome.md` | 1.1 | `welcome` | `1.1 Welcome to the Guide` |
| `02-getting-started` | 2 | `getting-started` | `2. Getting Started` |
| `01-installation.md` | 2.1 | `installation` | `2.1 Installation` |

> [!NOTE]
> Titles are extracted automatically from the first `# Heading` in each Markdown file. If no heading is present, the slug is formatted as the title.

---

## Frontmatter Support

`mdbind` tolerates YAML frontmatter blocks, allowing manuscript files to be shared seamlessly with SSG collections:

```markdown
---
title: "Installation Guide"
description: "How to set up the toolchain on Linux and macOS"
date: "2026-10-08"
---

# Installation Guide

Follow the instructions below to install the binaries...
```

The frontmatter is automatically stripped before Markdown rendering, ensuring clean HTML output.

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "ecosystem-docs"
  kind: book
  version: "0.1.0"

# Source directory containing markdown manuscripts
input: "content"

build:
  output: ".krewire/build"
  base: "/"
```

---

## Standalone CLI Usage (`mdbind`)

`mdbind` can also be run standalone without `kiw`:

```bash
# Build book to dist/
mdbind build --input content --output dist

# Serve book locally with live re-rendering
mdbind serve --input content --port 8080
```
