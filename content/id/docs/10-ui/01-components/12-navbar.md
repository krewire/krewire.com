---
title: "Navbar"
description: "Header aplikasi utama dengan logo brand, tautan navigasi, dan pemicu drawer."
date: "2026-11-09"
---

# Navbar

Kategori: **Navigasi & Struktur**

Header aplikasi utama dengan logo brand, tautan navigasi, dan pemicu drawer.

---

## Contoh Penggunaan Dasar

Gunakan `<Navbar />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Navbar class="kustom-navbar">
  <!-- Konten slot atau properti -->
</Navbar>
```

---

## Properti & Atribut

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

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
