---
title: "Text"
description: "Pembungkus tipografi untuk teks paragraf isi, teks pengantar, dan teks diredam."
date: "2026-11-09"
---

# Text

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Text.kiw`  
Sumber: `github.com/krewire/ui/components/Text.kiw`

Pembungkus tipografi untuk teks paragraf isi, teks pengantar, dan teks diredam.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Text.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Text.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Text />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Text class="kustom-text">
  <!-- Konten slot atau properti -->
</Text>
```

---

## 3. Properti & Atribut yang Didukung

- `color`: Mengonfigurasi atribut atau perilaku `color`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `weight`: Mengonfigurasi atribut atau perilaku `weight`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
