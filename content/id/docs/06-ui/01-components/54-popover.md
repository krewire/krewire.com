---
title: "Popover"
description: "Kontainer popover kontekstual mengambang yang terhubung ke elemen target."
date: "2026-11-09"
---

# Popover

Kategori: **Lapisan Mengambang & Dialog**

Kontainer popover kontekstual mengambang yang terhubung ke elemen target.

---

## Contoh Penggunaan Dasar

Gunakan `<Popover />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Popover class="kustom-popover">
  <!-- Konten slot atau properti -->
</Popover>
```

---

## Properti & Atribut

- `content`: Mengonfigurasi atribut atau perilaku `content`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `trigger`: Mengonfigurasi atribut atau perilaku `trigger`.
- `triggerText`: Mengonfigurasi atribut atau perilaku `triggerText`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
