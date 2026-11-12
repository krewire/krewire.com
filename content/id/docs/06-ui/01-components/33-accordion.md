---
title: "Accordion"
description: "Panel konten lipat buka-tutup dengan aksesibilitas keyboard lengkap."
date: "2026-11-09"
---

# Accordion

Kategori: **Tampilan Data & Konten**

Panel konten lipat buka-tutup dengan aksesibilitas keyboard lengkap.

---

## Contoh Penggunaan Dasar

Gunakan `<Accordion />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Accordion class="kustom-accordion">
  <!-- Konten slot atau properti -->
</Accordion>
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
