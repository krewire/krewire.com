---
title: "Tooltip"
description: "Petunjuk tooltip informasi yang muncul saat disorot atau difokuskan."
date: "2026-11-09"
---

# Tooltip

Kategori: **Lapisan Mengambang & Dialog**

Petunjuk tooltip informasi yang muncul saat disorot atau difokuskan.

---

## Contoh Penggunaan Dasar

Gunakan `<Tooltip />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Tooltip class="kustom-tooltip">
  <!-- Konten slot atau properti -->
</Tooltip>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `placement`: Mengonfigurasi atribut atau perilaku `placement`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
