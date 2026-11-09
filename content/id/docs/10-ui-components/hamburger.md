---
title: "Hamburger"
description: "Tombol menu hamburger 3 baris teranimasi untuk membuka drawer mobile."
date: "2026-11-09"
---

# Hamburger

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Hamburger.kiw`  
Sumber: `github.com/krewire/ui/components/Hamburger.kiw`

Tombol menu hamburger 3 baris teranimasi untuk membuka drawer mobile.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Hamburger.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Hamburger.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Hamburger />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Hamburger class="kustom-hamburger">
  <!-- Konten slot atau properti -->
</Hamburger>
```

---

## 3. Properti & Atribut yang Didukung

- `Class`: Mengonfigurasi atribut atau perilaku `Class`.
- `Controls`: Mengonfigurasi atribut atau perilaku `Controls`.
- `DataDrawerToggle`: Mengonfigurasi atribut atau perilaku `DataDrawerToggle`.
- `Expanded`: Mengonfigurasi atribut atau perilaku `Expanded`.
- `ID`: Mengonfigurasi atribut atau perilaku `ID`.
- `Label`: Mengonfigurasi atribut atau perilaku `Label`.
- `controls`: Mengonfigurasi atribut atau perilaku `controls`.
- `dataDrawerToggle`: Mengonfigurasi atribut atau perilaku `dataDrawerToggle`.
- `expanded`: Mengonfigurasi atribut atau perilaku `expanded`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
