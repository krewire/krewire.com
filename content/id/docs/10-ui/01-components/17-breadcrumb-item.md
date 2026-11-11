---
title: "BreadcrumbItem"
description: "Segmen individual di dalam jejak navigasi breadcrumb."
date: "2026-11-09"
---

# BreadcrumbItem

Kategori: **Navigasi & Struktur**

Segmen individual di dalam jejak navigasi breadcrumb.

---

## Contoh Penggunaan Dasar

Gunakan `<BreadcrumbItem />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<BreadcrumbItem class="kustom-breadcrumb-item">
  <!-- Konten slot atau properti -->
</BreadcrumbItem>
```

---

## Properti & Atribut

- `current`: Mengonfigurasi atribut atau perilaku `current`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `target`: Mengonfigurasi atribut atau perilaku `target`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
