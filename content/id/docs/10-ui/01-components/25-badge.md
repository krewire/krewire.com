---
title: "Badge"
description: "Lencana mikro pil dengan varian warna untuk tag, jumlah, dan status indikator."
date: "2026-11-09"
---

# Badge

Kategori: **Umpan Balik & Status**  
Berkas Komponen: `components/Badge.kiw`  
Sumber: `github.com/krewire/ui/components/Badge.kiw`

Lencana mikro pil dengan varian warna untuk tag, jumlah, dan status indikator.

---

## Pratinjau Langsung

<div data-component-preview="Badge" data-slug="badge"></div>

Bingkai di atas memuat `Badge.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `Badge.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Badge.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<Badge />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<Badge class="kustom-badge">
  <!-- Konten slot atau properti -->
</Badge>
```

---

## 3. Properti & Atribut yang Didukung

- `dot`: Mengonfigurasi atribut atau perilaku `dot`.
- `size`: Mengonfigurasi atribut atau perilaku `size`.
- `text`: Mengonfigurasi atribut atau perilaku `text`.
- `variant`: Mengonfigurasi atribut atau perilaku `variant`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
