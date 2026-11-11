---
title: "SearchInput"
description: "Kolom input pencarian khusus dengan ikon cari dan tombol hapus instan."
date: "2026-11-09"
---

# SearchInput

Kategori: **Formulir & Masukan**

Kolom input pencarian khusus dengan ikon cari dan tombol hapus instan.

---

## Contoh Penggunaan Dasar

Gunakan `<SearchInput />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<SearchInput class="kustom-search-input">
  <!-- Konten slot atau properti -->
</SearchInput>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `name`: Mengonfigurasi atribut atau perilaku `name`.
- `placeholder`: Mengonfigurasi atribut atau perilaku `placeholder`.
- `shortcut`: Mengonfigurasi atribut atau perilaku `shortcut`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
