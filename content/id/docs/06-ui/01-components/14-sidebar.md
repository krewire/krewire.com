---
title: "Sidebar"
description: "Bilah navigasi samping collapsible untuk dokumentasi dan dasbor aplikasi."
date: "2026-11-09"
---

# Sidebar

Kategori: **Navigasi & Struktur**

Bilah navigasi samping collapsible untuk dokumentasi dan dasbor aplikasi.

---

## Contoh Penggunaan Dasar

Gunakan `<Sidebar />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Sidebar class="kustom-sidebar">
  <!-- Konten slot atau properti -->
</Sidebar>
```

---

## Properti & Atribut

- `head`: Mengonfigurasi atribut atau perilaku `head`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `menu`: Mengonfigurasi atribut atau perilaku `menu`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
