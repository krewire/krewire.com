---
title: "Timeline"
description: "Garis waktu kronologis vertikal untuk jejak aktivitas dan riwayat peristiwa."
date: "2026-11-09"
---

# Timeline

Kategori: **Tampilan Data & Konten**

Garis waktu kronologis vertikal untuk jejak aktivitas dan riwayat peristiwa.

---

## Contoh Penggunaan Dasar

Gunakan `<Timeline />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Timeline class="kustom-timeline">
  <!-- Konten slot atau properti -->
</Timeline>
```

---

## Properti & Atribut

- `item`: Mengonfigurasi atribut atau perilaku `item`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
