---
title: "Radio"
description: "Tombol radio pilihan tunggal dengan bulatan bergaya pop-brutalist."
date: "2026-11-09"
---

# Radio

Kategori: **Formulir & Masukan**  
Berkas Komponen: `components/Radio.kiw`  
Sumber: `github.com/krewire/ui/components/Radio.kiw`

Tombol radio pilihan tunggal dengan bulatan bergaya pop-brutalist.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Radio.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Radio.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Radio />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Radio class="kustom-radio">
  <!-- Konten slot atau properti -->
</Radio>
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
