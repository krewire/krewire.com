---
title: "Skeleton"
description: "Placeholder kilau teranimasi saat memuat konten asinkron."
date: "2026-11-09"
---

# Skeleton

Kategori: **Umpan Balik & Status**

Placeholder kilau teranimasi saat memuat konten asinkron.

---

## Contoh Penggunaan Dasar

Gunakan `<Skeleton />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Skeleton class="kustom-skeleton">
  <!-- Konten slot atau properti -->
</Skeleton>
```

---

## Properti & Atribut

- `height`: Mengonfigurasi atribut atau perilaku `height`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.
- `width`: Mengonfigurasi atribut atau perilaku `width`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
