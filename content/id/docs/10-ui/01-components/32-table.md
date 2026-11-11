---
title: "Table"
description: "Tabel data responsif dengan gaya bergaris, bingkai, dan sorot baris."
date: "2026-11-09"
---

# Table

Kategori: **Tampilan Data & Konten**

Tabel data responsif dengan gaya bergaris, bingkai, dan sorot baris.

---

## Contoh Penggunaan Dasar

Gunakan `<Table />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Table class="kustom-table">
  <!-- Konten slot atau properti -->
</Table>
```

---

## Properti & Atribut

- `bordered`: Mengonfigurasi atribut atau perilaku `bordered`.
- `cell`: Mengonfigurasi atribut atau perilaku `cell`.
- `h`: Mengonfigurasi atribut atau perilaku `h`.
- `headers`: Mengonfigurasi atribut atau perilaku `headers`.
- `hover`: Mengonfigurasi atribut atau perilaku `hover`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `rows`: Mengonfigurasi atribut atau perilaku `rows`.
- `striped`: Mengonfigurasi atribut atau perilaku `striped`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
