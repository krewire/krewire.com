---
title: "Dokumentasi — Jalur Belajar Komprehensif"
description: "Ikuti jalur belajar bertahap dari situs Krewire pertama Anda hingga workload produksi dan arsitektur platform."
date: "2026-11-12"
---

# Dokumentasi Krewire

> **Misi:** Misi Krewire adalah membuatnya mendekati gratis — baik dalam biaya finansial, beban operasional, maupun beban kognitif — bagi siapa pun untuk membangun, memiliki, dan mengoperasikan produk digital yang aman, berkelanjutan, dan berskala besar. Baca selengkapnya di [Mengapa Krewire](/overview/why-krewire).

**Krewire** adalah sekumpulan paket Go independen yang dijalankan melalui satu CLI (**`kiw`**) dan satu berkas konfigurasi (**`krewire.yaml`**). Anda menggunakannya untuk membangun situs web, dokumentasi, aplikasi web, CLI, dan — saat Anda membutuhkannya — worker, microservices, gateway, serta infrastruktur.

Anda **tidak** perlu membaca tentang arsitektur untuk menggunakannya. Ikuti tahapan di bawah ini secara berurutan, atau langsung menuju jalur yang sesuai dengan tujuan Anda.

---

## Jalur Belajar Sekilas

| Tahap | Tujuan | Baca |
| :--- | :--- | :--- |
| **0. Orientasi** | Memahami apa itu Krewire dan apakah cocok untuk Anda. | [1. Ringkasan](/overview) |
| **1. Sukses pertama** | Instal `kiw`, scaffold proyek, dan build sesuatu yang bisa dijalankan. | [2. Memulai](/getting-started) |
| **2. Bangun produk** | Pilih workload yang tepat dan pelajari siklus hidup serta perintahnya. | [3. Workload](/workloads) |
| **3. Pelajari fondasi** | Pahami model domain, siklus hidup, konfigurasi, validasi, dan resiliensi. | [4. Pustaka Inti & Fondasi](/core-libraries) |
| **4. Bangun antarmuka** | Tambahkan routing HTTP, wadah aplikasi, WASM, dan design system. | [5. Framework Web & UI](/web-ui) · [6. Krewire UI](/ui) |
| **5. Tambahkan data dan layanan** | Simpan data, cache, berkas, kirim email, dan integrasikan pembayaran. | [7. Data, Penyimpanan & Layanan](/data-services) |
| **6. Amankan dan skalakan** | Terapkan kontrol keamanan, identitas, cloud, dan kapabilitas platform AI. | [8. Keamanan & Platform Ekosistem](/security-platforms) |
| **7. Operasikan dan referensi** | Cari perintah, konfigurasi, dan playbook produksi. | [9. Referensi CLI](/cli-reference) · [10. Panduan & Praktik Terbaik](/guides) |
| **8. Pendalaman** | Pelajari invarian arsitektur yang menjaga ekosistem tetap komposabel. | [11. Pendalaman Arsitektur](/architecture) |

---

## Baru di Sini? Mulai dengan Jalur Pemula

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/quickstart">Quickstart — 5 menit <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Instal <code>kiw</code>, scaffold sebuah situs, ubah halaman, tambahkan komponen, jalankan server pengembangan, lalu build. Tidak perlu pengetahuan awal tentang Krewire.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/fullstack-app-tutorial">Aplikasi fullstack — 15 menit <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pengelola tugas dengan SQLite, halaman <code>.kiw</code>, komponen Krewire UI, form tervalidasi yang dilindungi CSRF, dan satu binary di akhir.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3</span>
      <h3 class="chapter-card-title">
        <a href="/architecture">Pendalaman Arsitektur <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Untuk platform engineer dan kontributor: piramida berlapis, DAG multi-repo, level cakupan, dan bagaimana data mengalir antara waktu build dan waktu request. Opsional bagi yang lain.
    </p>
  </div>

</div>

---

## Pilih Jalur Berdasarkan Tujuan

### "Saya ingin membangun situs web statis atau dokumentasi" — `site`, `book`

1. [Quickstart](/getting-started/quickstart) — situs pertama Anda dalam lima menit.
2. [DSL `.kiw`](/getting-started/dsl) — halaman, layout, komponen, scoped CSS.
3. [Workload Site](/workloads/site) atau [Workload Book](/workloads/book) — koleksi konten, bab Markdown, i18n.
4. [Komponen Krewire UI](/ui/components) — 66 komponen siap pakai.
5. [Deployment & Operasi](/guides/deployment-and-operations) — rilis hasil build statis.

### "Saya ingin membangun aplikasi web dengan API backend dan database" — `app`

1. [Quickstart](/getting-started/quickstart) — pelajari siklus kerja `kiw` pada situs kecil.
2. [Membangun Aplikasi Fullstack dalam 15 Menit](/getting-started/fullstack-app-tutorial) — SQLite, form, CSRF, satu binary.
3. [Workload App](/workloads/app) dan [Wadah aplikasi (`pkg/app`)](/web-ui/app).
4. [Web engine & perutean (`pkg/web`)](/web-ui/web) dan [Database & ORM (`pkg/orm`)](/data-services/database-and-orm).
5. [Kontrol keamanan](/security-platforms/security) dan [Autentikasi](/security-platforms/auth).
6. [Alur data: waktu build vs waktu request](/architecture/data-flow) — kapan merender secara statis dan kapan per request.

### "Saya ingin membangun platform enterprise dengan microservices" — `service`, `gateway`, `worker`

1. [Mengapa Krewire?](/overview/why-krewire) dan [matriks workload](/overview/krewire-workloads).
2. Workload [Service](/workloads/service), [Gateway](/workloads/gateway), dan [Worker](/workloads/worker).
3. [Sistem terdistribusi (`pkg/cloud`)](/security-platforms/cloud-distributed) dan [Resiliensi](/core-libraries/resilience).
4. [Pendalaman Arsitektur](/architecture) — piramida, DAG ekosistem, level cakupan.
5. [Deployment & Operasi](/guides/deployment-and-operations) dan [Hybrid PaaS & Cloud](/guides/hybrid-paas-and-cloud).

---

## Urutan Bacaan yang Disarankan per Peran

| Peran | Baca dengan urutan ini |
| :--- | :--- |
| **Pengembang frontend** | [Quickstart](/getting-started/quickstart) → [DSL `.kiw`](/getting-started/dsl) → [Komponen UI](/ui/components) → [Design system UI (`pkg/ui`)](/web-ui/ui) → [Web runtime & WASM](/web-ui/web-runtime) → [Workload Site](/workloads/site) |
| **Pengembang backend** | [Quickstart](/getting-started/quickstart) → [Tutorial aplikasi fullstack](/getting-started/fullstack-app-tutorial) → [`pkg/web`](/web-ui/web) → [`pkg/app`](/web-ui/app) → [Database & ORM](/data-services/database-and-orm) → [Validasi](/core-libraries/validation) → [Alur data](/architecture/data-flow) |
| **DevOps engineer** | [Instalasi](/getting-started/installation) → [Konfigurasi](/getting-started/configuration) → [Perintah CLI](/cli-reference/commands) → [Skrip & CI/CD](/cli-reference/scripting) → [Deployment & Operasi](/guides/deployment-and-operations) → [Hybrid PaaS & Cloud](/guides/hybrid-paas-and-cloud) |
| **Platform engineer** | [Ringkasan](/overview) → [Matriks workload](/overview/krewire-workloads) → [Pendalaman Arsitektur](/architecture) → [Model & primitif](/core-libraries/model) → [Kernel siklus hidup](/core-libraries/lifecycle) → [Open core & lisensi](/guides/open-core-and-licensing) |

---

## Semua Bab

| Bab | Isi yang akan Anda temukan |
| :--- | :--- |
| [1. Ringkasan](/overview) | Mengapa Krewire ada, filosofinya, dan delapan jenis workload. |
| [2. Memulai](/getting-started) | **Jalur pemula**: quickstart, tutorial fullstack, instalasi, konfigurasi, DSL `.kiw`. |
| [3. Workload](/workloads) | Satu panduan per workload: `app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`. |
| [4. Pustaka Inti & Fondasi](/core-libraries) | Model, siklus hidup, konfigurasi, validasi, resiliensi, pengujian. |
| [5. Framework Web & UI](/web-ui) | Router, wadah aplikasi, runtime WASM, dan TUI. |
| [6. Krewire UI](/ui) | Design system dan seluruh 66 komponen, dengan pratinjau langsung. |
| [7. Data, Penyimpanan & Layanan](/data-services) | ORM, penyimpanan objek, sistem berkas, cache, email, pembayaran. |
| [8. Keamanan & Platform Ekosistem](/security-platforms) | Kontrol keamanan, auth, identitas, cloud, agen AI. |
| [9. Referensi CLI](/cli-reference) | Setiap perintah `kiw` dan field `krewire.yaml`. |
| [10. Panduan & Praktik Terbaik](/guides) | Deployment, pengembangan berbantuan AI, hybrid cloud, lisensi open-core. |
| [11. Pendalaman Arsitektur](/architecture) | Platform engineering: piramida, DAG ekosistem, level cakupan, alur data, dan jalur belajar bertahap untuk tim. |
