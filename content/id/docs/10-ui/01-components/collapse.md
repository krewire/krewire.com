---
title: "Collapse"
description: "Kontainer lipat tunggal untuk menampilkan atau menyembunyikan konten."
date: "2026-11-09"
---

# Collapse

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Collapse.kiw`  
Sumber: `github.com/krewire/ui/components/Collapse.kiw`

Kontainer lipat tunggal untuk menampilkan atau menyembunyikan konten.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Collapse.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Collapse.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Collapse />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Collapse class="kustom-collapse">
  <!-- Konten slot atau properti -->
</Collapse>
```

---

## 3. Properti & Atribut yang Didukung

- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `open`: Mengonfigurasi atribut atau perilaku `open`.
- `summary`: Mengonfigurasi atribut atau perilaku `summary`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
