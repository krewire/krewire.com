---
title: "Button"
description: "Tombol interaktif inti dengan varian primary, secondary, accent, outline, dan ghost."
date: "2026-11-09"
---

# Button (`Button.kiw`)

Kategori: **Aksi & Tombol**  
Berkas Komponen: `components/Button.kiw`  
Sumber: `github.com/krewire/ui/components/Button.kiw`

Komponen `<Button />` adalah elemen interaktif fundamental dari design system Krewire UI. Komponen ini me-render elemen HTML `<button>` atau tautan `<a>` dengan styling pop-brutalist (border solid tegas, bayangan kotak kaku, dan animasi klik kaku).

---

## 1. Instalasi (Model Copy & Own)

Salin berkas `Button.kiw` langsung ke direktori `components/` proyek Anda:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Button.kiw ./components/
```

---

## 2. Contoh Penggunaan

### 2.1 Varian Warna

```html
<Button variant="primary">Aksi Utama</Button>
<Button variant="secondary">Aksi Sekunder</Button>
<Button variant="accent">Aksi Peringatan</Button>
<Button variant="outline">Garis Tepi (Outline)</Button>
<Button variant="ghost">Tombol Polos (Ghost)</Button>
```

### 2.2 Skala Ukuran

```html
<Button variant="primary" size="sm">Kecil (sm)</Button>
<Button variant="primary" size="md">Sedang (md - standar)</Button>
<Button variant="primary" size="lg">Besar (lg)</Button>
```

### 2.3 Dengan Ikon dan Tautan

```html
<!-- Tombol dengan ikon di depan -->
<Button variant="primary" icon="🚀">Luncurkan Kluster</Button>

<!-- Tombol yang berfungsi sebagai tautan tautan -->
<Button href="/docs" variant="outline" iconRight="→">
  Baca Dokumentasi
</Button>
```

### 2.4 Status Nonaktif dan Handler Klik

```html
<Button variant="primary" disabled=true>Status Nonaktif</Button>
<Button variant="accent" onClick="handleDeleteEnvironment()">Hapus</Button>
```

---

## 3. Properti yang Didukung

- `variant`: Gaya tampilan tombol (`primary`, `secondary`, `accent`, `outline`, `ghost`, default: `primary`).
- `size`: Token ukuran tombol (`sm`, `md`, `lg`, default: `md`).
- `href`: Jika diisi, merender tag `<a>` aksesibel dengan gaya tombol.
- `target`: Target tautan saat `href` digunakan (misal `_blank`).
- `icon`: Ikon atau teks emoji di sebelah kiri teks.
- `iconRight`: Ikon atau teks emoji di sebelah kanan teks.
- `text` / `label`: Teks tombol (alternatif selain isi slot).
- `type`: Nilai atribut type HTML (`button`, `submit`, `reset`, default: `button`).
- `disabled`: Menonaktifkan interaksi dan meredupkan opasitas tombol.
- `onClick`: String fungsi handler event klik JavaScript.
- `id`: ID opsional untuk elemen tombol.
- `ariaLabel`: Label aksesibilitas untuk pembaca layar.
- `class`: Kelas CSS tambahan yang ingin disematkan.
