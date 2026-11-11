---
title: "Select"
description: "Kolom pilihan dropdown select bergaya native."
date: "2026-11-09"
---

# Select

Kategori: **Formulir & Masukan**

Kolom pilihan dropdown select bergaya native.

---

## Contoh Penggunaan Dasar

Gunakan `<Select />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Select class="kustom-select">
  <!-- Konten slot atau properti -->
</Select>
```

---

## Properti & Atribut

- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `error`: Mengonfigurasi atribut atau perilaku `error`.
- `help`: Mengonfigurasi atribut atau perilaku `help`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `required`: Mengonfigurasi atribut atau perilaku `required`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
