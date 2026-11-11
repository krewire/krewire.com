---
title: "SideMenu"
description: "Daftar navigasi vertikal untuk bilah samping administratif."
date: "2026-11-09"
---

# SideMenu

Kategori: **Navigasi & Struktur**

Daftar navigasi vertikal untuk bilah samping administratif.

---

## Contoh Penggunaan Dasar

Gunakan `<SideMenu />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<SideMenu class="kustom-side-menu">
  <!-- Konten slot atau properti -->
</SideMenu>
```

---

## Properti & Atribut

- `g`: Mengonfigurasi atribut atau perilaku `g`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
