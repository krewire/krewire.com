---
title: "Open Core & Lisensi"
description: "Cara kerja batas open-core Krewire saat ini: repositori publik berlisensi MIT, klasifikasi tier di pkg/model, aturan desain yang melindungi inti gratis, dan hal yang belum didefinisikan."
date: "2026-11-09"
---

# Open Core & Lisensi

Krewire adalah ekosistem open-source dengan desain **open-core**: inti publik berlisensi MIT, dan layanan platform opsional dikembangkan secara terpisah. Panduan ini menjelaskan batas lisensi **sebagaimana adanya saat ini** dan memisahkannya dari hal yang belum didefinisikan.

---

## 1. Apa yang dilisensikan saat ini

| Lapisan | Repositori | Lisensi |
| :--- | :--- | :--- |
| Inti open-source | `krewire`, `krew`, `ui`, `krewire.com`, `krewire.github.io` | MIT (publik) |
| Layanan platform | `dash`, `cloud`, `auth`, `forge` | Layanan platform dengan portal terdedikasi (lihat situs masing-masing produk) |

Lihat berkas LICENSE di setiap repositori.

---

## 2. Klasifikasi tier (`pkg/model`)

`pkg/model` mendefinisikan empat tier berurutan — `free` → `pro` → `team` → `enterprise` — dan `kiw tier [--json] [--check <tier>]` melaporkan postur tier sebuah proyek serta dapat memastikan proyek tetap dalam suatu tier (misalnya `kiw tier --check free`).

```text
┌──────────────────────────────┐
│  4. enterprise               │
├──────────────────────────────┤
│  3. team                     │
├──────────────────────────────┤
│  2. pro                      │
├──────────────────────────────┤
│  1. free  ← seluruh inti     │
└──────────────────────────────┘
```

Monorepo inti mendeklarasikan `Tier = model.TierFree` di `version.go`: seluruh kode dalam `krewire/krewire` adalah 100% open source, dan seluruh paket adalah free-tier secara default.

| Tier | Status |
| :--- | :--- |
| **`free`** | Seluruh inti publik. Berlisensi MIT. |
| **`pro`** | Hanya didefinisikan sebagai klasifikasi. Belum ada paket, lisensi, atau harga yang dipublikasikan. |
| **`team`** | Hanya didefinisikan sebagai klasifikasi. Belum ada paket, lisensi, atau harga yang dipublikasikan. |
| **`enterprise`** | Hanya didefinisikan sebagai klasifikasi. Belum ada paket, lisensi, ketentuan kontrak, atau harga yang dipublikasikan. |

Harga, model lisensi komersial, dan ketentuan dukungan untuk tier yang lebih tinggi **belum didefinisikan**. Definisi apa pun di masa depan akan dipublikasikan, beserta pemilik dan tanggalnya, di dokumentasi publik proyek.

---

## 3. Aturan desain yang melindungi inti gratis

Aturan berikut adalah komitmen desain Krewire yang dinyatakan untuk batas open-core:

### Aturan 1: Tanpa Lock-in Runtime
Proyek yang hanya mereferensikan paket `free` harus dapat **di-build, di-deploy, dan dijalankan tanpa kunci lisensi atau panggilan pulang ke cloud**. Tidak ada masa uji coba atau batasan tersembunyi pada tier gratis. Panggilan jaringan yang memang dilakukan CLI inti tercantum di dokumentasi publik proyek.

### Aturan 2: Batas Pembungkus Aditif
Modul berbayar di masa depan boleh membungkus, memperkuat, atau mengoptimalkan modul `free`, tetapi **tidak boleh menjadi prasyarat bagi modul free**.

### Aturan 3: Hierarki Superset yang Ketat
Setiap tier adalah superset dari tier di bawahnya: `pro` mencakup seluruh `free`, `team` mencakup seluruh `pro` dan `free`, dan `enterprise` mencakup semua tier di bawahnya.
