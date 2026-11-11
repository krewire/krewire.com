---
title: "Dropdown"
description: "Menu dropdown tanpa dependensi yang dibangun di atas elemen native details dan summary."
date: "2026-11-09"
---

# Dropdown

Kategori: **Aksi & Tombol**

Menu dropdown tanpa dependensi yang dibangun di atas elemen native details dan summary.

---

## Contoh Penggunaan Dasar

Gunakan `<Dropdown />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Dropdown class="kustom-dropdown">
  <!-- Konten slot atau properti -->
</Dropdown>
```

---

## Properti & Atribut

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `icon`: Ikon opsional yang dirender di depan label pada tombol pemicu default.
- `item`: Mengonfigurasi atribut atau perilaku `item`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `trigger`: Mengonfigurasi atribut atau perilaku `trigger`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
