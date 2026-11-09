---
title: "List"
description: "Kontainer daftar berurutan dan tak berurutan dengan penataan rapi."
date: "2026-11-09"
---

# List

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/List.kiw`  
Sumber: `github.com/krewire/ui/components/List.kiw`

Kontainer daftar berurutan dan tak berurutan dengan penataan rapi.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `List.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/List.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<List />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<List class="kustom-list">
  <!-- Konten slot atau properti -->
</List>
```

---

## 3. Properti & Atribut yang Didukung

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `ordered`: Mengonfigurasi atribut atau perilaku `ordered`.
- `spacing`: Mengonfigurasi atribut atau perilaku `spacing`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
