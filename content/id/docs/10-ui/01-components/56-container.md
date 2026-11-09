---
title: "Container"
description: "Kontainer halaman terpusat responsif dengan batas lebar maksimal standar."
date: "2026-11-09"
---

# Container

Kategori: **Tata Letak Primitif**  
Berkas Komponen: `components/Container.kiw`  
Sumber: `github.com/krewire/ui/components/Container.kiw`

Kontainer halaman terpusat responsif dengan batas lebar maksimal standar.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Container.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Container.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Container />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Container class="kustom-container">
  <!-- Konten slot atau properti -->
</Container>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
