---
title: "Pagination"
description: "Bilah navigasi halaman dengan tombol sebelum, sesudah, dan tautan angka."
date: "2026-11-09"
---

# Pagination

Kategori: **Navigasi & Struktur**

Bilah navigasi halaman dengan tombol sebelum, sesudah, dan tautan angka.

---

## Contoh Penggunaan Dasar

Gunakan `<Pagination />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Pagination class="kustom-pagination">
  <!-- Konten slot atau properti -->
</Pagination>
```

---

## Properti & Atribut

- `currentPage`: Mengonfigurasi atribut atau perilaku `currentPage`.
- `hasNext`: Mengonfigurasi atribut atau perilaku `hasNext`.
- `hasPrev`: Mengonfigurasi atribut atau perilaku `hasPrev`.
- `nextUrl`: Mengonfigurasi atribut atau perilaku `nextUrl`.
- `p`: Mengonfigurasi atribut atau perilaku `p`.
- `pages`: Mengonfigurasi atribut atau perilaku `pages`.
- `prevUrl`: Mengonfigurasi atribut atau perilaku `prevUrl`.
- `totalPages`: Mengonfigurasi atribut atau perilaku `totalPages`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
