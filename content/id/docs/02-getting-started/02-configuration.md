---
title: "Konfigurasi"
description: "Spesifikasi dan panduan mendalam untuk krewire.yaml — konfigurasi devtool deklaratif terpadu."
date: "2026-09-29"
---

# Konfigurasi

Setiap proyek Krewire dikendalikan oleh satu berkas konfigurasi deklaratif: **`krewire.yaml`**.

`krewire.yaml` bertindak sebagai **Single Source of Truth (SSOT)** untuk metadata proyek, pengaturan pipeline kompilasi, server pengembangan lokal, dan skrip otomatisasi kustom.

---

## 1. Filosofi Arsitektur

Krewire mempertahankan pemisahan tanggung jawab (Separation of Concerns) yang ketat antara konfigurasi devtool dan logika runtime aplikasi:

1. **Kontrol Devtool (`krewire.yaml`):** Menginstruksikan CLI `kiw` tentang cara membangun, menjalankan, menguji, dan memaketkan proyek.
2. **Pengaturan Aplikasi:** Judul halaman web, layout, tag meta, dan styling komponen berada di dalam halaman, template, dan kode Go—tidak pernah diduplikasi di dalam YAML.
3. **Toolchain Terpisah (Decoupled):** Proyek dapat dikompilasi secara bersih menggunakan perintah standar `go build` bahkan jika `krewire.yaml` tidak ada, dengan kembali ke konvensi standar Go.

---

## 2. Referensi Skema Lengkap

Berikut adalah skema kanonikal lengkap dari `krewire.yaml` beserta nilai default produksinya:

```yaml
# =============================================================================
# Konfigurasi Devtool Krewire (krewire.yaml)
# =============================================================================

# Identitas Proyek & Jenis Workload
project:
  name: "my-service"            # Pengidentifikasi proyek (dibaca oleh site loader)
  kind: "site"                  # Workload: app | cli | site | book | worker | service | gateway | infra
  version: "v0.1.0"             # Versi semantik yang diinjeksikan sebagai `.Version` ke dalam halaman
  dirs:                         # Opsional: Menimpa lokasi direktori kanonikal
    web: "web"                  # Handler dan template web
    public: "public"            # Aset statis mentah
    internal: "internal"        # Paket domain privat
    cmd: "cmd"                  # Titik masuk executable

# Input Konten (level teratas)
input: "content"                # Direktori konten yang digunakan oleh pipeline book

# Pipeline Build Devtool
build:
  output: ".krewire/build"      # Direktori target output kompilasi
  base: "/"                     # URL dasar tempat situs disajikan
  include:                      # Pola glob untuk penyertaan konten
    - "**/*.md"
  exclude:                      # Pola glob untuk pengecualian
    - "**/README.md"
    - "**/readme.md"
  depends_on:                   # Perintah shell yang dijalankan sebelum build
    - "go generate ./..."

# Injeksi CSS/JS Otomatis (jenis site; aktif secara default)
auto_assets:
  enabled: true                 # Setel false untuk menautkan setiap aset secara manual di layouts/
  js_placement: "head"          # head (default, skrip first-paint) | body
  exclude:                      # Jangan pernah menginjeksi aset-aset ini secara otomatis
    - "*.min.css"
  order:                        # Menentukan urutan pemuatan aset tertentu
    "assets/theme.css":
      layer: "page"             # scoped | vendor | component | theme | book | page
  bundle: false                 # true menggabungkan CSS/JS ke dalam assets/krewire.{css,js}

# Pipeline Dokumentasi Book (pkg/book)
book:
  mount: "/docs/"               # Path mount dalam ruang URL output (mode hybrid)
  toc: true                     # Menghasilkan halaman daftar isi (table of contents) utama

# Runtime Klien WebAssembly (KWF-T4X9P)
wasm:
  entry: "./wasm"               # Titik masuk paket utama untuk GOOS=js
  name: "runtime"               # Nama berkas modul wasm yang dihasilkan

# Skrip Task Runner Kustom
scripts:
  lint: "golangci-lint run ./..."
  fmt: "gofmt -s -w ."
  test: "go test -race -v ./..."
  seed: "go run cmd/seed/main.go"
```

---

## 3. Rincian Konfigurasi

### 3.1 Blok Proyek (`project:`)

Bagian `project:` mendefinisikan identitas dan perilaku workload:

- **`kind` (Wajib):** Menetapkan jenis arsitektur workload. Konfigurasi ini menginstruksikan `kiw` mengenai engine mana yang harus dipanggil:
  - `app`: Monolit web fullstack dengan aset tersemat (`kiw run`)
  - `cli`: Alat command-line terminal atau TUI (`kiw run`)
  - `site`: Generator situs statis berbasis komponen `.kiw` ber-scope (`kiw build`)
  - `book`: Manuskrip dokumentasi Markdown melalui `pkg/book` (`kiw build --target book`)
  - `worker`: Pemroses tugas latar belakang asinkron (`kiw worker`)
  - `service`: Microservice dengan throughput tinggi (`kiw run`)
  - `gateway`: Reverse proxy deklaratif dengan middleware per rute (`kiw gateway`)
  - `infra`: Infrastruktur sebagai kode dalam Go (`kiw deploy --target infra`)

  Proyek bootstrap sebelum perancangan kerangka (pre-scaffold) **bukanlah** suatu jenis kind: `kiw info` melaporkan `none` dengan penanda `pre-init` hingga `kiw init` melengkapi variannya.
- **`dirs` (Opsional):** Memungkinkan penyesuaian tata letak direktori default jika mengintegrasikan Krewire ke dalam struktur repositori yang sudah ada.

---

### 3.2 Input Konten (`input:`)

- **`input`:** Direktori konten yang digunakan oleh pipeline buku dokumentasi (`pkg/book`). Nilai defaultnya adalah `content`. Dengan mengaktifkan `i18n:`, `kiw build` me-resolve `content/<locale>/docs/` per lokal dan hanya beralih ke nilai ini jika lokal tidak ditemukan.

---

### 3.3 Blok Build (`build:`)

Mengatur lokasi penempatan artefak dan berkas konten mana yang diproses oleh pipeline:

- **`output`:** Direktori tempat penampungan berkas statis atau aset hasil kompilasi. Nilai defaultnya adalah `.krewire/build`. Binary `app`/`cli` yang dikompilasi ditempatkan di `<output>/bin/<nama proyek>`.
- **`base`:** Awalan URL dasar tempat aset dan halaman disajikan. Nilai defaultnya adalah `/`. Untuk situs yang di-hosting di bawah sub-path (misalnya `https://example.com/blog/`), atur `base: "/blog/"`.
- **`include` & `exclude`:** Pola glob yang mengatur manuskrip Markdown atau berkas konten mana yang diproses. Jika `include` tidak diatur, nilainya default ke `**/*.md`; jika `exclude` tidak diatur, defaultnya adalah mengabaikan `README.md`/`readme.md`. Daftar kosong akan menonaktifkan filter tersebut.
- **`depends_on`:** Perintah shell yang dijalankan secara berurutan sebelum proses build dimulai.
- **`outputs`:** Nama runtime ke direktori output, untuk proyek polyglot.

Kunci level teratas `output:` dan `base:` adalah alias warisan (legacy). Nilai `build.*` bertingkat akan diprioritaskan jika keduanya disetel, dan flag `--output` / `--base` mengesampingkan keduanya. Deklarasikan setiap nilai sekali saja — jangan mengulanginya di kedua tingkatan.

---

### 3.4 Port Dev Server (`dev:`)

Blok `dev:` mendeklarasikan port server pengembangan lokal:

```yaml
dev:
  port: 8080
```

Alamat listen diselesaikan dengan urutan prioritas berikut: flag `--addr`, lalu `runtime.port`, `dev.port`, `server.port`, kemudian `services[].port` pertama, dan terakhir `:8080`:

```bash
kiw dev --addr :3000
kiw serve --addr :3000
```

---

### 3.5 Blok Auto Assets (`auto_assets:`)

Berlaku untuk jenis workload `site`. Secara default, proses build menginjeksikan `<link rel="stylesheet">` ke dalam `<head>` dan `<script src>` ke dalam `<head>` untuk setiap berkas `.css`/`.js` di bawah `public/assets/`, ditambah stylesheet scoped yang dihasilkan `assets/style.css`. Layout tidak perlu mencantumkannya secara manual, dan tag yang Anda tulis sendiri tidak akan diduplikasi. Mekanisme cache busting menggunakan `?v=<versi proyek>`.

- **`enabled`:** `false` mengembalikan ke tag aset yang sepenuhnya manual.
- **`js_placement`:** `head` (default) menjaga agar skrip tema berjalan sebelum first paint; `body` menempatkannya di akhir elemen `<body>`.
- **`exclude`:** Nama aset, nama dasar (base name), atau glob yang tidak boleh diinjeksikan.
- **`order`:** Mengunci suatu aset ke layer cascade tertentu ketika posisi default tidak sesuai dengan proyek Anda. Nama layer yang tidak dikenali akan dilaporkan sebagai peringatan build (warning) dan urutan default tetap dipertahankan, sehingga kesalahan ketik dapat terlihat tanpa menggagalkan proses build.
- **`bundle`:** `true` menggabungkan CSS dan JS yang memenuhi syarat ke dalam `assets/krewire.css` dan `assets/krewire.js`, mengubah permintaan per-aset menjadi satu berkas untuk masing-masing tipe.

Aset **tidak** diinjeksikan secara alfabetis — urutan alfabetis dapat memuat konten buku sebelum utilitas yang seharusnya menimpanya. Aset dimuat berdasarkan urutan cascade, sehingga setiap layer dapat menimpa layer sebelumnya:

| Layer | Memuat |
| :--- | :--- |
| `scoped` | style yang dihasilkan dari `<style>` dalam komponen/layout |
| `vendor` | output pihak ketiga seperti CSS plugin (Tailwind) |
| `component` | styling pustaka komponen (Krewire UI) |
| `theme` | variabel tema dan penimpaan desain last-mile |
| `book` | styling pipeline konten (pkg/book) |
| `page` | penimpaan spesifik per-halaman — selalu terakhir |

```yaml
auto_assets:
  enabled: true
  js_placement: "body"
  exclude:
    - "vendor.js"
  order:
    "assets/theme.css":
      layer: "page"   # memindahkan tema ke posisi last-mile
```

---

### 3.6 Task Runner Scripts (`scripts:`)

Krewire menyertakan task runner terintegrasi, menghilangkan kebutuhan akan `Makefile` atau alat eksternal tambahan. Setiap tugas yang dideklarasikan dalam `scripts:` dapat dieksekusi langsung melalui `kiw run <task>`:

```yaml
scripts:
  lint: "golangci-lint run ./..."
  audit: "go run cmd/security-audit/main.go"
  generate: "go generate ./..."
```

Menjalankan tugas tertentu:

```bash
kiw run lint
kiw run audit
```

Tugas dieksekusi menggunakan shell sistem operasi native (`sh -c` di Unix/macOS, `cmd /C` di Windows).

---

## 4. Variabel Lingkungan (Environment Variables)

Konfigurasi runtime dapat ditimpa menggunakan variabel lingkungan tanpa perlu mengubah `krewire.yaml`:

| Variabel | Nilai | Deskripsi |
| :--- | :--- | :--- |
| `KIW_ENV` | `local`, `production`, `testing` | Menetapkan profil lingkungan aktif. |
| `KIW_DEBUG` | `true`, `false`, `1`, `0` | Mengaktifkan logging debug dan diagnostik terperinci. |

Contoh:

```bash
KIW_ENV=production KIW_DEBUG=false kiw run
```

---

## Langkah Berikutnya

Setelah konfigurasi Anda siap, lanjutkan ke [**2.3 Struktur Direktori →**](/getting-started/directory-structure) untuk mempelajari tata letak direktori standar bagi setiap jenis workload.
