---
title: "HStack"
description: "Kontainer flex horizontal dengan jarak seragam antar elemen anak."
date: "2026-11-09"
---

# HStack

Kategori: **Tata Letak Primitif**  
Berkas Komponen: `components/HStack.kiw`  
Sumber: `github.com/krewire/ui/components/HStack.kiw`

Kontainer flex horizontal dengan jarak seragam antar elemen anak.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `HStack.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/HStack.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<HStack />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<HStack class="kustom-h-stack">
  <!-- Konten slot atau properti -->
</HStack>
```

---

## 3. Properti & Atribut yang Didukung

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `gap`: Mengonfigurasi atribut atau perilaku `gap`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `justify`: Mengonfigurasi atribut atau perilaku `justify`.
- `noWrap`: Mengonfigurasi atribut atau perilaku `noWrap`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
