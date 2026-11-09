---
title: "Dokumentasi & Jalur Pembelajaran"
description: "Selamat datang di Krewire — Pustaka Go Modular untuk Setiap Workload. Panduan belajar komprehensif dari prinsip dasar hingga operasi cloud produksi."
date: "2026-10-24"
---

# Dokumentasi & Jalur Pembelajaran Komprehensif

Selamat datang di **Dokumentasi Krewire**.

> **Misi:** Misi Krewire adalah membuatnya mendekati gratis — baik dalam biaya finansial, beban operasional, maupun beban kognitif — bagi siapa pun untuk membangun, memiliki, dan mengoperasikan produk digital yang aman, berkelanjutan, dan berskala besar. Baca selengkapnya di [Tentang Krewire](/project).

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

### Tahap 4: Eksplorasi Mendalam Pustaka Domain Modular
Manfaatkan baterai modular mandiri Krewire yang dikelompokkan ke dalam empat klasifikasi domain:

1. **Inti & Fondasi:**
   - [**Model Workload & Primitif Domain (`pkg/model`, `pkg/errs`)**](/core-libraries/model) — Jenis workload, metadata proyek, taksonomi error, dan kontrak versi.
   - [**Supervisi Siklus Hidup & Kernel (`pkg/lifecycle`)**](/core-libraries/lifecycle) — Registri modul, pengurutan dependensi, dan graceful shutdown.
   - [**Konfigurasi Hierarkis (`pkg/config`)**](/core-libraries/config) — Pemuatan konfigurasi terstruktur dari flag CLI, env vars, dan file .env.
   - [**Validasi Input Komposabel (`pkg/validation`)**](/core-libraries/validation) — Validasi data komposabel tanpa refleksi untuk form dan API.
   - [**Pola Resiliensi & Circuit Breaker (`pkg/resilience`)**](/core-libraries/resilience) — Circuit breaker, retry backoff berjitter, dan isolasi kegagalan.
   - [**Utilitas Pengujian & Asersi (`pkg/testing`)**](/core-libraries/testing) — Harness asersi pengujian, pembantu HTTP, dan fixture mock.

2. **Framework Web & UI:**
   - [**Web Engine & Perutean HTTP (`pkg/web`)**](/web-ui/web) — Router radix tree berkinerja tinggi, pipeline middleware, dan helper render.
   - [**Wadah Aplikasi & Monolit (`pkg/app`)**](/web-ui/app) — Wadah dependency injection, manajemen siklus hidup aplikasi, dan signal draining.
   - [**Web Runtime & WASM (`pkg/runtime`)**](/web-ui/web-runtime) — Kompilasi Go ke WebAssembly di sisi klien, virtual DOM diffing, dan hidrasi SSR.
   - [**Komponen Web UI & Token Desain (`pkg/ui`)**](/web-ui/ui) — Komponen Single-File (.kiw), token desain CSS semantik, dan tema.
   - [**Katalog Komponen Krewire UI (66 Komponen)**](/ui/components) — Dokumentasi lengkap per komponen, tabel properti, slot, dan contoh interaktif.
   - [**Toolkit Terminal UI (`pkg/tui`)**](/web-ui/tui) — Antarmuka pengguna terminal, arsitektur Elm/Bubble Tea, dan pewarnaan ANSI.

3. **Data, Penyimpanan & Layanan:**
   - [**Database & Multi-Driver ORM (`pkg/orm`)**](/data-services/database-and-orm) — ORM multi-driver (SQLite, MySQL, Postgres, MongoDB, Elasticsearch).
   - [**Penyimpanan Objek Multi-Driver (`pkg/storage`)**](/data-services/storage) — Penyimpanan objek multi-driver (memori, disk lokal, S3 / MinIO).
   - [**Sistem Berkas Virtual & Penulisan Atomik (`pkg/fs`, `pkg/file`)**](/data-services/filesystem) — Batas sistem berkas virtual yang dapat diuji dan penulisan atomik crash-safe.
   - [**Cache Multi-Driver (`pkg/cache`, `pkg/redis`)**](/data-services/caching) — Caching in-memory dan Redis terdistribusi dengan TTL dan pembersihan invalidasi.
   - [**Mailer Transaksional (`pkg/mail`)**](/data-services/mail) — Pengiriman email transaksional multi-driver (SMTP, Resend, SendGrid, Postmark).
   - [**Commerce & Gateway Pembayaran Indonesia (`pkg/commerce`)**](/data-services/commerce-and-payments) — Midtrans QRIS, Virtual Account (BCA, Mandiri, BRI, BNI), dan kartu kredit.

4. **Keamanan & Platform Ekosistem:**
   - [**Kontrol Keamanan & Hardening (`pkg/sec`)**](/security-platforms/security) — Header pertahanan standar OWASP, token CSRF, proteksi SSRF, dan redaksi PII.
   - [**Primitif Autentikasi (`pkg/auth`)**](/security-platforms/auth) — Token JWT kriptografis, verifikasi claims, dan HTTP Basic Auth waktu konstan.
   - [**Platform Identitas & Autentikasi Enterprise (Krewire Auth)**](/security-platforms/identity-and-auth) — Manajemen identitas multi-tenant, PBKDF2, TOTP 2FA, SSO (OIDC/SAML2), dan RBAC.
   - [**Cloud & Sistem Terdistribusi (`pkg/cloud`)**](/security-platforms/cloud-distributed) — Penemuan klaster, koordinasi node, kunci terdistribusi, dan telemetri.
   - [**Platform PaaS Hibrida (Krewire Cloud)**](/security-platforms/cloud-paas) — Mesin PaaS open-core self-hosted dan kontrol plane cloud SaaS terkelola.
   - [**Orkestrasi AI & LLM (`pkg/ai`)**](/security-platforms/ai-and-agents) — Integrasi LLM multi-provider dan engine template prompt.
   - [**Mesin Agen AI Otonom (`krew`)**](/security-platforms/autonomous-agent) — Runtime agen otonom mandiri maupun tersemat.
   - [**Workspace Virtual Native AI (Krewire Dash)**](/security-platforms/virtual-workspace) — Orkestrasi ruang kerja multi-agen dan engine integrasi terpadu.

### Tahap 5: Workflow Harian & Otomasi CLI `kiw`
Tingkatkan produktivitas pengembang dengan suite pengujian otomatis, pemformatan, dan kustomisasi tugas kompilasi:

1. [**Katalog Perintah CLI**](/cli-reference/commands) — Referensi lengkap perintah `kiw new`, `init`, `dev`, `build`, `run`, `test`, `compat`, dan `release`.
2. [**Skema Konfigurasi**](/cli-reference/configuration) — Spesifikasi parameter lengkap berkas `krewire.yaml`.
3. [**Skrip Kustom & Otomasi**](/cli-reference/scripting) — Menentukan tugas pra-build dan hook otomatis.

### Tahap 6: Operasi Cloud, Produksi & Skala Enterprise
Skalakan aplikasi ke produksi, deploy ke PaaS hybrid mandiri atau terkelola, dan pelajari lisensi Open Core:

1. [**Panduan Deployment & Operasi**](/guides/deployment-and-operations) — Kontainerisasi binary mandiri, rolling deploy tanpa downtime, dan reverse proxy TLS otomatis Let's Encrypt.
2. **Krewire Cloud (`cloud`)** — Deploy ke engine PaaS Go-native mandiri atau SaaS terkelola (`kiw deploy --target infra` dengan `pkg/cloud/infra`).
3. [**Model Open Core & Lisensi**](/guides/open-core-and-licensing) — Jenjang lisensi 4 tingkat (*free*, *pro*, *team*, *enterprise*) dan kepatuhan kedaulatan data UU PDP.
