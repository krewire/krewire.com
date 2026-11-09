---
title: "Tooltip"
description: "Petunjuk tooltip informasi yang muncul saat disorot atau difokuskan."
date: "2026-11-09"
---

# Tooltip

Kategori: **Lapisan Mengambang & Dialog**  
Berkas Komponen: `components/Tooltip.kiw`  
Sumber: `github.com/krewire/ui/components/Tooltip.kiw`

Petunjuk tooltip informasi yang muncul saat disorot atau difokuskan.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Tooltip.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Tooltip.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Tooltip />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Tooltip class="kustom-tooltip">
  <!-- Konten slot atau properti -->
</Tooltip>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `placement`: Mengonfigurasi atribut atau perilaku `placement`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
