---
title: "Drawer"
description: "Panel geser mulus (*drawer*) untuk navigasi mobile dan panel samping."
date: "2026-11-09"
---

# Drawer

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Drawer.kiw`  
Sumber: `github.com/krewire/ui/components/Drawer.kiw`

Panel geser mulus (*drawer*) untuk navigasi mobile dan panel samping.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Drawer.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Drawer.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Drawer />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Drawer class="kustom-drawer">
  <!-- Konten slot atau properti -->
</Drawer>
```

---

## 3. Properti & Atribut yang Didukung

- `Body`: Mengonfigurasi atribut atau perilaku `Body`.
- `Class`: Mengonfigurasi atribut atau perilaku `Class`.
- `Footer`: Mengonfigurasi atribut atau perilaku `Footer`.
- `ID`: Mengonfigurasi atribut atau perilaku `ID`.
- `Side`: Mengonfigurasi atribut atau perilaku `Side`.
- `Title`: Mengonfigurasi atribut atau perilaku `Title`.
- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `footer`: Mengonfigurasi atribut atau perilaku `footer`.
- `head`: Mengonfigurasi atribut atau perilaku `head`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `side`: Mengonfigurasi atribut atau perilaku `side`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
