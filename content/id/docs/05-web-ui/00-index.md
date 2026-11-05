---
title: "Framework Web & UI"
description: "Katalog referensi pustaka web dan antarmuka pengguna Krewire: perutean HTTP radix, wadah aplikasi, runtime WebAssembly, komponen .kiw UI, dan TUI."
date: "2026-10-25"
---

# Framework Web & UI

Krewire menyediakan lapisan web dan presentasi yang kohesif, dirancang untuk latensi sub-milidetik, jejak memori minimal (~15MB RAM), dan eksekusi Go murni.

Framework ini mendukung pengalaman render di sisi server maupun interaktivitas sisi klien:
- **Perutean Radix Tree:** Mesin perutean tanpa alokasi dengan ekstraksi parameter dan pipeline middleware.
- **Arsitektur Monolit Terpadu:** Wadah aplikasi dengan injeksi dependensi otomatis dan supervisi siklus hidup.
- **Go WASM di Sisi Klien:** Jalankan Go murni di peramban web dengan diffing virtual DOM dan hidrasi SSR.
- **Komponen & Toolkit Terminal:** Token desain `.kiw` yang dapat digunakan ulang dan aplikasi terminal berbasis arsitektur Elm.

---

## Isi Chapter Ini

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.1</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/web">Web Engine & Perutean HTTP (pkg/web) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Router HTTP radix tree tanpa alokasi memori, pipeline middleware, format problem details RFC 7807, ekstraksi parameter, dan helper render.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.2</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/app">Wadah Aplikasi & Monolit (pkg/app) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Wadah dependency injection, manajemen siklus hidup aplikasi, dan supervisi proses yang anggun untuk monolit web.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.3</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/web-runtime">Web Runtime & WebAssembly (pkg/runtime) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Kompilasi WebAssembly Go di sisi klien, reaktivitas sinyal, rekonsiliasi virtual DOM, dan hidrasi render sisi server (SSR).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.4</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/ui">Komponen Web UI & Token Desain (pkg/ui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Komponen Single-File (.kiw) yang dapat digunakan kembali, token desain CSS semantik, skala tipografi, dan primitif komponen aksesibel.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">5.5</span>
      <h3 class="chapter-card-title">
        <a href="/web-ui/tui">Toolkit Terminal UI (pkg/tui) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Framework antarmuka terminal yang dibangun di atas arsitektur Elm/Bubble Tea, menampilkan pewarnaan ANSI, viewport scrolling, dan prompt interaktif.
    </p>
  </div>

</div>
