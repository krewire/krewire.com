---
title: "Skeleton"
description: "Placeholder kilau teranimasi saat memuat konten asinkron."
date: "2026-11-09"
---

# Skeleton

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Skeleton.kiw`  
Sumber: `github.com/krewire/ui/components/Skeleton.kiw`

Placeholder kilau teranimasi saat memuat konten asinkron.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Skeleton.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Skeleton.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Skeleton />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Skeleton class="kustom-skeleton">
  <!-- Konten slot atau properti -->
</Skeleton>
```

---

## 3. Properti & Atribut yang Didukung

- `height`: Mengonfigurasi atribut atau perilaku `height`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.
- `width`: Mengonfigurasi atribut atau perilaku `width`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
