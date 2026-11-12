---
title: "LangSwitch"
description: "Tombol pengalih bahasa untuk aplikasi multi-bahasa internasional."
date: "2026-11-09"
---

# LangSwitch

Kategori: **Aksi & Tombol**

Tombol pengalih bahasa untuk aplikasi multi-bahasa internasional.

---

## Contoh Penggunaan Dasar

Gunakan `<LangSwitch />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<LangSwitch class="kustom-lang-switch">
  <!-- Konten slot atau properti -->
</LangSwitch>
```

---

## Properti & Atribut

- `code`: Mengonfigurasi atribut atau perilaku `code`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
