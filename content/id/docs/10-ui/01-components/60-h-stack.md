---
title: "HStack"
description: "Kontainer flex horizontal dengan jarak seragam antar elemen anak."
date: "2026-11-09"
---

# HStack

Kategori: **Tata Letak Primitif**

Kontainer flex horizontal dengan jarak seragam antar elemen anak.

---

## Contoh Penggunaan Dasar

Gunakan `<HStack />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<HStack class="kustom-h-stack">
  <!-- Konten slot atau properti -->
</HStack>
```

---

## Properti & Atribut

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `gap`: Mengonfigurasi atribut atau perilaku `gap`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `justify`: Mengonfigurasi atribut atau perilaku `justify`.
- `noWrap`: Mengonfigurasi atribut atau perilaku `noWrap`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
