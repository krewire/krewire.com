---
title: "Tabs"
description: "Pengalih konten bertab yang aksesibel dengan dukungan status aktif persisten."
date: "2026-11-09"
---

# Tabs

Kategori: **Navigasi & Struktur**  
Berkas Komponen: `components/Tabs.kiw`  
Sumber: `github.com/krewire/ui/components/Tabs.kiw`

Pengalih konten bertab yang aksesibel dengan dukungan status aktif persisten.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Tabs.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Tabs.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Tabs />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Tabs class="kustom-tabs">
  <!-- Konten slot atau properti -->
</Tabs>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `panels`: Mengonfigurasi atribut atau perilaku `panels`.
- `tabs`: Mengonfigurasi atribut atau perilaku `tabs`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
