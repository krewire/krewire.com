---
title: "CodeWindow"
description: "Jendela pratinjau kode bergaya macOS dengan tombol traffic light dan tombol salin."
date: "2026-11-09"
---

# CodeWindow

Kategori: **Tampilan Data & Konten**

Jendela pratinjau kode bergaya macOS dengan tombol traffic light dan tombol salin.

---

## Contoh Penggunaan Dasar

Gunakan `<CodeWindow />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<CodeWindow class="kustom-code-window">
  <!-- Konten slot atau properti -->
</CodeWindow>
```

---

## Properti & Atribut

- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `code`: Mengonfigurasi atribut atau perilaku `code`.
- `context`: Mengonfigurasi atribut atau perilaku `context`.
- `get`: Mengonfigurasi atribut atau perilaku `get`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `json`: Mengonfigurasi atribut atau perilaku `json`.
- `map`: Mengonfigurasi atribut atau perilaku `map`.
- `new`: Mengonfigurasi atribut atau perilaku `new`.
- `run`: Mengonfigurasi atribut atau perilaku `run`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
