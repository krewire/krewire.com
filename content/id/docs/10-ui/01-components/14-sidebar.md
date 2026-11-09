---
title: "Sidebar"
description: "Bilah navigasi samping collapsible untuk dokumentasi dan dasbor aplikasi."
date: "2026-11-09"
---

# Sidebar

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Sidebar.kiw`  
Sumber: `github.com/krewire/ui/components/Sidebar.kiw`

Bilah navigasi samping collapsible untuk dokumentasi dan dasbor aplikasi.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Sidebar.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Sidebar.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Sidebar />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Sidebar class="kustom-sidebar">
  <!-- Konten slot atau properti -->
</Sidebar>
```

---

## 3. Properti & Atribut yang Didukung

- `head`: Mengonfigurasi atribut atau perilaku `head`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `menu`: Mengonfigurasi atribut atau perilaku `menu`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
