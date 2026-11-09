---
title: "Breadcrumb"
description: "Navigasi jejak hierarkis yang menunjukkan lokasi halaman saat ini."
date: "2026-11-09"
---

# Breadcrumb

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Breadcrumb.kiw`  
Sumber: `github.com/krewire/ui/components/Breadcrumb.kiw`

Navigasi jejak hierarkis yang menunjukkan lokasi halaman saat ini.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Breadcrumb.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Breadcrumb.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Breadcrumb />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Breadcrumb class="kustom-breadcrumb">
  <!-- Konten slot atau properti -->
</Breadcrumb>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `item`: Mengonfigurasi atribut atau perilaku `item`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
