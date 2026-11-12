---
title: "Flex"
description: "Kontainer tata letak fleksibel dengan konfigurasi arah, align, dan justify."
date: "2026-11-09"
---

# Flex

Kategori: **Tata Letak Primitif**

Kontainer tata letak fleksibel dengan konfigurasi arah, align, dan justify.

---

## Contoh Penggunaan Dasar

Gunakan `<Flex />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Flex class="kustom-flex">
  <!-- Konten slot atau properti -->
</Flex>
```

---

## Properti & Atribut

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `direction`: Mengonfigurasi atribut atau perilaku `direction`.
- `gap`: Mengonfigurasi atribut atau perilaku `gap`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `justify`: Mengonfigurasi atribut atau perilaku `justify`.
- `wrap`: Mengonfigurasi atribut atau perilaku `wrap`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
