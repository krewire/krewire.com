---
title: "ThemeSwitch"
description: "Tombol pengalih tema gelap dan terang dengan penyimpanan otomatis di localStorage."
date: "2026-11-09"
---

# ThemeSwitch

Kategori: **Aksi & Tombol**

Tombol pengalih tema gelap dan terang dengan penyimpanan otomatis di localStorage.

---

## Contoh Penggunaan Dasar

Gunakan `<ThemeSwitch />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<ThemeSwitch class="kustom-theme-switch">
  <!-- Konten slot atau properti -->
</ThemeSwitch>
```

---

## Properti & Atribut

- Tidak ada properti khusus (menerima konten slot default atau elemen anak).

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
