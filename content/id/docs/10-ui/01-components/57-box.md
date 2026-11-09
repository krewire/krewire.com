---
title: "Box"
description: "Kotak tata letak serbaguna dengan dukungan border, padding, dan bayangan pop-brutalist."
date: "2026-11-09"
---

# Box

Kategori: **Tata Letak Primitif**  
Berkas Komponen: `components/Box.kiw`  
Sumber: `github.com/krewire/ui/components/Box.kiw`

Kotak tata letak serbaguna dengan dukungan border, padding, dan bayangan pop-brutalist.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Box.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Box.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Box />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Box class="kustom-box">
  <!-- Konten slot atau properti -->
</Box>
```

---

## 3. Properti & Atribut yang Didukung

- `border`: Mengonfigurasi atribut atau perilaku `border`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `rounded`: Mengonfigurasi atribut atau perilaku `rounded`.
- `shadow`: Mengonfigurasi atribut atau perilaku `shadow`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
