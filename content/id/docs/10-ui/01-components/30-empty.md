---
title: "Empty"
description: "Placeholder status kosong dengan ikon ilustrasi, pesan, dan tombol aksi."
date: "2026-11-09"
---

# Empty

Kategori: **Umpan Balik & Status**

Placeholder status kosong dengan ikon ilustrasi, pesan, dan tombol aksi.

---

## Contoh Penggunaan Dasar

Gunakan `<Empty />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Empty class="kustom-empty">
  <!-- Konten slot atau properti -->
</Empty>
```

---

## Properti & Atribut

- `action`: Mengonfigurasi atribut atau perilaku `action`.
- `actionText`: Mengonfigurasi atribut atau perilaku `actionText`.
- `description`: Mengonfigurasi atribut atau perilaku `description`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
