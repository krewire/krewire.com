---
title: "Hamburger"
description: "Tombol menu hamburger 3 baris teranimasi untuk membuka drawer mobile."
date: "2026-11-09"
---

# Hamburger

Kategori: **Navigasi & Struktur**

Tombol menu hamburger 3 baris teranimasi untuk membuka drawer mobile.

---

## Contoh Penggunaan Dasar

Gunakan `<Hamburger />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Hamburger class="kustom-hamburger">
  <!-- Konten slot atau properti -->
</Hamburger>
```

---

## Properti & Atribut

- `Class`: Mengonfigurasi atribut atau perilaku `Class`.
- `Controls`: Mengonfigurasi atribut atau perilaku `Controls`.
- `DataDrawerToggle`: Mengonfigurasi atribut atau perilaku `DataDrawerToggle`.
- `Expanded`: Mengonfigurasi atribut atau perilaku `Expanded`.
- `ID`: Mengonfigurasi atribut atau perilaku `ID`.
- `Label`: Mengonfigurasi atribut atau perilaku `Label`.
- `controls`: Mengonfigurasi atribut atau perilaku `controls`.
- `dataDrawerToggle`: Mengonfigurasi atribut atau perilaku `dataDrawerToggle`.
- `expanded`: Mengonfigurasi atribut atau perilaku `expanded`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
