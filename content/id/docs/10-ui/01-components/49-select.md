---
title: "Select"
description: "Kolom pilihan dropdown select bergaya native."
date: "2026-11-09"
---

# Select

Kategori: **Formulir & Masukan**  
Berkas Komponen: `components/Select.kiw`  
Sumber: `github.com/krewire/ui/components/Select.kiw`

Kolom pilihan dropdown select bergaya native.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Select.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Select.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Select />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Select class="kustom-select">
  <!-- Konten slot atau properti -->
</Select>
```

---

## 3. Properti & Atribut yang Didukung

- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `error`: Mengonfigurasi atribut atau perilaku `error`.
- `help`: Mengonfigurasi atribut atau perilaku `help`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `required`: Mengonfigurasi atribut atau perilaku `required`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
