---
title: "Container"
description: "Kontainer halaman terpusat responsif dengan batas lebar maksimal standar."
date: "2026-11-09"
---

# Container

Kategori: **Tata Letak Primitif**

Kontainer halaman terpusat responsif dengan batas lebar maksimal standar.

---

## Contoh Penggunaan Dasar

Gunakan `<Container />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Container class="kustom-container">
  <!-- Konten slot atau properti -->
</Container>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
