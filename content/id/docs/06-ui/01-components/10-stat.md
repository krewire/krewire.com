---
title: "Stat"
description: "Kartu metrik indikator kinerja utama (KPI) dengan angka dan indikator tren."
date: "2026-11-09"
---

# Stat

Kategori: **Tampilan Data & Konten**

Kartu metrik indikator kinerja utama (KPI) dengan angka dan indikator tren.

---

## Contoh Penggunaan Dasar

Gunakan `<Stat />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Stat class="kustom-stat">
  <!-- Konten slot atau properti -->
</Stat>
```

---

## Properti & Atribut

- `change`: Mengonfigurasi atribut atau perilaku `change`.
- `help`: Mengonfigurasi atribut atau perilaku `help`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `label`: Mengonfigurasi atribut atau perilaku `label`.
- `subtitle`: Mengonfigurasi atribut atau perilaku `subtitle`.
- `trend`: Mengonfigurasi atribut atau perilaku `trend`.
- `unit`: Mengonfigurasi atribut atau perilaku `unit`.
- `value`: Mengonfigurasi atribut atau perilaku `value`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
