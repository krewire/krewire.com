---
title: "Card"
description: "Permukaan kartu terangkat dengan slot header, media, isi, dan footer."
date: "2026-11-09"
---

# Card

Kategori: **Tampilan Data & Konten**

Komponen `<Card />` menyediakan permukaan kontainer terangkat dengan pop-border khas (`2.5px solid var(--border)`), bayangan kotak kaku, serta slot fleksibel untuk header, media gambar, isi kartu, dan footer tombol aksi.

---

## Pratinjau Langsung

<div data-component-preview="Card" data-slug="card"></div>

Bingkai di atas memuat `Card.kiw` asli dari Krewire UI, yang sudah di-render oleh `kiw build` untuk setiap varian, dalam tema terang dan gelap secara berdampingan. Gunakan tombol pilihan untuk mengganti varian, lalu salin kode sumber komponen atau contoh penggunaannya. **Buka di playground** menampilkannya bersama komponen inti lainnya.

---

## Contoh Penggunaan

### Kartu Standar dengan Judul dan Badge

```html
<Card title="Database Cluster" subtitle="PostgreSQL 17.2 · Produksi" badge="Sehat">
  <p>Kluster beroperasi pada beban optimal dengan 6 replika worker.</p>
  <div slot="footer" class="flex justify-end gap-2">
    <Button variant="ghost" size="sm">Lihat Log</Button>
    <Button variant="outline" size="sm">Cadangkan</Button>
  </div>
</Card>
```

### Kartu Unggulan dengan Header dan Media Kustom

```html
<Card variant="featured">
  <div slot="header" class="p-2 border-b">
    <h4 class="font-bold">Enterprise Cloud Gateway</h4>
  </div>
  <div slot="media">
    <img src="/assets/images/architecture.png" alt="Diagram Arsitektur" />
  </div>
  <p>Gateway API edge throughput tinggi dengan verifikasi token dan pembatasan laju request.</p>
  <div slot="footer">
    <Button variant="primary" size="md">Konfigurasi Gateway</Button>
  </div>
</Card>
```

---

## Properti & Atribut

### Properti
- `title`: Teks judul kartu di area header.
- `subtitle`: Teks deskripsi sekunder di bawah judul.
- `badge`: Lencana status mikro di sudut kanan header kartu.
- `variant`: Varian visual kartu (misal `default`, `featured`).
- `class`: Kelas CSS kustom tambahan pada elemen utama kartu.

### Slot
- `header` / `head`: Kustom markup untuk menggantikan container judul bawaan.
- `media` / `image`: Slot untuk menyematkan gambar banner visual kartu.
- `body` / default: Area konten utama kartu.
- `footer` / `actions`: Kontainer aksi tombol di bagian bawah kartu.
