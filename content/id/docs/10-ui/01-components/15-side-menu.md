---
title: "SideMenu"
description: "Daftar navigasi vertikal untuk bilah samping administratif."
date: "2026-11-09"
---

# SideMenu

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/SideMenu.kiw`  
Sumber: `github.com/krewire/ui/components/SideMenu.kiw`

Daftar navigasi vertikal untuk bilah samping administratif.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `SideMenu.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/SideMenu.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<SideMenu />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<SideMenu class="kustom-side-menu">
  <!-- Konten slot atau properti -->
</SideMenu>
```

---

## 3. Properti & Atribut yang Didukung

- `g`: Mengonfigurasi atribut atau perilaku `g`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
