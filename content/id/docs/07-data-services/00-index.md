---
title: "Data, Penyimpanan & Layanan"
description: "Katalog referensi pustaka persistensi data dan layanan backend Krewire: multi-driver ORM, penyimpanan objek, sistem berkas virtual, caching, mailer, dan pembayaran Indonesia."
date: "2026-10-25"
---

# Data, Penyimpanan & Layanan

Layanan digital modern membutuhkan pola integrasi dan persistensi data yang tangguh dan modular. Krewire memisahkan backend data dari logika bisnis menggunakan abstraksi agnostik-driver:

- **Arsitektur Multi-Driver Fleksibel:** Beralih antara penyimpanan memori, disk lokal, dan engine terdistribusi tanpa mengubah kode aplikasi.
- **Keandalan Atomik:** Sistem berkas virtual dan writer atomik memastikan operasi tahan terhadap pemutusan proses mendadak atau kegagalan daya.
- **Lokalisasi Utama:** Dukungan bawaan untuk infrastruktur perbankan Indonesia (BCA, Mandiri, BRI, BNI) dan pembayaran QRIS nasional.

---

## Isi Chapter Ini

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.1</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/database-and-orm">Database & Multi-Driver ORM (pkg/orm) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Query builder yang fasih dan lapisan database multi-driver yang mendukung SQLite, MySQL, PostgreSQL, MongoDB, dan Elasticsearch disertai migrasi otomatis.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.2</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/storage">Penyimpanan Objek Multi-Driver (pkg/storage) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Antarmuka penyimpanan objek terpadu dengan driver yang dapat ditukar untuk pengujian in-memory, volume disk lokal, dan penyimpanan cloud kompatibel S3 (MinIO, AWS, Cloudflare R2).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.3</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/filesystem">Sistem Berkas Virtual & Berkas Atomik (pkg/fs, pkg/file) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Abstraksi sistem berkas virtual yang terisolasi dan dapat diuji (io/fs, memfs) serta primitif penulisan berkas atomik yang aman dari kerusakan.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.4</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/caching">Cache Multi-Driver (pkg/cache, pkg/redis) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Cache LRU in-memory berkecepatan tinggi dan klien cache Redis terdistribusi dengan kedaluwarsa TTL, invalidasi atomik, dan pengelompokan berbasis tag.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.5</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/mail">Mailer Transaksional (pkg/mail) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pengirim email transaksional multi-driver dengan rendering template bawaan, dukungan lampiran, dan integrasi untuk SMTP, Resend, SendGrid, serta Postmark.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">7.6</span>
      <h3 class="chapter-card-title">
        <a href="/data-services/commerce-and-payments">Commerce & Pembayaran Indonesia (pkg/commerce) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Primitif pemesanan e-commerce, alur keranjang belanja, dan gateway pembayaran asli Indonesia: Midtrans, QRIS, Virtual Account, dan e-wallet.
    </p>
  </div>

</div>
