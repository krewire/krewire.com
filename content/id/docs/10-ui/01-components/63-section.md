---
title: "Section"
description: "Pembungkus bagian halaman semantik selebar layar dengan token padding vertikal."
date: "2026-11-09"
---

# Section

Kategori: **Tata Letak Primitif**

Pembungkus bagian halaman semantik selebar layar dengan token padding vertikal.

---

## Contoh Penggunaan Dasar

Gunakan `<Section />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Section class="kustom-section">
  <!-- Konten slot atau properti -->
</Section>
```

---

## Properti & Atribut

- `content`: Mengonfigurasi atribut atau perilaku `content`.
- `desc`: Mengonfigurasi atribut atau perilaku `desc`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
