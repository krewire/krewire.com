---
title: "NavMenu"
description: "Menu navigasi horizontal dengan menu dropdown interaktif."
date: "2026-11-09"
---

# NavMenu

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/NavMenu.kiw`  
Sumber: `github.com/krewire/ui/components/NavMenu.kiw`

Menu navigasi horizontal dengan menu dropdown interaktif.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `NavMenu.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/NavMenu.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<NavMenu />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<NavMenu class="kustom-nav-menu">
  <!-- Konten slot atau properti -->
</NavMenu>
```

---

## 3. Properti & Atribut yang Didukung

- `dropdown`: Mengonfigurasi atribut atau perilaku `dropdown`.
- `links`: Mengonfigurasi atribut atau perilaku `links`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
