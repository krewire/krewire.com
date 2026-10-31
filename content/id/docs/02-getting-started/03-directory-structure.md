---
title: "Struktur Direktori"
description: "Anatomi dan konvensi tata letak direktori proyek Krewire di semua varian workload."
date: "2026-09-29"
---

# Struktur Direktori

Krewire menetapkan struktur proyek standar yang terprediksi, dirancang untuk berkembang secara mulus dari prototipe satu berkas menjadi monolit modular tingkat enterprise.

---

## 1. Anatomi Proyek Standar

Berikut adalah tata letak berkas standar terpadu dari proyek Krewire yang komprehensif:

```text
my-project/
├── krewire.yaml          # Konfigurasi devtool & pipeline build
├── go.mod                # Definisi modul Go
├── go.sum                # Checksum dependensi
│
├── pages/                # Rute berbasis berkas (berkas komponen .kiw)
│   ├── index.kiw         # Menyajikan "/"
│   ├── about.kiw         # Menyajikan "/about"
│   └── 404.html          # Halaman error kustom
│
├── layouts/              # Pembungkus halaman bersama (.kiw)
│   ├── Base.kiw          # Layout dokumen HTML default
│   └── Landing.kiw       # Layout halaman landing full-bleed
│
├── components/           # Komponen UI ber-scope yang dapat digunakan kembali (.kiw)
│   ├── Header.kiw        # Bilah navigasi
│   ├── Footer.kiw        # Bagian footer
│   └── CodeBlock.kiw     # Komponen syntax highlighter
│
├── content/              # Manuskrip & buku dokumentasi (pkg/book)
│   ├── 01-overview/      # Direktori Bab 1
│   │   ├── index.md      # Halaman utama bab
│   │   └── 01-arch.md    # Sub-bab 1.1
│   └── 02-guides/        # Direktori Bab 2
│
├── public/               # Aset statis mentah yang disajikan langsung
│   ├── favicon.svg       # Favicon browser
│   ├── robots.txt        # Konfigurasi web crawler
│   ├── sitemap.xml       # Indeks mesin pencari
│   └── assets/           # Skrip klien, stylesheet, dan gambar
│
├── cmd/                  # Titik masuk executable Go
│   └── my-service/
│       └── main.go       # Titik masuk service
│
├── internal/             # Logika aplikasi privat (ditegakkan oleh compiler Go)
│   ├── domain/           # Entitas bisnis inti & model
│   ├── handler/          # Handler permintaan HTTP & RPC
│   └── service/          # Domain service & aturan bisnis
│
└── .krewire/             # Status toolchain & artefak build (diabaikan oleh git)
    ├── build/            # Output kompilasi default
    └── .kiw-build-manifest
```

---

## 2. Peran & Tanggung Jawab Direktori

### `pages/` (Perutean Berbasis Berkas / File-Based Routing)
Berkas di dalam `pages/` dipetakan langsung ke rute URL publik:
- `pages/index.kiw` → `https://example.com/`
- `pages/pricing.kiw` → `https://example.com/pricing`
- `pages/blog/first-post.kiw` → `https://example.com/blog/first-post`

Setiap berkas halaman `.kiw` merangkum frontmatter YAML, markup template HTML, dan blok `<style>` yang memiliki scope lokal.

---

### `layouts/` (Cangkang Dokumen)
Layout menyediakan cangkang luar HTML (`<!doctype html>`, `<head>`, `<nav>`, `<footer>`) yang membungkus konten halaman melalui `{{.Content}}`. Halaman mendeklarasikan layout yang diinginkan di frontmatter:

```yaml
---
title: "Dashboard"
layout: AppLayout
---
```

---

### `components/` (Blok UI Ber-Scope yang Dapat Digunakan Kembali)
Komponen adalah berkas `.kiw` mandiri yang disusun menggunakan sintaksis mirip JSX `<ComponentName />` atau helper template `{{component "Name" .}}` yang kompatibel. Style di dalam komponen secara otomatis diisolasi (scoped) saat waktu kompilasi tanpa ada tabrakan CSS.

---

### `content/` (Engine Manuskrip)
Menyimpan manuskrip dokumentasi Markdown yang dikompilasi oleh `pkg/book`. Awalan numerik (`01-`, `02-`) menentukan urutan membaca bab dan sub-bab pada bilah navigasi samping (sidebar) yang dihasilkan.

---

### `public/` (Aset Statis Tanpa Kompilasi)
Berkas di dalam `public/` disajikan apa adanya tanpa pra-pemrosesan:
- `public/favicon.svg` disajikan di `/favicon.svg`.
- `public/robots.txt` disajikan di `/robots.txt`.
- `public/assets/` berisi skrip sisi klien dan stylesheet CSS. Berkas yang berakhiran `.css` atau `.js` ditautkan ke setiap halaman secara otomatis — tidak memerlukan tag `<link>` atau `<script>` manual di dalam layout.

---

### `cmd/` & `internal/` (Inti Aplikasi Go)
- **`cmd/<name>/main.go`:** Titik masuk utama untuk monolit fullstack (`app`), utilitas baris perintah (`cli`), atau worker latar belakang (`worker`).
- **`internal/`:** Paket privat yang ditegakkan langsung oleh compiler Go. Kode di dalam `internal/` tidak dapat diimpor oleh modul eksternal, menjaga enkapsulasi arsitektur yang bersih (`KWF-5ZHQV`).

---

### `.krewire/` (Artefak Build)
Dibuat secara otomatis oleh `kiw dev` dan `kiw build`:
- **`.krewire/build/`:** Menyimpan output produksi hasil kompilasi yang siap untuk hosting statis atau pemaketan citra Docker.
- Direktori ini dikelola sepenuhnya oleh `kiw` dan harus selalu ditambahkan ke dalam `.gitignore` Anda.

---

## 3. Tata Letak Khusus untuk Masing-Masing Workload

Saat membuat proyek baru melalui `kiw new <project> [flags]`, `kiw` akan menyiapkan struktur direktori yang disesuaikan untuk workload tersebut:

### 1. Situs Statis (`kiw new my-site --site`)
```text
my-site/
├── krewire.yaml
├── go.mod
├── pages/
│   └── index.kiw
├── layouts/
│   └── Base.kiw
├── components/
└── public/
```

### 2. Buku Dokumentasi (`kiw new my-docs --book`)
```text
my-docs/
├── krewire.yaml
├── go.mod
└── content/
    └── docs/
        ├── 01-introduction.md
        └── 02-getting-started.md
```

### 3. Monolit Fullstack (`kiw new my-app --app`)
```text
my-app/
├── krewire.yaml
├── go.mod
├── cmd/
│   └── my-app/
│       └── main.go
├── web/
│   ├── router.go
│   └── handlers/
├── internal/
│   └── domain/
└── public/
```

### 4. CLI Terminal (`kiw new my-cli --cli`)
```text
my-cli/
├── krewire.yaml
├── go.mod
├── cmd/
│   └── my-cli/
│       └── main.go
└── internal/
    └── commands/
```

---

## 4. Menyesuaikan Lokasi Direktori

Jika Anda mengintegrasikan Krewire ke dalam repositori lama yang memiliki konvensi folder non-standar, Anda dapat menimpanya di dalam `krewire.yaml`:

```yaml
project:
  name: "custom-app"
  kind: "app"
  dirs:
    web: "src/web"
    public: "static"
    internal: "pkg"
    cmd: "entrypoints"
```

Kompiler `kiw` akan me-resolve semua path terhadap pemetaan ini secara otomatis.

---

## Rangkuman

Dengan pemahaman mengenai instalasi, konfigurasi, dan tata letak direktori, Anda siap membangun aplikasi yang tangguh di seluruh varian workload Krewire.

Lanjutkan ke [**2.4 Pengembangan Agentik →**](/getting-started/agentic-development) untuk melengkapi proyek Anda dengan agen coding AI otonom.
