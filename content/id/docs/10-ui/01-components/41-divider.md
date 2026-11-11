---
title: "Divider"
description: "Garis pemisah horizontal atau vertikal dengan label teks opsional."
date: "2026-11-09"
---

# Divider

Kategori: **Tampilan Data & Konten**

Garis pemisah horizontal atau vertikal dengan label teks opsional.

---

## Contoh Penggunaan Dasar

Gunakan `<Divider />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Divider class="kustom-divider">
  <!-- Konten slot atau properti -->
</Divider>
```

---

## Properti & Atribut

- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `vertical`: Mengonfigurasi atribut atau perilaku `vertical`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
