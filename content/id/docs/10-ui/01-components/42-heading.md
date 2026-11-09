---
title: "Heading"
description: "Pembungkus tipografi heading semantik dari h1 hingga h6."
date: "2026-11-09"
---

# Heading

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Heading.kiw`  
Sumber: `github.com/krewire/ui/components/Heading.kiw`

Pembungkus tipografi heading semantik dari h1 hingga h6.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Heading.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Heading.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Heading />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Heading class="kustom-heading">
  <!-- Konten slot atau properti -->
</Heading>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `level`: Mengonfigurasi atribut atau perilaku `level`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
