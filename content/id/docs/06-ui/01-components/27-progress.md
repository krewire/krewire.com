---
title: "Progress"
description: "Indikator bilah progres kemajuan berstatus pasti maupun indeterminat."
date: "2026-11-09"
---

# Progress

Kategori: **Umpan Balik & Status**

Indikator bilah progres kemajuan berstatus pasti maupun indeterminat.

---

## Contoh Penggunaan Dasar

Gunakan `<Progress />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Progress class="kustom-progress">
  <!-- Konten slot atau properti -->
</Progress>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `showValue`: Mengonfigurasi atribut atau perilaku `showValue`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
