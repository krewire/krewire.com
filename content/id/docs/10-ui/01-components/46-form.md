---
title: "Form"
description: "Pembungkus formulir aksesibel dengan method spoofing dan status pengiriman."
date: "2026-11-09"
---

# Form

Kategori: **Formulir & Masukan**  
Berkas Komponen: `components/Form.kiw`  
Sumber: `github.com/krewire/ui/components/Form.kiw`

Pembungkus formulir aksesibel dengan method spoofing dan status pengiriman.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Form.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Form.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Form />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Form class="kustom-form">
  <!-- Konten slot atau properti -->
</Form>
```

---

## 3. Properti & Atribut yang Didukung

- `action`: Mengonfigurasi atribut atau perilaku `action`.
- `enctype`: Mengonfigurasi atribut atau perilaku `enctype`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `method`: Mengonfigurasi atribut atau perilaku `method`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
