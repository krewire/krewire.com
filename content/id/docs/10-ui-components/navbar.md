---
title: "Navbar"
description: "Header aplikasi utama dengan logo brand, tautan navigasi, dan pemicu drawer."
date: "2026-11-09"
---

# Navbar

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Navbar.kiw`  
Sumber: `github.com/krewire/ui/components/Navbar.kiw`

Header aplikasi utama dengan logo brand, tautan navigasi, dan pemicu drawer.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Navbar.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Navbar.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Navbar />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Navbar class="kustom-navbar">
  <!-- Konten slot atau properti -->
</Navbar>
```

---

## 3. Properti & Atribut yang Didukung

- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `brand`: Mengonfigurasi atribut atau perilaku `brand`.
- `brandBadge`: Mengonfigurasi atribut atau perilaku `brandBadge`.
- `brandBadgeI18n`: Mengonfigurasi atribut atau perilaku `brandBadgeI18n`.
- `brandHref`: Mengonfigurasi atribut atau perilaku `brandHref`.
- `brandIcon`: Mengonfigurasi atribut atau perilaku `brandIcon`.
- `brandName`: Mengonfigurasi atribut atau perilaku `brandName`.
- `drawer`: Mengonfigurasi atribut atau perilaku `drawer`.
- `drawerBody`: Mengonfigurasi atribut atau perilaku `drawerBody`.
- `gitHubURL`: Mengonfigurasi atribut atau perilaku `gitHubURL`.
- `showLangToggle`: Mengonfigurasi atribut atau perilaku `showLangToggle`.
- `showSidebarToggle`: Mengonfigurasi atribut atau perilaku `showSidebarToggle`.
- `version`: Mengonfigurasi atribut atau perilaku `version`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
