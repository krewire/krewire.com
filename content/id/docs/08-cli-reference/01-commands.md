---
title: "Katalog Perintah"
description: "Katalog perintah lengkap untuk kiw: sintaks, flag, dukungan workload, dan contoh untuk setiap perintah pada CLI pengembang Krewire."
date: "2026-10-08"
---

# Katalog Perintah

CLI **`kiw`** menyediakan rangkaian perintah menyeluruh yang mencakup seluruh siklus hidup pengembangan perangkat lunak (SDLC)—mulai dari inisialisasi proyek hingga deployment ke lingkungan produksi.

---

## 1. Inisialisasi Proyek & Perancah (Scaffolding)

### `kiw new <nama>`

Membuat perancah proyek baru di direktori baru menggunakan templat yang telah dikonfigurasi sebelumnya.

```bash
kiw new my-app [--site | --app | --cli | --book | --worker | --service | --infra]
```

- `--site`: Generator situs statis (SSG) dengan Single-File Component `.kiw`.
- `--app`: Monolit web fullstack dengan `net/http` dan aset yang disematkan (embedded assets).
- `--cli`: Utilitas terminal dengan arsitektur Model-Update-View bergaya Elm.
- `--book`: Buku dokumentasi berbasis naskah Markdown dan `pkg/book`.
- `--worker`: Daemon pemroses pekerjaan latar belakang (background job worker) asinkron.
- `--service`: Layanan mikro (microservice) cloud-native.
- `--infra`: Infrastructure as Code dalam bahasa Go.

### `kiw init`

Melengkapi direktori proyek yang sudah ada dengan struktur workload Krewire:

```bash
kiw init [--site | --book | --cli | --worker | --service | --infra]
```

- `--template <git-url>`: Melakukan kloning repositori starter template eksternal ke dalam direktori kosong.

### `kiw make <kind> <nama>`

Membuat perancah artefak aplikasi secara cepat mengikuti konvensi idiomatik:

```bash
kiw make component <Name>       # Membuat perancah components/<Name>.kiw
kiw make page <name>            # Membuat perancah pages/<name>.kiw
kiw make handler <Name>         # Membuat perancah internal/handlers/<name>.go
kiw make migration <name>       # Membuat perancah migrations/YYYYMMDDHHMMSS_<name>.sql
kiw make model <Name>           # Membuat perancah internal/models/<name>.go
kiw make middleware <Name>      # Membuat perancah internal/middleware/<name>.go
kiw make service <Name>         # Membuat perancah internal/services/<name>.go
kiw make job <Name>             # Membuat perancah internal/jobs/<name>_job.go
```

Mendukung sintaks titik dua (misalnya `kiw make:component Card` atau `kiw make component:Card`) dan flag eksplisit (`--kind`, `--name`).

---

## 2. Pengembangan & Eksekusi

### `kiw dev`

Menjalankan server pengembangan lokal dengan pemantauan berkas otomatis (file watching) dan live reload:

```bash
kiw dev [--port 8080]
```

- Memantau direktori `pages/`, `layouts/`, `components/`, `public/`, serta berkas sumber Go.
- Me-render ulang templat secara otomatis atau me-restart proses server Go dalam waktu `<10ms`.

### `kiw run [task|args]`

Menjalankan titik masuk (entry point) Go atau tugas kustom yang didefinisikan dalam `krewire.yaml`:

```bash
kiw run [nama-tugas] [-- args...]
```

Contoh:
```bash
kiw run                # Menjalankan entry point Go default
kiw run build-css      # Menjalankan tugas 'build-css' dari krewire.yaml
kiw run -- --verbose   # Meneruskan flag '--verbose' ke aplikasi
```

### `kiw serve`

Menjalankan server HTTP lokal untuk meninjau output statis (workload `site` atau `book`) atau mendengarkan port web:

```bash
kiw serve [--port 3000] [--dir .krewire/build]
```

### `kiw dashboard`

Meluncurkan dasbor pengembang lokal yang interaktif dan server telemetri observabilitas real-time:

```bash
kiw dashboard [--port 4000] [--host 127.0.0.1]
```

- **UI Interaktif (`/`)**: Dasbor web bertema gelap untuk memeriksa halaman yang terdeteksi, handler, layanan mikro terkonfigurasi, dan metrik memori secara langsung.
- **Telemetri Real-Time (`/api/events`)**: Aliran Server-Sent Events (SSE) yang mengirimkan jumlah goroutine, alokasi memori, dan statistik GC setiap 2 detik.
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

### `kiw build`

Mengompilasi proyek saat ini menjadi artefak siap rilis sesuai dengan jenis workload-nya:

```bash
kiw build [--target site|book|binary|wasm|infra] [--output <path>]
```

- **`site`**: Menghasilkan HTML statis teroptimasi dan CSS terlingkup (scoped CSS) ke dalam direktori `.krewire/build`.
- **`book`**: Mengompilasi naskah Markdown menjadi portal dokumentasi multi-bab.
- **`app`/`cli`/`service`/`worker`**: Mengompilasi binary Go tunggal dengan aset yang disematkan.
- **`wasm`**: Mengompilasi bundle klien `pkg/runtime` menjadi modul `.wasm` beserta lem (glue) JS.
- **`infra`**: Mengompilasi dan memverifikasi rencana topologi infrastruktur cloud.

### `kiw deploy`

Memvalidasi proyek, menjalankan quality gate, menyiapkan artefak distribusi ke `.krewire/dist/`, dan memublikasikannya ke lingkungan target:

```bash
kiw deploy [--target binary|systemd|gh-pages|infra] [--branch <name>] [--remote <name>] [--dry-run]
```

- `--dry-run`: Meninjau langkah-langkah deployment tanpa memublikasikan atau mengubah status cloud.
- `--target binary`: Mengompilasi dan menyiapkan berkas eksekusi mandiri ke direktori distribusi.
- `--target systemd`: Menghasilkan dan mengonfigurasi berkas unit service Linux systemd (`/etc/systemd/system/<app>.service`) lengkap dengan user kustom, direktori kerja, dan kebijakan auto-restart.
- `--target gh-pages`: Melakukan commit dan push output statis ke GitHub Pages (branch `gh-pages`).
- `--target infra`: Melakukan provisi sumber daya cloud yang dideklarasikan melalui adaptor penyedia (provider).

### `kiw bundle`

Membundel seluruh aset CSS dan JS yang memenuhi syarat menjadi `assets/krewire.css` dan `assets/krewire.js`:

```bash
kiw bundle [--minify] [--output <dir>]
```

- `--minify` / `-m`: mengaktifkan minifikasi CSS dan JS.
- `--output <dir>` / `-o`: direktori output tujuan (bawaan: diambil dari `krewire.yaml` atau `.krewire/build`).

---

## 4. Kualitas Kode & Pengujian

### `kiw test`

Menjalankan unit test dan integration test di seluruh paket Go:

```bash
kiw test [--filter <pattern>] [--watch] [--coverage] [--verbose] [--json]
```

Menjalankan `go test ./...` dengan output berwarna standar serta laporan ringkasan pengujian.

- `--filter <pattern>` (alias `--run`): hanya menjalankan test yang cocok dengan pola nama atau berkas.
- `--watch` / `-w`: memantau berkas dan menjalankan ulang test saat ada perubahan.
- `--coverage`: menjalankan dengan coverage (`go test -cover`); `--coverprofile <file>` menulis profil coverage.
- `--verbose` / `-v`, `--json`, `--reporter verbose|dot|json`: pengaturan output.
- `--update` / `-u`: memperbarui snapshot/golden file; `--list`: hanya menampilkan daftar test.
- `--count <n>`: meneruskan `go test -count` (mis. `1` untuk menonaktifkan cache test).

### `kiw vet`

Menjalankan analisis kode statis dan audit keamanan opsional:

```bash
kiw vet [--sec] [--vuln]
```

Menjalankan penganalisis resmi Go `go vet ./...`, mendeteksi potensi kerusakan kode, kode yang tidak terpakai (dead code), dan tag struct yang tidak valid.

- `--sec`: pemeriksaan postur keamanan — pemindaian secrets, higiene konfigurasi, dan cakupan `.gitignore`.
- `--vuln`: audit kerentanan melalui `govulncheck` dan verifikasi checksum.

### `kiw fmt`

Memeriksa atau memformat kode sumber dan berkas komponen `.kiw`:

```bash
kiw fmt [--write]
```

- Tanpa `--write`: Menampilkan berkas yang belum terformat dan keluar dengan kode `1` jika pemformatan diperlukan (berguna untuk gerbang CI).
- Dengan `--write`: Memformat seluruh berkas Go dan berkas komponen secara langsung di tempat sesuai dengan konvensi kanonikal.

---

## 5. Ekosistem & Manajemen

### `kiw info`

Menampilkan jenis workload yang terdeteksi, konfigurasi aktif, dan diagnostik lingkungan Go:

```bash
kiw info
```

*Contoh Output:*
```text
Project:      storefront (v0.1.0)
Kind:         app (Fullstack Monolith)
Config:       /home/reasvyn/workspace/storefront/krewire.yaml
Output:       bin/server
Go Version:   go1.27.1 (linux/amd64)
```

### `kiw version`

Mencetak versi CLI `kiw` dan paket Krewire yang mendasarinya:

```bash
kiw version
```

### `kiw compat`

Memvalidasi kompatibilitas versi di semua deklarasi modul dalam workspace:

```bash
kiw compat
```

### `kiw tier`

Memeriksa tier open-core proyek dan memvalidasi invarian batas open-core:

```bash
kiw tier [--json] [--check <tier>]
```

- `--json`: menampilkan postur tier dalam JSON terstruktur.
- `--check <tier>`: memastikan proyek tidak melebihi tier tertentu (mis. `--check free`).

### `kiw doctor`

Mendiagnosa lingkungan developer dan kesiapan proyek:

```bash
kiw doctor [--json]
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

### `kiw boost install`

Melengkapi repositori proyek apa pun dengan sistem AI agent **Krewire Boost**:

```bash
kiw boost install [path]
```

Memasang definisi agen otonom, slash command kustom, dan alur kerja SDLC otonom ke dalam direktori `.agents/` dan `AGENTS.md`.

### `kiw add <package[@version]>`

Menambahkan modul Go atau plugin ekosistem pilihan (seperti Tailwind CSS):

```bash
kiw add twcss@latest
kiw add tailwindcss
kiw add github.com/krewire/krewire@v0.1.0
```

- **Nol JS secara Default**: Krewire secara default tidak memerlukan JavaScript serta menggunakan token CSS mandiri.
- **Teknologi Pilihan (Opt-In)**: Saat menambahkan `twcss`, `kiw` akan menginisialisasi `tailwind.config.js` secara otomatis dan mengaktifkan runner compiler Tailwind (binary CLI mandiri di `bin/tailwindcss`, Bun, PNPM, atau fallback npm/npx).

### `kiw remove <package>`

Menghapus paket atau menonaktifkan plugin pilihan:

```bash
kiw remove twcss
```
