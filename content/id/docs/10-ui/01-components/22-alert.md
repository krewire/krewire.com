---
title: "Alert"
description: "Banner pemberitahuan kontekstual dengan status info, success, warning, dan error."
date: "2026-11-09"
---

# Alert

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Alert.kiw`  
Sumber: `github.com/krewire/ui/components/Alert.kiw`

Banner pemberitahuan kontekstual dengan status info, success, warning, dan error.

---

## Pratinjau Langsung

<div data-component-preview="Alert" data-slug="alert"></div>

Bingkai di atas memuat `Alert.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Alert.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Alert.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Alert />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Alert class="kustom-alert">
  <!-- Konten slot atau properti -->
</Alert>
```

---

## 3. Properti & Atribut yang Didukung

- `action`: Mengonfigurasi atribut atau perilaku `action`.
- `actions`: Mengonfigurasi atribut atau perilaku `actions`.
- `dismissible`: Mengonfigurasi atribut atau perilaku `dismissible`.
- `icon`: Mengonfigurasi atribut atau perilaku `icon`.
- `message`: Mengonfigurasi atribut atau perilaku `message`.
- `title`: Mengonfigurasi atribut atau perilaku `title`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
