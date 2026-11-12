---
title: "Head"
description: "Penyuntik tag head standar untuk metadata, OpenGraph, dan favicon."
date: "2026-11-09"
---

# Head

Kategori: **Navigasi & Struktur**

Penyuntik tag head standar untuk metadata, OpenGraph, dan favicon.

---

## Contoh Penggunaan Dasar

Gunakan `<Head />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Head class="kustom-head">
  <!-- Konten slot atau properti -->
</Head>
```

---

## Properti & Atribut

- `assetversion`: Mengonfigurasi atribut atau perilaku `assetversion`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
