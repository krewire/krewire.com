---
title: "Arsitektur & Disiplin Rekayasa"
description: "Invarian arsitektur inti: dependensi ketat ke arah bawah, Single Source of Truth (SSOT), dan Single Responsibility Principle."
date: "2026-10-24"
---

# Arsitektur & Disiplin Rekayasa

Basis kode Krewire diatur oleh invarian arsitektur yang tidak dapat ditawar, sebagaimana didefinisikan dalam `ADR-6CAB8C`.

## 1. Dependensi Ketat Hanya ke Arah Bawah

Aliran dependensi berjalan ketat ke arah bawah:

```text
templates/boost (Templat agen AI)
       │
       ▼
tools/kiw (CLI Pengembang)
       │
       ▼
pkg/* (Pustaka modular: web, orm, cache, auth, sec, storage...)
       │
       ▼
pkg/model, pkg/errs (Primitif domain & taksonomi error)
```

- `pkg/model` hanya bergantung pada pustaka standar Go dan `pkg/errs`.
- Tidak ada paket pustaka yang boleh **sekali pun** mengimpor `kiw`.
- Proyek pengguna yang dihasilkan bergantung pada pustaka Krewire, tidak pernah pada `kiw`.

## 2. Single Source of Truth (SSOT)

Nilai yang harus disepakati oleh seluruh ekosistem dideklarasikan **satu kali** di dalam paket yang memilikinya:

- **Dilarang:** Menyalin tabel literal atau daftar konstanta ke beberapa berkas berbeda.
- **Dilarang:** Menurunkan ulang logika pengurutan yang sudah memiliki canonical resolver yang ada.
- **Dilarang:** Menduplikasi string versi di markdown ketika nilai tersebut telah dideklarasikan di dalam kode.
- **Wajib:** Satu pemilik per fakta. Ekspor fungsi resolver alih-alih mengimplementasikannya kembali di dalam paket konsumen.

## 3. Invarian Go Murni (`CGO_ENABLED=0`)

Setiap pustaka dan binary workload harus dapat dikompilasi silang (cross-compile) secara bersih ke Linux, macOS (`darwin/arm64`), dan Windows (`windows/amd64`) tanpa memerlukan toolchain compiler C.
