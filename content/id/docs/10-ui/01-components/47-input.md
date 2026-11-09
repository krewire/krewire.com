---
title: "Input"
description: "Kolom input teks standar dengan label mengambang, status error, dan teks bantuan."
date: "2026-11-09"
---

# Input

Kategori: **Formulir & Masukan**  
Berkas Komponen: `components/Input.kiw`  
Sumber: `github.com/krewire/ui/components/Input.kiw`

Kolom input teks standar dengan label mengambang, status error, dan teks bantuan.

---

## Pratinjau Langsung

<div data-component-preview="Input" data-slug="input"></div>

Bingkai di atas memuat `Input.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Input.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Input.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Input />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Input class="kustom-input">
  <!-- Konten slot atau properti -->
</Input>
```

---

## 3. Properti & Atribut yang Didukung

- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `error`: Mengonfigurasi atribut atau perilaku `error`.
- `help`: Mengonfigurasi atribut atau perilaku `help`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `placeholder`: Mengonfigurasi atribut atau perilaku `placeholder`.
- `required`: Mengonfigurasi atribut atau perilaku `required`.
- `type`: Mengonfigurasi atribut atau perilaku `type`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
