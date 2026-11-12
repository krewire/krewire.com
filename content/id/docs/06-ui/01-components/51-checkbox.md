---
title: "Checkbox"
description: "Kotak centang aksesibel berkontras tinggi dengan gaya pop-border kaku."
date: "2026-11-09"
---

# Checkbox

Kategori: **Formulir & Masukan**

Kotak centang aksesibel berkontras tinggi dengan gaya pop-border kaku.

---

## Contoh Penggunaan Dasar

Gunakan `<Checkbox />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Checkbox class="kustom-checkbox">
  <!-- Konten slot atau properti -->
</Checkbox>
```

---

## Properti & Atribut

- `checked`: Mengonfigurasi atribut atau perilaku `checked`.
- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `required`: Mengonfigurasi atribut atau perilaku `required`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
