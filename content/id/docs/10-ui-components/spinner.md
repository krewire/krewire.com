---
title: "Spinner"
description: "Pemutar loader SVG teranimasi yang aksesibel untuk operasi yang sedang berjalan."
date: "2026-11-09"
---

# Spinner

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Spinner.kiw`  
Sumber: `github.com/krewire/ui/components/Spinner.kiw`

Pemutar loader SVG teranimasi yang aksesibel untuk operasi yang sedang berjalan.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Spinner.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Spinner.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Spinner />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Spinner class="kustom-spinner">
  <!-- Konten slot atau properti -->
</Spinner>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
