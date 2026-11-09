---
title: "ButtonGroup"
description: "Kelompok tombol bersambung dengan border terpadu dan status aktif."
date: "2026-11-09"
---

# ButtonGroup

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/ButtonGroup.kiw`  
Sumber: `github.com/krewire/ui/components/ButtonGroup.kiw`

Kelompok tombol bersambung dengan border terpadu dan status aktif.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `ButtonGroup.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/ButtonGroup.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<ButtonGroup />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<ButtonGroup class="kustom-button-group">
  <!-- Konten slot atau properti -->
</ButtonGroup>
```

---

## 3. Properti & Atribut yang Didukung

- `ariaLabel`: Mengonfigurasi atribut atau perilaku `ariaLabel`.
- `item`: Mengonfigurasi atribut atau perilaku `item`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
