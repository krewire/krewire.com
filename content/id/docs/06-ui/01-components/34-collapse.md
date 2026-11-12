---
title: "Collapse"
description: "Kontainer lipat tunggal untuk menampilkan atau menyembunyikan konten."
date: "2026-11-09"
---

# Collapse

Kategori: **Tampilan Data & Konten**

Kontainer lipat tunggal untuk menampilkan atau menyembunyikan konten.

---

## Contoh Penggunaan Dasar

Gunakan `<Collapse />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Collapse class="kustom-collapse">
  <!-- Konten slot atau properti -->
</Collapse>
```

---

## Properti & Atribut

- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `open`: Mengonfigurasi atribut atau perilaku `open`.
- `summary`: Mengonfigurasi atribut atau perilaku `summary`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
