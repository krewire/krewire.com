---
title: "Kbd"
description: "Indikator visual tombol keyboard untuk pintasan dan tombol hotkey."
date: "2026-11-09"
---

# Kbd

Kategori: **Tampilan Data & Konten**

Indikator visual tombol keyboard untuk pintasan dan tombol hotkey.

---

## Contoh Penggunaan Dasar

Gunakan `<Kbd />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Kbd class="kustom-kbd">
  <!-- Konten slot atau properti -->
</Kbd>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
