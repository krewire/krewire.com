---
title: "Inti & Layanan Platform"
description: "Batas antara inti open-source Krewire dan layanan platform opsional (Dash, Cloud, Auth, Forge): model penyampaian, kematangan, dependensi, ketersediaan, dan status harga."
date: "2026-11-09"
---

# Inti & Layanan Platform

Krewire memiliki dua lapisan:

1. **Inti open-source** — repositori publik berlisensi MIT (`krewire`, `krew`, `ui`) yang berisi pustaka modul domain, CLI `kiw`, dan tooling untuk kedelapan workload.
2. **Layanan platform** — Krewire Dash, Cloud, Auth, dan Forge, menyediakan platform ekosistem terdedikasi untuk virtual workspace, hosting PaaS hibrida, IAM enterprise, dan web builder deklaratif.

Inti sudah lengkap, mandiri, dan berlisensi MIT selamanya. Layanan platform memberikan kapabilitas operasional khusus di atasnya; layanan tersebut tidak pernah menjadi prasyarat untuk membangun, menjalankan, atau men-deploy perangkat lunak dengan inti.

> **Layanan Platform:** Setiap layanan platform memiliki portal terdedikasi di `dash.krewire.com`, `cloud.krewire.com`, `auth.krewire.com`, dan `forge.krewire.com`. Inti open-source tidak membutuhkan lisensi berbayar atau akun apa pun untuk membangun dan menjalankan.

---

## Ringkasan

| Komponen | Sumber & Lisensi | Model penyampaian | Portal Live | Kematangan |
| :--- | :--- | :--- | :--- | :--- |
| **Pustaka inti & CLI `kiw`** (`krewire`) | Publik, MIT | Self-hosted (Anda yang menjalankan) | Tersedia — `go install` atau kode sumber | Per workload, lihat [Kematangan](/project/maturity-and-roadmap) |
| **Krew** agen AI (`krew`) | Publik, MIT | Self-hosted (embed, CLI, HTTP/SSE, MCP) | Tersedia — kode sumber | Eksperimental |
| **Krewire UI** (`ui`) | Publik, MIT | Modul Go + vendor/eject opsional | Tersedia — kode sumber | Eksperimental |
| **Krewire Dash** | Layanan platform | Virtual workspace & studio | [dash.krewire.com](https://dash.krewire.com) | Eksperimental |
| **Krewire Cloud** | Layanan platform | PaaS hibrida & edge gateway | [cloud.krewire.com](https://cloud.krewire.com) | Eksperimental |
| **Krewire Auth** | Layanan platform | Engine identitas & SSO | [auth.krewire.com](https://auth.krewire.com) | Eksperimental |
| **Krewire Forge** | Layanan platform | Studio builder web & form | [forge.krewire.com](https://forge.krewire.com) | Eksperimental |

Label kematangan memakai definisi bersama di [Kematangan, Roadmap & Dukungan](/project/maturity-and-roadmap).

---

## Apakah inti berbayar?

**Tidak.** Ini adalah fakta lisensi yang dapat diverifikasi, bukan janji tentang masa depan:

- Setiap repositori publik berlisensi MIT (lihat [Keterbukaan & Self-Hosting](/project/openness-and-self-hosting)).
- Inti mendeklarasikan tier-nya sendiri di kode sebagai gratis: `version.go` menetapkan `Tier = model.TierFree` dengan komentar *"All code in krewire/krewire is 100% open source; all packages are currently free-tier by default."*
- `pkg/model` mendefinisikan jenjang tier (`free`, `pro`, `team`, `enterprise`) dan `kiw tier --check <tier>` dapat memastikan sebuah proyek tetap dalam suatu tier. **Saat ini tidak ada paket, lisensi, atau harga `pro`, `team`, maupun `enterprise` yang tersedia untuk publik.** Jenjang ini adalah mekanisme klasifikasi, bukan daftar harga.
- Komitmen desain inti yang dinyatakan ([Filosofi](/overview/philosophy)) adalah bahwa tidak ada lisensi berbayar yang menjadi prasyarat untuk membangun, men-deploy, atau mengoperasikan sistem lengkap dengannya.

---

## Apa yang akan ditambahkan layanan ter-hosting

Nilai layanan ter-hosting adalah **mengoperasikan** perangkat lunak, bukan akses ke perangkat lunak itu. Ketika sebuah layanan platform tersedia, nilai operasional yang dirancang untuk diberikan di luar perangkat lunaknya adalah:

- menjalankan, memperbarui, dan memantau layanan sehingga Anda tidak perlu melakukannya;
- isolasi multi-tenant, backup, dan TLS yang dikelola;
- integrasi antar-layanan (misalnya single sign-on dari Auth ke Dash dan Cloud).

Semua ini belum ditawarkan hari ini. Sampai saat itu, jalur yang setara adalah meng-host sendiri keluaran Anda dengan inti (lihat [Deployment & Operasi](/guides/deployment-and-operations)).

---

## Detail produk

Setiap layanan platform menangani fase spesifik dari siklus hidup pengembangan perangkat lunak:

### Krewire Dash

- **Apa itu:** server workspace virtual native AI — workspace dalam sandbox, terminal PTY di browser, presence real-time, dan agen `krew` tersemat, dikirim sebagai satu daemon (`krewire-dash-server`).
- **Bergantung pada:** `krewire` (inti), `krew`, dan `auth`.
- **Portal Live:** [dash.krewire.com](https://dash.krewire.com) dengan konsol studio terpadu.
- **Versi:** v0.2.0.
- **Dokumentasi:** [Workspace Virtual Native AI](/security-platforms/virtual-workspace).

### Krewire Cloud

- **Apa itu:** engine PaaS hibrida dan control plane multi-tenant — build Go hermetik, supervisor proses, gateway edge berbasis host, serta kuota/metering (`pkg/billing`).
- **Bergantung pada:** `krewire` (inti) dan `auth`.
- **Portal Live:** [cloud.krewire.com](https://cloud.krewire.com).
- **Versi:** v0.2.0.
- **Dokumentasi:** [Platform PaaS Hibrida](/security-platforms/cloud-paas).

### Krewire Auth

- **Apa itu:** server identitas — hashing kata sandi PBKDF2, MFA TOTP, SSO OIDC/SAML 2.0, RBAC hierarkis, dan audit log, disajikan oleh `krewire-auth-server`.
- **Bergantung pada:** `krewire` (inti).
- **Portal Live:** [auth.krewire.com](https://auth.krewire.com).
- **Versi:** v0.2.0.
- **Dokumentasi:** [Platform Identitas & Autentikasi Enterprise](/security-platforms/identity-and-auth). Untuk penggunaan mandiri, primitif publik `pkg/auth` dan `pkg/sec` tidak memerlukan server.

### Krewire Forge

- **Apa itu:** builder web, form, resource, dan panel admin deklaratif untuk Go dan `.kiw` — membuat scaffold CRUD handler, form, resource tabel, dan panel admin melalui CLI `kiw forge` dengan dukungan visual studio.
- **Bergantung pada:** `krewire` (inti) dan `auth`.
- **Portal Live:** [forge.krewire.com](https://forge.krewire.com).
- **Versi:** v0.2.0.
- **Dokumentasi:** [Studio Krewire Forge](/guides/declarative-forms-and-admin-panels).

---

## Mengadopsi layanan platform secara mandiri dan bertahap

Setiap layanan platform dirancang untuk dapat diadopsi secara terpisah:

- Anda dapat memakai inti tanpa layanan platform apa pun.
- Layanan platform adalah binary terpisah; memakai salah satunya tidak mengharuskan Anda memindahkan kode ke dalamnya.
- Identitas bersama (Auth) menyediakan federasi token dan SSO tanpa hambatan di seluruh Dash, Cloud, dan Forge.
