---
title: "Kbd"
description: "Indikator visual tombol keyboard untuk pintasan dan tombol hotkey."
date: "2026-11-09"
---

# Kbd

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Kbd.kiw`  
Sumber: `github.com/krewire/ui/components/Kbd.kiw`

Indikator visual tombol keyboard untuk pintasan dan tombol hotkey.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Kbd.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Kbd.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Kbd />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Kbd class="kustom-kbd">
  <!-- Konten slot atau properti -->
</Kbd>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
