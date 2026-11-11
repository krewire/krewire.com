---
title: "Keterbukaan & Self-Hosting"
description: "Repositori Krewire mana yang publik dan dengan lisensi apa, kapabilitas mana yang berjalan tanpa layanan yang di-host Krewire, cara meng-host sendiri keluaran Anda, dan bagaimana Krewire hidup berdampingan dengan runtime eksternal."
date: "2026-11-09"
---

# Keterbukaan & Self-Hosting

Halaman ini memungkinkan Anda memeriksa — tanpa harus memercayai materi pemasaran — apa saja yang dapat Anda **periksa, gunakan, modifikasi, host sendiri, dan operasikan secara mandiri**. Setiap pernyataan ditautkan ke lisensi, repositori, atau referensi teknis. Fakta diverifikasi terhadap repositori pada **2026-11-09**.

---

## Ketersediaan repositori & lisensi

| Repositori | Isi | Visibilitas | Lisensi |
| :--- | :--- | :--- | :--- |
| `krewire/krewire` | Monorepo inti: pustaka `pkg/*`, CLI `kiw`, template. | Publik | MIT |
| `krewire/krew` | Mesin agen AI otonom (embed, CLI, HTTP/SSE, MCP). | Publik | MIT |
| `krewire/ui` | Design system dan komponen UI `.kiw`. | Publik | MIT |
| `krewire/krewire.com` | Situs dan dokumentasi ini. | Publik | MIT |
| `krewire/krewire.github.io` | Kode sumber portal komunitas. | Publik | MIT |
| `krewire/dash` | Krewire Dash — virtual workspace & konsol studio. | Layanan platform | Portal platform terdedikasi |
| `krewire/cloud` | Krewire Cloud — control plane PaaS hibrida. | Layanan platform | Portal platform terdedikasi |
| `krewire/auth` | Krewire Auth — engine identitas, SSO, MFA, RBAC. | Layanan platform | Portal platform terdedikasi |
| `krewire/forge` | Krewire Forge — builder deklaratif web, form & admin panel. | Layanan platform | Portal platform terdedikasi |

Inti open-source adalah semua yang berada di baris **publik berlisensi MIT**. Layanan platform menyediakan kapabilitas operasional khusus di atas fondasi inti. Jalur kontribusi untuk setiap repositori publik dijelaskan di [Kontribusi & Komunitas](/project/contributing).

---

## Apa yang berjalan tanpa layanan yang di-host Krewire

Kapabilitas inti publik berikut berjalan sepenuhnya di mesin Anda atau infrastruktur Anda sendiri. Tidak ada yang memerlukan akun Krewire, kunci lisensi, atau server yang dioperasikan Krewire.

| Kapabilitas | Perintah / paket | Akses jaringan yang dibutuhkan |
| :--- | :--- | :--- |
| Scaffold proyek untuk kedelapan workload | `kiw new`, `kiw init` | Tidak ada (template bawaan); opsi `--template <URL git>` meng-clone dari URL pilihan Anda. |
| Pengembangan dan pratinjau lokal | `kiw dev`, `kiw serve`, `kiw dash` | Tidak ada. |
| Membangun situs, buku, binary | `kiw build` | Tidak ada selain mengambil modul Go yang Anda gunakan. |
| Pengujian dan verifikasi | `kiw test`, `kiw vet`, `kiw compat` | Tidak ada. |
| Deploy ke target Anda sendiri | `kiw deploy --target` (`binary`, `systemd`, `gh-pages`, `infra`) | Hanya target yang Anda pilih (misalnya push ke branch GitHub Pages Anda sendiri). |
| Memakai pustaka di program Go mana pun | `github.com/krewire/krewire/pkg/...` | Tidak ada — paket Go standar. |
| Menjalankan agen AI | `krew` | Penyedia LLM pilihan Anda; dengan Ollama (`OLLAMA_HOST`) inferensi dapat berjalan sepenuhnya lokal. |

**Panggilan jaringan yang memang ada di CLI inti** (diverifikasi dari kode sumber):

- `kiw self-update` meminta data ke GitHub Releases API untuk `krewire/krewire`.
- [Skrip instalasi](https://krewire.com/scripts/install.sh) mengunduh dari GitHub Releases jika ada rilis, dan jika tidak beralih ke `go install` (Go module proxy). Belum ada GitHub Release yang dipublikasikan, sehingga saat ini jalur fallback yang dipakai dan Go 1.27.1+ diperlukan.
- Kode sumber CLI mendeklarasikan endpoint default di bawah `api.krewire.com` untuk integrasi platform di masa depan; per v0.2.0 tidak ada perintah `kiw` inti yang memanggilnya.

**Satu catatan pada pustaka:** `pkg/auth.NewClient` adalah klien untuk server Krewire Auth dan memakai default `https://auth.krewire.com` bila URL tidak diberikan. Untuk penggunaan yang sepenuhnya mandiri, primitif stateless di paket yang sama — `SignJWT`/parsing JWT dan middleware `BasicAuth` — tidak memerlukan server. Lihat [Primitif Autentikasi](/security-platforms/auth).

---

## Meng-host sendiri keluaran Anda

Krewire tidak meng-host aplikasi Anda; Anda yang meng-host-nya. Keluaran `kiw build` / `kiw deploy` berupa:

- **Berkas statis** untuk `site` dan `book` — dapat disajikan oleh web server statis mana pun atau GitHub Pages. Situs ini contohnya: dibangun dengan `kiw build` dan disajikan oleh kontainer `nginx:alpine` standar.
- **Satu binary Go** untuk `app`, `cli`, `worker`, `service`, dan `gateway` — dapat dijalankan di host Linux/macOS mana pun; `kiw deploy --target systemd` menghasilkan unit systemd.

**Kebutuhan operasional:**

| Kebutuhan | Detail |
| :--- | :--- |
| Toolchain build | Go 1.27.1 atau lebih baru (semua modul publik mendeklarasikan `go 1.27.1`). |
| Platform installer | Linux dan macOS pada amd64/arm64 ([skrip instalasi](https://krewire.com/scripts/install.sh)). |
| Target `systemd` | Host Linux dengan systemd. |
| `krew` dari kode sumber | `go.mod`-nya saat ini memakai `replace` lokal untuk modul inti, sehingga bangun di dalam Go workspace multi-repositori (`go work`), seperti yang dilakukan CI situs ini. |

Lihat [Deployment & Operasi](/guides/deployment-and-operations) untuk panduan langkah demi langkah.

---

## Layanan platform

Krewire Dash, Cloud, Auth, dan Forge bersifat **opsional**. Tidak ada perintah atau pustaka inti publik yang bergantung padanya untuk membangun, menjalankan, atau men-deploy perangkat lunak Anda. Peran, arsitektur, dan portal live masing-masing didokumentasikan di [Inti & Layanan Platform](/project/core-and-platform-services).

---

## Interoperabilitas dan koeksistensi

- **Berdampingan dengan stack lain.** Workload Krewire dapat berada di direktorinya sendiri di dalam repositori Node.js, Laravel, atau PHP. `kiw info` mendeteksi proyek host tersebut tetapi tidak mengubah `package.json`, `composer.json`, `node_modules/`, atau `vendor/`.
- **Antarmuka standar.** Workload berkomunikasi melalui HTTP atau input/output standar; keluarannya berupa binary dan berkas statis biasa.
- **Infrastruktur Anda, pilihan Anda.** Target deploy adalah host Anda sendiri (binary/systemd), GitHub Pages, atau infrastruktur Anda sendiri melalui `pkg/cloud/infra`.
