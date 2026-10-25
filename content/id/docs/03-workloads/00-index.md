---
title: "Workload Mendalam"
description: "Panduan komprehensif untuk 8 jenis workload resmi Krewire: app, cli, site, book, worker, service, gateway, dan infra. Pelajari arsitektur dan model eksekusinya."
date: "2026-10-08"
---

# Workload Mendalam

Krewire memberikan pengalaman pengembang yang terpadu di delapan jenis workload kelas satu. Alih-alih menyambungkan pustaka dan runtime yang berbeda secara terpisah-pisah, setiap workload didefinisikan dalam satu berkas konfigurasi (**`krewire.yaml`**) dan diorkestrasikan melalui satu developer tool (**`kiw`**).

Pemilihan jenis workload memberi tahu kompiler, server dev lokal, dan pipeline deployment sembari tetap mempertahankan ergonomi Go standar dan dependensi menurun (*downward-only*) yang ketat.

---

## Isi Bab Ini

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.1</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/app">Monolit Web &amp; API (app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Aplikasi web fullstack binary tunggal dengan Server-Side Rendering (SSR), API REST/JSON, autentikasi sesi, dan integrasi database langsung.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.2</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/cli">CLI &amp; TUI Terminal (cli) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Membangun aplikasi baris perintah dan TUI terminal interaktif menggunakan <code>pkg/tui</code> dengan startup milidetik.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.3</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/site">Situs Statis / SSG (site) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Situs web cepat, ramah SEO, dan hemat biaya komputasi yang dikompilasi dari template komponen <code>.kiw</code> tanpa dependensi Node.js.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.4</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/book">Buku Dokumentasi (book) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Engine kompilasi manuskrip Markdown menjadi portal buku teknis lengkap dengan TOC hierarkis, navigasi pager, dan dukungan i18n.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.5</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/worker">Worker Latar Belakang (worker) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pemrosesan tugas latar belakang asinkron, antrean terdistribusi, penjadwal cron periodik, dan dead-letter queue.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.6</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/service">Microservices (service) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Layanan terdistribusi dengan penemuan layanan dinamis (service discovery), observabilitas OTel, dan API RPC/HTTP.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.7</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/infra">Infrastruktur sebagai Kode (infra) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Mendeklarasikan dan memprovisi sumber daya cloud modern dalam bahasa Go murni dengan rencana eksekusi dan verifikasi status.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">3.8</span>
      <h3 class="chapter-card-title">
        <a href="/workloads/gateway">Gateway API (gateway) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Reverse proxy berkinerja tinggi dengan konfigurasi per rute, pembatasan kuota (rate limiting), dan load balancing dinamis.
    </p>
  </div>

</div>
