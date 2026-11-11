---
title: "Spinner"
description: "Pemutar loader SVG teranimasi yang aksesibel untuk operasi yang sedang berjalan."
date: "2026-11-09"
---

# Spinner

Kategori: **Umpan Balik & Status**

Pemutar loader SVG teranimasi yang aksesibel untuk operasi yang sedang berjalan.

---

## Contoh Penggunaan Dasar

Gunakan `<Spinner />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Spinner class="kustom-spinner">
  <!-- Konten slot atau properti -->
</Spinner>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
