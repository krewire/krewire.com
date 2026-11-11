---
title: "Textarea"
description: "Kolom input teks multi-baris dengan dukungan penyesuaian ukuran otomatis."
date: "2026-11-09"
---

# Textarea

Kategori: **Formulir & Masukan**

Kolom input teks multi-baris dengan dukungan penyesuaian ukuran otomatis.

---

## Contoh Penggunaan Dasar

Gunakan `<Textarea />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Textarea class="kustom-textarea">
  <!-- Konten slot atau properti -->
</Textarea>
```

---

## Properti & Atribut

- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `error`: Mengonfigurasi atribut atau perilaku `error`.
- `help`: Mengonfigurasi atribut atau perilaku `help`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `placeholder`: Mengonfigurasi atribut atau perilaku `placeholder`.
- `required`: Mengonfigurasi atribut atau perilaku `required`.
- `rows`: Mengonfigurasi atribut atau perilaku `rows`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
