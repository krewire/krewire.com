---
title: "DAG Ekosistem"
description: "Grafik dependensi multi-repositori ekosistem Krewire: sisi modul Go saat kompilasi, protokol integrasi saat runtime, paralelisme produk, dan pemeriksaan yang menjaga grafik tetap asiklik."
date: "2026-11-09"
---

# DAG Ekosistem

<nav class="track-crumbs" aria-label="Jalur Pendalaman Arsitektur">
<span class="track-crumbs-label">Jalur pendalaman</span>
<ol>
<li><a href="/architecture">Ringkasan</a></li>
<li><a href="/architecture/layered-pyramid">1 · Piramida</a></li>
<li aria-current="step"><span>2 · DAG ekosistem</span></li>
<li><a href="/architecture/scope-levels">3 · Level cakupan</a></li>
<li><a href="/architecture/data-flow">4 · Alur data</a></li>
</ol>
</nav>

Krewire bukan satu repositori. Pustaka inti dan CLI, katalog UI, engine agen, layanan platform, dan situs web berada di repositori terpisah yang saling bergantung sebagai **directed acyclic graph (DAG)**. Halaman ini menampilkan grafik tersebut beserta aturan yang menjaganya tetap asiklik.

## 1. Baterai Arsitektur Ekosistem

Seluruh komponen, platform, dan workload berjalan di atas progresi baterai arsitektur terpadu, dikelola penuh oleh CLI `kiw` dan skema deklaratif `krewire.yaml`:

```text
Go Stdlib ──► Krewire Libraries ──► Krewire UI | Krewire Auth | Krewire Krew ──► Krewire Forge | Krewire Cloud ──► Krewire Dash
                                                                                                               (Workspace Terpadu)
```

## 2. Sisi Waktu Kompilasi (Impor Modul Go)

Sisi-sisi ini adalah dependensi modul Go: sebuah panah berarti "mengimpor paket dari". Data ini dibaca dari `go.mod` dan pernyataan impor setiap repositori pada 9 November 2026.

```text
                 krewire.com    krewire.github.io        (Layer 5 — built by kiw, no Go module)
                        │              │
                        └──────┬───────┘   uses kiw + kui at build time
                               ▼
     dash ──────┐        cloud ─────┐                    (Layer 4 — products & platform services)
      │  │      │         │  │      │
      │  └──► auth ◄──────┘  │      │        forge
      │         │            │      │          │
      └──► krew ◄────────────┘      │          │
            │   │                   │          │
            ▼   ▼                   ▼          ▼
        ┌──────────────── github.com/krewire/krewire ────────────────┐  (Layers 1–3 — core + kiw)
        │  tools/kiw · pkg/* · pkg/model · pkg/errs · pkg/lifecycle   │
        └─────────────────────────────┬──────────────────────────────┘
                                      ▼
                             github.com/krewire/ui  (kui)              (Layer 2 — no dependencies)
                                      ▼
                             Go standard library                       (Layer 0)
```

| Repositori | Modul | Impor dari ekosistem |
| :--- | :--- | :--- |
| `ui` (kui) | `github.com/krewire/ui` | tidak ada — hanya pustaka standar (`TestModuleHasNoExternalDependencies`) |
| `krewire` (inti + `kiw`) | `github.com/krewire/krewire` | `ui` (`pkg/ui/components` mengekspor ulang `embed.FS` milik kui) |
| `krew` (engine agen) | `github.com/krewire/krew` | `krewire` |
| `auth` | `github.com/krewire/auth` | `krewire` |
| `forge` | `github.com/krewire/forge` | `krewire` |
| `cloud` | `github.com/krewire/cloud` | `krewire`, `auth`, `krew` |
| `dash` | `github.com/krewire/dash` | `krewire`, `auth`, `krew` |
| `krewire.com`, `krewire.github.io` | — | tidak ada di Go; di-build dengan `kiw build`, yang merender komponen kui |

Dua sifat yang penting:

- **Grafiknya asiklik.** `kui` tidak mengimpor apa pun, inti hanya mengimpor `kui`, dan setiap produk mengimpor ke arah bawah. Tidak ada produk yang mengimpor produk lain: `dash` dan `forge` tidak berbagi kode Go.
- **CLI adalah daun bagi konsumen.** Produk mengimpor paket `pkg/*`; tidak ada yang mengimpor `tools/kiw`.

## 3. Integrasi Runtime (Protokol Terpisah)

Impor saat kompilasi dibatasi pada pustaka bersama. Sebagai gantinya, produk berkomunikasi dengan kapabilitas platform melalui protokol standar:

- **Identitas:** OpenID Connect / OAuth 2.0 dengan verifikasi JWKS publik terhadap `auth`.
- **API HTTP:** REST di balik reverse proxy edge atau [workload gateway](/workloads/gateway).
- **Event:** event bus asinkron atau stream Server-Sent Events.

`cloud` dan `auth` versi hosted bersifat opsional. `forge` dan `dash` berjalan sepenuhnya self-hosted; lihat [Inti & Layanan Platform](/project/core-and-platform-services) untuk mengetahui apa saja yang tersedia untuk publik saat ini.

## 4. Paralelisme Produk

| Produk | Tanggung Jawab | Bergantung pada yang lain? |
| :--- | :--- | :--- |
| **Krewire Forge** | Web, form, resource, dan panel builder deklaratif untuk Go dan `.kiw`. | Tidak |
| **Krewire Dash** | Workspace pengembang virtual berbasis AI, sandbox eksekusi, dan konsol kolaborasi. | Tidak |

Keduanya tidak saling bergantung, baik saat kompilasi maupun saat deployment, sehingga masing-masing dapat dirilis, di-fork, atau diganti secara mandiri.

## 5. Cara Grafik Dijaga Tetap Konsisten

| Pemeriksaan | Lokasi |
| :--- | :--- |
| Impor hanya ke arah bawah dan peringkat fondasi | `krewire/layering_test.go` |
| kui tidak memiliki dependensi modul | `ui/components/contract_test.go` |
| Portabilitas Go murni | `CGO_ENABLED=0 go build ./...` di CI |
| Pemformatan dan vet | `gofmt -l .`, `go vet ./...` |

Saat Anda menambahkan repositori atau impor baru antar-repositori, tempatkan terlebih dahulu pada sebuah lapisan, lalu pastikan sisi barunya mengarah ke bawah. Jika dua peer perlu bekerja sama, tambahkan protokol (API, token, event) — bukan impor.

---

**Sebelumnya:** [12.1 Piramida Berlapis](/architecture/layered-pyramid) · **Berikutnya:** [**12.3 Level Cakupan →**](/architecture/scope-levels)
