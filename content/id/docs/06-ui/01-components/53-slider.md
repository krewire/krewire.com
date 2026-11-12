---
title: "Slider"
description: "Input bilah geser rentang angka dengan gelembung nilai dinamis."
date: "2026-11-09"
---

# Slider

Kategori: **Formulir & Masukan**

Input bilah geser rentang angka dengan gelembung nilai dinamis.

---

## Contoh Penggunaan Dasar

Gunakan `<Slider />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Slider class="kustom-slider">
  <!-- Konten slot atau properti -->
</Slider>
```

---

## Properti & Atribut

- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `max`: Mengonfigurasi atribut atau perilaku `max`.
- `min`: Mengonfigurasi atribut atau perilaku `min`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `onInput`: Mengonfigurasi atribut atau perilaku `onInput`.
- `showValue`: Mengonfigurasi atribut atau perilaku `showValue`.
- `step`: Mengonfigurasi atribut atau perilaku `step`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
