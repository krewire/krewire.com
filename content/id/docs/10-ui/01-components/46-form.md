---
title: "Form"
description: "Pembungkus formulir aksesibel dengan method spoofing dan status pengiriman."
date: "2026-11-09"
---

# Form

Kategori: **Formulir & Masukan**

Pembungkus formulir aksesibel dengan method spoofing dan status pengiriman.

---

## Contoh Penggunaan Dasar

Gunakan `<Form />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Form class="kustom-form">
  <!-- Konten slot atau properti -->
</Form>
```

---

## Properti & Atribut

- `action`: Mengonfigurasi atribut atau perilaku `action`.
- `enctype`: Mengonfigurasi atribut atau perilaku `enctype`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `method`: Mengonfigurasi atribut atau perilaku `method`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
