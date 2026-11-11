---
title: "Alert"
description: "Banner pemberitahuan kontekstual dengan status info, success, warning, dan error."
date: "2026-11-09"
---

# Alert

Kategori: **Umpan Balik & Status**

Banner pemberitahuan kontekstual dengan status info, success, warning, dan error.

---

## Pratinjau Langsung

<div data-component-preview="Alert" data-slug="alert"></div>

Bingkai di atas memuat `Alert.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## Contoh Penggunaan Dasar

Gunakan `<Alert />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Alert class="kustom-alert">
  <!-- Konten slot atau properti -->
</Alert>
```

---

## Properti & Atribut

- `action`: Mengonfigurasi atribut atau perilaku `action`.
- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `dismissible`: Mengonfigurasi atribut atau perilaku `dismissible`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `message`: Mengonfigurasi atribut atau perilaku `message`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
