---
title: "Tag"
description: "Label kata kunci yang dapat dihapus atau diklik dengan varian warna."
date: "2026-11-09"
---

# Tag

Kategori: **Umpan Balik & Status**

Label kata kunci yang dapat dihapus atau diklik dengan varian warna.

---

## Contoh Penggunaan Dasar

Gunakan `<Tag />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Tag class="kustom-tag">
  <!-- Konten slot atau properti -->
</Tag>
```

---

## Properti & Atribut

- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `removable`: Mengonfigurasi atribut atau perilaku `removable`.
- `selected`: Mengonfigurasi atribut atau perilaku `selected`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
