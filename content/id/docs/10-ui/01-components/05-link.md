---
title: "Link"
description: "Tautan hiperteks aksesibel dengan transisi garis bawah dan keamanan target."
date: "2026-11-09"
---

# Link

Kategori: **Aksi & Tombol**

Tautan hiperteks aksesibel dengan transisi garis bawah dan keamanan target.

---

## Contoh Penggunaan Dasar

Gunakan `<Link />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Link class="kustom-link">
  <!-- Konten slot atau properti -->
</Link>
```

---

## Properti & Atribut

- `external`: Mengonfigurasi atribut atau perilaku `external`.
- `href`: Mengonfigurasi atribut atau perilaku `href`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
