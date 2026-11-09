---
title: "Krewire UI & Design System"
description: "Design system pop-brutalist tanpa dependensi eksternal dan arsitektur komponen untuk ekosistem Krewire."
date: "2026-11-09"
---

# Krewire UI (`github.com/krewire/ui`)

Jalur import: `github.com/krewire/ui`  
Direktori Komponen: `/id/docs/ui/components`  
Stylesheet Design System: `/styles/krewire-ui.css`

**Krewire UI (`kui`)** adalah design system neo-brutalist ("pop-brutalist") tanpa dependensi eksternal dan pustaka Single-File Component (`.kiw`) yang dibangun secara native untuk ekosistem Go Krewire.

Seluruh komponen mengikuti filosofi **shadcn/ui "Copy & Own"**: Anda menyalin berkas `.kiw` langsung ke direktori `./components/` proyek Anda, memberi Anda kepemilikan penuh tanpa dependensi runtime tersembunyi atau risiko breaking changes dari upstream.

---

## Apa yang Ada di Bab Ini?

1. [**Direktori Komponen (`/ui/components`)**](/id/docs/ui/components) — Katalog referensi komprehensif untuk seluruh 66 Single-File Component (.kiw), termasuk tombol, kartu, dialog modal, bilah navigasi, formulir, tampilan data, dan primitif tata letak.
2. **Token Desain & Pop-Brutalism** — Palet warna semantik kontras tinggi, border solid 2.5px, bayangan kotak kaku (*hard offset drop shadows* `4px 4px 0px var(--border)`), dan inversi mode gelap yang aksesibel.
3. **Arsitektur Tanpa Dependensi (Zero-Dependency)** — Kompilasi pure Go (`CGO_ENABLED=0`), bebas dari beban runtime npm atau Node.js, dan distribusi native via `embed.FS` pustaka standar.
