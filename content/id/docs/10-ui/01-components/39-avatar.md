---
title: "Avatar"
description: "Avatar profil pengguna yang menampilkan gambar dengan inisial cadangan otomatis."
date: "2026-11-09"
---

# Avatar

Kategori: **Tampilan Data & Konten**  
Berkas Komponen: `components/Avatar.kiw`  
Sumber: `github.com/krewire/ui/components/Avatar.kiw`

Avatar profil pengguna yang menampilkan gambar dengan inisial cadangan otomatis.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Avatar.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Avatar.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Avatar />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Avatar class="kustom-avatar">
  <!-- Konten slot atau properti -->
</Avatar>
```

---

## 3. Properti & Atribut yang Didukung

- `alt`: Mengonfigurasi atribut atau perilaku `alt`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `initials`: Mengonfigurasi atribut atau perilaku `initials`.
- `shape`: Mengonfigurasi atribut atau perilaku `shape`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `src`: Mengonfigurasi atribut atau perilaku `src`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
