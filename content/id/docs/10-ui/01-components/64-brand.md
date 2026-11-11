---
title: "Brand"
description: "Logo brand ekosistem lengkap dengan lambang ikon dan badge versi."
date: "2026-11-09"
---

# Brand

Kategori: **Navigasi & Struktur**

Logo brand ekosistem lengkap dengan lambang ikon dan badge versi.

---

## Contoh Penggunaan Dasar

Gunakan `<Brand />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Brand class="kustom-brand">
  <!-- Konten slot atau properti -->
</Brand>
```

---

## Properti & Atribut

- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `badgeI18n`: Mengonfigurasi atribut atau perilaku `badgeI18n`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `version`: Mengonfigurasi atribut atau perilaku `version`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
