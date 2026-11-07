---
title: "Arsitektur & Disiplin Rekayasa"
description: "Invarian arsitektur inti: Piramida 6 Tingkat, dependensi ketat ke arah bawah, pemisahan peer via Bridge/Gateway, Single Source of Truth (SSOT), dan Single Responsibility Principle."
date: "2026-11-07"
---

# Arsitektur & Disiplin Rekayasa

Basis kode Krewire diatur oleh invarian arsitektur yang tidak dapat ditawar, sebagaimana didefinisikan dalam `ADR-6CAB8C` dan `ADR-E8P9D`.

## 1. Piramida Arsitektur (Dependensi Ketat Hanya ke Arah Bawah)

Seluruh ekosistem Krewire terstruktur secara ketat dalam bentuk **Piramida**:

```text
▲
│                       [ Tier 5: Applications & Portals ]
│               krewire.com · dash (client workspace) · starters · templates
│                                        │
│               [ Tier 4: Enterprise Platforms & Workload Engines ]
│            cloud (PaaS)  ◄───[GATEWAY / BRIDGE]───►  panel (Server Admin)
│                 │                                         │
│                 └───►  auth (SSO / OIDC) ◄────────────────┘
│                                        │
│                     [ Tier 3: Developer Driver & CLI ]
│                                tools/kiw (kiw)
│                                        │
│                 [ Tier 2: Modular Domain Libraries (pkg/*) ]
│                 web · ui · kiw-dsl · orm · cache · resilience · tui
│                                        │
│                    [ Tier 1: Core Foundation & Kernel ]
│               pkg/model · pkg/lifecycle · pkg/errs · pkg/version
│                                        │
│                 [ Tier 0: Go Standard Library & OS Primitives ]
▼
```

- **Dependensi Ketat Menurun**: Aliran dependensi berjalan eksklusif ke arah bawah (`Tier N → Tier < N`). Dependensi ke atas (`Tier N → Tier > N`) dilarang keras dan ditegakkan secara otomatis melalui AST checking CI (`layering_test.go`).
- **Isolasi Kernel Inti**: `pkg/model` (Tier 1) hanya bergantung pada pustaka standar Go dan `pkg/errs`. Layer ini tidak pernah mengimpor paket tingkat tinggi maupun perkakas pengembang.
- **Enkapsulasi CLI**: Tidak ada pustaka atau paket platform yang boleh **sekali pun** mengimpor `kiw` (Tier 3). Proyek pengguna yang dihasilkan bergantung pada pustaka Krewire, bukan pada `kiw`.

## 2. Pemisahan Peer & Pola Bridge/Gateway

Platform dan pustaka yang berada pada tingkat yang sama adalah **rekan sejajar (peers)** dan wajib menjaga independensi lateral:

- **Bebas Ketergantungan Peer (Zero Peer Coupling)**: Platform setara pada Tier 4 (`cloud`, `panel`, `auth`, `dash`, `krew`) dilarang keras mengimpor kode sumber atau internal privat satu sama lain secara langsung.
- **Mediasi Wajib Melalui Bridge & Gateway**: Setiap interaksi antar-platform atau lintas batas domain wajib dimediasi melalui Bridge atau Gateway standar:
  - **API Gateway & Reverse Proxy**: Merutekan lalu lintas HTTP REST dan gRPC dengan otorisasi rute, proteksi TLS, dan rate limiting.
  - **Federated Identity Bridge**: Standar OpenID Connect (OIDC), OAuth2 Bearer Tokens, dan validasi kunci publik JWKS (tanpa berbagi kredensial database langsung).
  - **Event / Message Bus**: Pengiriman pesan asinkron Pub/Sub untuk penyiaran event terdesentralisasi.

## 3. Single Source of Truth (SSOT)

Nilai yang harus disepakati oleh seluruh ekosistem dideklarasikan **satu kali** di dalam paket yang memilikinya:

- **Dilarang:** Menyalin tabel literal atau daftar konstanta ke beberapa berkas berbeda.
- **Dilarang:** Menurunkan ulang logika pengurutan yang sudah memiliki canonical resolver yang ada.
- **Dilarang:** Menduplikasi string versi di markdown ketika nilai tersebut telah dideklarasikan di dalam kode.
- **Wajib:** Satu pemilik per fakta. Ekspor fungsi resolver alih-alih mengimplementasikannya kembali di dalam paket konsumen.

## 4. Invarian Go Murni (`CGO_ENABLED=0`)

Setiap pustaka dan binary workload harus dapat dikompilasi silang (cross-compile) secara bersih ke Linux, macOS (`darwin/arm64`), dan Windows (`windows/amd64`) tanpa memerlukan toolchain compiler C.
