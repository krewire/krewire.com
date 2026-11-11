---
title: "Toggle"
description: "Sakelar toggle biner untuk pengubahan pengaturan secara langsung."
date: "2026-11-09"
---

# Toggle

Kategori: **Aksi & Tombol**

Sakelar toggle biner untuk pengubahan pengaturan secara langsung.

---

## Contoh Penggunaan Dasar

Gunakan `<Toggle />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Toggle class="kustom-toggle">
  <!-- Konten slot atau properti -->
</Toggle>
```

---

## Properti & Atribut

- `activeIcon`: Mengonfigurasi atribut atau perilaku `activeIcon`.
- `disabled`: Mengonfigurasi atribut atau perilaku `disabled`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `inactiveIcon`: Mengonfigurasi atribut atau perilaku `inactiveIcon`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `onClick`: Mengonfigurasi atribut atau perilaku `onClick`.
- `pressed`: Mengonfigurasi atribut atau perilaku `pressed`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
