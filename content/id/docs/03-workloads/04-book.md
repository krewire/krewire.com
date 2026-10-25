---
title: "Buku Dokumentasi (book)"
description: "Kompilasi manuskrip Markdown menjadi portal dokumentasi berbentuk buku yang responsif dengan engine book bawaan Krewire, pengurutan bab otomatis, pencarian, dan tema."
date: "2026-10-08"
---

# Buku Dokumentasi (`book`)

Workload **`book`** mengompilasi folder-folder manuskrip Markdown menjadi portal dokumentasi yang terstruktur, responsif, dan berbentuk buku.

Ditenagai oleh engine bawaan Krewire **`pkg/book`** (`github.com/krewire/krewire/pkg/book`), engine buku beroperasi tanpa dependensi eksternal apa pun dan mengompilasi ratusan halaman dokumentasi menjadi HTML statis dalam waktu kurang dari **15 milidetik**.

---

## Karakteristik Arsitektur

1. **Pustaka Standar Go Murni:** Tanpa script npm, tanpa runtime Python, dan tanpa dependensi Node.js.
2. **Pengurutan Bab Otomatis:** File dan direktori dengan awalan numerik (`01-overview`, `02-guide`) secara otomatis diubah menjadi bab dan sub-bab bernomor dengan URL slug yang rapi.
3. **Navigasi & Pencarian Responsif:** Dilengkapi sidebar yang dapat dilipat, filter pencarian bab instan, pintasan navigasi keyboard (`/` untuk mencari), dan breadcrumb.
4. **Penulisan Langsung Instan (Live Authoring):** `kiw dev` atau `kiw serve` menyediakan server HTTP lokal yang merender ulang file Markdown secara dinamis saat browser dimuat ulang.
5. **Tema Gelap & Terang:** Pengalih tema bawaan dengan deteksi preferensi OS otomatis dan persistensi `localStorage`.

---

## Scaffolding Proyek `book`

```bash
kiw new my-book --book
cd my-book
```

Tata letak direktori:

```text
my-book/
├── krewire.yaml          # Konfigurasi proyek
└── content/              # File manuskrip dokumentasi
    ├── 01-introduction/
    │   ├── index.md      # Ikhtisar Bab 1
    │   ├── 01-welcome.md # Sub-bab 1.1
    │   └── 02-setup.md   # Sub-bab 1.2
    ├── 02-architecture/
    │   ├── index.md      # Ikhtisar Bab 2
    │   └── 01-layers.md  # Sub-bab 2.1
    └── 03-api/
        └── index.md      # Ikhtisar Bab 3
```

---

## Aturan Pengurutan Bab

Engine buku menggunakan aturan penamaan deterministik untuk menghasilkan navigasi yang ramah pengguna:

| Directory or File Name | Number | Slug | Display Title |
| :--- | :--- | :--- | :--- |
| `01-introduction` | 1 | `introduction` | `1. Introduction` |
| `01-welcome.md` | 1.1 | `welcome` | `1.1 Welcome to the Guide` |
| `02-getting-started` | 2 | `getting-started` | `2. Getting Started` |
| `01-installation.md` | 2.1 | `installation` | `2.1 Installation` |

> [!NOTE]
> Judul diekstraksi secara otomatis dari `# Heading` pertama di setiap file Markdown. Jika tidak ada heading, slug akan diformat sebagai judul.

---

## Dukungan Frontmatter

Compiler buku menerima blok frontmatter YAML, sehingga file manuskrip dapat digunakan bersama dengan koleksi SSG secara mulus:

```markdown
---
title: "Panduan Instalasi"
description: "Cara menyiapkan toolchain di Linux dan macOS"
date: "2026-10-08"
---

# Panduan Instalasi

Ikuti petunjuk di bawah untuk menginstal binary...
```

Frontmatter akan dihapus otomatis sebelum perenderan Markdown untuk memastikan keluaran HTML yang bersih.

---

## Konfigurasi (`krewire.yaml`)

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

## Penggunaan CLI Bawaan (`kiw`)

Build dan sajikan buku Anda langsung dengan `kiw`:

```bash
# Build buku ke direktori output
kiw build

# Sajikan buku secara lokal dengan live re-rendering
kiw serve
```
