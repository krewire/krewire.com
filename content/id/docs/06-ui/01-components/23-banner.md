---
title: "Banner"
description: "Bilah pengumuman tingkat atas yang mencolok dengan tombol penutup."
date: "2026-11-09"
---

# Banner

Kategori: **Umpan Balik & Status**

Bilah pengumuman tingkat atas yang mencolok dengan tombol penutup.

---

## Contoh Penggunaan Dasar

Gunakan `<Banner />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Banner class="kustom-banner">
  <!-- Konten slot atau properti -->
</Banner>
```

---

## Properti & Atribut

- `action`: Mengonfigurasi atribut atau perilaku `action`.
- `badge`: Mengonfigurasi atribut atau perilaku `badge`.
- `cta`: Mengonfigurasi atribut atau perilaku `cta`.
- `dismissible`: Mengonfigurasi atribut atau perilaku `dismissible`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `message`: Mengonfigurasi atribut atau perilaku `message`.
- `target`: Mengonfigurasi atribut atau perilaku `target`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
