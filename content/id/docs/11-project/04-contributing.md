---
title: "Kontribusi & Komunitas"
description: "Cara berpartisipasi dalam ekosistem Krewire — kode, dokumentasi, terjemahan, contoh, pengujian, triase, dan dukungan komunitas — beserta alur kerja, standar, kanal, dan gambaran jujur tentang tata kelola saat ini."
date: "2026-11-09"
---

# Kontribusi & Komunitas

Krewire dimaksudkan untuk dibangun oleh penggunanya, bukan hanya dipakai. Halaman ini adalah titik masuk bagi siapa pun yang ingin berpartisipasi — baik Anda menulis Go, menulis dokumentasi, menerjemahkan, menguji, maupun membantu orang lain.

---

## Cara berkontribusi

Anda tidak harus menulis kode untuk memberi kontribusi yang berarti.

| Kontribusi | Bentuknya | Tempat |
| :--- | :--- | :--- |
| **Kode** | Memperbaiki bug, meningkatkan paket, atau memperluas `kiw`. | [`krewire/krewire`](https://github.com/krewire/krewire), [`krewire/krew`](https://github.com/krewire/krew), [`krewire/ui`](https://github.com/krewire/ui) |
| **Dokumentasi** | Mengoreksi, memperjelas, atau memperluas halaman dokumentasi. Dokumentasi berada di `content/<locale>/docs/` sebagai Markdown. | [`krewire/krewire.com`](https://github.com/krewire/krewire.com) |
| **Terjemahan** | Menjaga konsistensi halaman bahasa Inggris dan Indonesia; menerjemahkan halaman baru. | [`krewire/krewire.com`](https://github.com/krewire/krewire.com) (`content/id/`, `lang/id/`) |
| **Contoh & starter** | Meningkatkan template proyek atau menulis contoh kerja. | `templates/` di [`krewire/krewire`](https://github.com/krewire/krewire) |
| **Pengujian** | Mencoba sebuah workload dari awal sampai akhir di platform Anda dan melaporkan yang rusak, beserta langkah reproduksinya. | Issue di repositori terkait |
| **Triase issue** | Mereproduksi bug yang dilaporkan, meminta detail yang kurang, dan menyarankan label (`bug`, `documentation`, `good first issue`, …). | Issue di repositori publik mana pun |
| **Dukungan komunitas** | Menjawab pertanyaan dan berbagi apa yang Anda bangun. | [Discussions di `krewire/krewire.com`](https://github.com/krewire/krewire.com/discussions) |

Kontribusi saat ini difokuskan pada repositori open-source inti (`krewire`, `krew`, `ui`, `krewire.com`).

---

## Menemukan kontribusi pertama

1. Cari issue berlabel [`good first issue`](https://github.com/search?q=org%3Akrewire+is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22&type=issues) atau [`help wanted`](https://github.com/search?q=org%3Akrewire+is%3Aissue+is%3Aopen+label%3A%22help+wanted%22&type=issues) di seluruh organisasi. **Per 2026-11-09 belum ada issue terbuka yang memakai kedua label tersebut**, sehingga daftar ini mungkin kosong.
2. Jika belum ada yang berlabel, pilih sesuatu yang Anda temukan sendiri — halaman dokumentasi yang membingungkan, terjemahan yang hilang, contoh yang gagal — dan buka issue yang menjelaskannya sebelum memulai pekerjaan besar.
3. Perbaikan dokumentasi adalah jalan masuk tercepat: setiap halaman dokumentasi di situs ini berpadanan dengan satu berkas Markdown di `content/en/docs/` dan kembarannya dalam bahasa Indonesia di `content/id/docs/`.

---

## Alur kerja dan standar

Berikut diambil dari `CONTRIBUTING.md` dan `AGENTS.md` setiap repositori:

1. **Buat branch** dari `main` dengan branch berumur pendek yang deskriptif (misalnya `fix/ssg-scoped-css`).
2. **Commit** memakai [Conventional Commits](https://www.conventionalcommits.org/) (`feat`, `fix`, `docs`, `refactor`, `test`, `chore`), dalam langkah kecil dan terfokus.
3. **Jalankan quality gate** sebelum membuka pull request:

   ```bash
   gofmt -l .        # harus tidak mencetak apa pun
   go vet ./...
   go test ./...
   kiw compat        # setelah perubahan version.go
   ```

   Untuk situs ini: `kiw build` harus berhasil, `python3 scripts/check-i18n-parity.py` harus lolos untuk dokumentasi, dan `python3 scripts/sync-i18n.py` harus dijalankan setelah mengubah `lang/*/site.json`.
4. **Buka pull request** ke `main`. Perubahan ditinjau sebelum di-merge.
5. **Tulis dalam bahasa Inggris** di repositori (kode, commit, dokumentasi); konten bahasa Indonesia ditempatkan di berkas locale `id`.

> Untuk panduan setup pengembang lokal, gunakan panduan publik [Instalasi](/getting-started/installation).

---

## Kanal komunitas

| Kanal | Gunakan untuk | Status (2026-11-09) |
| :--- | :--- | :--- |
| [GitHub Issues](https://github.com/krewire) (setiap repositori publik) | Bug, permintaan fitur, masalah dokumentasi. | Aktif |
| [GitHub Discussions di `krewire/krewire.com`](https://github.com/krewire/krewire.com/discussions) | Pertanyaan, ide, berbagi karya. | Diaktifkan; belum ada diskusi |
| Email: [hello@krewire.com](mailto:hello@krewire.com) | Kontak umum. | Kontak yang dipublikasikan |
| Instagram: [@krewire](https://instagram.com/krewire) | Pengumuman. | Kanal sosial, bukan kanal dukungan |

Kerentanan keamanan **tidak** boleh dilaporkan di issue publik; lihat [Kematangan, Roadmap & Dukungan](/project/maturity-and-roadmap) untuk jalur pelaporan saat ini.

---

## Tata kelola dan maintainership saat ini

Agar tidak menyiratkan struktur yang belum ada, berikut kondisi saat ini apa adanya:

- **Pengambilan keputusan:** setiap repositori dirawat oleh Krewire Team; perubahan diputuskan melalui diskusi issue dan review pull request oleh para maintainer tersebut.
- **Tata kelola formal:** belum ada piagam tata kelola, komite pengarah, proses voting RFC, atau jenjang kontributor yang berjalan.
- **Kode etik:** belum ada yang diadopsi. Sampai ada, kontributor diminta bersikap saling menghormati dan konstruktif.
- **Sponsorship:** belum ada program sponsorship yang aktif; GitHub Sponsors belum diaktifkan untuk organisasi ini.
- **Bagaimana tanggung jawab dapat berkembang:** kontributor yang memberi kontribusi berkualitas secara konsisten dapat diundang maintainer untuk membantu triase atau review. Proses maintainership tertulis **belum didefinisikan**; ketika sudah, proses itu akan dipublikasikan di halaman ini.

Beberapa halaman komunitas (misalnya kode sumber portal komunitas di [`krewire/krewire.github.io`](https://github.com/krewire/krewire.github.io)) menggambarkan jenjang kontributor atau tingkatan sponsorship. Per 2026-11-09 hal tersebut belum berjalan; halaman inilah yang mencerminkan kondisi saat ini.

*Pemilik konten: Krewire Team. Tinjau halaman ini setiap kali tata kelola, kanal, atau panduan kontribusi berubah, dan pada setiap rilis situs yang di-tag.*
