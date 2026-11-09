---
title: "Link"
description: "Tautan hiperteks aksesibel dengan transisi garis bawah dan keamanan target."
date: "2026-11-09"
---

# Link

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/Link.kiw`  
Sumber: `github.com/krewire/ui/components/Link.kiw`

Tautan hiperteks aksesibel dengan transisi garis bawah dan keamanan target.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Link.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Link.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Link />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Link class="kustom-link">
  <!-- Konten slot atau properti -->
</Link>
```

---

## 3. Properti & Atribut yang Didukung

- `external`: Mengonfigurasi atribut atau perilaku `external`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
