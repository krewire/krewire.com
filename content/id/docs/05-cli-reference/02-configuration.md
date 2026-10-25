---
title: "Skema Konfigurasi (krewire.yaml)"
description: "Referensi kanonikal untuk krewire.yaml: identitas proyek, input build dan konten, dev server, blok book/SSG/i18n, serta setiap bagian spesifik workload."
date: "2026-10-25"
---

# Skema Konfigurasi (`krewire.yaml`)

Krewire beroperasi di bawah aturan **Satu Konfigurasi (One Config)**: semua pengaturan di seluruh delapan jenis workload dideklarasikan dalam satu berkas tunggal—**`krewire.yaml`**.

Tidak ada berkas konfigurasi terpisah seperti `ssg.yaml` atau berkas spesifik perkakas lainnya. Konfigurasi memiliki tipe data kuat, divalidasi terhadap aturan skema (`pkg/validation`), dan dimuat menggunakan `pkg/config` (pemilik skema: `tools/kiw/internal/config`).

---

## Ringkasan Skema Kanonikal

```yaml
# 1. Identitas Proyek
project:
  name: "my-service"            # Nama unik proyek (pengenal aman-DNS)
  kind: "service"               # app, cli, site, book, worker, service, gateway, infra
  version: "v0.1.0"             # Versi produk, disuntikkan sebagai .Version ke halaman

title: "My Service"             # Judul situs/buku
author: "Krewire Team"          # Penulis situs/buku
env: "local"                    # local, production, testing (dapat ditimpa oleh KIW_ENV)

# 2. Input Konten
input: "content"                # Direktori konten untuk workload site dan book

# 3. Pipeline Build
build:
  output: ".krewire/build"      # Direktori output (default .krewire/build)
  base: "/"                     # URL dasar tempat situs disajikan
  include: ["**/*.md"]          # Pola glob konten yang disertakan
  exclude: ["**/README.md"]     # Pola glob konten yang diabaikan setelah disertakan
  depends_on: []                # Perintah shell yang dijalankan sebelum build

# 4. Server Pengembangan Lokal
dev:
  port: 8080                    # Port server dev (integer; ditimpa oleh --addr)

# 5. Pengaturan Khusus Workload (hanya blok yang cocok dengan project.kind yang berlaku)

# kind: worker
worker:
  concurrency: 10               # Jumlah maksimum worker konkuren
  queues: ["default", "email"]  # Nama-nama antrean yang dikonsumsi worker ini

# kind: service
service:
  name: "user-service"
  port: 8080

# kind: gateway
gateway:
  host: ""
  port: 8080
  routes:
    - path: /api
      prefix: true
      strip_prefix: true
      target_url: http://localhost:8081
      timeout: 10s
    - path: /users
      service: user-service
      middleware: [cors]
      rate_limit: 100

# kind: infra
infra:
  provider: "aws"               # aws, gcp, kubernetes
  region: "ap-southeast-1"
  state: "s3://krewire-infra-state"
```

---

## Referensi Bagian Skema

### `project`

| Kunci | Tipe | Wajib | Deskripsi |
| :--- | :--- | :---: | :--- |
| `name` | String | Ya | Nama proyek atau layanan. |
| `kind` | String | Ya | Jenis workload (`app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`). |
| `version` | String | Tidak | Versi rilis semantik (`vX.Y.Z`); digabungkan ke dalam versi efektif yang digunakan oleh setiap pipeline. |

Proyek bootstrap pra-perancah (pre-scaffold) **bukan** sebuah kind: `kiw info` melaporkan `none` dengan penanda `pre-init` hingga `kiw init` memilih suatu varian.

### Metadata tingkat atas

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `title` | String | nama modul | Judul situs atau buku. |
| `author` | String | — | Penulis yang ditampilkan pada halaman buku dan metadata. |
| `version` | String | `project.version` | Versi produk yang disuntikkan sebagai `.Version`. |
| `env` | String | `local` | Lingkungan target (`local`, `production`, `testing`). |
| `debug` | Bool | `false` | Mengaktifkan mode debug (`KIW_DEBUG` menimpa ini). |
| `input` | String | `content` | Direktori konten untuk workload site dan book. |
| `nav` | Array | `[]` | Tautan navigasi (`text` + `url`). |
| `footer` | String | — | Teks footer. |
| `scripts` | Map | `{}` | Tugas shell bernama yang dijalankan dengan `kiw run <tugas>`. |

### `build`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `output` | String | `.krewire/build` | Direktori tempat artefak build ditulis. |
| `base` | String | `/` | Awalan URL dasar untuk tautan dan aset statis. |
| `include` | Array | `["**/*.md"]` | Pola glob konten yang disertakan; daftar kosong menonaktifkan filter. |
| `exclude` | Array | catatan README | Pola glob konten yang dihapus setelah proses penyertaan. |
| `depends_on` | Array | `[]` | Perintah shell yang dijalankan berurutan sebelum build. |
| `outputs` | Map | `{}` | Nama runtime → direktori output, untuk proyek poliglot. |

Kunci tingkat atas `output:` dan `base:` adalah alias lama (legacy); nilai bersarang `build.*` diprioritaskan saat keduanya diatur, dan flag `--output` / `--base` menimpa keduanya. Deklarasikan nilai sekali saja—jangan mengulanginya di kedua tingkatan.

### `dev`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `port` | Int | `8080` | Port server dev lokal; flag `--addr` diprioritaskan dibanding ini. |

### `book`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `mount` | String | — | Sub-path yang ditempati buku dalam output dan ruang URL (misalnya `/docs/`). |
| `toc` | Bool | auto | Memaksa halaman daftar isi akar (table-of-contents) yang dihasilkan aktif atau nonaktif. |

### `i18n`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `default_locale` | String | `en` | Lokal yang disajikan pada root situs. |
| `fallback_locale` | String | `en` | Lokal cadangan yang digunakan saat kunci terjemahan tidak ditemukan. |
| `base_path` | String | `lang` | Direktori yang menyimpan kamus `<locale>/site.json`. |
| `locales` | Array | `[]` | Lokal yang akan di-build (`content/<locale>/docs/` → `.krewire/build/<locale>/docs/`). |

### `ssg`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `description` | String | — | Disuntikkan ke dalam data templat setiap halaman. |
| `layouts` | Array | `[]` | Layout bernama (`name`, `body`, `style`, atau shell `ui`). |
| `components` | Array | `[]` | Komponen bernama (`name`, `body`, `style`). |
| `pages` | Array | `[]` | Halaman output (`path`, `title`, `layout`, `root`, `data`). |
| `assets` | Map | `{}` | Jalur output ke isi berkas. |
| `auto_assets` | Blok | on | Penyuntikan CSS/JS; lihat [Konfigurasi](/id/docs/01-getting-started/02-configuration). |

Proyek dengan direktori `pages/` otomatis di-build sebagai `site` tanpa memerlukan blok ini. Kunci `ssg:` hanya diperlukan untuk situs yang sepenuhnya deklaratif tanpa adanya berkas halaman di disk.

### `wasm`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `entry` | String | `./wasm` | Paket utama yang dikompilasi dengan `GOOS=js GOARCH=wasm`. |
| `name` | String | `runtime` | Nama modul output. |

WebAssembly sisi klien adalah **target pustaka (library target)**, bukan jenis workload: lakukan build dengan `kiw build --target wasm` di dalam proyek `app` atau `site`.

### `worker`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `concurrency` | Int | `1` | Jumlah maksimum worker yang berjalan serentak. |
| `queues` | Array | `[]` | Nama antrean yang dikonsumsi oleh worker ini. |

### `service`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `name` | String | `project.name` | Nama service yang digunakan untuk registrasi dan discovery. |
| `port` | Int | `8080` | Port yang di-bind oleh service. |

`services` menerima daftar dengan struktur yang sama saat satu proyek mendeklarasikan beberapa service sekaligus.

### `gateway`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `host` | String | `""` | Alamat pengikatan (bind address, kosong berarti semua antarmuka). |
| `port` | Int | `8080` | Port yang di-bind oleh gateway (`GATEWAY_ADDR` menimpa keduanya). |
| `routes` | Array | `[]` | Tabel perutean; lihat di bawah. |

Kunci untuk rute:

| Kunci | Tipe | Deskripsi |
| :--- | :--- | :--- |
| `path` | String | Jalur pasti, atau awalan saat `prefix: true`. |
| `method` | String | Filter metode HTTP opsional. |
| `service` | String | Upstream yang di-resolve melalui service registry. |
| `target_url` | String | URL upstream tetap (alternatif untuk `service`). |
| `prefix` | Bool | Memperlakukan `path` sebagai pencocokan awalan (prefix match). |
| `strip_prefix` | Bool | Menghapus `path` dari jalur permintaan upstream. |
| `middleware` | Array | Nama middleware yang didaftarkan pada gateway melalui `WithMiddleware`. |
| `rate_limit` | Int | Permintaan per detik untuk rute ini (diberlakukan). |
| `timeout` | String | Batas waktu per rute, misal `10s` (diberlakukan; jika terlampaui → `504`). |

### `infra`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `provider` | String | `aws` | Penyedia cloud target (`aws`, `gcp`, `kubernetes`). |
| `region` | String | — | Wilayah/region penyedia. |
| `state` | String | — | Lokasi backend status penyimpanan (state). |
| `resources` | Array | `[]` | Sumber daya terdeklarasi yang diterapkan oleh `kiw deploy --target infra`. |

### `runtime`

| Kunci | Tipe | Nilai Default | Deskripsi |
| :--- | :--- | :--- | :--- |
| `namespace` | String | — | Awalan untuk variabel lingkungan Krewire (misal `go` → `GO_*`). |
| `env_file` | String | `.env` | Jalur ke berkas lingkungan yang dimuat saat startup. |
| `port` | Int | — | Port runtime default; port pertama yang dikonsultasikan saat menentukan alamat dengar. |

---

## Lapisan Lingkungan & Dotenv (`.env`)

`kiw` memuat berkas lingkungan dari direktori akar proyek (jalur dari `runtime.env_file`, default `.env`) tanpa menimpa variabel yang telah disetel sebelumnya di lingkungan proses:

```bash
# .env
DATABASE_URL=postgres://user:pass@localhost:5432/db
LOG_LEVEL=debug
```

`KIW_ENV` dan `KIW_DEBUG` dibaca langsung dari lingkungan dan memiliki prioritas lebih tinggi dibandingkan kunci `env:` dan `debug:` dalam `krewire.yaml`.

Port server dev **bukan** sebuah variabel lingkungan. Alamat dengar ditentukan dengan urutan resolusi berikut:

1. Flag `--addr` (`kiw serve --addr :3000`).
2. `runtime.port`, kemudian `dev.port`, lalu `server.port`, disusul `services[].port` pertama.
3. `:8080`.
