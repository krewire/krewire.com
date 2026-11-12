---
title: "ActionMenu"
description: "Menu dropdown aksi kontekstual ringkas untuk baris tabel dan kartu."
date: "2026-11-09"
---

# ActionMenu

Kategori: **Aksi & Tombol**

Menu dropdown aksi kontekstual ringkas untuk baris tabel dan kartu.

---

## Contoh Penggunaan Dasar

Gunakan `<ActionMenu />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<ActionMenu class="kustom-action-menu">
  <!-- Konten slot atau properti -->
</ActionMenu>
```

---

## Properti & Atribut

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `item`: Mengonfigurasi atribut atau perilaku `item`.
- `trigger`: Mengonfigurasi atribut atau perilaku `trigger`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
