---
title: "Heading"
description: "Pembungkus tipografi heading semantik dari h1 hingga h6."
date: "2026-11-09"
---

# Heading

Kategori: **Tampilan Data & Konten**

Pembungkus tipografi heading semantik dari h1 hingga h6.

---

## Contoh Penggunaan Dasar

Gunakan `<Heading />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Heading class="kustom-heading">
  <!-- Konten slot atau properti -->
</Heading>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `level`: Mengonfigurasi atribut atau perilaku `level`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
