---
title: "Section"
description: "Pembungkus bagian halaman semantik selebar layar dengan token padding vertikal."
date: "2026-11-09"
---

# Section

Kategori: **Tata Letak Primitif**  
Berkas Komponen: `components/Section.kiw`  
Sumber: `github.com/krewire/ui/components/Section.kiw`

Pembungkus bagian halaman semantik selebar layar dengan token padding vertikal.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Section.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Section.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Section />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Section class="kustom-section">
  <!-- Konten slot atau properti -->
</Section>
```

---

## 3. Properti & Atribut yang Didukung

- `content`: Mengonfigurasi atribut atau perilaku `content`.
- `desc`: Mengonfigurasi atribut atau perilaku `desc`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
