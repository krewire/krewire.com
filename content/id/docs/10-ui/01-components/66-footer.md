---
title: "Footer"
description: "Footer situs multi-kolom komprehensif dengan brand, tautan navigasi, dan hak cipta."
date: "2026-11-09"
---

# Footer

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Footer.kiw`  
Sumber: `github.com/krewire/ui/components/Footer.kiw`

Footer situs multi-kolom komprehensif dengan brand, tautan navigasi, dan hak cipta.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Footer.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Footer.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Footer />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Footer class="kustom-footer">
  <!-- Konten slot atau properti -->
</Footer>
```

---

## 3. Properti & Atribut yang Didukung

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

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
