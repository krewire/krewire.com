---
title: "Checkbox"
description: "Kotak centang aksesibel berkontras tinggi dengan gaya pop-border kaku."
date: "2026-11-09"
---

# Checkbox

Kategori: **Formulir & Masukan**  
Berkas Komponen: `components/Checkbox.kiw`  
Sumber: `github.com/krewire/ui/components/Checkbox.kiw`

Kotak centang aksesibel berkontras tinggi dengan gaya pop-border kaku.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Checkbox.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Checkbox.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Checkbox />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Checkbox class="kustom-checkbox">
  <!-- Konten slot atau properti -->
</Checkbox>
```

---

## 3. Properti & Atribut yang Didukung

- `checked`: Mengonfigurasi atribut atau perilaku `checked`.
- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `required`: Mengonfigurasi atribut atau perilaku `required`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
