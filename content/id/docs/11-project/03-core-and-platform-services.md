---
title: "Inti & Layanan Platform"
description: "Batas antara inti open-source Krewire dan layanan platform opsional (Dash, Cloud, Auth, Forge): model penyampaian, kematangan, dependensi, ketersediaan, dan status harga."
date: "2026-11-09"
---

# Inti & Layanan Platform

Krewire memiliki dua lapisan:

1. **Inti open-source** — repositori publik berlisensi MIT (`krewire`, `krew`, `ui`) yang berisi pustaka, CLI `kiw`, dan tooling untuk kedelapan workload.
2. **Layanan platform opsional** — Krewire Dash, Cloud, Auth, dan Forge, yang dikembangkan di repositori **privat**.

Inti sudah lengkap dengan sendirinya. Layanan platform dimaksudkan untuk menambah kemudahan operasional di atasnya; layanan tersebut tidak pernah menjadi prasyarat untuk membangun, menjalankan, atau men-deploy perangkat lunak dengan inti.

> **Status per 2026-11-09:** tidak satu pun dari keempat layanan platform tersedia untuk publik — tidak ada pendaftaran publik, unduhan, ataupun endpoint ter-hosting. Subdomain `dash.krewire.com`, `cloud.krewire.com`, `auth.krewire.com`, dan `forge.krewire.com` saat ini menyajikan endpoint atau pengalihan masing-masing. Belum ada harga, paket, atau lisensi komersial yang dipublikasikan.

---

## Ringkasan

| Komponen | Sumber | Model penyampaian | Ketersediaan saat ini | Kematangan |
| :--- | :--- | :--- | :--- | :--- |
| **Pustaka inti & CLI `kiw`** (`krewire`) | Publik, MIT | Self-hosted (Anda yang menjalankan) | Tersedia — `go install` atau kode sumber | Per workload, lihat [Kematangan](/project/maturity-and-roadmap) |
| **Krew** agen AI (`krew`) | Publik, MIT | Self-hosted (embed, CLI, HTTP/SSE, MCP) | Tersedia — kode sumber | Eksperimental |
| **Krewire UI** (`ui`) | Publik, MIT | Komponen salin-dan-miliki | Tersedia — kode sumber | Eksperimental |
| **Krewire Dash** | Privat | Rencana: layanan ter-hosting | Belum tersedia untuk publik | Eksperimental |
| **Krewire Cloud** | Privat | Rencana: hibrida — SaaS terkelola dan engine self-hosted | Belum tersedia untuk publik | Eksperimental |
| **Krewire Auth** | Privat | Rencana: layanan identitas ter-hosting dan server mandiri | Belum tersedia untuk publik | Eksperimental |
| **Krewire Forge** | Privat | Rencana: engine self-hosted & visual studio | Belum tersedia untuk publik | Eksperimental |

Label kematangan memakai definisi bersama di [Kematangan, Roadmap & Dukungan](/project/maturity-and-roadmap). Model penyampaian "Rencana" berasal dari README masing-masing repositori privat; itu menggambarkan maksud desain, bukan layanan yang dapat Anda gunakan hari ini.

---

## Apakah inti berbayar?

**Tidak.** Ini adalah fakta lisensi yang dapat diverifikasi, bukan janji tentang masa depan:

- Setiap repositori publik berlisensi MIT (lihat [Keterbukaan & Self-Hosting](/project/openness-and-self-hosting)).
- Inti mendeklarasikan tier-nya sendiri di kode sebagai gratis: [`version.go`](https://github.com/krewire/krewire/blob/main/version.go) menetapkan `Tier = model.TierFree` dengan komentar *"All code in krewire/krewire is 100% open source; all packages are currently free-tier by default."*
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

Deskripsi di bawah merangkum README setiap repositori privat. Fungsionalitasnya tidak dapat diperiksa publik selama kode sumbernya privat.

### Krewire Dash

- **Apa itu:** server workspace virtual native AI — workspace dalam sandbox, terminal PTY di browser, presence real-time, dan agen `krew` tersemat, dikirim sebagai satu daemon (`krewire-dash-server`).
- **Bergantung pada:** `krewire` (inti), `krew`, dan `auth` (versi minimum dideklarasikan di `version.go`-nya).
- **Model penyampaian:** direncanakan sebagai layanan ter-hosting di `dash.krewire.com`. Belum aktif.
- **Versi:** v0.2.0 (git tag di repositori privat).
- **Dokumentasi:** [Workspace Virtual Native AI](/security-platforms/virtual-workspace).

### Krewire Cloud

- **Apa itu:** engine PaaS hibrida dan control plane multi-tenant — build Go hermetik, supervisor proses, gateway edge berbasis host, serta kuota/metering (`pkg/billing`) yang dirancang untuk penagihan SaaS di masa depan.
- **Bergantung pada:** `krewire` (inti) dan `auth`.
- **Model penyampaian:** direncanakan hibrida — layanan multi-tenant terkelola dan engine self-hosted (`krewire-cloud-server`). Belum tersedia untuk publik dalam bentuk apa pun.
- **Versi:** v0.2.0 (git tag di repositori privat).
- **Dokumentasi:** [Platform PaaS Hibrida](/security-platforms/cloud-paas).

### Krewire Auth

- **Apa itu:** server identitas — hashing kata sandi PBKDF2, MFA TOTP, SSO OIDC/SAML 2.0, RBAC hierarkis, dan audit log, disajikan oleh `krewire-auth-server`.
- **Bergantung pada:** `krewire` (inti).
- **Model penyampaian:** direncanakan sebagai layanan identitas ter-hosting (`auth.krewire.com`) dan server mandiri. Belum tersedia untuk publik.
- **Versi:** v0.2.0 (git tag di repositori privat).
- **Dokumentasi:** [Platform Identitas & Autentikasi Enterprise](/security-platforms/identity-and-auth). Untuk penggunaan mandiri, primitif publik `pkg/auth` dan `pkg/sec` tidak memerlukan server.

### Krewire Forge

- **Apa itu:** builder web, form, resource, dan panel admin deklaratif untuk Go dan `.kiw` — membuat scaffold CRUD handler, form, resource tabel, dan panel admin melalui CLI `kiw forge` dengan dukungan visual studio opsional.
- **Bergantung pada:** `krewire` (inti) dan `auth`.
- **Model penyampaian:** direncanakan sebagai tooling self-hosted dan daemon visual studio (`kiw forge studio`). Belum didistribusikan untuk publik.
- **Versi:** `version.go` mendeklarasikan 0.2.0; belum ada tag rilis.
- **Dokumentasi:** belum ada halaman dokumentasi khusus.

---

## Mengadopsi layanan platform secara mandiri dan bertahap

Ketika layanan platform tersedia, masing-masing dirancang untuk diadopsi secara terpisah:

- Anda dapat memakai inti tanpa layanan platform apa pun — saat ini itulah satu-satunya pilihan.
- Layanan platform adalah repositori dan binary terpisah; memakai salah satunya tidak mengharuskan Anda memindahkan kode ke dalamnya.
- Identitas bersama (Auth) adalah dependensi yang dideklarasikan oleh Dash, Cloud, dan Forge; dependensi itu akan didokumentasikan di sini sebelum salah satunya ditawarkan.

---

## Harga dan ketersediaan: kepemilikan dan proses pembaruan

- **Harga saat ini:** belum ada yang dipublikasikan.
- **Pemilik:** Krewire Team (maintainer [`krewire/krewire.com`](https://github.com/krewire/krewire.com)).
- **Proses pembaruan:** informasi ketersediaan atau harga apa pun akan ditambahkan ke halaman ini melalui pull request, diberi tanggal, dan ditinjau sebagai bagian dari proses rilis situs yang di-tag yang dijelaskan di [Kematangan, Roadmap & Dukungan](/project/maturity-and-roadmap). Informasi yang tidak ada di halaman ini tidak boleh dianggap berlaku.
