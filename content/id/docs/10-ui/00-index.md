---
title: "Krewire UI & Design System"
description: "Design system pop-brutalist tanpa dependensi eksternal dan arsitektur komponen untuk ekosistem Krewire."
date: "2026-11-09"
---

# Krewire UI

**Krewire UI (`kui`)** adalah design system neo-brutalist ("pop-brutalist") tanpa dependensi eksternal dan pustaka komponen yang dibangun secara native untuk ekosistem Krewire.

Komponen dapat digunakan langsung melalui dua authoring surface:
- **Go-Native:** Widget dan node bertipe aman secara programatik (`github.com/krewire/ui/widget`).
- **`.kiw` DSL:** Single-File Component deklaratif (`import { Button } from '@krewire/ui'`).
- **Berbasis Paket:** Dipasang langsung via `kiw add @krewire/ui` tanpa mengotori direktori lokal, dengan opsi `kiw pkg eject` bila kustomisasi sumber lokal diperlukan.

---

## Apa yang Ada di Bab Ini?

1. [**Direktori Komponen (`/ui/components`)**](/id/docs/ui/components) — Katalog referensi komprehensif untuk seluruh 66 Single-File Component (.kiw), termasuk tombol, kartu, dialog modal, bilah navigasi, formulir, tampilan data, dan primitif tata letak.
2. **Token Desain & Pop-Brutalism** — Palet warna semantik kontras tinggi, border solid 2.5px, bayangan kotak kaku (*hard offset drop shadows* `4px 4px 0px var(--border)`), dan inversi mode gelap yang aksesibel.
3. **Arsitektur Tanpa Dependensi (Zero-Dependency)** — Kompilasi pure Go (`CGO_ENABLED=0`), bebas dari beban runtime npm atau Node.js, dan distribusi native via `embed.FS` pustaka standar.
4. [**Playground Komponen (`/ui/playground`)**](https://krewire.com/ui/playground) — Pratinjau langsung tema terang/gelap untuk `Button`, `Badge`, `Input`, `Card`, `Modal`, `Tabs`, dan `Alert` dengan pilihan varian dan ukuran serta salin sekali klik untuk kode sumber `.kiw` dan contoh penggunaannya. Pratinjau yang sama disematkan di bagian atas halaman masing-masing komponen.
