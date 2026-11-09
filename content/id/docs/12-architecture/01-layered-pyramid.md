---
title: "Piramida Berlapis & Invarian"
description: "Invarian arsitektur inti: piramida Layer 0–5, dependensi ketat ke arah bawah, pemisahan peer melalui Bridge/Gateway, Single Source of Truth (SSOT), dan aturan Go murni."
date: "2026-11-09"
---

# Piramida Berlapis & Invarian

<nav class="track-crumbs" aria-label="Jalur Pendalaman Arsitektur">
<span class="track-crumbs-label">Jalur pendalaman</span>
<ol>
<li><a href="/architecture">Ringkasan</a></li>
<li aria-current="step"><span>1 · Piramida</span></li>
<li><a href="/architecture/ecosystem-dag">2 · DAG ekosistem</a></li>
<li><a href="/architecture/scope-levels">3 · Level cakupan</a></li>
<li><a href="/architecture/data-flow">4 · Alur data</a></li>
</ol>
</nav>

Basis kode Krewire diatur oleh invarian arsitektur yang tidak dapat ditawar di seluruh paket dan workload. Invarian ini ditegakkan oleh tes, bukan sekadar konvensi.

## 1. Piramida Ekosistem (Dependensi Ketat ke Arah Bawah)

Dilihat dari ekosistem secara keseluruhan, setiap repositori berada pada satu lapisan piramida (sumber: `krewire/docs/foundation/ecosystem-dag.md`):

```text
Layer 5: Applications & Portals            krewire.com · krewire.github.io
   ↓
Layer 4: Platform Engines & Products       Dash (workspace) ∥ Panel (server admin)
                                           Cloud (PaaS) ∥ Auth (IAM / SSO)   ← horizontal capabilities
   ↓
Layer 3: Developer Driver & CLI            tools/kiw (kiw)
   ↓
Layer 2: UI Foundation & Domain Libraries  kui (.kiw components) · pkg/* (web, orm, sec, …)
   ↓
Layer 1: Core Foundation & Kernel          pkg/model · pkg/errs · pkg/version · pkg/lifecycle
   ↓
Layer 0: Go Standard Library & OS
```

- **Ketat hanya ke arah bawah:** sebuah lapisan hanya boleh mengimpor lapisan di bawahnya (`Layer N → Layer < N`). Impor ke atas dilarang dan terdeteksi oleh `layering_test.go`.
- **Isolasi kernel:** `pkg/model` hanya bergantung pada pustaka standar dan `pkg/errs`.
- **Enkapsulasi CLI:** tidak ada pustaka atau paket platform yang boleh mengimpor `kiw` (`tools/kiw`). Proyek pengguna yang dihasilkan bergantung pada pustaka Krewire, tidak pernah pada CLI.

## 2. Di Dalam Repositori Inti

Jika diperbesar ke `github.com/krewire/krewire`, gagasan yang sama diterapkan sebagai tangga paket yang berperingkat (sumber: `krewire/docs/foundation/architecture.md`):

| Layer | Isi | Boleh mengimpor |
| :--- | :--- | :--- |
| 5 | Blueprint scaffolding (`templates/*`) | Layer 0–4 |
| 4 | Perkakas pengembang (`tools/kiw`) | Layer 0–3 |
| 3 | Pustaka domain (`pkg/web`, `pkg/sec`, `pkg/orm`, `pkg/ui`, …) | Layer 0–2 dan paket domain sejajar, secara asiklik |
| 2 | Eksekusi & supervisi (`pkg/lifecycle`, `pkg/env`, `pkg/log`) | Layer 0–1 |
| 1 | Inti deklaratif (`pkg/errs`, `pkg/model`, `pkg/version`) | Pustaka standar (`model`, `version` boleh mengimpor `errs`) |
| 0 | Pustaka standar Go | — |

Enam invarian menjaga tangga ini tetap konsisten:

| # | Invarian | Ditegakkan oleh |
| :--- | :--- | :--- |
| 1 | **Arah hanya ke bawah** — setiap sisi impor mengarah ke lapisan yang lebih rendah, atau tetap asiklik di dalam Layer 3. | `TestFoundationLayering`, `TestUIDoesNotImportWeb` di `layering_test.go` |
| 2 | **Lantai dasar** — `pkg/errs` hanya mengimpor pustaka standar; `model` dan `version` hanya menambahkan `errs`. | `pkg/errs/dependency_test.go`, `pkg/model/dependency_test.go` |
| 3 | **Tanpa nama ke atas** — paket fondasi tidak berisi string literal yang menyebut paket di atasnya. | `TestFoundationNamesNoHigherPackage` |
| 4 | **Inti hermetis** — `pkg/model` tidak memiliki impor pihak ketiga maupun impor Krewire lainnya. | `TestModel_StdlibBoundary` |
| 5 | **Isolasi perkakas** — tidak ada apa pun di bawah `pkg/` atau `templates/` yang mengimpor `tools/kiw`. | `TestPackagesNeverImportKiw` |
| 6 | **Baterai opt-in** — paket khusus workload (`cloud/worker`, `web/ssg`, …) tidak dapat diimpor oleh jenis lain kecuali di-opt-in. | `ViolatesOptIn` di `pkg/model/project.go` |

## 3. Pemisahan Peer & Pola Bridge/Gateway

Komponen pada lapisan yang sama adalah **peer** (rekan sejajar) dan tetap independen secara lateral:

- **Tanpa keterikatan peer:** produk paralel (`dash` dan `panel`) tidak pernah mengimpor paket Go atau source tree privat satu sama lain.
- **Mediasi bridge & gateway:** komunikasi antar-peer, atau dengan kapabilitas horizontal, melalui antarmuka standar:
  - **API gateway / reverse proxy** untuk lalu lintas HTTP dan gRPC dengan otorisasi per rute dan rate limiting.
  - **Identitas terfederasi** — OpenID Connect, bearer token OAuth 2.0, dan validasi JWKS publik, tanpa berbagi kredensial database.
  - **Event bus / stream SSE** untuk notifikasi asinkron.
- **Tanpa keterikatan wajib:** `cloud` dan `auth` versi hosted tidak pernah menjadi prasyarat; `dash` dan `panel` dapat berjalan self-hosted tanpa keduanya.

Halaman [DAG Ekosistem](/architecture/ecosystem-dag) mencantumkan sisi modul yang benar-benar ada di antara repositori saat ini.

## 4. Single Source of Truth (SSOT)

Nilai yang harus disepakati oleh seluruh ekosistem dideklarasikan **satu kali**, di dalam paket yang memilikinya:

- **Dilarang:** menyalin tabel literal atau daftar konstanta ke beberapa berkas.
- **Dilarang:** menurunkan ulang logika pengurutan yang sudah memiliki resolver kanonikal.
- **Dilarang:** menduplikasi string versi di Markdown ketika nilainya sudah dideklarasikan di kode.
- **Wajib:** satu pemilik per fakta; ekspor resolver-nya alih-alih mengimplementasikannya ulang di sisi konsumen.

## 5. Go Murni (`CGO_ENABLED=0`)

Setiap pustaka dan binary workload harus dapat dikompilasi silang (cross-compile) ke Linux, macOS (`darwin/arm64`), dan Windows (`windows/amd64`) tanpa toolchain C. CI memeriksanya dengan `CGO_ENABLED=0 go build ./...`. Saat aplikasi membutuhkan driver database, pilih driver Go murni — [tutorial fullstack](/getting-started/fullstack-app-tutorial) menggunakan `modernc.org/sqlite` tepat karena alasan ini.

---

**Berikutnya:** [**12.2 DAG Ekosistem →**](/architecture/ecosystem-dag)
