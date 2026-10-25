---
title: "Pengembangan Agentik"
description: "Lengkapi proyek Krewire Anda dengan agen AI otonom menggunakan Krewire Boost — konstitusi terpadu, skill agen, dan quality gate otomatis."
date: "2026-09-29"
---

# Pengembangan Agentik

Rekayasa perangkat lunak modern bergerak cepat menuju kolaborasi antara pengembang manusia dan agen coding AI otonom. Namun, agen AI tanpa panduan sering kali berhalusinasi tentang arsitektur, melanggar konvensi, atau memasukkan dependensi yang berlebihan (bloat).

Krewire menyelesaikan tantangan ini melalui **Krewire Boost** (`templates/boost` di `github.com/krewire/krewire`)—sebuah framework agentik dan standar konfigurasi tersemat yang membekali agen AI dengan pengetahuan mendalam dan intrinsik tentang ekosistem Krewire, 8 workload-nya, CLI `kiw`, serta quality gate otomatis yang ketat.

---

## 1. Apa itu Krewire Boost?

**Krewire Boost** didistribusikan sebagai modul Go yang tersemat langsung di dalam CLI `kiw`. Hanya dengan satu perintah, Anda dapat melakukan bootstrap sistem agen yang lengkap ke dalam repositori Krewire mana pun:

```bash
kiw boost install
```

### Kemampuan Inti:

1. **Konstitusi Terpadu (`AGENTS.md`):** Mengodekan tabel deteksi 8 workload, matriks perintah lengkap `kiw`, batasan arsitektur (SRP, SoC), serta aturan proyek yang tidak boleh dilanggar (immutable).
2. **Daftar Sub-agen Terspesialisasi:** 15+ persona agen khusus (`scout`, `plan`, `build`, `reviewer`, `tester`, `debugger`, `refactor`, `docs`, `security`, `deploy`).
3. **Perintah Slash Terstruktur:** 16 perintah standar (`/kickoff`, `/spec`, `/plan`, `/review`, `/test`, `/fix`, `/commit`, `/deploy`).
4. **Skill yang Sadar Konteks (Context-Aware Skills):** 20+ skill prosedural yang dapat digunakan kembali (`context-awareness`, `quality-gate`, `wasm-runtime`, `infra-provision`, `worker-queue`).
5. **Bebas Keterikatan Vendor (Zero Vendor Lock-In):** Berfungsi langsung secara native dengan OpenCode, dan mudah diadaptasi untuk Claude Code, Cursor, Windsurf, Copilot, serta Codex.

---

## 2. Instalasi & Scaffolding

Untuk menginstal Boost ke dalam proyek Krewire Anda saat ini, jalankan:

```bash
kiw boost install .
```

*Output terminal:*
```text
2026/09/29 21:00:00 INFO installing Krewire Boost into .
created AGENTS.md
created opencode.json
created .agents/agents/build.md
created .agents/agents/scout.md
created .agents/agents/reviewer.md
created .agents/agents/tester.md
created .agents/commands/kickoff.md
created .agents/commands/spec.md
created .agents/skills/quality-gate/SKILL.md
2026/09/29 21:00:00 INFO Boost installation complete: 42 files written
```

### Flag CLI:

| Flag | Deskripsi |
| :--- | :--- |
| `kiw boost install [path]` | Menginstal berkas template ke direktori target (default ke direktori saat ini). |
| `--force` | Menimpa berkas agen terkelola yang ada tanpa konfirmasi interaktif. |
| `--dry-run` | Menampilkan pratinjau berkas yang akan dibuat tanpa menulisnya ke disk. |

---

## 3. Tata Letak Direktori Sistem Agen

Setelah terinstal, repositori Anda akan memiliki struktur `.agents/` yang terstandarisasi:

```text
my-project/
├── AGENTS.md                   # Konstitusi & aturan Krewire terpadu
├── opencode.json               # Konfigurasi plugin & agen OpenCode
└── .agents/
    ├── agents/                 # Agen otonom spesifik peran
    │   ├── build.md            # Orkestrator implementasi utama
    │   ├── scout.md            # Eksplorasi & pemetaan basis kode
    │   ├── plan.md             # Arsitektur & desain sistem
    │   ├── reviewer.md         # Tinjauan kualitas, keamanan & gaya kode
    │   ├── tester.md           # Otomatisasi pengujian per-workload
    │   ├── debugger.md         # Analisis akar penyebab & triase insiden
    │   └── deploy.md           # CI/CD dan deployment produksi
    │
    ├── commands/               # Perintah slash untuk pengembang
    │   ├── kickoff.md          # /kickoff — memetakan stack dan struktur
    │   ├── spec.md             # /spec — membuat dokumen desain yang dapat diverifikasi
    │   ├── test.md             # /test — menjalankan test suite workload
    │   ├── review.md           # /review — memicu audit sebelum commit
    │   └── commit.md           # /commit — membuat conventional commit
    │
    └── skills/                 # Paket kemampuan prosedural
        ├── context-awareness/  # Kecerdasan situasional 4-lapisan
        ├── quality-gate/       # Verifikasi lint, test, dan SAST
        └── agent-workflow/     # Kontrak handoff & siklus hidup agen
```

---

## 4. Alur Kerja Agentik 6 Langkah

Ketika agen AI beroperasi di dalam proyek Krewire Boost, agen tersebut menjalankan alur kerja enam langkah yang ketat dan deterministik:

```text
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│ 1. Pahami    │ ──> │ 2. Petakan   │ ──> │ 3. Rancang   │
│(Baca & Tanya)│     │  (/kickoff)  │     │(/spec /plan) │
└──────────────┘     └──────────────┘     └──────┬───────┘
                                                 │
┌──────────────┐     ┌──────────────┐     ┌──────▼───────┐
│ 6. Ringkas   │ <── │ 5. Verifikasi│ <── │4. Implementasi
│(Lapor Gate)  │     │  (kiw test)  │     │(agen build)  │
└──────────────┘     └──────────────┘     └──────────────┘
```

### 1. Pahami (Understand)
Agen membaca instruksi pengguna secara menyeluruh. Jika terdeteksi adanya ambiguitas atau batasan yang belum jelas, agen akan berhenti sejenak dan mengajukan pertanyaan terarah daripada menebak atau membuat asumsi tanpa verifikasi.

### 2. Petakan (Map - `/kickoff`)
Pada sesi pertama di sebuah proyek, pengembang atau agen memicu `/kickoff`. Sub-agen `scout` menginventarisasi jenis proyek (`app`, `site`, `cli`, `worker`, dll.), memetakan konfigurasi `krewire.yaml`, mencari titik masuk (entry points), dan mencatat konteks ini untuk sesi-sesi berikutnya.

### 3. Rancang & Susun Spesifikasi (Plan & Spec - `/spec`, `/plan`)
Fitur penting atau refactor besar tidak pernah diimplementasikan secara membabi buta. Agen menyusun draf spesifikasi teknis ringkas di `docs/specs/` yang mendefinisikan:
- Sasaran (goals) dan non-sasaran (non-goals)
- Kontrak API atau perubahan skema
- Kriteria penerimaan yang terukur
- Rencana pengujian dan strategi rollback

### 4. Implementasi (Implement - agen `build`)
Agen `build` mengeksekusi perubahan kode minimal yang mematuhi konvensi. Untuk tugas domain yang khusus, agen ini mendelegasikannya ke sub-agen domain:

| Kebutuhan Khusus | Sub-agen yang Dipanggil |
| :--- | :--- |
| Migrasi database atau kueri ORM | `service` |
| UI komponen `.kiw` ber-scope | `build` |
| Tabel rute dan middleware gateway | `gateway` |
| Definisi infrastruktur cloud | `infra` |
| Antrean latar belakang dan jadwal cron | `worker` |

### 5. Verifikasi (Automated Quality Gate)
Agen mengeksekusi rangkaian perintah verifikasi:
```bash
kiw test && kiw vet
```
Agen dilarang keras melaporkan suatu tugas selesai jika pengujian gagal atau masih terdapat kode yang belum diformat dengan rapi.

### 6. Ringkas (Summarize)
Agen memberikan ringkasan perubahan yang ringkas dan dapat diverifikasi, sitasi berkas (`file:///...`), serta hasil verifikasi quality gate.

---

## 5. Menggunakan Boost dalam Praktik

### Memulai Proyek (Kickoff)

Buka proyek Anda di OpenCode atau lingkungan berbasis AI pilihan Anda dan jalankan perintah:

```text
/kickoff
```

Agen akan memindai `krewire.yaml`, memeriksa dependensi Go, mendeteksi bahwa workload-nya adalah `site`, dan merespons:

```text
✓ Identified Krewire project: "my-first-app" (workload: site)
✓ Build command: kiw build
✓ Dev command: kiw dev (port 8080)
✓ Ready for tasks. Use /plan to design or /test to verify.
```

### Pengembangan Fitur Berbasis Persyaratan

Untuk merancang fitur registrasi pengguna baru:

```text
/plan user-registration
```

Agen akan menghasilkan rencana persyaratan dan implementasi terstruktur yang menguraikan skema request, batasan, penanganan error, dan skenario pengujian sebelum menyentuh kode produksi.

---

## 6. Asisten AI yang Didukung

| Asisten | Tingkat Integrasi | Cara Kerjanya |
| :--- | :--- | :--- |
| **OpenCode** | **Tier 1 (Native Plug-and-Play)** | Membaca `opencode.json` secara otomatis, memuat `.agents/skills`, dan menyediakan perintah native `/`. |
| **Claude Code** | **Tier 2 (Didukung via Referensi)** | Membaca `AGENTS.md` dan mengeksekusi perintah CLI via sesi bash terminal. |
| **Cursor / Windsurf** | **Tier 2 (Impor Aturan)** | Mengacu pada `AGENTS.md` dari `.cursorrules` atau `.windsurfrules`. |
| **GitHub Copilot** | **Tier 2 (Instruksi Workspace)** | Membaca `AGENTS.md` dari root workspace secara otomatis. |

---

## Langkah Berikutnya

Sekarang setelah Anda mengetahui cara mempercepat pengembangan menggunakan agen AI, saatnya mempelajari bahasa templating komponen:

Lanjutkan ke [**2.5 DSL (.kiw) →**](/getting-started/dsl) untuk mempelajari arsitektur komponen Krewire.
