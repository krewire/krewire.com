---
title: "Katalog Perintah"
description: "Katalog perintah lengkap untuk kiw: sintaks, flag, dukungan workload, dan contoh untuk setiap perintah pada CLI pengembang Krewire."
date: "2026-10-08"
---

# Katalog Perintah

CLI **`kiw`** menyediakan rangkaian perintah menyeluruh yang mencakup seluruh siklus hidup pengembangan perangkat lunak (SDLC)—mulai dari inisialisasi proyek hingga deployment ke lingkungan produksi.

---

## 1. Inisialisasi Proyek & Perancah (Scaffolding)

### `krew new <nama>`

Membuat perancah proyek baru di direktori baru menggunakan templat yang telah dikonfigurasi sebelumnya.

```bash
krew new my-app [--site | --app | --cli | --book | --worker | --service | --infra]
```

- `--site`: Generator situs statis (SSG) dengan Single-File Component `.kiw`.
- `--app`: Monolit web fullstack dengan `net/http` dan aset yang disematkan (embedded assets).
- `--cli`: Utilitas terminal dengan arsitektur Model-Update-View bergaya Elm.
- `--book`: Buku dokumentasi berbasis naskah Markdown dan `pkg/book`.
- `--worker`: Daemon pemroses pekerjaan latar belakang (background job worker) asinkron.
- `--service`: Layanan mikro (microservice) cloud-native.
- `--infra`: Infrastructure as Code dalam bahasa Go.

### `krew init`

Melengkapi direktori proyek yang sudah ada dengan struktur workload Krewire:

```bash
krew init [--site | --book | --cli | --worker | --service | --infra]
```

- `--template <git-url>`: Melakukan kloning repositori starter template eksternal ke dalam direktori kosong.

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
krew make chapter <name>         # Membuat naskah bab buku (manuscript)
krew make doc <name>             # Membuat artikel dokumentasi
```

Mendukung sintaks titik dua (misalnya `krew make:component Card` atau `krew make component:Card`) dan flag eksplisit (`--kind`, `--name`). Mendukung notasi komponen majemuk bertitik (`krew make component Card.Header` secara otomatis membuat `components/Card/Header.kiw` untuk pemanggilan komponen majemuk `<Card.Header />`).

### `kiw migrate <up [n] | down [n] | redo | status | create <name>>`

Mengelola migrasi skema database dengan jaminan transaksi aman dan advisory lock:

```bash
kiw migrate create create_users   # Membuat berkas migrasi skema Go baru
kiw migrate up                    # Menerapkan seluruh migrasi yang tertunda
kiw migrate up 1                  # Menerapkan satu migrasi berikutnya
kiw migrate down                  # Mengembalikan batch migrasi terakhir
kiw migrate redo                  # Rollback dan terapkan ulang secara atomik
kiw migrate status                # Menampilkan tabel status migrasi
```

Flag:

- `--dsn <dsn>`: Override string koneksi database.
- `--driver <driver>`: Driver database (`sqlite`, `postgres`, `mysql`).
- `--dir <path>`: Direktori migrasi (default: `migrations`).
- `--lock-timeout <duration>`: Batas waktu perolehan lock (default: `30s`).

### `kiw seed [--class <nama>] [--dsn <dsn>] [--driver <driver>] [--dir <dir>]`
Alias: `kiw db:seed`

Mengisi database dengan data pengujian dan status inisial sistem via seeder Go dan model factory:

```bash
kiw seed                       # Menjalankan DatabaseSeeder
kiw seed --class UserSeeder    # Menjalankan seeder tertentu
kiw db:seed                    # Alias kompatibel dengan Laravel
```

---

## 2. Pengembangan & Eksekusi

### `krew dev`

Menjalankan server pengembangan lokal dengan pemantauan berkas otomatis (file watching) dan live reload:

```bash
krew dev [--port 8080]
```

- Memantau direktori `pages/`, `layouts/`, `components/`, `public/`, serta berkas sumber Go.
- Me-render ulang templat secara otomatis atau me-restart proses server Go dalam waktu `<10ms`.

### `krew run [task|args]`

Menjalankan titik masuk (entry point) Go atau tugas kustom yang didefinisikan dalam `krewire.yaml`:

```bash
krew run [nama-tugas] [-- args...]
```

Contoh:
```bash
krew run                # Menjalankan entry point Go default
krew run build-css      # Menjalankan tugas 'build-css' dari krewire.yaml
krew run -- --verbose   # Meneruskan flag '--verbose' ke aplikasi
```

### `krew serve`

Menjalankan server HTTP lokal untuk meninjau output statis (workload `site` atau `book`) atau mendengarkan port web:

```bash
krew serve [--port 3000] [--dir .krewire/build]
```

### `kiw dash` (alias: `kiw dashboard`)

Meluncurkan dasbor pengembang lokal yang interaktif dan server telemetri observabilitas real-time, atau mengelola virtual developer workspace berbasis AI dan sandbox eksekusi terisolasi:

```bash
kiw dash [--port 4000] [--host 127.0.0.1]
kiw dash start --sandbox=go-1.27
kiw dash list
kiw dash status
kiw dash console [--port 9000]
kiw dash stop [id]
```

- **UI Interaktif (`/`)**: Dasbor web bertema gelap untuk memeriksa halaman yang terdeteksi, handler, layanan mikro terkonfigurasi, dan metrik memori secara langsung.
- **Telemetri Real-Time (`/api/events`)**: Aliran Server-Sent Events (SSE) yang mengirimkan jumlah goroutine, alokasi memori, dan statistik GC setiap 2 detik.
- **Virtual Workspace Sandbox**:
  - `start [--sandbox=<image>]`: Menyiapkan dan melampirkan ke sandbox developer workspace terisolasi.
  - `list`: Menampilkan daftar developer workspace dan sandbox aktif.
  - `status`: Memeriksa konektivitas, kuota, dan kesehatan server Dash.
  - `console`: Membuka atau mencetak URL Web IDE console.
  - `stop [id]`: Menghentikan sandbox yang sedang berjalan.
- **Endpoint REST & MCP**:
  - `GET /api/info`: Bentuk proyek, versi toolchain Go, layanan, dan tabel rute.
  - `GET /api/metrics`: Alokasi memori langsung dan penghitung goroutine.
  - `GET /api/health`: Probe status kesehatan untuk kontainer lokal dan otomasi.
  - `GET /mcp/sse`: Server native Model Context Protocol (MCP) untuk AI coding agent.

### `kiw mcp`

Menjalankan server native Model Context Protocol (MCP) untuk AI coding agent (Antigravity, Cursor, Claude Desktop, OpenCode):

```bash
kiw mcp --stdio                      # Transportasi standar stdio untuk AI IDE
kiw mcp --port 4000 --host 127.0.0.1 # Transportasi mandiri HTTP/SSE
```

- **Resource MCP**:
  - `krewire://project/info`: Metadata proyek, kind yang terdeteksi, versi Go, dan layanan yang terkonfigurasi.
  - `krewire://project/routes`: Semua halaman `.kiw`, artikel Markdown, dan HTTP handler Go yang ditemukan.
  - `krewire://runtime/metrics`: Alokasi memori langsung, memori sistem, goroutine, dan statistik GC.
  - `krewire://runtime/health`: Probe status kesehatan berkala.
- **Tools MCP**:
  - `kiw_inspect_project`: Memeriksa konfigurasi proyek, bentuk arsitektur, dan lingkungan.
  - `kiw_list_routes`: Mengekstrak seluruh pohon perutean (routing tree).
  - `kiw_make_scaffold`: Menghasilkan komponen, halaman, handler, migrasi, model, middleware, dan service.
  - `kiw_get_metrics`: Probe telemetri real-time untuk audit performa berbasis AI.

### `kiw worker`

Menjalankan worker latar belakang asinkron dan pemroses antrean (queue processor):

```bash
kiw worker [--queue memory|file|kv] [--concurrency <N>] [--dlq list|clear]
```

- `--queue`: Backend penyimpanan antrean (`memory`, `file`, atau `kv`).
- `--concurrency`: Jumlah worker pekerjaan konkuren (default diambil dari `krewire.yaml` atau jumlah CPU).
- `--dlq list`: Memeriksa pekerjaan yang gagal dan masuk ke dead-letter queue.
- `--dlq clear`: Mengosongkan dan membersihkan seluruh pekerjaan di dead-letter queue.

### `kiw gateway`

Menjalankan API gateway dan reverse proxy performa tinggi untuk layanan mikro:

```bash
kiw gateway [--port 8080] [--host 0.0.0.0] [--routes "/api/orders=http://localhost:8081,/api/users=http://localhost:8082"]
```

- Meneruskan rute berdasarkan konfigurasi `gateway:` di `krewire.yaml` atau flag eksplisit `--routes`.
- Dilengkapi circuit breaking bawaan, distributed rate limiting, CORS, distributed tracing header, dan penyeimbangan beban round-robin.

---

## 3. Kompilasi & Deployment

### `krew build`

Mengompilasi proyek saat ini menjadi artefak siap rilis sesuai dengan jenis workload-nya:

```bash
krew build [--target site|book|binary|wasm|infra] [--output <path>]
```

- **`site`**: Menghasilkan HTML statis teroptimasi dan CSS terlingkup (scoped CSS) ke dalam direktori `.krewire/build`.
- **`book`**: Mengompilasi naskah Markdown menjadi portal dokumentasi multi-bab.
- **`app`/`cli`/`service`/`worker`**: Mengompilasi binary Go tunggal dengan aset yang disematkan.
- **`wasm`**: Mengompilasi bundle klien `pkg/runtime` menjadi modul `.wasm` beserta lem (glue) JS.
- **`infra`**: Mengompilasi dan memverifikasi rencana topologi infrastruktur cloud.

### `krew deploy`

Memvalidasi proyek, menjalankan quality gate, menyiapkan artefak distribusi ke `.krewire/dist/`, dan memublikasikannya ke lingkungan target:

```bash
krew deploy [--target binary|systemd|gh-pages|infra] [--branch <name>] [--remote <name>] [--dry-run]
```

- `--dry-run`: Meninjau langkah-langkah deployment tanpa memublikasikan atau mengubah status cloud.
- `--target binary`: Mengompilasi dan menyiapkan berkas eksekusi mandiri ke direktori distribusi.
- `--target systemd`: Menghasilkan dan mengonfigurasi berkas unit service Linux systemd (`/etc/systemd/system/<app>.service`) lengkap dengan user kustom, direktori kerja, dan kebijakan auto-restart.
- `--target gh-pages`: Melakukan commit dan push output statis ke GitHub Pages (branch `gh-pages`).
- `--target infra`: Melakukan provisi sumber daya cloud yang dideklarasikan melalui adaptor penyedia (provider).

### `krew bundle`

Membundel seluruh aset CSS dan JS yang memenuhi syarat menjadi `assets/krewire.css` dan `assets/krewire.js`:

```bash
krew bundle [--minify] [--output <dir>]
```

- `--minify` / `-m`: mengaktifkan minifikasi CSS dan JS.
- `--output <dir>` / `-o`: direktori output tujuan (bawaan: diambil dari `krewire.yaml` atau `.krewire/build`).

---

## 4. Kualitas Kode & Pengujian

### `krew test`

Menjalankan unit test dan integration test di seluruh paket Go:

```bash
krew test [--filter <pattern>] [--watch] [--coverage] [--verbose] [--json]
```

Menjalankan `go test ./...` dengan output berwarna standar serta laporan ringkasan pengujian.

- `--filter <pattern>` (alias `--run`): hanya menjalankan test yang cocok dengan pola nama atau berkas.
- `--watch` / `-w`: memantau berkas dan menjalankan ulang test saat ada perubahan.
- `--coverage`: menjalankan dengan coverage (`go test -cover`); `--coverprofile <file>` menulis profil coverage.
- `--verbose` / `-v`, `--json`, `--reporter verbose|dot|json`: pengaturan output.
- `--update` / `-u`: memperbarui snapshot/golden file; `--list`: hanya menampilkan daftar test.
- `--count <n>`: meneruskan `go test -count` (mis. `1` untuk menonaktifkan cache test).

### `krew vet`

Menjalankan analisis kode statis dan audit keamanan opsional:

```bash
krew vet [--sec] [--vuln]
```

Menjalankan penganalisis resmi Go `go vet ./...`, mendeteksi potensi kerusakan kode, kode yang tidak terpakai (dead code), dan tag struct yang tidak valid.

- `--sec`: pemeriksaan postur keamanan — pemindaian secrets, higiene konfigurasi, dan cakupan `.gitignore`.
- `--vuln`: audit kerentanan melalui `govulncheck` dan verifikasi checksum.

### `krew fmt`

Memeriksa atau memformat kode sumber dan berkas komponen `.kiw`:

```bash
krew fmt [--write]
```

- Tanpa `--write`: Menampilkan berkas yang belum terformat dan keluar dengan kode `1` jika pemformatan diperlukan (berguna untuk gerbang CI).
- Dengan `--write`: Memformat seluruh berkas Go dan berkas komponen secara langsung di tempat sesuai dengan konvensi kanonikal.

---

## 5. Ekosistem & Manajemen

### `krew info`

Menampilkan jenis workload yang terdeteksi, konfigurasi aktif, dan diagnostik lingkungan Go:

```bash
krew info
```

*Contoh Output:*
```text
Project:      storefront (v0.1.0)
Kind:         app (Fullstack Monolith)
Config:       /home/reasvyn/workspace/storefront/krewire.yaml
Output:       bin/server
Go Version:   go1.27.1 (linux/amd64)
```

### `krew version`

Mencetak versi CLI `kiw` dan paket Krewire yang mendasarinya:

```bash
krew version
```

### `krew compat`

Memvalidasi kompatibilitas versi di semua deklarasi modul dalam workspace:

```bash
krew compat
```

### `krew tier`

Memeriksa tier open-core proyek dan memvalidasi invarian batas open-core:

```bash
krew tier [--json] [--check <tier>]
```

- `--json`: menampilkan postur tier dalam JSON terstruktur.
- `--check <tier>`: memastikan proyek tidak melebihi tier tertentu (mis. `--check free`).

### `krew doctor`

Mendiagnosa lingkungan developer dan kesiapan proyek:

```bash
krew doctor [--json]
```

Memeriksa prasyarat toolchain, tata letak workspace, dan prasyarat yang harus dipenuhi proyek sebelum dapat di-build. `--json` menghasilkan laporan siap-mesin untuk CI.

### `kiw ws <perintah>`

Mengelola multi-paket Go workspace (`go.work`):

```bash
kiw ws info         # Menampilkan anggota workspace dan modul yang aktif
kiw ws list         # Mendaftar path dari semua paket
kiw ws sync         # Menyinkronkan dependensi go.work dan go.mod
kiw ws exec <cmd>   # Menjalankan perintah di setiap modul workspace
```

### `krew boost install`

Melengkapi repositori proyek apa pun dengan sistem AI agent otonom **Krewire Boost**:

```bash
krew boost install [path]
```

Memasang:
- 21 persona agen AI otonom ke dalam direktori `.agents/`
- Aturan alur kerja SDLC 6-fase ke dalam direktori `.workflow/`
- Definisi slash command (`/plan`, `/build`, `/test`, `/review`, `/deploy`, `/boost`)
- Panduan operasional workspace ke dalam `AGENTS.md`

### `krew add <package[@version]>`

Menambahkan modul Go, dependensi eksternal, atau plugin frontend ekosistem pilihan:

```bash
krew add svelte                     # Pasang kompilator pulau Svelte 5 dan plugin
krew add vue                        # Pasang kompilator pulau Vue 3 dan plugin
krew add twcss                      # Pasang runner kompilator Tailwind CSS
krew add tailwindcss                # Alias untuk twcss
krew add github.com/krewire/krewire@v0.2.0
```

- **Nol JS secara Default**: Krewire secara default tidak memerlukan JavaScript serta menggunakan token CSS mandiri.
- **Teknologi Pilihan (Opt-In)**: Saat menambahkan `twcss`, `kiw` akan menginisialisasi `tailwind.config.js` secara otomatis dan mengaktifkan runner compiler Tailwind (binary CLI mandiri di `bin/tailwindcss`, Bun, PNPM, atau fallback npm/npx).
- **Pulau Reaktif**: Menambahkan `svelte` atau `vue` mengonfigurasi bundler untuk hidrasi pulau progresif (`data-kiw-framework="svelte|vue"`).

### `krew remove <package>`

Menghapus paket atau menonaktifkan plugin pilihan:

```bash
krew remove twcss
```

### `krew update [packages...] [flags]`

Memperbarui dependensi proyek ke versi kompatibel terbaru. Jika dijalankan tanpa argumen, akan memperbarui seluruh dependensi langsung maupun transitif di `go.mod` serta merapikan file modul:

```bash
krew update                      # Perbarui seluruh dependensi di proyek
krew update --dry-run            # Tinjau pembaruan yang tersedia tanpa mengubah file
krew update krewire ui           # Perbarui paket ekosistem ke @latest
krew update github.com/foo/bar   # Perbarui dependensi eksternal tertentu
```

- `--dry-run`: Melihat versi pembaruan yang tersedia tanpa menyentuh disk.
- `--tidy`: Menjalankan `go mod tidy` secara otomatis setelah unduhan selesai (default `true`).
- Otomatis memperbarui dependensi `package.json` melalui `npm update` jika terdapat integrasi plugin frontend.

### `kiw self-update [flags]`

Memeriksa dan memperbarui binary devtool `kiw` ke rilis stabil terbaru:

```bash
kiw self-update          # Perbarui binary kiw ke versi stabil terbaru
kiw self-update --check  # Cek ketersediaan pembaruan tanpa mengunduh
kiw self-update --force  # Paksa instalasi ulang versi terbaru
kiw self-update --dry-run # Simulasikan proses pembaruan
```

- `--check`: Memeriksa GitHub Releases untuk mendeteksi rilis stabil baru.

> **Catatan (2026-11-09):** belum ada GitHub Release yang dipublikasikan untuk `krewire/krewire`, sehingga `kiw self-update` saat ini mengandalkan fallback toolchain Go (Go 1.27.1+).

- `--force`: Memasang ulang atau memperbarui meskipun sudah berada di versi terbaru.
- `--version <ver>`: Memasang versi target semver tertentu.
- Melakukan penggantian biner atomik di Linux, macOS, dan Windows dengan fallback otomatis ke toolchain Go.

### `krew ui <subcommand>`

Kelola komponen design system dan template Pop-Brutalist `.kiw`:

```bash
krew ui list               # Tampilkan daftar semua komponen di Krewire UI
krew ui add <name>         # Pasang komponen langsung ke ./components/
krew ui diff <name>        # Bandingkan komponen lokal dengan template upstream
```

### `kiw pkg <init|pack|add|list|export>`

Kelola paket komponen mandiri `.kiw` dan kit UI yang dapat didistribusikan secara independen (ADR-K7P9L). Memungkinkan pengemasan modular, pengarsipan tar.gz (`.kpkg`) dengan verifikasi integritas SHA-256, pemasangan lokal ke `./.krewire/pkg/`, dan injeksi stylesheet otomatis tanpa dependensi npm:

```bash
kiw pkg init @acme/ui              # Membuat perancah paket mandiri dengan kiw.pkg.json
kiw pkg pack ./my-kit              # Mengemas paket ke dalam arsip distributable .kpkg
kiw pkg add ./dist/ui-1.0.0.kpkg   # Memasang arsip paket ke ./.krewire/pkg/
kiw pkg add ./local-kit            # Memasang dari direktori lokal ke ./.krewire/pkg/
kiw pkg list                       # Menampilkan daftar paket komponen yang terpasang di ./.krewire/pkg
kiw pkg list --json                # Inventaris JSON paket terpasang untuk mesin/CI
kiw pkg export Card                # Memvalidasi dan mengekspor komponen mandiri
```

- **Manifest (`kiw.pkg.json`)**: Mendeklarasikan metadata paket, pemetaan path komponen yang diekspor, dan stylesheet (`styles`).
- **Namespace & Aliasing**: Komponen dari paket terpasang dapat dipanggil menggunakan namespace paket (misalnya `<acme:Button />`) atau diimpor langsung via `<script> import Button from '@acme/ui'; </script>`.
- **Integritas & Keamanan**: `kiw pkg pack` dan `kiw pkg add` menegakkan checksum SHA-256 dan batas `sec.SafePath` untuk mencegah kerentanan ZipSlip dan traversal direktori.
