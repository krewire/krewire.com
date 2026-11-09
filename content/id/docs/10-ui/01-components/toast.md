---
title: "Toast"
description: "Notifikasi mengambang ringan (*toast*) dengan pewaktu tutup otomatis."
date: "2026-11-09"
---

# Toast

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Toast.kiw`  
Sumber: `github.com/krewire/ui/components/Toast.kiw`

Notifikasi mengambang ringan (*toast*) dengan pewaktu tutup otomatis.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Toast.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Toast.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Toast />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Toast class="kustom-toast">
  <!-- Konten slot atau properti -->
</Toast>
```

---

## 3. Properti & Atribut yang Didukung

- `dismissible`: Mengonfigurasi atribut atau perilaku `dismissible`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `id`: Mengonfigurasi atribut atau perilaku `id`.
- `message`: Mengonfigurasi atribut atau perilaku `message`.
- `position`: Mengonfigurasi atribut atau perilaku `position`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
