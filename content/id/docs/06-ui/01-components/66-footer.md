---
title: "Footer"
description: "Footer situs multi-kolom komprehensif dengan brand, tautan navigasi, dan hak cipta."
date: "2026-11-09"
---

# Footer

Kategori: **Navigasi & Struktur**

Footer situs multi-kolom komprehensif dengan brand, tautan navigasi, dan hak cipta.

---

## Contoh Penggunaan Dasar

Gunakan `<Footer />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Footer class="kustom-footer">
  <!-- Konten slot atau properti -->
</Footer>
```

---

## Properti & Atribut

- `bottom`: Mengonfigurasi atribut atau perilaku `bottom`.
- `brand`: Mengonfigurasi atribut atau perilaku `brand`.
- `brandBadge`: Mengonfigurasi atribut atau perilaku `brandBadge`.
- `brandIcon`: Mengonfigurasi atribut atau perilaku `brandIcon`.
- `brandName`: Mengonfigurasi atribut atau perilaku `brandName`.
- `brandTagline`: Mengonfigurasi atribut atau perilaku `brandTagline`.
- `columns`: Mengonfigurasi atribut atau perilaku `columns`.
- `copyrightText`: Mengonfigurasi atribut atau perilaku `copyrightText`.
- `gitHubURL`: Mengonfigurasi atribut atau perilaku `gitHubURL`.
- `showEcosystemFooter`: Mengonfigurasi atribut atau perilaku `showEcosystemFooter`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
