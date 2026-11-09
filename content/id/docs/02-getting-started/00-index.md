---
title: "Memulai (Getting Started)"
description: "Jalur pemula: instal kiw, bangun situs pertama Anda dalam lima menit, lalu aplikasi fullstack dengan SQLite dan form. Tanpa perlu latar belakang arsitektur."
date: "2026-11-09"
---

# Memulai (Getting Started)

<nav class="track-crumbs" aria-label="Jalur pemula">
<span class="track-crumbs-label">Jalur pemula</span>
<ol>
<li><a href="/">Pilih jalur</a></li>
<li aria-current="step"><span>Memulai</span></li>
<li><a href="/getting-started/quickstart">1 · Quickstart</a></li>
<li><a href="/getting-started/fullstack-app-tutorial">2 · Aplikasi fullstack</a></li>
<li><a href="/getting-started/dsl">3 · DSL .kiw</a></li>
</ol>
</nav>

Bab ini adalah **jalur pemula**. Bab ini mengasumsikan Anda sedikit mengenal Go dan terminal — tidak lebih. Anda tidak perlu memahami lapisan arsitektur, grafik dependensi, level cakupan, atau ADR untuk membangun dengan Krewire; semua itu ada di bab terpisah [Pendalaman Arsitektur](/architecture) untuk platform engineer dan kontributor.

---

## 1. Jalur Pemula

Ikuti secara berurutan. Setiap perintah di dalamnya telah dijalankan terhadap CLI `kiw` versi terkini.

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.1</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/quickstart">Quickstart (5 menit) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Instal <code>kiw</code>, scaffold situs statis, ubah halaman, tambahkan komponen buatan Anda sendiri, jalankan server pengembangan, dan hasilkan build produksi.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">2.2</span>
      <h3 class="chapter-card-title">
        <a href="/getting-started/fullstack-app-tutorial">Membangun Aplikasi Fullstack (15 menit) <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pengelola tugas dengan <code>kiw new --app</code>, SQLite melalui <code>pkg/orm</code>, halaman <code>.kiw</code> dengan komponen Krewire UI, form tervalidasi yang dilindungi CSRF, dan build satu binary.
    </p>
  </div>

</div>

Setelah menyelesaikan jalur ini, simpan halaman referensi bab ini agar mudah dijangkau:

| # | Halaman | Baca saat Anda ingin… |
| :--- | :--- | :--- |
| 2.3 | [Instalasi](/getting-started/installation) | menginstal di OS lain, mengunci versi, atau menyiapkan shell completion. |
| 2.4 | [Konfigurasi](/getting-started/configuration) | memahami setiap field di `krewire.yaml`. |
| 2.5 | [Struktur Direktori](/getting-started/directory-structure) | mengetahui letak berkas untuk setiap workload. |
| 2.6 | [Pengembangan Berbantuan AI](/getting-started/agentic-development) | berkolaborasi dengan agen koding AI pada proyek Krewire. |
| 2.7 | [DSL `.kiw`](/getting-started/dsl) | mempelajari sintaks komponen selengkapnya: props, slot, scoped CSS, skrip. |

---

## 2. Prasyarat

| Persyaratan | Versi | Periksa dengan |
| :--- | :--- | :--- |
| **Go** | 1.27.1 atau lebih baru | `go version` |
| **Git** | versi terbaru apa pun | `git --version` |
| **OS** | Linux, macOS, atau Windows (PowerShell atau WSL2) | — |

Krewire **tidak memerlukan Node.js, npm, maupun bundler**. Jika Go belum terpasang atau versinya lebih lama dari 1.27.1, instal terlebih dahulu dari [go.dev/dl](https://go.dev/dl).

---

## 3. Workload Mana yang Sebaiknya Saya Mulai?

Setiap proyek Krewire memiliki satu *jenis* (kind). Pilih baris yang sesuai dengan apa yang Anda bangun — Anda dapat berkembang ke jenis lain nanti tanpa memindahkan komponen Anda.

| Saya ingin membangun… | Jenis | Scaffold | Pelajari lebih lanjut |
| :--- | :--- | :--- | :--- |
| Situs web, landing page, atau blog | `site` | `kiw new my-site --site` | [Workload Site](/workloads/site) |
| Dokumentasi atau buku | `book` | `kiw new my-docs --book` | [Workload Book](/workloads/book) |
| Aplikasi web dengan database dan form | `app` | `kiw new my-app --app` | [Workload App](/workloads/app) |
| Alat command-line | `cli` | `kiw new my-tool --cli` | [Workload CLI](/workloads/cli) |
| Pemroses job latar belakang | `worker` | `kiw new my-worker --worker` | [Workload Worker](/workloads/worker) |
| Microservice HTTP | `service` | `kiw new my-svc --service` | [Workload Service](/workloads/service) |
| Reverse proxy di depan layanan | `gateway` | `kiw new my-gw --gateway` | [Workload Gateway](/workloads/gateway) |
| Definisi infrastruktur cloud | `infra` | `kiw new my-infra --infra` | [Workload Infra](/workloads/infra) |

Belum yakin? Mulailah dengan `--site` (Quickstart) atau `--app` (tutorial Fullstack).

---

## 4. Perintah Sehari-hari

| Perintah | Fungsinya |
| :--- | :--- |
| `kiw new <name> --site` | Membuat proyek baru (`--app`, `--book`, `--cli`, … untuk jenis lainnya). |
| `kiw dev` | Menjalankan proyek secara lokal dengan rebuild otomatis saat disimpan (`http://localhost:8080`). |
| `kiw build` | Build produksi: berkas statis di `.krewire/build/` untuk situs dan buku, binary di `.krewire/build/bin/` untuk aplikasi. |
| `kiw serve --addr :3000` | Pratinjau hasil build statis melalui HTTP. |
| `kiw make component Card` | Membuat berkas awal (`component`, `page`, `model`, `migration`, `handler`, …). |
| `kiw test` | Menjalankan tes Go proyek. |
| `kiw doctor` | Mendiagnosis toolchain dan workspace Anda. |
| `kiw help <command>` | Menampilkan opsi dari perintah apa pun. |

Daftar lengkapnya ada di [Referensi Perintah CLI](/cli-reference/commands).

---

## 5. Kendala Umum bagi Pemula

**`kiw: command not found`** — direktori instalasi tidak ada di `PATH` Anda. Tambahkan `export PATH="$HOME/.local/bin:$PATH"` (installer shell) atau `export PATH="$(go env GOPATH)/bin:$PATH"` (`go install`) ke `~/.zshrc` atau `~/.bashrc`, lalu buka terminal baru.

**`address already in use`** — port 8080 sedang dipakai program lain. Ubah `dev.port` di `krewire.yaml`, atau pratinjau build di alamat lain dengan `kiw serve --addr :3000`.

**`requires go >= 1.27.1`** — versi Go Anda lebih lama dari yang ditargetkan Krewire. Instal Go terbaru dari [go.dev/dl](https://go.dev/dl); paket bawaan distribusi sering tertinggal beberapa versi.

---

## Langkah Selanjutnya

- Mulai jalurnya: [**2.1 Quickstart →**](/getting-started/quickstart)
- Ingin membaca gambaran besarnya terlebih dahulu? Lihat [**1. Ringkasan →**](/overview).
