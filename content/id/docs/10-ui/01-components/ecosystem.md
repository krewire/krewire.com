---
title: "Ecosystem"
description: "Komponen grid interaktif untuk menampilkan arsitektur 10 repositori Krewire."
date: "2026-11-09"
---

# Ecosystem

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Ecosystem.kiw`  
Sumber: `github.com/krewire/ui/components/Ecosystem.kiw`

Komponen grid interaktif untuk menampilkan arsitektur 10 repositori Krewire.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Ecosystem.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Ecosystem.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Ecosystem />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Ecosystem class="kustom-ecosystem">
  <!-- Konten slot atau properti -->
</Ecosystem>
```

---

## 3. Properti & Atribut yang Didukung

- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `header`: Mengonfigurasi atribut atau perilaku `header`.
- `lead`: Mengonfigurasi atribut atau perilaku `lead`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
