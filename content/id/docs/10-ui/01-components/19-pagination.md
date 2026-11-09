---
title: "Pagination"
description: "Bilah navigasi halaman dengan tombol sebelum, sesudah, dan tautan angka."
date: "2026-11-09"
---

# Pagination

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Pagination.kiw`  
Sumber: `github.com/krewire/ui/components/Pagination.kiw`

Bilah navigasi halaman dengan tombol sebelum, sesudah, dan tautan angka.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Pagination.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Pagination.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Pagination />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Pagination class="kustom-pagination">
  <!-- Konten slot atau properti -->
</Pagination>
```

---

## 3. Properti & Atribut yang Didukung

- `currentPage`: Mengonfigurasi atribut atau perilaku `currentPage`.
- `hasNext`: Mengonfigurasi atribut atau perilaku `hasNext`.
- `hasPrev`: Mengonfigurasi atribut atau perilaku `hasPrev`.
- `nextUrl`: Mengonfigurasi atribut atau perilaku `nextUrl`.
- `p`: Mengonfigurasi atribut atau perilaku `p`.
- `pages`: Mengonfigurasi atribut atau perilaku `pages`.
- `prevUrl`: Mengonfigurasi atribut atau perilaku `prevUrl`.
- `totalPages`: Mengonfigurasi atribut atau perilaku `totalPages`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
