---
title: "Kematangan, Roadmap & Dukungan"
description: "Label kematangan yang konsisten untuk workload dan produk Krewire beserta buktinya, roadmap yang memisahkan pekerjaan yang sudah rilis, sedang berjalan, diusulkan, dan ditunda, ekspektasi pemeliharaan dan dukungan yang realistis, serta tempat melaporkan bug dan masalah keamanan."
date: "2026-11-09"
---

# Kematangan, Roadmap & Dukungan

Halaman ini menetapkan ekspektasi yang realistis tentang apa yang sudah siap, apa yang masih eksperimental, apa yang direncanakan, dan dukungan seperti apa yang tersedia. Status terakhir diverifikasi pada **2026-11-09** terhadap repositori.

> **Status seluruh proyek:** Krewire masih **pra-1.0**. Tag terbaru monorepo inti adalah `v0.2.0` (2026-11-05), dan belum ada GitHub Release dengan binary unduhan yang dipublikasikan. Berdasarkan [spesifikasi rilis dan kompatibilitas](https://github.com/krewire/krewire/blob/main/docs/foundation/release-and-compatibility.md) proyek, rilis `0.y.0` dapat memuat perubahan yang tidak kompatibel, yang didokumentasikan di [Panduan Upgrade](/overview/upgrade-guide).

---

## Label kematangan

Satu set label dipakai untuk workload **dan** produk. Tiga definisi pertama diambil dari [matriks kematangan workload](https://github.com/krewire/krewire/blob/main/docs/workloads/maturity-matrix.md) di repositori inti.

| Label | Definisi |
| :--- | :--- |
| **Stabil (Stable)** | Dicakup pengujian otomatis (`go test ./...`), berjalan di workload produksi (misalnya krewire.com atau CLI `kiw` itu sendiri), Go murni (`CGO_ENABLED=0`), dengan dokumentasi dan scaffolding starter. |
| **Layak Pakai (Beta)** | Pipeline dan driver inti berfungsi serta dicakup unit test; cocok untuk penggunaan single-node, self-hosted, atau produksi awal; ada keterbatasan yang diketahui. |
| **Eksperimental (Alpha)** | Antarmuka inti dan prototipe sudah ada; verifikasi otomatis dan produksi yang lebih luas masih berjalan. Komponen tanpa bukti produksi publik tidak diberi label di atas ini. |
| **Direncanakan** | Dijelaskan sebagai niat (roadmap, README) tetapi belum diimplementasikan dalam versi yang di-tag. |

**Penting:** "Stabil" menggambarkan kesiapan implementasi. Ini **bukan** jaminan kompatibilitas API — sampai `v1.0.0`, versi minor masih dapat mengubah API publik.

---

## Kematangan workload

| Workload | Label | Bukti | Catatan verifikasi di situs ini |
| :--- | :--- | :--- | :--- |
| `site` | Stabil | Situs ini dibangun dengan `kiw build` ([repositori](https://github.com/krewire/krewire.com)). | Terverifikasi. |
| `book` | Stabil | Dokumentasi ini (Inggris dan Indonesia) dirender oleh `pkg/book`. | Terverifikasi. |
| `cli` | Stabil | CLI `kiw` sendiri adalah workload `cli` ([kode sumber](https://github.com/krewire/krewire/tree/main/tools/kiw)). | Terverifikasi. |
| `app` | Stabil | Matriks upstream mengutip kompilasi SSR + WebAssembly dan pengujiannya. | Bukti upstream tidak mengutip deployment produksi publik. |
| `gateway` | Stabil | Matriks upstream mengutip gateway `api.krewire.com`. | Gateway merespons secara publik, tetapi upstream `/auth`, `/cloud`, dan `/dash` mengembalikan HTTP 502 pada 2026-11-09. |
| `worker` | Layak Pakai (Beta) | Pemrosesan antrean dan graceful shutdown diuji di `tools/kiw`. | Adapter broker terdistribusi tercantum sebagai rencana di upstream. |
| `service` | Layak Pakai (Beta) | Runtime REST dengan tracing, circuit breaker, rate limiting. | Pembuatan gRPC memerlukan toolchain eksternal `protoc`. |
| `infra` | Layak Pakai (Beta) | Perencanaan resource (`kiw build --plan`) dan rekonsiliasi state. | Driver provider menargetkan Docker lokal dan systemd Linux. |

## Kematangan produk

| Komponen | Label | Bukti |
| :--- | :--- | :--- |
| CLI `kiw` & pustaka inti | Per workload (di atas) | Kode sumber publik, tag `v0.2.0`. |
| Krew (`krew`) | Eksperimental | Kode sumber publik, tag `v0.2.0`; belum ada deployment produksi publik yang dikutip. |
| Krewire UI (`ui`) | Eksperimental | Kode sumber publik, belum ada tag; dipakai oleh komponen situs ini. |
| Workspace (Dash), Cloud, Auth | Eksperimental | Kode sumber privat, tag `v0.2.0`; belum tersedia untuk publik. |
| Panel | Eksperimental | Kode sumber privat; `version.go` mendeklarasikan 0.2.0, belum ada tag; belum tersedia untuk publik. |

Model penyampaian dan ketersediaan ada di [Inti & Layanan Platform](/project/core-and-platform-services).

---

## Roadmap

Definisi status:

- **Sudah rilis** — ada dalam versi yang di-tag, atau sudah di-merge ke `main` (ditandai "belum dirilis" sampai di-tag).
- **Sedang berjalan** — ada issue atau pull request terbuka di repositori publik.
- **Diusulkan** — dipublikasikan sebagai niat di README atau dokumen repositori; belum dijadwalkan.
- **Ditunda** — secara eksplisit ditangguhkan.

### Sudah rilis

- `v0.2.0` inti (2026-11-05): CLI `kiw` dengan scaffolding untuk kedelapan workload, pustaka `pkg/*`, pipeline `site`/`book` yang membangun situs ini, serta tooling rilis/kompatibilitas (`kiw release`, `kiw compat`).
- Belum dirilis di `main` `krewire/krewire` setelah `v0.2.0`: `kiw self-update`/`kiw update`, spesifikasi rilis dengan pengujian integrasi workload dan dokumentasi interoperabilitas, serta sub-bab bertingkat di `pkg/book` ([commit](https://github.com/krewire/krewire/compare/v0.2.0...main)).

### Sedang berjalan

| Butir | Pelacakan |
| :--- | :--- |
| Situs: misi, audiens, keterbukaan, batas inti/layanan, perjalanan komunitas, laporan dampak, kematangan & roadmap (bab ini) | [Issue `krewire/krewire.com` #15–#21](https://github.com/krewire/krewire.com/issues) |
| Kontrak sistem UI dan batas integrasinya dengan `.kiw` | [`krewire/ui` #1](https://github.com/krewire/ui/issues/1) |
| Portal komunitas: jalur dari penemuan ke proyek pertama; informasi sponsorship dan dukungan | [`krewire/krewire.github.io` #8, #9](https://github.com/krewire/krewire.github.io/issues) |

### Diusulkan

- **Celah implementasi inti** (dari [matriks kematangan](https://github.com/krewire/krewire/blob/main/docs/workloads/maturity-matrix.md)): build tanpa konfigurasi yang berhasil untuk kedelapan starter (P0); koordinasi worker terdistribusi (P1); endpoint Prometheus standar untuk `service` dan `gateway` (P1).
- **Krew** ([roadmap README](https://github.com/krewire/krew#readme)): memori hierarkis, tracing langkah OpenTelemetry, dan identitas agen (dinyatakan untuk 2027); mesh antar-agen dan LTS `v1.0.0` (dinyatakan untuk 2028).
- **Workspace (Dash)** (roadmap README privat): isolasi MicroVM dan snapshot (dinyatakan untuk 2027); pengeditan kolaboratif dan LTS `v1.0.0` (dinyatakan untuk 2028).

Tahun pada butir yang diusulkan adalah target yang dinyatakan di README tersebut, bukan komitmen dari halaman ini.

### Ditunda

Belum ada butir yang tercatat sebagai ditunda.

### Celah yang diketahui (belum masuk roadmap)

- Belum ada GitHub Release dengan binary `kiw` siap pakai; installer beralih ke `go install` (Go 1.27.1+).
- Belum ada `SECURITY.md`, kode etik, `CHANGELOG`, atau hasil uji CI publik untuk repositori inti.

---

## Ekspektasi pemeliharaan dan dukungan

- **Siapa yang merawat:** Krewire Team, berdasarkan upaya terbaik (best-effort).
- **Tanpa dukungan yang dijamin:** tidak ada SLA, tidak ada jaminan waktu respons, dan tidak ada layanan dukungan berbayar.
- **Belum ada rilis dukungan jangka panjang (LTS).** Hanya versi terbaru di `main` yang dirawat; perbaikan tidak di-backport ke tag lama.
- **Kompatibilitas:** pra-1.0, rilis minor dapat mengubah API; perubahan yang tidak kompatibel diharapkan disertai catatan migrasi di [Panduan Upgrade](/overview/upgrade-guide). Kompatibilitas modul diperiksa dengan `kiw compat`.
- **Perbaikan keamanan:** ditangani saat dilaporkan; belum ada linimasa pengungkapan formal atau kebijakan versi yang didukung.

---

## Tempat melapor

| Apa | Di mana |
| :--- | :--- |
| Bug dan permintaan fitur | GitHub Issues di repositori publik terkait: [`krewire`](https://github.com/krewire/krewire/issues), [`krew`](https://github.com/krewire/krew/issues), [`ui`](https://github.com/krewire/ui/issues), [`krewire.com`](https://github.com/krewire/krewire.com/issues) |
| Pertanyaan dan ide | [GitHub Discussions di `krewire/krewire.com`](https://github.com/krewire/krewire.com/discussions) |
| Riwayat rilis | [Tag `krewire/krewire`](https://github.com/krewire/krewire/tags) — catatan rilis tertulis belum dipublikasikan |
| Standar kontribusi | [Kontribusi & Komunitas](/project/contributing) |

### Melaporkan masalah keamanan

Mohon **jangan** membuka issue publik untuk dugaan kerentanan. Sampai kebijakan keamanan khusus tersedia, kirim email ke [hello@krewire.com](mailto:hello@krewire.com) dengan subjek `SECURITY` dan sertakan repositori yang terdampak, versi atau commit, serta langkah reproduksinya.

Per 2026-11-09, belum ada repositori yang memiliki `SECURITY.md`, dan pelaporan kerentanan privat GitHub belum diaktifkan di repositori publik. Bagian ini akan diperbarui ketika salah satunya berubah.

---

## Cara halaman ini dijaga tetap mutakhir

- **Pemilik:** Krewire Team (maintainer [`krewire/krewire.com`](https://github.com/krewire/krewire.com)).
- **Kapan:** ditinjau sebagai bagian dari setiap rilis situs ini yang di-tag (lihat "Release checklist" di `AGENTS.md` repositori ini), dan setiap kali kematangan workload, ketersediaan produk, atau butir roadmap berubah.
- **Bagaimana:** verifikasi ulang tag (`git tag`), visibilitas repositori, matriks kematangan upstream, issue terbuka, dan endpoint aktif; perbarui tabel dan tanggal "terakhir diverifikasi" di bagian atas halaman ini.
