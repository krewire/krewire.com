---
title: "Dokumentasi & Jalur Pembelajaran"
description: "Selamat datang di Krewire — Pustaka Go Modular untuk Setiap Workload. Panduan belajar komprehensif dari prinsip dasar hingga operasi cloud produksi."
date: "2026-10-24"
---

# Dokumentasi & Jalur Pembelajaran Komprehensif

Selamat datang di **Dokumentasi Krewire**.

**Krewire** adalah ekosistem **SDLC digital menyeluruh** open-source yang disajikan sebagai **Pustaka Modular**: paket Go independen, satu tujuan yang Anda komposisi di balik satu CLI (**`kiw`**) dan satu konfigurasi terpadu (**`krewire.yaml`**) untuk membangun, menguji, dan men-deploy **delapan workload berbeda**:

1. **Monolit Web Fullstack (`app`)**
2. **Antarmuka Terminal & TUI (`cli`)**
3. **Situs Statis Berkecepatan Tinggi (`site`)**
4. **Buku Dokumentasi Terstruktur (`book`)**
5. **Worker Latar Belakang & Antrean Tugas (`worker`)**
6. **Microservices Terdistribusi (`service`)**
7. **Infrastruktur Cloud sebagai Kode (`infra`)**
8. **Gateway API (`gateway`)**

---

## Alur Belajar yang Disarankan (Learning Path)

Untuk menguasai ekosistem Krewire secara efisien—baik Anda pengembang individu maupun tim rekayasa skala enterprise—ikuti jalur pembelajaran 6 tahapan berikut:

```text
[Tahap 1: Model Mental & Arsitektur] ───► [Tahap 2: Onboarding & Mulai Membangun]
                  │                                           │
                  ▼                                           ▼
[Tahap 3: Spesialisasi Workload]     ───► [Tahap 4: Eksplorasi Paket Domain]
                  │                                           │
                  ▼                                           ▼
[Tahap 5: Workflow Harian & CLI kiw] ───► [Tahap 6: Operasi Cloud & Skala Produksi]
```

### Tahap 1: Model Mental & Fondasi Arsitektur
Pahami mengapa Krewire mengadopsi pustaka modular alih-alih framework monolitik, Tiga Pilar (*Secure, Sustainable, Scalable*), dan Single Source of Truth (SSOT).

1. [**Mengapa Krewire?**](/id/docs/overview/why-krewire) — Efisiensi sumber daya Go yang ekstrem (~15MB RAM), cold start sub-milidetik, bebas kelelahan JS/bundler, dan kedaulatan data lokal (UU PDP).
2. [**Filosofi Inti**](/overview/philosophy) — Tiga Pilar (*Secure, Sustainable, Scalable*), stdlib-first, dan piramida dependensi menurun (*downward-only*).
3. [**Matriks Workload**](/overview/krewire-workloads) — Memahami 8 tipe proyek dan bagaimana mereka berevolusi secara progresif tanpa penulisan ulang (*rewrite*).

### Tahap 2: Onboarding & Memulai Proyek
Siapkan toolchain Go lokal, kuasai CLI tunggal `kiw`, struktur direktori standar, dan DSL komponen `.kiw`.

1. [**Instalasi**](/getting-started/installation) — Menyiapkan Go 1.27.1+ dan memasang CLI `kiw`.
2. [**Konfigurasi Terpadu**](/getting-started/configuration) — Mengelola seluruh pengaturan proyek lewat satu berkas `krewire.yaml`.
3. [**Struktur Direktori**](/getting-started/directory-structure) — Tata letak basis kode standar (`cmd/`, `pkg/`, `pages/`, `.krewire/`).
4. [**DSL `.kiw`**](/getting-started/dsl) — Komponen file tunggal dengan YAML frontmatter, scoped `<style>`, `<markdown>`, dan sintaks komponen ala JSX.
5. [**Pengembangan Berbantuan AI (Agentic)**](/getting-started/agentic-development) — Kolaborasi dengan agen AI menggunakan starter template Boost dan disiplin konteks `AGENTS.md`.

### Tahap 3: Spesialisasi Beban Kerja (Workload)
Bangun aplikasi nyata dengan mendalami jenis workload yang sesuai dengan kebutuhan Anda:

1. **Dokumentasi & Konten:**
   - [**Buku Dokumentasi (`book`)**](/workloads/book) — Kompilasi manuskrip markdown super cepat menjadi situs buku dengan `pkg/book`.
   - [**Situs Statis (`site`)**](/workloads/site) — Situs SSG cepat dan ramah SEO yang dibuat dengan Go murni.
2. **Alat Bantu Terminal:**
   - [**CLI Interaktif & TUI (`cli`)**](/workloads/cli) — Membangun antarmuka terminal interaktif dengan `pkg/tui`.
3. **Layanan Web & Fullstack:**
   - [**Monolit Web & API (`app`)**](/workloads/app) — Aplikasi web throughput tinggi dengan `pkg/web` dan `pkg/app`.
   - [**Gateway API (`gateway`)**](/workloads/gateway) — Reverse proxy deklaratif dengan middleware per rute, rate limiting, dan service discovery.
4. **Sistem Terdistribusi & Cloud:**
   - [**Worker Latar Belakang (`worker`)**](/workloads/worker) — Antrean tugas asinkron, cron scheduler, dan dead-letter queue.
   - [**Microservices (`service`)**](/workloads/service) — Layanan tangguh dengan dynamic service discovery dan tracing.
   - [**Infrastruktur sebagai Kode (`infra`)**](/workloads/infra) — Provisi cloud deklaratif dengan rencana eksekusi dan penguncian status.

### Tahap 4: Eksplorasi Paket Domain & Baterai
Manfaatkan baterai modular mandiri Krewire untuk menangani kebutuhan inti aplikasi:

- [**Model, Siklus Hidup & Error (`pkg/model`, `pkg/lifecycle`, `pkg/errs`)**](/packages/model-and-lifecycle) — Jenis workload, supervisi siklus hidup, taksonomi error, dan kontrak versi.
- [**Web & HTTP Routing (`pkg/web`, `pkg/app`)**](/packages/web-and-app) — Multiplexer HTTP, pipeline middleware, dan container Dependency Injection.
- [**Keamanan & Otentikasi (`pkg/sec`, `pkg/auth`)**](/packages/security-and-auth) — Middleware keamanan standar OWASP, hashing PBKDF2, TOTP 2FA, SSO (OIDC/SAML2), dan RBAC hierarkis.
- [**Antarmuka Pengguna (`pkg/ui`, `pkg/tui`)**](/packages/ui-and-tui) — Komponen pakai-ulang, token desain, dan tema gelap/terang.
- [**Penyimpanan & Cache (`pkg/storage`, `pkg/cache`, `pkg/redis`)**](/packages/storage-and-fs) — Object storage multi-driver (memory, disk, S3) dan caching Redis.
- [**Database, Mail & Commerce (`pkg/orm`, `pkg/mail`, `pkg/commerce`)**](/packages) — ORM multi-driver (SQLite, MySQL, Postgres, MongoDB, Elasticsearch), pengiriman email (SMTP, Resend, Sendgrid), dan pembayaran lokal Indonesia (Midtrans QRIS/VA).
- [**Resiliensi & Jaminan Mutu (`pkg/resilience`, `pkg/testing`)**](/packages/resilience-and-testing) — Circuit breaker, retry exponential backoff, dan assertion testing harness.
- [**Orkestrasi AI (`pkg/ai`)**](/packages/ai-and-agents) — Integrasi multi-provider LLM dan manajemen prompt.

### Tahap 5: Workflow Harian & Otomasi CLI `kiw`
Tingkatkan produktivitas pengembang dengan suite pengujian otomatis, pemformatan, dan kustomisasi tugas kompilasi:

1. [**Katalog Perintah CLI**](/cli-reference/commands) — Referensi lengkap perintah `kiw new`, `init`, `dev`, `build`, `run`, `test`, `compat`, dan `release`.
2. [**Skema Konfigurasi**](/cli-reference/configuration) — Spesifikasi parameter lengkap berkas `krewire.yaml`.
3. [**Skrip Kustom & Otomasi**](/cli-reference/scripting) — Menentukan tugas pra-build dan hook otomatis.

### Tahap 6: Operasi Cloud, Produksi & Skala Enterprise
Skalakan aplikasi ke produksi, deploy ke PaaS hybrid mandiri atau terkelola, dan pelajari lisensi Open Core:

1. [**Panduan Deployment & Operasi**](/guides/deployment-and-operations) — Kontainerisasi binary mandiri, rolling deploy tanpa downtime, dan reverse proxy TLS otomatis Let's Encrypt.
2. **Krewire Cloud (`cloud`)** — Deploy ke engine PaaS Go-native mandiri atau SaaS terkelola (`kiw deploy --cloud`).
3. [**Model Open Core & Lisensi**](/guides/open-core-and-licensing) — Jenjang lisensi 4 tingkat (*free*, *pro*, *team*, *enterprise*) dan kepatuhan kedaulatan data UU PDP.
