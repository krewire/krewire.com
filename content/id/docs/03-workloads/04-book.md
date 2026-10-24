---
title: "Documentation Book (book)"
description: "Compile Markdown manuscripts into responsive, book-shaped documentation portals with Krewire native book engine, automated chapter ordering, search, and themes."
date: "2026-10-08"
---

# Documentation Book (`book`)

The **`book`** workload compiles folders of Markdown manuscripts into structured, responsive, book-shaped documentation portals.

Powered by Krewire's native **`pkg/book`** engine (`github.com/krewire/krewire/pkg/book`), the book engine operates with zero external dependencies and compiles hundreds of documentation pages into static HTML in less than **15 milliseconds**.

---

## Architectural Characteristics

1. **Pure Go Standard Library:** Zero npm scripts, no Python runtime, and no Node.js dependencies.
2. **Automated Chapter Ordering:** Files and directories with numeric prefixes (`01-overview`, `02-guide`) are automatically converted into numbered chapters and subchapters with clean URL slugs.
3. **Responsive Navigation & Search:** Built-in collapsible sidebar, quick chapter search filter, keyboard shortcut navigation (`/` to search), and breadcrumbs.
4. **Instant Live Authoring:** `kiw dev` or `kiw serve` provides a local HTTP server that re-renders Markdown files dynamically on browser refresh.
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

Engine book menggunakan aturan penamaan deterministik untuk membuat navigasi yang ramah pembaca:

| Directory or File Name | Number | Slug | Display Title |
| :--- | :--- | :--- | :--- |
| `01-introduction` | 1 | `introduction` | `1. Introduction` |
| `01-welcome.md` | 1.1 | `welcome` | `1.1 Welcome to the Guide` |
| `02-getting-started` | 2 | `getting-started` | `2. Getting Started` |
| `01-installation.md` | 2.1 | `installation` | `2.1 Installation` |

> [!NOTE]
> Judul diekstraksi otomatis dari `# Heading` pertama di setiap file Markdown. Jika tidak ada heading, slug akan diformat sebagai judul.

---

## Frontmatter Support

Kompiler book mendukung blok frontmatter YAML sehingga file manuskrip dapat digunakan bersama dengan koleksi SSG:

```markdown
---
title: "Panduan Instalasi"
description: "Cara menyiapkan toolchain di Linux dan macOS"
date: "2026-10-08"
---

# Panduan Instalasi

Ikuti petunjuk di bawah untuk menginstal binary...
```

Frontmatter akan dihapus otomatis sebelum perenderan Markdown untuk menghasilkan output HTML yang bersih.

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "ecosystem-docs"
  kind: book
  version: "0.1.0"

# Direktori sumber manuskrip markdown
input: "content"

build:
  output: ".krewire/build"
  base: "/"
```

---

## Penggunaan CLI Native (`kiw`)

Build dan serve buku Anda langsung menggunakan `kiw`:

```bash
# Build book ke direktori output
kiw build

# Serve book secara lokal dengan live re-rendering
kiw serve
```
