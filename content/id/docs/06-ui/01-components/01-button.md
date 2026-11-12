---
title: "Button"
description: "Tombol interaktif inti dengan varian primary, secondary, accent, outline, dan ghost."
date: "2026-11-09"
---

# Button

Kategori: **Aksi & Tombol**

Komponen `<Button />` adalah elemen interaktif fundamental dari design system Krewire UI. Komponen ini me-render elemen HTML `<button>` atau tautan `<a>` dengan styling pop-brutalist (border solid tegas, bayangan kotak kaku, dan animasi klik kaku).

---

## Pratinjau Langsung

<div data-component-preview="Button" data-slug="button"></div>

Bingkai di atas memuat `Button.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## Contoh Penggunaan

### Varian Warna

```html
<Button variant="primary">Aksi Utama</Button>
<Button variant="secondary">Aksi Sekunder</Button>
<Button variant="accent">Aksi Peringatan</Button>
<Button variant="outline">Garis Tepi (Outline)</Button>
<Button variant="ghost">Tombol Polos (Ghost)</Button>
```

### Skala Ukuran

```html
<Button variant="primary" size="sm">Kecil (sm)</Button>
<Button variant="primary" size="md">Sedang (md - standar)</Button>
<Button variant="primary" size="lg">Besar (lg)</Button>
```

### Dengan Ikon dan Tautan

```html
<!-- Tombol dengan ikon di depan -->
<Button variant="primary" icon="🚀">Luncurkan Kluster</Button>

<!-- Tombol yang berfungsi sebagai tautan tautan -->
<Button href="/docs" variant="outline" iconRight="→">
  Baca Dokumentasi
</Button>
```

### Status Nonaktif dan Handler Klik

```html
<Button variant="primary" disabled=true>Status Nonaktif</Button>
<Button variant="accent" onClick="handleDeleteEnvironment()">Hapus</Button>
```

---

## Properti & Atribut

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
