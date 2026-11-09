---
title: "Head"
description: "Penyuntik tag head standar untuk metadata, OpenGraph, dan favicon."
date: "2026-11-09"
---

# Head

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Head.kiw`  
Sumber: `github.com/krewire/ui/components/Head.kiw`

Penyuntik tag head standar untuk metadata, OpenGraph, dan favicon.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Head.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Head.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Head />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Head class="kustom-head">
  <!-- Konten slot atau properti -->
</Head>
```

---

## 3. Properti & Atribut yang Didukung

- `assetversion`: Mengonfigurasi atribut atau perilaku `assetversion`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
