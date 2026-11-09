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
| [`krewire/krewire`](https://github.com/krewire/krewire) | Monorepo inti: pustaka `pkg/*`, CLI `kiw`, template. | Publik | [MIT](https://github.com/krewire/krewire/blob/main/LICENSE) (berkas `LICENSE` tingkat paket juga MIT) |
| [`krewire/krew`](https://github.com/krewire/krew) | Mesin agen AI otonom (embed, CLI, HTTP/SSE, MCP). | Publik | [MIT](https://github.com/krewire/krew/blob/main/LICENSE) |
| [`krewire/ui`](https://github.com/krewire/ui) | Design system dan komponen UI `.kiw`. | Publik | [MIT](https://github.com/krewire/ui/blob/main/LICENSE) |
| [`krewire/krewire.com`](https://github.com/krewire/krewire.com) | Situs dan dokumentasi ini. | Publik | [MIT](https://github.com/krewire/krewire.com/blob/main/LICENSE) |
| [`krewire/krewire.github.io`](https://github.com/krewire/krewire.github.io) | Kode sumber portal komunitas. | Publik | [MIT](https://github.com/krewire/krewire.github.io/blob/main/LICENSE) |
| `krewire/dash` | Krewire Workspace (Dash) — server workspace virtual. | **Privat** | Kode sumber tidak tersedia untuk publik; tidak ada pemberian lisensi publik. |
| `krewire/cloud` | Krewire Cloud — control plane PaaS hibrida. | **Privat** | Kode sumber tidak tersedia untuk publik; tidak ada pemberian lisensi publik. |
| `krewire/auth` | Krewire Auth — server identitas, SSO, MFA, RBAC. | **Privat** | Kode sumber tidak tersedia untuk publik; tidak ada pemberian lisensi publik. |
| `krewire/panel` | Krewire Panel — panel administrasi server. | **Privat** | Kode sumber tidak tersedia untuk publik; tidak ada pemberian lisensi publik. |
| `krewire/internal` | Pusat dokumentasi internal (ADR, roadmap). | **Privat** | Tidak tersedia untuk publik. |

Inti open-source adalah semua yang berada di baris **publik**. Jalur kontribusi untuk setiap repositori publik dijelaskan di [Kontribusi & Komunitas](/project/contributing).

---

## Apa yang berjalan tanpa layanan yang di-host Krewire

Kapabilitas inti publik berikut berjalan sepenuhnya di mesin Anda atau infrastruktur Anda sendiri. Tidak ada yang memerlukan akun Krewire, kunci lisensi, atau server yang dioperasikan Krewire.

| Kapabilitas | Perintah / paket | Akses jaringan yang dibutuhkan |
| :--- | :--- | :--- |
| Scaffold proyek untuk kedelapan workload | `kiw new`, `kiw init` | Tidak ada (template bawaan); opsi `--template <URL git>` meng-clone dari URL pilihan Anda. |
| Pengembangan dan pratinjau lokal | `kiw dev`, `kiw serve`, `kiw dashboard` | Tidak ada. |
| Membangun situs, buku, binary | `kiw build` | Tidak ada selain mengambil modul Go yang Anda gunakan. |
| Pengujian dan verifikasi | `kiw test`, `kiw vet`, `kiw compat` | Tidak ada. |
| Deploy ke target Anda sendiri | `kiw deploy --target` (`binary`, `systemd`, `gh-pages`, `infra`) | Hanya target yang Anda pilih (misalnya push ke branch GitHub Pages Anda sendiri). |
| Memakai pustaka di program Go mana pun | `github.com/krewire/krewire/pkg/...` | Tidak ada — paket Go standar. |
| Menjalankan agen AI | `krew` | Penyedia LLM pilihan Anda; dengan Ollama (`OLLAMA_HOST`) inferensi dapat berjalan sepenuhnya lokal ([README krew](https://github.com/krewire/krew#readme)). |

**Panggilan jaringan yang memang ada di CLI inti** (diverifikasi dari kode sumber):

- `kiw self-update` meminta data ke GitHub Releases API untuk `krewire/krewire`.
- [Skrip instalasi](https://krewire.com/scripts/install.sh) mengunduh dari GitHub Releases jika ada rilis, dan jika tidak beralih ke `go install` (Go module proxy). Belum ada GitHub Release yang dipublikasikan, sehingga saat ini jalur fallback yang dipakai dan Go 1.27.1+ diperlukan.
- Kode sumber CLI mendeklarasikan endpoint default di bawah `api.krewire.com` untuk integrasi platform di masa depan; per v0.2.0 tidak ada perintah `kiw` inti yang memanggilnya.

**Satu catatan pada pustaka:** `pkg/auth.NewClient` adalah klien untuk server Krewire Auth dan memakai default `https://auth.krewire.com` bila URL tidak diberikan. Server Auth bersifat privat dan belum ada layanan Auth ter-hosting yang tersedia untuk publik (lihat [Inti & Layanan Platform](/project/core-and-platform-services)). Untuk penggunaan yang sepenuhnya mandiri, primitif stateless di paket yang sama — `SignJWT`/parsing JWT dan middleware `BasicAuth` — tidak memerlukan server. Lihat [Primitif Autentikasi](/security-platforms/auth).

---

## Meng-host sendiri keluaran Anda

Krewire tidak meng-host aplikasi Anda; Anda yang meng-host-nya. Keluaran `kiw build` / `kiw deploy` berupa:

- **Berkas statis** untuk `site` dan `book` — dapat disajikan oleh web server statis mana pun atau GitHub Pages. Situs ini contohnya: dibangun dengan `kiw build` dan disajikan oleh kontainer `nginx:alpine` standar (lihat [`Dockerfile`](https://github.com/krewire/krewire.com/blob/main/Dockerfile) di repositori ini).
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

## Layanan ter-hosting opsional

Krewire Workspace (Dash), Cloud, Auth, dan Panel bersifat **opsional**. Tidak ada perintah atau pustaka inti publik yang bergantung padanya untuk membangun, menjalankan, atau men-deploy perangkat lunak Anda. Visibilitas, model penyampaian, dan ketersediaannya saat ini didokumentasikan di [Inti & Layanan Platform](/project/core-and-platform-services). Singkatnya: keempatnya dikembangkan di repositori privat dan belum ada yang tersedia untuk publik sebagai layanan ter-hosting per 2026-11-09.

---

## Interoperabilitas dan koeksistensi

- **Berdampingan dengan stack lain.** Workload Krewire dapat berada di direktorinya sendiri di dalam repositori Node.js, Laravel, atau PHP. `kiw info` mendeteksi proyek host tersebut tetapi tidak mengubah `package.json`, `composer.json`, `node_modules/`, atau `vendor/` ([panduan interoperabilitas](https://github.com/krewire/krewire/blob/main/docs/guides/interoperability-and-migration.md), dicakup oleh pengujian seperti `TestDetectHost_NodeJS` dan `TestDetectHost_Laravel`).
- **Antarmuka standar.** Workload berkomunikasi melalui HTTP atau input/output standar; keluarannya berupa binary dan berkas statis biasa.
- **Infrastruktur Anda, pilihan Anda.** Target deploy adalah host Anda sendiri (binary/systemd), GitHub Pages, atau infrastruktur Anda sendiri melalui `pkg/cloud/infra`.

*Pemilik konten: Krewire Team. Verifikasi ulang tabel di atas pada setiap rilis yang di-tag; visibilitas repositori dan lisensi adalah sumber kebenarannya.*
