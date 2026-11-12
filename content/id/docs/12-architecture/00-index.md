---
title: "Pendalaman Arsitektur"
description: "Jalur platform engineering: piramida dependensi berlapis, DAG ekosistem multi-repo, lima level cakupan, alur data antara waktu build dan waktu request, serta jalur belajar bertahap untuk tim."
date: "2026-11-09"
---

# Pendalaman Arsitektur (Platform Engineering)

<nav class="track-crumbs" aria-label="Jalur Pendalaman Arsitektur">
<span class="track-crumbs-label">Jalur pendalaman</span>
<ol>
<li aria-current="step"><span>Ringkasan</span></li>
<li><a href="/architecture/layered-pyramid">1 · Piramida</a></li>
<li><a href="/architecture/ecosystem-dag">2 · DAG ekosistem</a></li>
<li><a href="/architecture/scope-levels">3 · Level cakupan</a></li>
<li><a href="/architecture/data-flow">4 · Alur data</a></li>
</ol>
</nav>

Bab ini ditujukan bagi **platform engineer, arsitek, dan kontributor** repositori Krewire. Bab ini menjelaskan aturan yang menjaga puluhan paket independen dan beberapa repositori produk tetap dapat dikomposisi: kode mana yang boleh mengimpor kode mana, bagaimana repositori saling bergantung, bagaimana cakupan sistem ditentukan, dan di mana data masuk ke sebuah halaman.

> **Tips:** **Sedang membangun aplikasi atau situs?** Anda dapat melewati bab ini sepenuhnya. [Jalur pemula](/getting-started) dan panduan workload ditulis sedemikian rupa sehingga tidak ada konsep di sini yang diwajibkan. Kembalilah saat Anda merancang platform multi-layanan, berkontribusi ke inti, atau perlu menalar perbedaan rendering waktu build dan waktu request.

---

## Halaman dalam Jalur Ini

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.1</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/layered-pyramid">Piramida Berlapis &amp; Invarian <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Layer 0–5, impor ketat hanya ke arah bawah, pemisahan peer melalui bridge dan gateway, Single Source of Truth, dan aturan Go murni.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.2</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/ecosystem-dag">DAG Ekosistem <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Bagaimana repositori (<code>krewire</code>, <code>ui</code>, <code>krew</code>, <code>cloud</code>, <code>auth</code>, <code>dash</code>, <code>panel</code>, situs-situs) saling bergantung saat kompilasi dan terintegrasi saat runtime.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.3</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/scope-levels">Level Cakupan <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Workspace → Module → Domain → Service → Unit: lima level yang dikodifikasi di <code>pkg/model</code> dan bagaimana spesifikasi serta tes menargetkannya.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">12.4</span>
      <h3 class="chapter-card-title">
        <a href="/architecture/data-flow">Alur Data: SSG vs Runtime <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Bagaimana data mencapai view <code>.kiw</code> saat waktu build (<code>site</code>, <code>book</code>) dan saat waktu request (<code>app</code>), dengan resep handler Go dan pola fetch di sisi klien.
    </p>
  </div>

</div>

---

## Jalur Belajar Bertahap untuk Tim

Tim rekayasa yang mengadopsi Krewire di beberapa produk dapat mengikuti jalur enam tahap ini. Jalur ini mencakup seluruh dokumentasi, jadi perlakukan sebagai kurikulum, bukan prasyarat.

```text
[Stage 1: Core Mental Model] ───► [Stage 2: Developer Onboarding]
             │                                   │
             ▼                                   ▼
[Stage 3: Workload Mastery]   ───► [Stage 4: Domain Packages Deep-Dive]
             │                                   │
             ▼                                   ▼
[Stage 5: Daily Dev & CLI]    ───► [Stage 6: Cloud, Operations & Scale]
```

### Tahap 1: Model Mental Inti & Arsitektur
1. [**Mengapa Krewire?**](/overview/why-krewire) — Efisiensi sumber daya Go, binary tunggal, tanpa toolchain npm/bundler, dan kedaulatan data.
2. [**Filosofi Inti**](/overview/philosophy) — Tiga Pilar (*Secure, Sustainable, Scalable*) dan konvensi stdlib-first.
3. [**Matriks Workload**](/overview/krewire-workloads) — Delapan jenis proyek dan bagaimana mereka dikomposisi tanpa penulisan ulang.
4. [**Piramida Berlapis**](/architecture/layered-pyramid), [**DAG Ekosistem**](/architecture/ecosystem-dag), dan [**Level Cakupan**](/architecture/scope-levels) — aturan-aturan dalam bab ini.

### Tahap 2: Onboarding Pengembang
1. [**Jalur pemula**](/getting-started) — [Quickstart](/getting-started/quickstart) dan [Tutorial aplikasi fullstack](/getting-started/fullstack-app-tutorial).
2. [**Instalasi**](/getting-started/installation), [**Konfigurasi**](/getting-started/configuration), dan [**Struktur Direktori**](/getting-started/directory-structure).
3. [**DSL `.kiw`**](/getting-started/dsl) dan [**Pengembangan Berbantuan AI**](/getting-started/agentic-development).

### Tahap 3: Spesialisasi Workload
- **Konten:** [Book](/workloads/book), [Site](/workloads/site).
- **Alat terminal:** [CLI & TUI](/workloads/cli).
- **Web:** [App](/workloads/app), [Gateway](/workloads/gateway).
- **Sistem terdistribusi:** [Worker](/workloads/worker), [Service](/workloads/service), [Infra](/workloads/infra).

### Tahap 4: Pustaka Domain
- **Fondasi:** [Model](/core-libraries/model), [Siklus hidup](/core-libraries/lifecycle), [Konfigurasi](/core-libraries/config), [Validasi](/core-libraries/validation), [Resiliensi](/core-libraries/resilience), [Pengujian](/core-libraries/testing).
- **Web & UI:** [Web engine](/web-ui/web), [Wadah aplikasi](/web-ui/app), [Web runtime & WASM](/web-ui/web-runtime), [UI](/web-ui/ui), [Komponen UI](/ui/components), [TUI](/web-ui/tui).
- **Data:** [Database & ORM](/data-services/database-and-orm), [Penyimpanan](/data-services/storage), [Sistem berkas](/data-services/filesystem), [Caching](/data-services/caching), [Email](/data-services/mail), [Commerce & pembayaran](/data-services/commerce-and-payments).
- **Keamanan & platform:** [Keamanan](/security-platforms/security), [Auth](/security-platforms/auth), [Identitas](/security-platforms/identity-and-auth), [Sistem terdistribusi](/security-platforms/cloud-distributed), [Hybrid PaaS](/security-platforms/cloud-paas), [Orkestrasi AI](/security-platforms/ai-and-agents), [Agen otonom](/security-platforms/autonomous-agent), [Workspace virtual](/security-platforms/virtual-workspace).

### Tahap 5: Workflow Harian & CLI
1. [**Referensi Perintah CLI**](/cli-reference/commands)
2. [**Referensi Skema Konfigurasi**](/cli-reference/configuration)
3. [**Skrip Kustom & Otomasi**](/cli-reference/scripting)

### Tahap 6: Operasi, Cloud & Batas Open-Core
1. [**Deployment & Operasi**](/guides/deployment-and-operations)
2. [**Open Core & Lisensi**](/guides/open-core-and-licensing)

---

## Sumber

Aturan di halaman-halaman ini merangkum spesifikasi di repositori inti, yang tetap menjadi acuan resmi:

- `krewire/docs/foundation/architecture.md` — lapisan monorepo dan enam invarian lapisan.
- `krewire/docs/foundation/ecosystem-dag.md` — DAG ekosistem serta batas waktu kompilasi vs runtime.
- `krewire/pkg/model/scope.go` — hierarki cakupan dalam bentuk kode.
- `krewire/layering_test.go` — tes yang menegakkan pelapisan.
