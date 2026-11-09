---
title: "ThemeSwitch"
description: "Tombol pengalih tema gelap dan terang dengan penyimpanan otomatis di localStorage."
date: "2026-11-09"
---

# ThemeSwitch

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/ThemeSwitch.kiw`  
Sumber: `github.com/krewire/ui/components/ThemeSwitch.kiw`

Tombol pengalih tema gelap dan terang dengan penyimpanan otomatis di localStorage.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `ThemeSwitch.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/ThemeSwitch.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<ThemeSwitch />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<ThemeSwitch class="kustom-theme-switch">
  <!-- Konten slot atau properti -->
</ThemeSwitch>
```

---

## 3. Properti & Atribut yang Didukung

- Tidak ada properti khusus (menerima konten slot default atau elemen anak).

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
