---
title: "ButtonGroup"
description: "Kelompok tombol bersambung dengan border terpadu dan status aktif."
date: "2026-11-09"
---

# ButtonGroup

Kategori: **Aksi & Tombol**

Kelompok tombol bersambung dengan border terpadu dan status aktif.

---

## Contoh Penggunaan Dasar

Gunakan `<ButtonGroup />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<ButtonGroup class="kustom-button-group">
  <!-- Konten slot atau properti -->
</ButtonGroup>
```

---

## Properti & Atribut

- `ariaLabel`: Mengonfigurasi atribut atau perilaku `ariaLabel`.
- `item`: Mengonfigurasi atribut atau perilaku `item`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
