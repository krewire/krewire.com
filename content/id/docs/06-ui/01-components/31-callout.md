---
title: "Callout"
description: "Kotak catatan berbingkai untuk tips, peringatan, dan catatan dokumentasi."
date: "2026-11-09"
---

# Callout

Kategori: **Umpan Balik & Status**

Kotak catatan berbingkai untuk tips, peringatan, dan catatan dokumentasi.

---

## Contoh Penggunaan Dasar

Gunakan `<Callout />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Callout class="kustom-callout">
  <!-- Konten slot atau properti -->
</Callout>
```

---

## Properti & Atribut

- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `type`: Mengonfigurasi atribut atau perilaku `type`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
