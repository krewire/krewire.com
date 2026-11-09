---
title: "Stack"
description: "Kontainer tata letak terarah dengan pengaturan jarak fleksibel."
date: "2026-11-09"
---

# Stack

Kategori: **Tata Letak Primitif**  
Berkas Komponen: `components/Stack.kiw`  
Sumber: `github.com/krewire/ui/components/Stack.kiw`

Kontainer tata letak terarah dengan pengaturan jarak fleksibel.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Stack.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Stack.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Stack />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Stack class="kustom-stack">
  <!-- Konten slot atau properti -->
</Stack>
```

---

## 3. Properti & Atribut yang Didukung

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `direction`: Mengonfigurasi atribut atau perilaku `direction`.
- `gap`: Mengonfigurasi atribut atau perilaku `gap`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `justify`: Mengonfigurasi atribut atau perilaku `justify`.
- `wrap`: Mengonfigurasi atribut atau perilaku `wrap`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
