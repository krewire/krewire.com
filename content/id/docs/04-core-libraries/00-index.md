---
title: "Pustaka Inti & Fondasi"
description: "Katalog referensi pustaka inti Krewire: model workload, supervisi siklus hidup, konfigurasi, validasi, resiliensi, dan harness pengujian."
date: "2026-10-25"
---

# Pustaka Inti & Fondasi

Tingkat fondasi Krewire menetapkan jaminan arsitektur, aturan domain, dan supervisi siklus hidup yang dibutuhkan oleh seluruh workload.

Pustaka pada lapisan ini mematuhi prinsip desain yang ketat:
- **Dependensi Searah ke Bawah:** Primitif domain dan aturan bisnis berada di `pkg/model` dan `pkg/errs`. Pustaka lapisan atas tidak boleh menyuntikkan dependensi ke atas.
- **Single Source of Truth (SSOT):** Fakta, konstanta, dan aturan validasi dimiliki oleh tepat satu paket dan diimpor oleh modul lainnya.
- **Tanpa Ketergantungan Eksternal yang Tidak Perlu:** Jika pustaka standar Go sudah mencukupi, Krewire mengembangkannya secara langsung tanpa menambah rantai dependensi pihak ketiga.

---

## Isi Chapter Ini

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.1</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/model">Model Workload & Primitif Domain (pkg/model, pkg/errs) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Model domain dasar: delapan jenis workload, metadata proyek, hierarki scope, taksonomi error, deteksi lingkungan, dan kontrak versi.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.2</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/lifecycle">Supervisi Siklus Hidup & Kernel (pkg/lifecycle) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Control plane imperatif Krewire: supervisor kernel, registri modul, pengurutan dependensi, dan penghentian proses secara anggun (graceful shutdown).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.3</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/config">Konfigurasi Hierarkis (pkg/config) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pemuatan konfigurasi deterministik dengan urutan preseden dari flag CLI, variabel lingkungan, dan berkas .env bertingkat.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.4</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/validation">Validasi Input Komposabel (pkg/validation) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Aturan validasi data komposabel tanpa refleksi untuk payload API, formulir, dan skema logika bisnis.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.5</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/resilience">Pola Resiliensi & Circuit Breaker (pkg/resilience) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Circuit breaker, backoff retry eksponensial dengan jitter, batas waktu deadline, dan isolasi kegagalan untuk operasi terdistribusi.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">4.6</span>
      <h3 class="chapter-card-title">
        <a href="/core-libraries/testing">Utilitas Pengujian & Asersi (pkg/testing) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Utilitas asersi cepat, perekam respons HTTP, fixture server efemeral, dan pembantu pengujian snapshot/golden.
    </p>
  </div>

</div>
