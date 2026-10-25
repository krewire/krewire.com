---
title: "Workload Krewire"
description: "Panduan mendalam tentang 8 jenis workload resmi Krewire: app, cli, site, book, worker, service, gateway, dan infra."
date: "2026-09-29"
---

# Workload Krewire

Krewire menyatukan delapan jenis arsitektur workload di bawah satu antarmuka baris perintah (**`kiw`**) dan satu berkas konfigurasi devtool deklaratif (**`krewire.yaml`**).

Tipe workload dideklarasikan di dalam berkas `krewire.yaml` proyek Anda melalui kolom `project.kind`. `kiw` menggunakan konfigurasi ini untuk menjalankan kompiler, server pengembangan lokal, dan pipeline deployment yang tepat:

```yaml
project:
  name: "my-service"
  kind: "service" # app, cli, site, book, worker, service, gateway, infra
  version: "0.1.0"
```

---

## Matriks 8 Jenis Workload

| Workload | Nama Kind | Kasus Penggunaan Utama | Eksekusi CLI | Paket Kunci |
| :--- | :--- | :--- | :--- | :--- |
| **Monolit Fullstack** | `app` | Aplikasi web, SSR, API JSON, sesi, akses database | `kiw run`, `kiw dev` | `pkg/web`, `pkg/app`, `pkg/ui` |
| **CLI & TUI Terminal** | `cli` | Utilitas baris perintah, TUI pengembang interaktif | `kiw run`, `kiw new my-cli --cli` | `pkg/tui` |
| **Situs Statis (SSG)** | `site` | Landing page, situs pemasaran, blog, portofolio | `kiw build`, `kiw dev` | `pkg/web/ssg`, `pkg/kiw` |
| **Buku Teknis & Dokumentasi** | `book` | Dokumentasi multi-bab, buku panduan, manual software | `kiw build --target book`, `kiw dev` | `pkg/book` |
| **Worker Latar Belakang** | `worker` | Antrean tugas asinkron, jadwal cron, retry, DLQ | `kiw worker` | `pkg/cloud/worker` |
| **Microservice** | `service` | API RPC/HTTP terdistribusi, registry layanan, gateway | `kiw run`, `kiw dashboard` | `pkg/cloud/service` |
| **Infrastruktur Cloud (IaC)** | `infra` | Provisi sumber daya cloud multi-provider (AWS/K8s) | `kiw deploy --target infra` | `pkg/cloud/infra` |
| **Gateway API** | `gateway` | Reverse proxy deklaratif, middleware per rute, load balancing | `kiw gateway`, `kiw run` | `pkg/web/gateway` |

---

## 1. Monolit Fullstack (`app`)

Workload `app` menghasilkan binary statis Go mandiri yang menyematkan aset web statis, template, migrasi database, dan routing HTTP:

### Karakteristik

- **Nol Dependensi Eksternal:** Dibangun di atas `net/http` dan `embed.FS` bawaan Go. Tidak memerlukan server web eksternal (seperti Apache atau Nginx) untuk pengembangan lokal maupun deployment dasar.
- **Engine Template Terpadu:** Server-Side Rendering (SSR) via `html/template` atau kompilasi komponen `.kiw`.
- **API REST & JSON:** Routing ekspresif dengan group, pengikatan parameter (binding), dan format standar RFC 7807 Problem Details.
- **Siap Menjadi Monolit Modular (`KWF-5ZHQV`):** Batasan domain internal yang bersih memungkinkan ekstraksi service di masa depan tanpa penulisan ulang kode.

### Pembuatan & Menjalankan

```bash
kiw new my-app --app
cd my-app
kiw run
```

---

## 2. CLI & TUI Terminal (`cli`)

Workload `cli` dirancang untuk membangun utilitas baris perintah, alat bantu pengembang, dan antarmuka pengguna terminal (TUI) yang interaktif:

### Karakteristik

- **Ditenagai `pkg/tui`:** Rendering TUI reaktif, komponen tabel interaktif, progress bar, dan penanganan input keyboard.
- **Parsing Argumen Cepat:** Penanganan argumen, flag, dan sub-perintah yang intuitif dan cepat.
- **Startup Cepat:** Waktu cold start instan (<2 milidetik) menjamin kenyamanan saat digunakan dalam skrip bash atau otomatisasi CI.

---

## 3. Situs Statis / SSG (`site`)

Workload `site` mengompilasi halaman dan komponen `.kiw` menjadi HTML, CSS, dan aset statis murni yang ramah mesin pencari (SEO) tanpa memerlukan JavaScript di sisi klien:

### Karakteristik

- **Nol JavaScript di Client:** Menghasilkan dokumen web statis yang sangat ringan dengan skor performa Lighthouse 100/100.
- **Komponen `.kiw` yang Kuat:** Menulis komponen file tunggal dengan scoped style dan desugaring komponen seperti JSX.
- **Hot-Reloading Cepat:** Server pengembangan lokal `kiw dev` memantau perubahan berkas dan memuat ulang halaman secara otomatis dalam milidetik.

---

## 4. Buku Dokumentasi (`book`)

Workload `book` mengompilasi direktori manuskrip Markdown multi-bab menjadi portal dokumentasi teknis yang indah dan responsif:

### Karakteristik

- **Struktur Bab Otomatis:** Menghasilkan daftar isi hierarkis (TOC), tautan bab sebelumnya/berikutnya (pager), dan pencarian instan.
- **Dukungan Multibahasa (i18n):** Pemetaan konten terstruktur berdasarkan lokal (`content/en/docs/`, `content/id/docs/`).
- **Tema Terang & Gelap:** Pemilihan tema otomatis yang menyesuaikan preferensi sistem pengguna atau toggle manual.

---

## 5. Worker Latar Belakang (`worker`)

Workload `worker` dirancang untuk memproses tugas-tugas berat di latar belakang secara asinkron:

### Karakteristik

- **Antrean Tugas & Jadwal Cron:** Mendukung penjadwalan periodik dan penanganan antrean tugas dengan jaminan eksekusi.
- **Penanganan Kegagalan yang Andal:** Dilengkapi mekanisme retry dengan exponential backoff dan Dead Letter Queue (DLQ).
- **Driver Fleksibel:** Dapat menggunakan antrean in-memory lokal untuk pengembangan atau Redis/SQS di lingkungan produksi.

---

## 6. Microservices (`service`)

Workload `service` memfasilitasi pembangunan layanan terdistribusi dengan tata kelola jaringan yang tangguh:

### Karakteristik

- **Penemuan Layanan Dinamis (Service Discovery):** Registrasi dan penemuan layanan otomatis tanpa konfigurasi manual yang rumit.
- **Observabilitas Bawaan:** Tracing terdistribusi dan metrik performa terintegrasi menggunakan standar OpenTelemetry.
- **Protokol Standar:** Komunikasi antar-layanan menggunakan HTTP/JSON atau gRPC berkinerja tinggi.

---

## 7. Infrastruktur Cloud (`infra`)

Workload `infra` memungkinkan pendefinisian dan provisi sumber daya cloud secara deklaratif menggunakan bahasa Go:

### Karakteristik

- **Penyedia Agnostik:** Mengelola infrastruktur di penyedia cloud publik terkemuka maupun klaster Kubernetes lokal.
- **Rencana Eksekusi Aman:** Menghasilkan ringkasan rencana perubahan sumber daya sebelum dieksekusi secara nyata (`plan` dan `apply`).
- **Penguncian Status:** Mengamankan file status infrastruktur untuk mencegah modifikasi konkuren yang merusak.

---

## 8. Gateway API (`gateway`)

Workload `gateway` bertindak sebagai pintu gerbang masuk terpadu untuk mengarahkan lalu lintas ke berbagai layanan backend:

### Karakteristik

- **Reverse Proxy Deklaratif:** Konfigurasi rute, upstream backend, dan penyeimbangan beban (load balancing).
- **Middleware Terpusat:** Penegakan autentikasi, rate limiting, logging audit, dan proteksi CORS di tingkat gerbang.
- **Integrasi Service Mesh:** Terhubung langsung dengan registry penemuan layanan dari workload `service`.

---

## Rangkuman

Kedelapan tipe workload ini berbagi satu filosofi desain yang sama: mereka dapat dikomposisikan secara bertahap tanpa perombakan arsitektur besar.

Lanjutkan ke: [**1.4 Panduan Pembaruan →**](/overview/upgrade-guide) atau lompat ke bab [**2. Memulai (Getting Started) →**](/getting-started).
