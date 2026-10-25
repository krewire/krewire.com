---
title: "Referensi CLI"
description: "Panduan komprehensif untuk CLI kiw: arsitektur, flag global, variabel environment, kode keluar, dan alur kerja pengembang harian."
date: "2026-10-08"
---

# Referensi CLI

CLI **`kiw`** adalah pintu masuk tunggal yang menyatukan seluruh ekosistem Krewire. Alat ini menggerakkan kedelapan jenis proyek (`app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`) di balik satu binary dan satu berkas konfigurasi (**`krewire.yaml`**).

`kiw` dibangun secara langsung menggunakan pustaka `pkg/tui` dan `pkg/term` milik Krewire sendiri, memastikan nol dependensi eksternal, waktu mulai instan (<2 milidetik), dan disiplin kode keluar POSIX yang konsisten.

---

## Isi Bab Ini

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.1</span>
      <h3 class="chapter-card-title">
        <a href="/cli-reference/commands">Katalog Perintah <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Dokumentasi terperinci untuk setiap perintah <code>kiw</code>: flag, sintaks penggunaan, workload yang didukung, dan perilaku eksekusinya.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.2</span>
      <h3 class="chapter-card-title">
        <a href="/cli-reference/configuration">Skema Konfigurasi (krewire.yaml) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Spesifikasi menyeluruh skema <code>krewire.yaml</code>: metadata proyek, target build, port server dev, dan blok konfigurasi khusus workload.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.3</span>
      <h3 class="chapter-card-title">
        <a href="/cli-reference/scripting">Otomasi Tugas &amp; Skrip <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Mendefinisikan perintah otomasi kustom, hook pra-build, dan tugas harian tanpa memerlukan Make atau pustaka task runner eksternal.
    </p>
  </div>

</div>
