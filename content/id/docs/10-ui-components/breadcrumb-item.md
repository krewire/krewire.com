---
title: "BreadcrumbItem"
description: "Segmen individual di dalam jejak navigasi breadcrumb."
date: "2026-11-09"
---

# BreadcrumbItem

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/BreadcrumbItem.kiw`  
Sumber: `github.com/krewire/ui/components/BreadcrumbItem.kiw`

Segmen individual di dalam jejak navigasi breadcrumb.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `BreadcrumbItem.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/BreadcrumbItem.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<BreadcrumbItem />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<BreadcrumbItem class="kustom-breadcrumb-item">
  <!-- Konten slot atau properti -->
</BreadcrumbItem>
```

---

## 3. Properti & Atribut yang Didukung

- `current`: Mengonfigurasi atribut atau perilaku `current`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `target`: Mengonfigurasi atribut atau perilaku `target`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
