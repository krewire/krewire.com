---
title: "Table"
description: "Tabel data responsif dengan gaya bergaris, bingkai, dan sorot baris."
date: "2026-11-09"
---

# Table

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Table.kiw`  
Sumber: `github.com/krewire/ui/components/Table.kiw`

Tabel data responsif dengan gaya bergaris, bingkai, dan sorot baris.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Table.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Table.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Table />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Table class="kustom-table">
  <!-- Konten slot atau properti -->
</Table>
```

---

## 3. Properti & Atribut yang Didukung

- `bordered`: Mengonfigurasi atribut atau perilaku `bordered`.
- `cell`: Mengonfigurasi atribut atau perilaku `cell`.
- `h`: Mengonfigurasi atribut atau perilaku `h`.
- `headers`: Mengonfigurasi atribut atau perilaku `headers`.
- `hover`: Mengonfigurasi atribut atau perilaku `hover`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `rows`: Mengonfigurasi atribut atau perilaku `rows`.
- `striped`: Mengonfigurasi atribut atau perilaku `striped`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
