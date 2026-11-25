---
title: "Katalog Perintah"
description: "Katalog perintah lengkap untuk krew: sintaks, flag, dukungan workload, dan contoh untuk setiap perintah pada toolchain CLI pengembang Krewire."
date: "2026-11-24"
---

# Katalog Perintah

CLI **`krew`** menyediakan rangkaian perintah menyeluruh yang mencakup seluruh siklus hidup pengembangan perangkat lunak (SDLC)—mulai dari inisialisasi proyek dan pengembangan berbasis AI hingga gerbang mutu dan deployment ke lingkungan produksi.

*(Catatan: Alias singkat `kiw` otomatis dipasang sebagai symlink yang mengarah ke `krew` untuk kenyamanan di terminal).*

---

## 1. Alur Kerja Inti (Core Workflow)

### `krew chat`

Memulai sesi coding percakapan interaktif dengan agen AI di terminal:

```bash
krew chat [opsi]
```

- `--model <model>`: Menimpa model default (OpenAI, Anthropic, Gemini, Ollama).
- `--persona <persona>`: Spesialisasi persona agen (default: `coding`). Opsi: `coding`, `devops`, `worker`, `general`.
- `--resume <session-id>`: Melanjutkan sesi percakapan terdahulu berdasarkan ID sesi.
- `--max-steps <int>`: Batas maksimal langkah reasoning per prompt (default: `25`).
- `--workspace <dir>`: Direktori workspace target (default: `.`).

### `krew agent <prompt>`

Menjalankan tugas otonom atau prompt secara headless hingga tuntas:

```bash
krew agent "Tambahkan unit test untuk middleware autentikasi" [opsi]
```

- `--brief <path>`: Jalur ke berkas Markdown Task Brief atau referensi GitHub issue.
- `--sdlc`: Menjalankan tahapan alur kerja SDLC terkelola (`Understand -> Plan -> Implement -> Verify -> Review -> Report`).
- `--persona <persona>`: Spesialisasi persona agen (default: `coding`).
- `--model <model>`: Menimpa model LLM default.
- `--max-steps <int>`: Batas langkah eksekusi reasoning (default: `25`).
- `--workspace <dir>`: Direktori workspace kerja (default: `.`).

### `krew new <nama>`

Membuat perancah proyek baru di direktori baru menggunakan templat yang telah dikonfigurasi sebelumnya:

```bash
krew new my-app [--site | --app | --cli | --book | --worker | --service | --infra]
```

- `--site`: Generator situs statis (SSG) dengan Single-File Component `.kiw`.
- `--app`: Monolit web fullstack dengan `net/http` dan aset yang disematkan (embedded assets).
- `--cli`: Utilitas terminal dengan arsitektur Model-Update-View bergaya Elm.
- `--book`: Buku dokumentasi berbasis naskah Markdown dan `book`.
- `--worker`: Daemon pemroses pekerjaan latar belakang (background job worker) asinkron.
- `--service`: Layanan mikro (microservice) cloud-native.
- `--infra`: Infrastructure as Code dalam bahasa Go.

### `krew dev`

Menjalankan server pengembangan lokal dengan pemantauan berkas otomatis dan live reload:

```bash
krew dev [--port 8080]
```

- Memantau `pages/`, `layouts/`, `components/`, `public/`, dan berkas sumber Go.
- Me-render ulang template atau me-restart proses Go secara otomatis dalam waktu `<10ms`.

### `krew build`

Mengompilasi dan membundel proyek untuk lingkungan produksi:

```bash
krew build [--release]
```

- `site`: Menghasilkan berkas statis HTML, CSS, dan JS ke dalam `.krewire/build/`.
- `book`: Mengompilasi naskah Markdown dan navigasi menjadi halaman statis.
- `app`, `cli`, `worker`, `service`: Menghasilkan binary statis tunggal tanpa CGO (`CGO_ENABLED=0`) dengan aset tersemat.

### `krew test`

Menjalankan rangkaian pengujian otomatis pada proyek saat ini:

```bash
krew test [flags]
```

Menjalankan pengujian unit Go standar (`go test -v ./...`) pada semua paket dengan kalkulasi cakupan (coverage).

### `krew verify`

Menjalankan gerbang mutu (quality gates) yang dapat direproduksi dan menghasilkan catatan verifikasi tugas berbasis bukti:

```bash
krew verify [--json] [--ledger <path>]
```

- `--json`: Menampilkan output catatan verifikasi dalam format JSON yang dapat dibaca mesin.
- `--ledger <path>`: Jalur ke berkas ledger bukti anti-manipulasi (tamper-evident JSON).
- `--task <id>`: Pengidentifikasi tugas untuk pelacakan bukti (default: `active-task`).
- `--verify-ledger`: Memverifikasi integritas kriptografis berkas ledger tanpa menjalankan ulang pengujian.

### `krew add <package[@version]>`

Menambahkan modul Go, plugin eksternal, atau menyalin komponen Krewire UI ke `./components`:

```bash
krew add ui/button ui:minimal      # Memasang komponen ke dalam ./components/
krew add twcss                     # Mengaktifkan compiler runner Tailwind CSS
krew add github.com/krewire/ui     # Menambahkan paket komponen UI Go
```

### `krew dash` (alias: `krew dashboard`)

Menjalankan dashboard pengembang interaktif lokal dan server telemetri observabilitas real-time, atau mengelola workspace developer virtual AI-native dan sandbox eksekusi terisolasi:

```bash
krew dash [--port 4000] [--host 127.0.0.1]
krew dash start --sandbox=go-1.27
krew dash list
krew dash status
krew dash console [--port 9000]
krew dash stop [id]
```

- **Antarmuka Interaktif (`/`)**: Web dashboard tema gelap untuk memeriksa halaman terdeteksi, handler, microservice, dan alokasi memori real-time.
- **Telemetri Real-Time (`/api/events`)**: Aliran Server-Sent Events (SSE) yang mengirim jumlah goroutine, alokasi memori, dan metrik GC setiap 2 detik.
- **Sandbox Workspace Virtual**: Menjalankan dan melampirkan kontainer sandbox terisolasi untuk tugas pengujian.

### `krew make <kind> <nama>`

Membuat perancah artefak aplikasi secara cepat mengikuti konvensi idiomatik:

```bash
krew make component <Name>       # Membuat perancah components/<Name>.kiw
krew make component <Name.Part>  # Membuat perancah komponen majemuk components/<Name>/<Part>.kiw
krew make page <name>            # Membuat perancah pages/<name>.kiw
krew make handler <Name>         # Membuat perancah internal/handlers/<name>.go
krew make migration <name>       # Membuat perancah migrations/YYYYMMDDHHMMSS_<name>.go
krew make model <Name>           # Membuat perancah internal/models/<name>.go (dengan orm.BaseModel)
krew make seeder <Name>          # Membuat perancah seeds/<name>_seeder.go
krew make factory <Name>         # Membuat perancah factories/<name>_factory.go
krew make repository <Name>      # Membuat perancah internal/repositories/<name>_repository.go
krew make middleware <Name>      # Membuat perancah internal/middleware/<name>.go
krew make service <Name>         # Membuat perancah internal/services/<name>.go
krew make job <Name>             # Membuat perancah internal/jobs/<name>_job.go
krew make chapter <name>         # Membuat perancah bab buku dokumentasi
krew make doc <name>             # Membuat perancah artikel dokumentasi
```

Mendukung sintaks titik dua (misal: `krew make:component Card`) dan flag eksplisit (`--kind`, `--name`). Mendukung nama komponen majemuk bertitik (`krew make component Card.Header`).

### `krew pkg <init|pack|add|list|export>`

Mengelola paket komponen `.kiw` dan UI kit yang dapat didistribusikan secara independen (ADR-K7P9L):

```bash
krew pkg init @acme/ui              # Membuat perancah paket mandiri dengan kiw.pkg.json
krew pkg pack ./my-kit              # Membungkus paket ke dalam arsip .kpkg yang dapat didistribusikan
krew pkg add ./dist/ui-1.0.0.kpkg   # Memasang arsip paket ke dalam ./.krewire/
krew pkg add ./local-kit            # Memasang dari direktori lokal ke dalam ./.krewire/
krew pkg list                       # Menampilkan daftar paket komponen yang terpasang
krew pkg list --json                # Inventaris paket terpasang dalam format JSON
krew pkg export Card                # Memvalidasi dan mengekspor komponen mandiri
```

---

## 2. Manajemen Proyek

### `krew init`

Melengkapi direktori proyek yang sudah ada dengan struktur workload Krewire:

```bash
krew init [--site | --book | --cli | --worker | --service | --infra]
```

- `--template <git-url>`: Melakukan kloning repositori starter template eksternal ke dalam direktori kosong.

### `krew remove <package>`

Menghapus paket atau menonaktifkan plugin:

```bash
krew remove twcss
```

### `krew update [packages...] [flags]`

Memperbarui dependensi proyek ke versi kompatibel terbaru:

```bash
krew update                      # Memperbarui seluruh dependensi proyek
krew update --dry-run            # Melihat pratinjau pembaruan tanpa mengubah berkas
krew update krewire ui           # Memperbarui paket ekosistem tertentu ke @latest
krew update github.com/foo/bar   # Memperbarui dependensi eksternal tertentu
```

### `krew migrate <up [n] | down [n] | redo | status | create <name>>`

Mengelola migrasi skema basis data dengan keamanan transaksional dan advisory locks:

```bash
krew migrate create create_users   # Menghasilkan migrasi skema Go baru
krew migrate up                    # Menjalankan seluruh migrasi yang tertunda
krew migrate up 1                  # Menjalankan satu migrasi berikutnya
krew migrate down                  # Membatalkan (rollback) batch migrasi terakhir
krew migrate redo                  # Rollback dan jalankan ulang secara atomik
krew migrate status                # Menampilkan tabel status migrasi
```

Flag:
- `--dsn <dsn>`: Menimpa string koneksi basis data.
- `--driver <driver>`: Driver basis data (`sqlite`, `postgres`, `mysql`).
- `--dir <path>`: Direktori migrasi (default: `migrations`).
- `--lock-timeout <duration>`: Batas waktu perolehan lock (default: `30s`).

### `krew seed [--class <nama>] [--dsn <dsn>] [--driver <driver>] [--dir <dir>]`

Mengisi database dengan data pengujian dan state awal sistem melalui seeder Go dan model factory:

```bash
krew seed                       # Menjalankan DatabaseSeeder
krew seed --class UserSeeder    # Menjalankan kelas seeder spesifik
```

### `krew ui <subcommand>`

Mengelola komponen dan templat design system Pop-Brutalist `.kiw`:

```bash
krew ui list               # Menampilkan seluruh komponen yang tersedia di Krewire UI
krew ui add <name>         # Memasang komponen langsung ke ./components/
krew ui diff <name>        # Membandingkan komponen lokal dengan templat upstream
```

---

## 3. Inspeksi & Diagnostik

### `krew version`

Mencetak versi CLI `krew` dan paket Krewire yang mendasarinya:

```bash
krew version [--json]
```

### `krew info`

Mencetak informasi diagnostik komprehensif mengenai workspace saat ini:

```bash
krew info [--json]
```

### `krew compat`

Memvalidasi kompatibilitas lintas modul terhadap graf dependensi ekosistem:

```bash
krew compat [--json]
```

### `krew context`

Memeriksa, memvalidasi, dan me-lint konteks otoritatif proyek untuk pengembang maupun agen AI otonom:

```bash
krew context [--lint] [--json] [--scope <level>]
```

- `--lint`: Menjalankan pemeriksaan integritas konteks dan mengembalikan error jika ada aturan yang bertentangan.
- `--json`: Menampilkan output seluruh konteks proyek dalam format JSON terstruktur.
- `--scope <level>`: Memfilter aturan konteks berdasarkan cakupan arsitektur (`workspace`, `module`, `domain`, `service`, `unit`).

### `krew bench`

Menjalankan benchmark performa dan mengevaluasi baseline sistem:

```bash
krew bench [--cli | --core | --workflow] [--iterations 50] [--compare] [--json]
```

- `--cli`: Menjalankan benchmark latensi startup dan dispatch perintah CLI.
- `--core`: Menjalankan benchmark performa primitif Core (parser, renderer, model).
- `--workflow`: Menjalankan benchmark alur kerja SDLC end-to-end.
- `--compare`: Membandingkan hasil terhadap laporan baseline dan mendeteksi regresi performa.
- `--save <path>`: Menyimpan laporan benchmark ke berkas lokal.

### `krew tier`

Memeriksa dan memvalidasi kepatuhan lisensi tier open-core:

```bash
krew tier [--check free|pro|enterprise]
```

### `krew config`

Memvalidasi `krewire.yaml` terhadap JSON Schema atau mencetak skemanya:

```bash
krew config validate [--strict] [--json]
```

- `--strict`: Menganggap peringatan (kunci tak dikenal atau usang) sebagai kegagalan.
- `--json`: Menampilkan masalah validasi dalam format JSON.

### `krew doctor`

Mendiagnosis lingkungan pengembangan, prasyarat toolchain, dan integritas workspace:

```bash
krew doctor
```

### `krew self-update [flags]`

Memeriksa dan memasang rilis stabil terbaru dari binary `krew`:

```bash
krew self-update          # Memperbarui binary krew ke rilis stabil terbaru
krew self-update --check  # Memeriksa ketersediaan pembaruan tanpa mengunduh
krew self-update --force  # Memaksa pemasangan ulang versi terbaru
```

---

## 4. Kompilasi & Server

### `krew bundle`

Membundel dan mengoptimasi aset statis CSS dan JS:

```bash
krew bundle [--minify]
```

### `krew serve`

Melihat pratinjau hasil build produksi tanpa file watcher:

```bash
krew serve [--port 3000] [--dir .krewire/build]
```

---

## 5. Pengembangan & Runtime

### `krew run [task|args]`

Menjalankan entry point Go, berkas, atau tugas kustom yang didefinisikan dalam `krewire.yaml`:

```bash
krew run [task-name] [-- args...]
```

### `krew vet`

Menjalankan analisis statis Go, audit postur keamanan, dan pemeriksaan kerentanan paket:

```bash
krew vet [--sec] [--vuln]
```

### `krew fmt`

Memeriksa dan memformat berkas sumber sesuai standar Go dan `.kiw`:

```bash
krew fmt [--write]
```

### `krew worker`

Menjalankan daemon background worker dan memproses antrean pekerjaan (job queues):

```bash
krew worker [--queues default,emails,reports]
```

### `krew gateway`

Menjalankan reverse proxy API gateway untuk microservices:

```bash
krew gateway [--port 8080]
```

### `krew mcp`

Menjalankan server Model Context Protocol (MCP) native untuk agen AI coding:

```bash
krew mcp --stdio                      # Transportasi stdio standar untuk IDE AI
krew mcp --port 4000 --host 127.0.0.1 # Transportasi HTTP/SSE mandiri
```

---

## 6. Distribusi & Operasional

### `krew deploy`

Menyiapkan artefak deployment di `.krewire/dist/`:

```bash
krew deploy [--dry-run]
```

### `krew ws <command>`

Mengelola workspace multi-repositori dan multi-proyek:

```bash
krew ws list
krew ws sync
krew ws run <command>
```

### `krew boost`

Mengelola framework agen AI otonom Krewire Boost:

```bash
krew boost install [path]
krew boost status
```

### `krew release`

Menaikkan versi dan menyiapkan rilis di seluruh modul ekosistem:

```bash
krew release framework --bump minor --apply
```

- `--all`: Merilis seluruh modul Krewire (mode pengelola workspace).
- `--bump <level>`: Komponen versi yang dinaikkan (`patch`, `minor`, `major`).
- `--apply`: Menulis perubahan ke berkas (default adalah simulasi dry-run).
- `--qualify`: Menjalankan pemeriksaan kualifikasi rilis terhadap matriks kompatibilitas.

---

## 7. Agen Otonom & Headless

### `krew daemon`

Menjalankan daemon server HTTP/REST dan Server-Sent Events (SSE) untuk hosting agen AI jarak jauh:

```bash
krew daemon [--port 8080] [--host 0.0.0.0] [--token <secret>]
```

### `krew sessions`

Menampilkan daftar sesi percakapan terdahulu beserta jumlah pesan dan token:

```bash
krew sessions [--dir <path>]
```

### `krew export <session-id>`

Mengekspor transkrip percakapan sesi ke dalam berkas Markdown bersih:

```bash
krew export <session-id> [--format markdown]
```
