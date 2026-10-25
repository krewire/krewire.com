---
title: "Memulai (Getting Started)"
description: "Panduan onboarding komprehensif ke Krewire: prasyarat, verifikasi sistem, quickstart 5 menit, siklus hidup pengembangan, dan peta bab."
date: "2026-09-29"
---

# Memulai (Getting Started)

Selamat datang di **Panduan Memulai Krewire**.

Bab ini adalah manual onboarding lengkap Anda untuk ekosistem Krewire. Panduan ini menuntun Anda mulai dari menyiapkan environment kerja, memahami toolchain pengembang, membuat perancah (*scaffolding*) dan menjalankan proyek pertama Anda, hingga mengorganisasi basis kode sesuai standar arsitektur Go yang teruji.

Baik Anda sedang menulis portal dokumentasi statis yang ringan, membuat utilitas terminal interaktif, membangun monolit web fullstack, atau mengoperasikan microservices berkinerja tinggi, panduan ini membekali Anda dengan alur kerja fundamental untuk membangun dan merilis perangkat lunak tanpa kelelahan toolchain.

---

## 1. Isi Bab Ini

Bab **Memulai** ini terbagi ke dalam lima sub-bab mendalam yang dirancang untuk memberikan penguasaan penuh atas toolchain Krewire:

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.1</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/installation">Instalasi <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Instruksi pemasangan lengkap untuk Linux, macOS, dan Windows. Pelajari skrip shell otomatis, instalasi dari source Go, validasi checksum kriptografis, dan auto-completion shell.
    </p>
    <ul class="chapter-card-list">
      <li>Installer shell otomatis & validasi checksum</li>
      <li><code>go install</code> & rilis binary mandiri</li>
      <li>Tab-completion untuk Bash, Zsh, dan Fish</li>
    </ul>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.2</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/configuration">Konfigurasi Terpadu <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Eksplorasi mendalam berkas <code>krewire.yaml</code>: metadata proyek, target build, port server, pemetaan direktori aset, dan skrip otomasi tugas.
    </p>
    <ul class="chapter-card-list">
      <li>Spesifikasi skema kanonikal YAML</li>
      <li>Override variabel environment (<code>KIW_ENV</code>, <code>KIW_DEBUG</code>)</li>
      <li>Otomasi tugas & skrip dev kustom</li>
    </ul>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.3</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/directory-structure">Struktur Direktori <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pelajari tata letak direktori standar Krewire yang memisahkan boundary domain secara ketat dan mengikuti konvensi Go kanonikal.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.4</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/agentic-development">Pengembangan Berbantuan AI <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Alur kerja kolaborasi bersama agen koding otonom menggunakan disiplin single-source-of-truth dan template Boost.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.5</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/dsl">Komponen DSL (.kiw) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Membuat antarmuka komponen file tunggal dengan frontmatter YAML, scoped style, blok markdown, dan sintaks mirip JSX.
    </p>
  </div>

</div>

---

## 2. Quickstart 5 Menit: Proyek Pertama Anda

Siapkan proyek web statis pertama Anda dalam waktu kurang dari 5 menit menggunakan CLI `kiw`:

```bash
# 1. Pasang CLI kiw
curl -fsSL https://krewire.com/scripts/install.sh | sh

# 2. Buat proyek baru
kiw new my-site --site
cd my-site

# 3. Jalankan server pengembangan lokal dengan live-reload
kiw dev
```

Buka peramban Anda di `http://localhost:8080`. Setiap perubahan berkas pada direktori `pages/` akan langsung diperbarui secara instan.

---

## Langkah Selanjutnya

Lanjutkan eksplorasi langkah demi langkah:

- Pelajari petunjuk lengkap di [**2.1 Instalasi →**](/getting-started/installation)
- Atau pahami manajemen konfigurasi di [**2.2 Konfigurasi Terpadu →**](/getting-started/configuration)
