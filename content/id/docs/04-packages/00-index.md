---
title: "Paket & Pustaka"
description: "Katalog referensi pustaka domain modular Krewire: batas dependensi menurun, invarian SSOT, dan pola integrasi."
date: "2026-10-24"
---

# Paket & Pustaka

Krewire dihadirkan sebagai **Pustaka Modular**: paket-paket Go independen dengan tujuan tunggal yang Anda komposisikan, bukan sebuah framework monolitik kaku yang harus Anda warisi.

Setiap paket di dalam katalog `krewire/pkg` mematuhi batasan arsitektur yang ketat:
- **Dependensi Menurun Secara Ketat (Strict Downward Dependencies):** Primitif domain dan aturan bisnis inti berada di `pkg/model` dan `pkg/errs`. Pustaka tingkat tinggi tidak boleh menyuntikkan dependensi ke atas.
- **Satu Sumber Kebenaran (SSOT - Single Source of Truth):** Fakta, konstanta, dan aturan validasi dimiliki tepat oleh satu paket dan hanya diimpor oleh pihak lain.
- **Pustaka Standar Pertama:** Nol dependensi pihak ketiga yang tidak perlu. Jika pustaka standar Go menyediakan kapabilitas (`net/http`, `context`, `log/slog`, `embed`), Krewire memperluasnya secara harmonis.

---

## Katalog Paket Domain

| Paket | Fungsi & Tanggung Jawab Utama | Lapisan |
| :--- | :--- | :--- |
| **`pkg/model`** | Model domain, definisi workload, aturan entitas, tipe terstruktur | Fondasi Domain |
| **`pkg/errs`** | Taksonomi error standar, kode keluar POSIX, RFC 7807 | Error & Status |
| **`pkg/lifecycle`** | Manajemen siklus hidup proses, graceful shutdown, healthcheck | Mekanika Runtime |
| **`pkg/web`** | HTTP router berkinerja tinggi, pipeline middleware, context binding | Lapisan Web |
| **`pkg/app`** | Wadah monolit web, dependency injection, integrasi SSR | Web Monolith |
| **`pkg/sec`** | Header keamanan OWASP, validasi input, sanitasi, proteksi CSRF | Keamanan |
| **`pkg/auth`** | Hashing kata sandi, token sesi, MFA TOTP, OIDC/SAML2, RBAC | Identitas |
| **`pkg/ui`** | Token desain, komponen UI pakai ulang, tema terang/gelap | Presentasi Web |
| **`pkg/tui`** | Komponen terminal interaktif, tabel reaktif, bilah progres | Terminal UI |
| **`pkg/storage`** | Abstraksi penyimpanan berkas & objek (memory, disk, S3) | File Storage |
| **`pkg/cache`** | Caching in-memory berkecepatan tinggi dengan TTL dan invalidasi | Cache |
| **`pkg/resilience`** | Circuit breaker, retry dengan exponential backoff, timeout | Keandalan Sistem |
| **`pkg/testing`** | Harness pengujian, fixture assertion, mock HTTP | Jaminan Mutu |
| **`pkg/book`** | Engine generator situs dokumentasi dan buku teknis | Dokumentasi |
| **`pkg/kiw`** | Parser komponen file tunggal `.kiw` dan kompiler DSL | DSL Compiler |
