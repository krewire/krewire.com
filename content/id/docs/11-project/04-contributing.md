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
| **Kode** | Memperbaiki bug, meningkatkan paket, atau memperluas `kiw`. | Monorepo inti (`krewire`), AI agent (`krew`), UI kit (`ui`) |
| **Dokumentasi** | Mengoreksi, memperjelas, atau memperluas halaman dokumentasi. Dokumentasi berada di `content/<locale>/docs/` sebagai Markdown. | Repositori dokumentasi (`krewire.com`) |
| **Terjemahan** | Menjaga konsistensi halaman bahasa Inggris dan Indonesia; menerjemahkan halaman baru. | Repositori dokumentasi (`content/id/`, `lang/id/`) |
| **Contoh & starter** | Meningkatkan template proyek atau menulis contoh kerja. | Direktori `templates/` di `krewire` |
| **Pengujian** | Mencoba sebuah workload dari awal sampai akhir di platform Anda dan melaporkan yang rusak, beserta langkah reproduksinya. | Issue di repositori terkait |
| **Triase issue** | Mereproduksi bug yang dilaporkan, meminta detail yang kurang, dan menyarankan label (`bug`, `documentation`, `good first issue`, …). | Issue di repositori publik mana pun |
| **Dukungan komunitas** | Menjawab pertanyaan dan berbagi apa yang Anda bangun. | GitHub Discussions |

Kontribusi saat ini difokuskan pada repositori open-source inti (`krewire`, `krew`, `ui`, `krewire.com`).

---

## Menemukan kontribusi pertama

1. Cari issue berlabel `good first issue` atau `help wanted` di seluruh organisasi.
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

- **GitHub Issues:** Laporan bug dan usulan fitur di repositori publik Krewire.
- **GitHub Discussions:** Tanya jawab, ide, dan diskusi arsitektur di `krewire/krewire.com`.
- **Email:** Kontak umum via [hello@krewire.com](mailto:hello@krewire.com).

Kerentanan keamanan **tidak** boleh dilaporkan di issue publik; lihat [Kematangan, Roadmap & Dukungan](/project/maturity-and-roadmap) untuk jalur pelaporan keamanan.

---

## Tata kelola dan maintainership

- **Pengambilan keputusan:** Repositori dirawat oleh Krewire Team; perubahan ditinjau dan digabungkan melalui pull request.
- **Standar:** Seluruh kode harus lolos quality gate otomatis (`gofmt`, `go vet`, `go test ./...`) dengan CGO dinonaktifkan (`CGO_ENABLED=0`).
- **Pengembangan:** Kontributor yang konsisten dapat diundang untuk membantu merawat dan meninjau domain serta paket tertentu.
