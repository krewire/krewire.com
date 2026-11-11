---
title: "DSL (.kiw)"
description: "Panduan mendalam tentang bahasa templating komponen .kiw: YAML frontmatter, ekspresi Go html/template, CSS ber-scope, skrip klien, dan Markdown tersemat."
date: "2026-09-29"
---

# DSL (.kiw)

**DSL Komponen `.kiw`** (`pkg/kiw`) adalah bahasa templating komponen berbasis berkas terpadu milik Krewire. Bahasa ini menghadirkan ergonomi Single-File Component (SFC) modern—serupa dengan Astro, Svelte, atau Vue—langsung ke ekosistem Go tanpa memerlukan Node.js, `npm`, webpack, ataupun kerumitan build JavaScript.

Setiap berkas `.kiw` merangkum metadata, markup, scoped style, dan interaksi sisi klien di dalam satu berkas yang mudah dibaca, yang dikompilasi menjadi pohon template Go berkinerja tinggi serta aset statis hanya dalam hitungan milidetik.

---

## 1. Anatomi Berkas `.kiw`

Sebuah komponen kanonikal `.kiw` terdiri dari empat bagian opsional yang terpisah:

```html
---
# 1. YAML Frontmatter
title: "User Profile"
layout: Base
roles: ["admin", "editor"]
---

<!-- 2. Body Template (Go html/template) -->
<article class="profile-card">
  <h2>{{.Title}}</h2>
  <p>Status: <span class="badge">Active</span></p>

  <!-- 3. Markdown Tersemat (Opsional) -->
  <markdown>
  ### Bio
  Pengembang fullstack yang antusias dengan **Go** dan rekayasa sistem.
  </markdown>
</article>

<!-- 4. Scoped CSS (Diisolasi Otomatis) -->
<style>
.profile-card {
  border: var(--pop-border, 2px solid #0B1F3B);
  border-radius: 12px;
  padding: 1.5rem;
  background: var(--bg-2, #f0ede5);
}
h2 {
  color: var(--primary, #39D353);
  margin-top: 0;
}
.badge {
  font-weight: bold;
  text-transform: uppercase;
}
</style>

<!-- 5. Skrip Klien (Diekstraksi Otomatis) -->
<script hydrate="idle">
console.log("Profile card mounted and active.");
</script>
```

---

## 2. YAML Frontmatter

Bagian atas dari setiap halaman atau berkas layout dapat mendeklarasikan frontmatter yang dibatasi oleh tiga tanda hubung (`---`). Frontmatter diparsing ke dalam struktur data Go bertipe melalui `gopkg.in/yaml.v3`:

```yaml
---
title: "Panduan Dokumentasi"
description: "Pelajari cara membangun portal dokumentasi."
layout: Base
date: "2026-09-29"
draft: false
tags:
  - golang
  - architecture
author:
  name: "Tim Krewire"
  email: "team@krewire.com"
---
```

### Kunci Frontmatter Khusus:

| Kunci | Tujuan | Digunakan Di |
| :--- | :--- | :--- |
| `title` | Judul halaman yang diinjeksikan ke dalam tag `<title>` layout dan header OpenGraph. | `pages/` |
| `layout` | Nama pembungkus layout di dalam `layouts/` (misalnya `layout: Base` membungkus dengan `layouts/Base.kiw`). | `pages/` |
| `description` | Deskripsi meta untuk mesin pencari dan kartu media sosial. | `pages/` |
| `draft` | Jika disetel ke `true`, halaman tidak akan dimasukkan ke dalam build produksi. | `pages/` |
| `date` | Timestamp publikasi atau modifikasi berformat ISO 8601. | `pages/`, `content/` |

Semua kunci kustom yang dideklarasikan dalam frontmatter dapat diakses langsung di dalam ekspresi template melalui root context: `{{.Title}}`, `{{.Tags}}`, `{{.Author.Name}}`.

---

## 3. Body Template & Ekspresi Go

Body dari berkas `.kiw` menggunakan sintaksis standar `html/template` bawaan Go. Ini menyediakan fitur escaping HTML yang sadar konteks (context-aware) untuk mencegah celah Cross-Site Scripting (XSS) secara default.

### Ekspresi Umum:

#### Interpolasi Data
```html
<h1>{{.Title}}</h1>
<p>Penulis: {{.Author.Name}}</p>
```

#### Render Bersyarat (Conditional Rendering)
```html
{{if .IsAdmin}}
  <a href="/admin/dashboard" class="btn btn-danger">Admin Panel</a>
{{else if .IsSubscriber}}
  <span class="badge">Subscriber</span>
{{else}}
  <a href="/signup" class="btn">Daftar Sekarang</a>
{{end}}
```

#### Looping pada Slice dan Map
```html
<ul class="tag-list">
  {{range .Tags}}
    <li class="tag-item">{{.}}</li>
  {{else}}
    <li class="empty">Tidak ada tag tersedia</li>
  {{end}}
</ul>
```

#### Helper Komposisi Kelas (`class`)
Untuk meringkas blok modifikasi bersyarat yang panjang di dalam atribut class tanpa melewati escaping kontekstual HTML:

```html
<!-- Mengomposisikan token dasar + modifier bersyarat + prop class tambahan -->
<a href="{href}" class="{{class "btn" .Variant "primary" .Size "md" .Class}}">
  {body}
</a>
```

Helper `class` secara otomatis menangani resolusi prefix (misalnya `btn-primary`, `btn-md`), kembali ke nilai default saat props kosong atau diabaikan, merapikan spasi (whitespace), dan mengembalikan string polos sehingga escaping atribut standar `html/template` tetap terjaga dengan ketat.

#### Injeksi Konten Layout
Pada komponen layout (`layouts/*.kiw`), konten halaman yang telah dirender diinjeksikan menggunakan direktif `{{.Content}}`:

```html
<!doctype html>
<html lang="id">
<head>
  <meta charset="utf-8">
  <title>{{.Title}}</title>
</head>
<body>
  <header><Header /></header>
  <main id="main">
    {{.Content}}
  </main>
  <footer><Footer /></footer>
</body>
</html>
```

Tidak diperlukan tag `<link>` atau `<script>` manual untuk CSS/JS situs: setiap berkas di `public/assets/` yang berakhiran `.css` atau `.js` akan diinjeksikan secara otomatis — stylesheet ke dalam `<head>`, skrip ke dalam `<head>` — dengan kueri cache-busting `?v=<version>` yang diambil dari versi proyek. Menulis tag secara manual tetap didukung; pipeline tidak akan menduplikasi aset yang sudah Anda tautkan sendiri. Lihat `auto_assets:` di [`krewire.yaml`](configuration) untuk mematikan injeksi atau mengecualikan berkas tertentu.

---

## 4. Komposisi Komponen yang Dapat Digunakan Kembali

Komponen yang terletak di direktori `components/` (atau dipasang di `./.krewire/pkg/`, atau disematkan dari Krewire UI) dapat diinstansiasi secara dinamis di dalam halaman, layout, atau komponen lainnya menggunakan tag komponen kustom:

```html
<!-- Pemanggilan berbasis tag dengan atribut sebagai props -->
<Button Variant="primary" Href="/docs/getting-started">Mulai Sekarang</Button>
<Alert Type="warning" Message="Pembaruan sistem sedang berlangsung" />
```

Bentuk ekuivalen menggunakan helper template Go adalah `{{component "ComponentName" .}}`; kedua bentuk ini menghasilkan pohon komponen yang persis sama.

### 4.1 Komponen Majemuk (Notasi Titik / Compound)

Komponen dapat diorganisasikan ke dalam subdirektori dan dipanggil menggunakan notasi majemuk bertitik:

```html
<Card>
  <Card.Header title="Status Deployment" />
  <Card.Body>
    <p>Semua layanan mikro beroperasi normal.</p>
  </Card.Body>
  <Card.Footer>
    <Button Variant="ghost">Tutup</Button>
  </Card.Footer>
</Card>
```

Scaffold komponen majemuk langsung menggunakan CLI:
```bash
kiw make component Card.Header     # menghasilkan components/Card/Header.kiw
kiw make component Card.Body       # menghasilkan components/Card/Body.kiw
kiw make component Card.Footer     # menghasilkan components/Card/Footer.kiw
```

### 4.2 Komponen Bernamespace & Vendor

Paket antarmuka pengguna yang dipasang dari `./.krewire/pkg/` atau design system bawaan dapat dipanggil menggunakan sintaks namespace titik dua:

```html
<ui:Button variant="primary">Jalankan</ui:Button>
<acme:MetricCard label="Permintaan Aktif" value=1420 />
```

### 4.3 Slot Outlet & Konten Cadangan (Fallback)

Konten anak yang diteruskan di dalam tag berpasangan mengisi slot default (`{body}`). Komponen dapat mendeklarasikan outlet slot bernama dengan konten cadangan default yang dirender secara otomatis bila pemanggil tidak menyediakan konten:

```html
<!-- components/Card/Header.kiw -->
<div class="card-header">
  <slot name="title">
    <h3>Judul Header Default</h3>
  </slot>
  <slot name="actions"></slot>
</div>
```

Pemanggil dapat menyediakan konten slot menggunakan `<template slot="...">` atau `<slot:...>`:

```html
<Card.Header>
  <template slot="title">Metrik Klaster Kustom</template>
  <template slot="actions">
    <Button size="sm">Segarkan</Button>
  </template>
</Card.Header>
```

### 4.4 Inline Snippet & Makro Lokal

Fragmen template yang dapat digunakan kembali dapat dideklarasikan dan dirender di dalam berkas `.kiw` yang sama tanpa perlu membuat berkas eksternal:

```html
{#snippet userBadge(name, role)}
  <span class="badge {role}">{name}</span>
{/snippet}

<div class="team-roster">
  {@render userBadge("Alice", "admin")}
  {@render userBadge("Bob", "contributor")}
</div>
```

Snippet dikonversi secara bersih ke dalam definisi dan pemanggilan template Go ber-scope.

### 4.5 Impor Komponen Deklaratif

Komponen dapat mendeklarasikan impor di dalam blok `<script>` atau frontmatter:

```html
<script>
  import Header from './Header.kiw';
  import { Badge } from '@krewire/ui';
</script>

<Header />
<Badge variant="success">Online</Badge>
```

### 4.6 Paket Komponen Mandiri (`.kpkg`)

Pustaka komponen yang dapat digunakan kembali dapat didistribusikan dan dipasang sebagai paket mandiri tanpa dependensi npm:
- **Manifest (`kiw.pkg.json`)**: Mendeklarasikan metadata paket, pemetaan komponen yang diekspor, dan stylesheet.
- **Format Arsip (`.kpkg`)**: Arsip tarball gzipped dengan verifikasi integritas SHA-256.
- **Direktori Pemasangan (`./.krewire/pkg/`)**: Repositori paket proyek lokal.

```bash
kiw pkg init @acme/ui              # Men-scaffold direktori paket dengan kiw.pkg.json
kiw pkg pack ./acme-ui             # Mengemas ke arsip acme-ui-0.1.0.kpkg
kiw pkg add ./acme-ui-0.1.0.kpkg   # Memasang ke ./.krewire/pkg/acme-ui
kiw pkg list                       # Menampilkan daftar paket dan komponen yang diekspor
```

Stylesheet yang dideklarasikan di `kiw.pkg.json` (`styles: "./styles/theme.css"`) secara otomatis dikompilasi dan disuntikkan ke dalam pipeline aset situs.

### Contoh: Komponen dengan Props Ber-Scope

Definisikan `components/Alert.kiw`:

```html
<div class="alert alert-{{.Type}}">
  <span class="icon">ℹ</span>
  <div class="message">{{.Message}}</div>
</div>

<style>
.alert {
  display: flex;
  gap: 0.75rem;
  padding: 1rem;
  border-radius: 8px;
}
.alert-info { background: #e0f2fe; color: #0369a1; }
.alert-warning { background: #fef3c7; color: #92400e; }
</style>
```

Instansiasi `Alert` di dalam `pages/index.kiw`:

```html
<Alert Type="warning" Message="Pemeliharaan dijadwalkan malam ini pukul 23:00 UTC" />
```

### Island Architecture & Hidrasi Sisi Klien

Krewire mengimplementasikan **Island Architecture** (`pkg/runtime/mount` dan `pkg/runtime/boot`) yang memungkinkan hidrasi sisi klien secara selektif untuk komponen interaktif, sementara bagian halaman lainnya tetap berupa HTML statis 100% tanpa JavaScript.

Anda mendeklarasikan perilaku hidrasi menggunakan direktif `client:*` langsung pada tag komponen:

```html
<!-- Hidrasi langsung saat halaman dimuat -->
<Counter client:load initial=0 />

<!-- Hidrasi setelah thread utama browser menganggur/idle (requestIdleCallback) -->
<SearchDialog client:idle placeholder="Cari dokumentasi..." />

<!-- Hidrasi hanya ketika elemen masuk ke dalam viewport (IntersectionObserver) -->
<CommentThread client:visible threadId="kw-99" />
```

**Cara Kerjanya:**
1. **Server Rendering (SSR/SSG)**: Selama proses build, Krewire me-render komponen menjadi HTML dan membungkusnya dalam wadah penanda ringan (`data-kiw-mount="Counter" data-kiw-hydrate="load" data-kiw-props='{"initial":0}'`).
2. **Pratinjau Instan**: Browser menampilkan konten lengkap dengan segera tanpa adanya pergeseran tata letak (CLS: 0).
3. **Hidrasi Progresif**: Runtime WebAssembly Krewire (`pkg/runtime`) memindai penanda mount, menginisialisasi fiber VDOM dengan DOM yang sudah ada, dan memasang event reaktif tanpa me-render ulang teks server.

---

## 5. CSS Ber-Scope Otomatis

Style yang didefinisikan di dalam blok `<style>` secara **otomatis diisolasi ke komponen terkait** pada saat proses build. Anda tidak perlu khawatir tentang tabrakan nama class CSS.

### Cara Kerja Scoping:

1. Saat mem-parsing komponen (misalnya `components/Card.kiw`), Krewire menyisipkan atribut unik ke elemen root komponen:
   ```html
   <div class="card" data-kiw-component="Card">...</div>
   ```
2. Kompiler CSS (`pkg/web/ssg`) menulis ulang selektor Anda agar terikat ketat pada scope komponen tersebut:
   ```css
   /* CSS yang Anda tulis */
   .card { padding: 20px; }
   h3 { font-size: 1.25rem; }

   /* Output CSS yang dikompilasi di assets/style.css */
   [data-kiw-component="Card"] .card, [data-kiw-component="Card"].card { padding: 20px; }
   [data-kiw-component="Card"] h3, [data-kiw-component="Card"]h3 { font-size: 1.25rem; }
   ```
3. Semua style komponen digabungkan dan diminifikasi menjadi satu berkas `assets/style.css`, menghilangkan permintaan CSS yang memblokir rendering halaman.

---

## 6. Skrip Klien & Direktif Hidrasi

Skrip di dalam tag `<script>` diekstraksi dari HTML dan dikompilasi menjadi berkas aset JavaScript statis. Ini menjamin Pemisahan Tanggung Jawab (Separation of Concerns) yang bersih serta kepatuhan ketat pada Kebijakan Keamanan Konten (Content Security Policy - CSP) tanpa risiko eksekusi skrip inline.

```html
<script hydrate="visible">
  const button = document.querySelector(".btn-interactive");
  button.addEventListener("click", () => {
    console.log("Tombol diklik!");
  });
</script>
```

### Tingkatan Hidrasi (Atribut `hydrate`):

| Direktif Hidrasi | Waktu Eksekusi | Kasus Penggunaan Ideal |
| :--- | :--- | :--- |
| `hydrate="load"` *(default)* | Dieksekusi langsung saat DOM dimuat (`DOMContentLoaded`). | Menu navigasi, toggle UI penting, pengubah tema. |
| `hydrate="idle"` | Dieksekusi saat thread utama browser menganggur (`requestIdleCallback`). | Analitik, telemetri non-kritis, widget latar belakang. |
| `hydrate="visible"` | Dieksekusi hanya saat elemen memasuki viewport (`IntersectionObserver`). | Carousel gambar, grafik interaktif yang dimuat malas (lazy-loaded), kolom komentar. |

---

## 7. Markdown Tersemat

Untuk bagian yang berfokus pada konten tekstual, Anda dapat menulis Markdown native langsung di dalam komponen `.kiw` menggunakan tag `<markdown>`:

```html
<section class="changelog">
  <h2>Catatan Rilis</h2>
  <markdown>
  ### Versi 0.1.0 (Rilis Perdana)
  - **Pustaka:** Arsitektur pustaka modular inti.
  - **Kiw CLI:** Scaffolding, dev server, dan pipeline build.
  - **Nol Bloat:** 100% pustaka standar Go, nol dependensi npm.
  </markdown>
</section>
```

Kompiler Krewire secara otomatis mengonversi blok Markdown menjadi HTML yang bersih dan tersanitasi pada waktu build menggunakan `pkg/kiw` dan `pkg/markdown`.

---

## 8. Helper Bawaan & Internasionalisasi (i18n)

Krewire menyediakan fungsi-fungsi helper template bawaan:

| Helper | Sintaksis | Deskripsi |
| :--- | :--- | :--- |
| Tag Komponen | `<Name />` | Me-render komponen turunan menggunakan sintaksis mirip JSX. |
| `component` | `{{component "Name" .}}` | Bentuk helper template yang kompatibel untuk me-render komponen turunan dengan konteks atau props dict. |
| `t` / `translate` | `{{t "home.welcome"}}` | Menerjemahkan kunci menggunakan bundle bahasa yang aktif (`lang/*.json`). |
| `tLocale` | `{{tLocale "id" "home.welcome"}}` | Menerjemahkan kunci untuk lokal tertentu. |
| `locales` | `{{locales}}` | Mengembalikan daftar semua kode bahasa yang dikonfigurasi. |

---

## Langkah Berikutnya

Setelah Anda menguasai DSL `.kiw` dan sistem komponen Krewire:

- Pelajari matriks [**1.3 Workload Krewire →**](/overview/krewire-workloads) untuk melihat bagaimana `.kiw` mendukung situs statis, buku, dan aplikasi web.
- Kembali ke [**2. Indeks Memulai (Getting Started) →**](/getting-started).
