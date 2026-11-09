---
title: "Callout"
description: "Kotak catatan berbingkai untuk tips, peringatan, dan catatan dokumentasi."
date: "2026-11-09"
---

# Callout

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Callout.kiw`  
Sumber: `github.com/krewire/ui/components/Callout.kiw`

Kotak catatan berbingkai untuk tips, peringatan, dan catatan dokumentasi.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Callout.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Callout.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Callout />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Callout class="kustom-callout">
  <!-- Konten slot atau properti -->
</Callout>
```

---

## 3. Properti & Atribut yang Didukung

- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `type`: Mengonfigurasi atribut atau perilaku `type`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
