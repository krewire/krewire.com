---
title: "Quickstart"
description: "Quickstart pemula lima menit: instal kiw, scaffold situs, ubah halaman, tambahkan komponen, jalankan server pengembangan, dan build untuk produksi."
date: "2026-11-09"
---

# Quickstart

<nav class="track-crumbs" aria-label="Jalur pemula">
<span class="track-crumbs-label">Jalur pemula</span>
<ol>
<li><a href="/">Pilih jalur</a></li>
<li><a href="/getting-started">Memulai</a></li>
<li aria-current="step"><span>1 · Quickstart</span></li>
<li><a href="/getting-started/fullstack-app-tutorial">2 · Aplikasi fullstack</a></li>
<li><a href="/getting-started/dsl">3 · DSL .kiw</a></li>
</ol>
</nav>

Dalam lima menit Anda akan menginstal CLI `kiw`, membuat situs statis, mengubah halaman, menulis komponen pertama Anda, dan menghasilkan build produksi. Anda hanya memerlukan Go 1.27.1+ dan terminal — tanpa Node.js, tanpa npm, dan tanpa pengetahuan tentang internal Krewire.

> **Catatan:** Setiap perintah dan berkas di halaman ini telah dijalankan terhadap `kiw` yang di-build dari `github.com/krewire/krewire` pada commit `33fe553` (November 2026). Baris output disalin dari eksekusi tersebut; timestamp dan path akan berbeda di mesin Anda.

---

## Langkah 1: Instal `kiw`

Di Linux atau macOS:

```bash
curl -fsSL https://krewire.com/scripts/install.sh | sh
```

Di platform apa pun dengan Go 1.27.1+ (termasuk Windows):

```bash
go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
```

Pastikan shell dapat menemukannya:

```bash
kiw version
```

Baris pertama berbunyi `kiw Krewire Devtool`, diikuti versi ekosistem dan versi Go Anda. Jika perintah tidak ditemukan, lihat [Instalasi → Mengonfigurasi PATH Anda](/getting-started/installation).

---

## Langkah 2: Buat Situs

```bash
kiw new my-site --site --title "My Site"
cd my-site
```

`kiw` mencetak berkas-berkas yang dibuatnya:

```text
created my-site/krewire.yaml
created my-site/pages/index.kiw
created my-site/layouts/Base.kiw
created my-site/components/Hero.kiw
created my-site/public/favicon.svg
created my-site/README.md
```

Kegunaan masing-masing berkas:

```text
my-site/
├── krewire.yaml         # project name, kind (site), build output, dev port
├── go.mod               # Go module (created by kiw new)
├── pages/index.kiw      # the home page → /
├── layouts/Base.kiw     # the HTML shell every page renders inside
├── components/Hero.kiw  # a reusable component, used as <Hero /> in index.kiw
└── public/favicon.svg   # copied as-is into the build
```

Berkas di `pages/` menjadi URL, berkas di `components/` menjadi tag, dan berkas di `layouts/` membungkus halaman. Itulah seluruh model mental untuk sebuah situs.

---

## Langkah 3: Jalankan Server Pengembangan

```bash
kiw dev
```

```text
INFO msg="dev server running" url=http://localhost:8080 ...
INFO msg="watching for changes" root=... interval=500ms
```

Buka <http://localhost:8080>. Biarkan server tetap berjalan untuk langkah-langkah berikutnya.

---

## Langkah 4: Ubah Halaman

Ganti isi `pages/index.kiw` dengan:

```html
---
title: My Site
layout: Base
---

<div class="hero-container">
  <h1>Hello from {{.Title}}</h1>
  <p>I edited this page and the browser reloaded by itself.</p>
  <Badge variant="primary" text="Built with Krewire" />
</div>

<style>
  .hero-container {
    max-width: 680px;
    margin: 4rem auto;
    text-align: center;
  }
  .hero-container h1 {
    font-size: 2.75rem;
    color: #39D353;
  }
</style>
```

Simpan. Terminal mencatat `site rebuilt successfully` dan peramban menampilkan judul yang baru.

- Blok di antara baris `---` adalah **frontmatter**: `title` tersedia sebagai `{{.Title}}`, `layout: Base` memilih `layouts/Base.kiw`.
- `<Badge … />` adalah komponen dari **Krewire UI** — 66 komponen bawaan tersedia di setiap situs tanpa perlu menginstal apa pun. Jelajahi semuanya di [katalog UI](/ui/components).
- Blok `<style>` bersifat **scoped**: aturannya hanya berlaku untuk markup halaman ini.

---

## Langkah 5: Tambahkan Komponen Anda Sendiri

Buat `components/Feature.kiw`:

```html
<div class="feature">
  <h3>{title}</h3>
  <p>{body}</p>
</div>

<style>
.feature {
  padding: 1rem 1.25rem;
  border: 1px solid #30363d;
  border-radius: 8px;
  text-align: left;
}
.feature h3 {
  margin: 0 0 0.5rem;
  color: #39D353;
}
</style>
```

`{title}` membaca atribut `title`; `{body}` adalah apa pun yang Anda letakkan di antara tag pembuka dan penutup. Gunakan di `pages/index.kiw`, di bawah badge:

```html
  <div class="features">
    <Feature title="Fast">Pages are plain HTML generated at build time.</Feature>
    <Feature title="Simple">One CLI, one config file, no npm.</Feature>
  </div>
```

lalu tambahkan `.features { display: grid; gap: 1rem; margin-top: 2rem; }` ke blok `<style>` halaman tersebut. Simpan, dan kedua kartu akan muncul.

> **Tips:** `kiw make component Feature` juga membuatkan `components/Feature.kiw` untuk Anda, sebagai titik awal untuk diedit.

---

## Langkah 6: Build untuk Produksi

Hentikan server pengembangan (`Ctrl+C`) lalu jalankan:

```bash
kiw build
```

```text
INFO building SSG site from file layout root=... output=.../my-site/.krewire/build
created index.html
created favicon.svg
created assets/krewire.css
```

`.krewire/build/` adalah situs web yang lengkap: HTML dan CSS biasa yang dapat disajikan oleh host statis apa pun (Nginx, Caddy, GitHub Pages, object bucket). Pratinjau persis seperti saat nanti disajikan:

```bash
kiw serve --addr :3000
```

lalu buka <http://localhost:3000>.

---

## Yang Telah Anda Pelajari

| Konsep | Letaknya |
| :--- | :--- |
| Halaman adalah berkas | `pages/*.kiw` → masing-masing satu URL |
| Layout membungkus halaman | `layout: Base` → `layouts/Base.kiw` |
| Komponen adalah tag | `components/Feature.kiw` → `<Feature>` |
| UI bawaan | `<Badge>`, `<Button>`, `<Card>`, … dari Krewire UI |
| Satu perintah per langkah | `kiw new` → `kiw dev` → `kiw build` → `kiw serve` |

## Langkah Selanjutnya

- Bangun sesuatu yang dinamis: [**2.2 Membangun Aplikasi Fullstack dalam 15 Menit →**](/getting-started/fullstack-app-tutorial)
- Pelajari sintaks komponen selengkapnya: [**2.7 DSL `.kiw` →**](/getting-started/dsl)
- Ingin menulis dokumentasi? Baca [Workload Book](/workloads/book).
