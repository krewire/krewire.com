---
title: "Grid"
description: "Kontainer CSS Grid dengan pengaturan jumlah kolom dan token jarak."
date: "2026-11-09"
---

# Grid

Kategori: **Tata Letak Primitif**  
Berkas Komponen: `components/Grid.kiw`  
Sumber: `github.com/krewire/ui/components/Grid.kiw`

Kontainer CSS Grid dengan pengaturan jumlah kolom dan token jarak.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Grid.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Grid.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Grid />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Grid class="kustom-grid">
  <!-- Konten slot atau properti -->
</Grid>
```

---

## 3. Properti & Atribut yang Didukung

- `cols`: Mengonfigurasi atribut atau perilaku `cols`.
- `gap`: Mengonfigurasi atribut atau perilaku `gap`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
