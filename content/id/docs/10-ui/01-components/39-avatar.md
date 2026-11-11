---
title: "Avatar"
description: "Avatar profil pengguna yang menampilkan gambar dengan inisial cadangan otomatis."
date: "2026-11-09"
---

# Avatar

Kategori: **Tampilan Data & Konten**

Avatar profil pengguna yang menampilkan gambar dengan inisial cadangan otomatis.

---

## Contoh Penggunaan Dasar

Gunakan `<Avatar />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Avatar class="kustom-avatar">
  <!-- Konten slot atau properti -->
</Avatar>
```

---

## Properti & Atribut

- `alt`: Mengonfigurasi atribut atau perilaku `alt`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `initials`: Mengonfigurasi atribut atau perilaku `initials`.
- `shape`: Mengonfigurasi atribut atau perilaku `shape`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `src`: Mengonfigurasi atribut atau perilaku `src`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
