---
title: "Code"
description: "Cuplikan kode monospace sebaris dengan style kontras tinggi."
date: "2026-11-09"
---

# Code

Kategori: **Tampilan Data & Konten**

Cuplikan kode monospace sebaris dengan style kontras tinggi.

---

## Contoh Penggunaan Dasar

Gunakan `<Code />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Code class="kustom-code">
  <!-- Konten slot atau properti -->
</Code>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
