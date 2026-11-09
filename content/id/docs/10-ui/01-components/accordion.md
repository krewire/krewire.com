---
title: "Accordion"
description: "Panel konten lipat buka-tutup dengan aksesibilitas keyboard lengkap."
date: "2026-11-09"
---

# Accordion

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Accordion.kiw`  
Sumber: `github.com/krewire/ui/components/Accordion.kiw`

Panel konten lipat buka-tutup dengan aksesibilitas keyboard lengkap.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Accordion.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Accordion.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Accordion />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Accordion class="kustom-accordion">
  <!-- Konten slot atau properti -->
</Accordion>
```

---

## 3. Properti & Atribut yang Didukung

- `item`: Mengonfigurasi atribut atau perilaku `item`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
