---
title: "Box"
description: "Kotak tata letak serbaguna dengan dukungan border, padding, dan bayangan pop-brutalist."
date: "2026-11-09"
---

# Box

Kategori: **Tata Letak Primitif**

Kotak tata letak serbaguna dengan dukungan border, padding, dan bayangan pop-brutalist.

---

## Contoh Penggunaan Dasar

Gunakan `<Box />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Box class="kustom-box">
  <!-- Konten slot atau properti -->
</Box>
```

---

## Properti & Atribut

- `border`: Mengonfigurasi atribut atau perilaku `border`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `rounded`: Mengonfigurasi atribut atau perilaku `rounded`.
- `shadow`: Mengonfigurasi atribut atau perilaku `shadow`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
