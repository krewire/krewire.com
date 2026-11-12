---
title: "Tabs"
description: "Pengalih konten bertab yang aksesibel dengan dukungan status aktif persisten."
date: "2026-11-09"
---

# Tabs

Kategori: **Navigasi & Struktur**

Pengalih konten bertab yang aksesibel dengan dukungan status aktif persisten.

---

## Pratinjau Langsung

<div data-component-preview="Tabs" data-slug="tabs"></div>

Bingkai di atas memuat `Tabs.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## Contoh Penggunaan Dasar

Gunakan `<Tabs />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Tabs class="kustom-tabs">
  <!-- Konten slot atau properti -->
</Tabs>
```

---

## Properti & Atribut

- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `panels`: Mengonfigurasi atribut atau perilaku `panels`.
- `tabs`: Mengonfigurasi atribut atau perilaku `tabs`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
