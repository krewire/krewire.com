---
title: "Progress"
description: "Indikator bilah progres kemajuan berstatus pasti maupun indeterminat."
date: "2026-11-09"
---

# Progress

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Progress.kiw`  
Sumber: `github.com/krewire/ui/components/Progress.kiw`

Indikator bilah progres kemajuan berstatus pasti maupun indeterminat.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Progress.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Progress.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Progress />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Progress class="kustom-progress">
  <!-- Konten slot atau properti -->
</Progress>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `showValue`: Mengonfigurasi atribut atau perilaku `showValue`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
