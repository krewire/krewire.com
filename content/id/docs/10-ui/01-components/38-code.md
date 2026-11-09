---
title: "Code"
description: "Cuplikan kode monospace sebaris dengan style kontras tinggi."
date: "2026-11-09"
---

# Code

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Code.kiw`  
Sumber: `github.com/krewire/ui/components/Code.kiw`

Cuplikan kode monospace sebaris dengan style kontras tinggi.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Code.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Code.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Code />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Code class="kustom-code">
  <!-- Konten slot atau properti -->
</Code>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
