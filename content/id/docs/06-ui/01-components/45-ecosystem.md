---
title: "Ecosystem"
description: "Komponen grid interaktif untuk menampilkan arsitektur 10 repositori Krewire."
date: "2026-11-09"
---

# Ecosystem

Kategori: **Tampilan Data & Konten**

Komponen grid interaktif untuk menampilkan arsitektur 10 repositori Krewire.

---

## Contoh Penggunaan Dasar

Gunakan `<Ecosystem />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Ecosystem class="kustom-ecosystem">
  <!-- Konten slot atau properti -->
</Ecosystem>
```

---

## Properti & Atribut

- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `lead`: Mengonfigurasi atribut atau perilaku `lead`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
