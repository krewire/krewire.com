---
title: "Brand"
description: "Logo brand ekosistem lengkap dengan lambang ikon dan badge versi."
date: "2026-11-09"
---

# Brand

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Brand.kiw`  
Sumber: `github.com/krewire/ui/components/Brand.kiw`

Logo brand ekosistem lengkap dengan lambang ikon dan badge versi.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Brand.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Brand.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Brand />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Brand class="kustom-brand">
  <!-- Konten slot atau properti -->
</Brand>
```

---

## 3. Properti & Atribut yang Didukung

- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `badgeI18n`: Mengonfigurasi atribut atau perilaku `badgeI18n`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `version`: Mengonfigurasi atribut atau perilaku `version`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
