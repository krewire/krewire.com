---
title: "Open Core & Lisensi"
description: "Tangga open core empat tingkat: free, pro, team, dan enterprise. Aturan arsitektur, batas komersial, dan nol ketergantungan runtime (zero lock-in)."
date: "2026-10-08"
---

# Open Core & Lisensi

Krewire adalah komunitas sumber terbuka (open-source) berbasis di Indonesia, yang beroperasi di bawah model bisnis **Open Core** yang transparan.

Kerangka kerja lisensi kami menjamin kebebasan penuh bagi para pembuat sistem (builders) sekaligus menyediakan fitur operasional, kepatuhan (compliance), dan perusahaan tingkat lanjut melalui tingkatan komersial.

---

## 1. Tangga Empat Tingkat (The Four-Rung Tier Ladder)

Ekosistem ini mengorganisasikan modul ke dalam empat tingkatan yang jelas:

```text
┌──────────────────────────────────────────────┐
│  4. Enterprise (B2B, air-gapped, custom SLA)  │
├──────────────────────────────────────────────┤
│  3. Team (Multi-tenant governance, audit)    │
├──────────────────────────────────────────────┤
│  2. Pro (Telemetri canggih, throughput tinggi)│
├──────────────────────────────────────────────┤
│  1. Free (Pustaka inti, CLI, SDLC lengkap)   │
└──────────────────────────────────────────────┘
```

| Tingkat | Model Lisensi | Harga | Target Audiens |
| :--- | :--- | :--- | :--- |
| **`free`** | Open Source Permisif (MIT) | $0 Selamanya | Pengembang individu, startup, proyek open-source. |
| **`pro`** | Komersial Source-Available | Publik Self-Serve | Tim berkembang yang membutuhkan driver yang diperkuat dan metrik mendalam. |
| **`team`** | Lisensi Komersial | Publik Organisasi | Perusahaan mid-market yang membutuhkan SSO, RBAC, dan kebijakan multi-tenant. |
| **`enterprise`** | Kontrak B2B Nego | Penawaran Kustom | Deployment air-gapped, kepatuhan kustom, dan SLA khusus 24/7. |

---

## 2. Aturan Arsitektur yang Tidak Dapat Dilanggar

Sebagaimana didokumentasikan dalam `internal/docs/open-core.md`, tiga aturan berikut melindungi integritas inti open-source:

### Aturan 1: Nol Ketergantungan Runtime (Zero Runtime Lock-in)
Proyek yang hanya merujuk pada paket `free` harus **dapat di-build, di-deploy, dan dijalankan tanpa lisensi kunci atau panggilan verifikasi cloud (call-homes)**. Tidak pernah ada masa uji coba (trial expiration) atau batasan tersembunyi pada tingkat free.

### Aturan 2: Batasan Pembungkus Aditif (Additive Wrapper Boundaries)
Modul berbayar boleh membungkus, memperkuat, atau mengoptimalkan modul `free`, tetapi modul berbayar **tidak boleh sekali pun menjadi prasyarat untuk modul free**.

### Aturan 3: Subset Hirarki Ketat (Strict Hierarchy Subsets)
Setiap tingkat merupakan superset ketat dari tingkat di bawahnya:
- `pro` mencakup semua fitur `free`.
- `team` mencakup semua fitur `pro` dan `free`.
- `enterprise` dinegosiasikan per organisasi dan tidak pernah dipublikasikan sebagai harga tetap.
