---
title: "Terminal"
description: "Pratinjau terminal interaktif baris tunggal dengan tombol salin perintah."
date: "2026-11-09"
---

# Terminal

Kategori: **Tampilan Data & Konten**

Pratinjau terminal interaktif baris tunggal dengan tombol salin perintah.

---

## Contoh Penggunaan Dasar

Gunakan `<Terminal />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Terminal class="kustom-terminal">
  <!-- Konten slot atau properti -->
</Terminal>
```

---

## Properti & Atribut

- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `command`: Mengonfigurasi atribut atau perilaku `command`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `prompt`: Mengonfigurasi atribut atau perilaku `prompt`.
- `promptColor`: Mengonfigurasi atribut atau perilaku `promptColor`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
