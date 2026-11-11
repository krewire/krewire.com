---
title: "Text"
description: "Pembungkus tipografi untuk teks paragraf isi, teks pengantar, dan teks diredam."
date: "2026-11-09"
---

# Text

Kategori: **Tampilan Data & Konten**

Pembungkus tipografi untuk teks paragraf isi, teks pengantar, dan teks diredam.

---

## Contoh Penggunaan Dasar

Gunakan `<Text />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Text class="kustom-text">
  <!-- Konten slot atau properti -->
</Text>
```

---

## Properti & Atribut

- `color`: Mengonfigurasi atribut atau perilaku `color`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `weight`: Mengonfigurasi atribut atau perilaku `weight`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
