---
title: "ActionMenu"
description: "Menu dropdown aksi kontekstual ringkas untuk baris tabel dan kartu."
date: "2026-11-09"
---

# ActionMenu

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/ActionMenu.kiw`  
Sumber: `github.com/krewire/ui/components/ActionMenu.kiw`

Menu dropdown aksi kontekstual ringkas untuk baris tabel dan kartu.

---

## 1. Instalasi (Model Copy & Own)

Salin berkas komponen `ActionMenu.kiw` langsung ke direktori `./components` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/ActionMenu.kiw ./components/
```

Setelah disalin, proyek Anda memiliki berkas komponen ini sepenuhnya. Anda bebas menyesuaikan markup HTML, logika, dan styling tanpa khawatir pembaruan upstream merusak aplikasi Anda.

---

## 2. Contoh Penggunaan Dasar

Gunakan `<ActionMenu />` langsung di dalam halaman `.kiw` atau template tata letak:

```html
<ActionMenu class="kustom-action-menu">
  <!-- Konten slot atau properti -->
</ActionMenu>
```

---

## 3. Properti & Atribut yang Didukung

- `align`: Mengonfigurasi atribut atau perilaku `align`.
- `item`: Mengonfigurasi atribut atau perilaku `item`.
- `trigger`: Mengonfigurasi atribut atau perilaku `trigger`.

- `class`: Kelas CSS kustom tambahan yang ingin ditambahkan pada elemen utama komponen.

---

## 4. Token Visual & Aksesibilitas

- **Stylesheet Design System**: Mewarisi token CSS terpusat dari `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Mode Gelap (Dark Mode)**: Warna permukaan dan teks otomatis dibalik saat `data-theme="dark"` atau class `.dark` diterapkan pada `:root`.
- **Zero-Bloat**: Dibangun secara native menggunakan elemen primitif HTML standar tanpa ketergantungan npm atau JavaScript berat.
