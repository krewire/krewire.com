---
title: "Paket & Pustaka"
description: "Katalog referensi pustaka domain modular Krewire: batas dependensi menurun, invarian SSOT, dan pola integrasi."
date: "2026-10-25"
---

# Paket & Pustaka

Krewire dihadirkan sebagai **Pustaka Modular**: paket-paket Go independen dengan tujuan tunggal yang Anda komposisikan, bukan sebuah framework monolitik kaku yang harus Anda warisi.

Setiap paket di dalam katalog `krewire/pkg` mematuhi batasan arsitektur yang ketat:
- **Dependensi Menurun Secara Ketat (Strict Downward Dependencies):** Primitif domain dan aturan bisnis inti berada di `pkg/model` dan `pkg/errs`. Pustaka tingkat tinggi tidak boleh menyuntikkan dependensi ke atas.
- **Satu Sumber Kebenaran (SSOT - Single Source of Truth):** Fakta, konstanta, dan aturan validasi dimiliki tepat oleh satu paket dan hanya diimpor oleh pihak lain.
- **Pustaka Standar Pertama:** Nol dependensi pihak ketiga yang tidak perlu. Jika pustaka standar Go menyediakan kapabilitas (`net/http`, `context`, `log/slog`, `embed`), Krewire memperluasnya secara harmonis.

---

## Apa yang Ada di Bab Ini?

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.1</span>
      <h3 class="chapter-card-title">
        <a href="/packages/model">Model Workload & Primitif Domain (pkg/model, pkg/errs) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Model domain fondasi: delapan jenis workload, metadata proyek, hierarki scope, aturan opt-in, taksonomi error, resolusi lingkungan, dan kontrak versi.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.2</span>
      <h3 class="chapter-card-title">
        <a href="/packages/lifecycle">Supervisi Siklus Hidup & Kernel (pkg/lifecycle) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Bidang kontrol imperatif Krewire: supervisor kernel, registri modul, pengurutan dependensi teratur, dan graceful shutdown.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.3</span>
      <h3 class="chapter-card-title">
        <a href="/packages/web">Web Engine & Perutean HTTP (pkg/web) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Router HTTP radix tree zero-allocation, pipeline middleware, problem details RFC 7807, ekstraksi parameter, dan helper render.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.4</span>
      <h3 class="chapter-card-title">
        <a href="/packages/app">Wadah Aplikasi & Monolit (pkg/app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Wadah dependency injection, manajemen siklus hidup aplikasi, dan supervisi proses graceful shutdown.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.5</span>
      <h3 class="chapter-card-title">
        <a href="/packages/security">Kontrol Keamanan & Hardening (pkg/sec) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Default keamanan tangguh, middleware OWASP Top 10, pencegahan CSRF, proteksi SSRF, dan redaksi PII.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.6</span>
      <h3 class="chapter-card-title">
        <a href="/packages/auth">Primitif Autentikasi (pkg/auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Primitif autentikasi kriptografis inti, tanda tangan JWT HMAC-SHA256, dan HTTP Basic Auth waktu konstan.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.7</span>
      <h3 class="chapter-card-title">
        <a href="/packages/ui">Komponen Web UI & Token Desain (pkg/ui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Komponen berkas tunggal (.kiw), token desain CSS semantik, dan engine peralihan tema terang/gelap.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.8</span>
      <h3 class="chapter-card-title">
        <a href="/packages/tui">Toolkit Terminal UI (pkg/tui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Antarmuka terminal interaktif, arsitektur Elm/Bubble Tea, pewarnaan ANSI, tabel, dan bilah progres.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.9</span>
      <h3 class="chapter-card-title">
        <a href="/packages/storage">Penyimpanan Objek Multi-Driver (pkg/storage) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Abstraksi penyimpanan objek seragam dengan driver memory, disk lokal, dan Amazon S3 / MinIO.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.10</span>
      <h3 class="chapter-card-title">
        <a href="/packages/filesystem">Sistem Berkas Virtual & Penulisan Atomik (pkg/fs, pkg/file) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Batas sistem berkas virtual yang dapat diuji, penulisan berkas atomik crash-safe, dan keamanan path.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.11</span>
      <h3 class="chapter-card-title">
        <a href="/packages/resilience">Pola Resiliensi & Circuit Breaker (pkg/resilience) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Komunikasi jaringan toleran kesalahan dengan retry backoff berjitter, batas waktu, dan circuit breaker.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.12</span>
      <h3 class="chapter-card-title">
        <a href="/packages/testing">Utilitas Pengujian & Asersi (pkg/testing) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Harness asersi pengujian, pembantu HTTP, asersi komponen UI, dan fixture mock.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.13</span>
      <h3 class="chapter-card-title">
        <a href="/packages/cloud-distributed">Cloud & Sistem Terdistribusi (pkg/cloud) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Antrean worker asinkron, penjadwal cron, registri microservice, dan skema Infrastructure as Code.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.14</span>
      <h3 class="chapter-card-title">
        <a href="/packages/config">Konfigurasi Hierarkis (pkg/config) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pemuatan konfigurasi aplikasi terstruktur dari flag CLI, env vars, file .env, dan tag struct.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.15</span>
      <h3 class="chapter-card-title">
        <a href="/packages/validation">Validasi Input Komposabel (pkg/validation) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Validasi struktur data komposabel tanpa overhead refleksi, menyediakan error terstruktur untuk API.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.16</span>
      <h3 class="chapter-card-title">
        <a href="/packages/ai-and-agents">AI & Agen Otonom (pkg/ai) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Integrasi LLM multi-provider, adapter OpenAI/Anthropic/Gemini, penyelesaian streaming, dan template prompt.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.17</span>
      <h3 class="chapter-card-title">
        <a href="/packages/database-and-orm">Database & Multi-Driver ORM (pkg/orm) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Fluent query builder, pemetaan entitas model, dan engine database multi-driver (SQLite, MySQL, Postgres, MongoDB, Elastic).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.18</span>
      <h3 class="chapter-card-title">
        <a href="/packages/caching">Cache Multi-Driver (pkg/cache, pkg/redis) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Caching performa tinggi in-memory dan Redis terdistribusi dengan TTL, invalidasi, dan connection pooling.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.19</span>
      <h3 class="chapter-card-title">
        <a href="/packages/mail">Mailer Transaksional (pkg/mail) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pengiriman email transaksional multi-driver yang mendukung SMTP, Resend, SendGrid, Postmark, dan mock in-memory.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.20</span>
      <h3 class="chapter-card-title">
        <a href="/packages/commerce-and-payments">Commerce & Gateway Pembayaran Indonesia (pkg/commerce) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Integrasi pembayaran Midtrans mendukung QRIS dinamis, Virtual Account bank nasional (BCA, Mandiri, BRI, BNI), dan kartu kredit.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.21</span>
      <h3 class="chapter-card-title">
        <a href="/packages/identity-and-auth">Platform Identitas & Autentikasi Enterprise (krewire/auth) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Manajemen identitas multi-tenant, hashing PBKDF2, MFA TOTP RFC 6238, SSO enterprise (OIDC/SAML2), dan RBAC hierarkis.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.22</span>
      <h3 class="chapter-card-title">
        <a href="/packages/cloud-paas">Platform PaaS Hibrida (krewire/cloud) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Mesin PaaS open-core self-hosted dan cloud SaaS terkelola dengan builder hermetis dan router edge reverse proxy.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.23</span>
      <h3 class="chapter-card-title">
        <a href="/packages/web-runtime">Web Runtime & WASM (pkg/runtime) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Kompilasi Go ke WebAssembly di sisi klien: virtual DOM diffing, generic hooks, registri komponen, dan hidrasi SSR.
    </p>
  </div>

</div>

---

## Hierarki Dependensi Ekosistem

Sebagaimana didokumentasikan dalam `internal/docs/foundation/architecture.md`, aliran dependensi berjalan ketat ke arah bawah:

```text
templates/boost (AI agent starters)
       │
       ▼
tools/kiw (Developer CLI)
       │
       ▼
pkg/* (Pustaka domain: web, sec, auth, ui, tui, storage, orm, cache, mail...)
       │
       ▼
pkg/model, pkg/lifecycle, pkg/errs (model domain, supervisi, error, environment, versi)
```

Tidak ada pustaka atau paket yang boleh mengimpor `kiw`. Aturan domain berada di `pkg/model`, dan mekanika runtime berada di `pkg/lifecycle`.
