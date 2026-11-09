---
title: "LangSwitch"
description: "Tombol pengalih bahasa untuk aplikasi multi-bahasa internasional."
date: "2026-11-09"
---

# LangSwitch

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/LangSwitch.kiw`  
Sumber: `github.com/krewire/ui/components/LangSwitch.kiw`

Tombol pengalih bahasa untuk aplikasi multi-bahasa internasional.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `LangSwitch.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/LangSwitch.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<LangSwitch />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<LangSwitch class="kustom-lang-switch">
  <!-- Konten slot atau properti -->
</LangSwitch>
```

---

## 3. Properti & Atribut yang Didukung

- `code`: Mengonfigurasi atribut atau perilaku `code`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
