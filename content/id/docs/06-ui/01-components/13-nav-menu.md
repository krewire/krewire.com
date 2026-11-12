---
title: "NavMenu"
description: "Menu navigasi horizontal dengan menu dropdown interaktif."
date: "2026-11-09"
---

# NavMenu

Kategori: **Navigasi & Struktur**

Menu navigasi horizontal dengan menu dropdown interaktif.

---

## Contoh Penggunaan Dasar

Gunakan `<NavMenu />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<NavMenu class="kustom-nav-menu">
  <!-- Konten slot atau properti -->
</NavMenu>
```

---

## Properti & Atribut

- `dropdown`: Mengonfigurasi atribut atau perilaku `dropdown`.
- `links`: Mengonfigurasi atribut atau perilaku `links`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
