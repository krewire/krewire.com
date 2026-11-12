---
title: "Toast"
description: "Notifikasi mengambang ringan (*toast*) dengan pewaktu tutup otomatis."
date: "2026-11-09"
---

# Toast

Kategori: **Umpan Balik & Status**

Notifikasi mengambang ringan (*toast*) dengan pewaktu tutup otomatis.

---

## Contoh Penggunaan Dasar

Gunakan `<Toast />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Toast class="kustom-toast">
  <!-- Konten slot atau properti -->
</Toast>
```

---

## Properti & Atribut

- `dismissible`: Mengonfigurasi atribut atau perilaku `dismissible`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `message`: Mengonfigurasi atribut atau perilaku `message`.
- `position`: Mengonfigurasi atribut atau perilaku `position`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
