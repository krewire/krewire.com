---
title: "Dropdown"
description: "Menu dropdown tanpa dependensi yang dibangun di atas elemen native details dan summary."
date: "2026-11-09"
---

# Dropdown

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/Dropdown.kiw`  
Sumber: `github.com/krewire/ui/components/Dropdown.kiw`

Menu dropdown tanpa dependensi yang dibangun di atas elemen native details dan summary.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Dropdown.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Dropdown.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Dropdown />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Dropdown class="kustom-dropdown">
  <!-- Konten slot atau properti -->
</Dropdown>
```

---

## 3. Properti & Atribut yang Didukung

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `item`: Mengonfigurasi atribut atau perilaku `item`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `trigger`: Mengonfigurasi atribut atau perilaku `trigger`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
