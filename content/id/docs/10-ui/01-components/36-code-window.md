---
title: "CodeWindow"
description: "Jendela pratinjau kode bergaya macOS dengan tombol traffic light dan tombol salin."
date: "2026-11-09"
---

# CodeWindow

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/CodeWindow.kiw`  
Sumber: `github.com/krewire/ui/components/CodeWindow.kiw`

Jendela pratinjau kode bergaya macOS dengan tombol traffic light dan tombol salin.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `CodeWindow.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/CodeWindow.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<CodeWindow />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<CodeWindow class="kustom-code-window">
  <!-- Konten slot atau properti -->
</CodeWindow>
```

---

## 3. Properti & Atribut yang Didukung

- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `code`: Mengonfigurasi atribut atau perilaku `code`.
- `context`: Mengonfigurasi atribut atau perilaku `context`.
- `get`: Mengonfigurasi atribut atau perilaku `get`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `json`: Mengonfigurasi atribut atau perilaku `json`.
- `map`: Mengonfigurasi atribut atau perilaku `map`.
- `new`: Mengonfigurasi atribut atau perilaku `new`.
- `run`: Mengonfigurasi atribut atau perilaku `run`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
