---
title: "Terminal"
description: "Pratinjau terminal interaktif baris tunggal dengan tombol salin perintah."
date: "2026-11-09"
---

# Terminal

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Terminal.kiw`  
Sumber: `github.com/krewire/ui/components/Terminal.kiw`

Pratinjau terminal interaktif baris tunggal dengan tombol salin perintah.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Terminal.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Terminal.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Terminal />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Terminal class="kustom-terminal">
  <!-- Konten slot atau properti -->
</Terminal>
```

---

## 3. Properti & Atribut yang Didukung

- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `command`: Mengonfigurasi atribut atau perilaku `command`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `prompt`: Mengonfigurasi atribut atau perilaku `prompt`.
- `promptColor`: Mengonfigurasi atribut atau perilaku `promptColor`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
