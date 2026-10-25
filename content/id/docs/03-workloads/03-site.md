---
title: "Generator Situs Statis (site)"
description: "Kompilasi situs statis berkinerja tinggi dengan perutean berbasis file, komponen .kiw berkas-tunggal, scoped CSS saat kompilasi, dan nol JavaScript sisi klien."
date: "2026-10-08"
---

# Generator Situs Statis (`site`)

Workload **`site`** menyediakan generator situs statis (SSG) super cepat yang dirancang untuk situs web pemasaran, landing page, blog, dan portal dokumentasi publik.

Ditenagai oleh `pkg/web/ssg` dan DSL `.kiw`, compiler `site` mengubah Single-File Components (SFC) menjadi HTML dan CSS semantik yang teroptimasi dengan **nol overhead JavaScript runtime secara default**.

---

## Karakteristik Arsitektur

1. **Perutean Berbasis File:** Halaman dipetakan langsung dari hierarki sistem file:
   - `pages/index.kiw` → `/index.html` (disajikan di `/`)
   - `pages/about.kiw` → `/about.html` (disajikan di `/about`)
   - `pages/blog/[slug].kiw` → Pembuatan slug dinamis dari frontmatter Markdown.
2. **Scoped CSS Saat Kompilasi:** Gaya visual yang dideklarasikan dalam blok `<style>` di-hash secara otomatis dan diberi cakupan (scoped) dengan atribut `data-kiw-component`. Tidak ada tabrakan CSS atau perlunya konvensi penamaan yang rumit.
3. **Bebas Bloat Langkah Build:** Mengompilasi puluhan halaman dalam waktu kurang dari 15 md menggunakan Go murni. Tidak memerlukan Node.js atau folder `node_modules`.
4. **Plugin Tailwind CSS:** Integrasi bawaan dengan Tailwind CSS: jika `tailwind.config.js` tersedia, `kiw` secara otomatis memicu pembuatan utility class yang diminifikasi.

---

## Scaffolding Proyek `site`

```bash
kiw new my-site --site
cd my-site
```

Tata letak direktori standar:

```text
my-site/
├── krewire.yaml          # Konfigurasi proyek
├── layouts/
│   └── Base.kiw          # Layout dokumen HTML
├── pages/
│   ├── index.kiw         # Halaman utama
│   └── about.kiw         # Halaman tentang
├── components/
│   ├── Navbar.kiw        # Navigasi pakai-ulang
│   └── Footer.kiw        # Footer pakai-ulang
├── public/               # File statis mentah yang disalin langsung
│   ├── favicon.svg
│   └── robots.txt
└── assets/
    └── tailwind.css      # Stylesheet Tailwind opsional
```

---

## Menulis Komponen (`pages/index.kiw`)

```html
---
title: "Modern Web Engineering in Go"
layout: Base
---

<section class="hero">
  <div class="badge">Ecosystem Release v0.1.0</div>
  <h1>Modular Go Libraries. Zero Bloat.</h1>
  <p>Compile production web surfaces directly with native Go tooling.</p>
  <a href="/about" class="btn">Learn More →</a>
</section>

<style>
.hero {
  padding: 4rem 1.5rem;
  text-align: center;
  background: var(--bg);
}
.badge {
  display: inline-block;
  padding: 0.25rem 0.75rem;
  background: var(--primary);
  color: var(--primary-content);
  font-weight: 700;
  border-radius: 999px;
  font-size: 0.85rem;
  margin-bottom: 1.5rem;
}
h1 {
  font-size: 2.75rem;
  font-weight: 900;
  letter-spacing: -0.03em;
  color: var(--fg);
}
p {
  color: var(--muted);
  font-size: 1.15rem;
  max-width: 600px;
  margin: 1rem auto 2rem;
}
.btn {
  display: inline-block;
  background: var(--primary);
  color: var(--primary-content);
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 700;
  text-decoration: none;
}
</style>
```

---

## Dokumen Layout (`layouts/Base.kiw`)

Layout membungkus halaman dengan metadata HTML, gaya visual global, dan slot konten:

```html
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <link rel="stylesheet" href="/assets/style.css">
</head>
<body>
  <Navbar />
  <main>{{.Content}}</main>
  <Footer />
</body>
</html>
```

---

## Pengembangan & Deployment

### Menjalankan Dev Server dengan Live Reload

```bash
kiw dev
```

Buka `http://localhost:8080`. Perubahan pada file `.kiw`, layout, dan komponen akan memicu reload instan dalam waktu `< 10ms`.

### Membangun Situs Statis

```bash
kiw build
