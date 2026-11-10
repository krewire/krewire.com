---
title: "Laporan Dampak & Keberlanjutan"
description: "Laporan berbasis bukti tentang kontribusi Krewire bagi ekosistem pengembang dan cara proyek ini dipertahankan: dimensi dampak, definisi metrik, periode pelaporan, sumber data, dan data yang secara eksplisit belum tersedia."
date: "2026-11-09"
---

# Laporan Dampak & Keberlanjutan

Laporan ini menjelaskan apa yang dikontribusikan Krewire dan bagaimana pekerjaan itu dipertahankan. Laporan ini sengaja dibuat ringkas: hanya mempublikasikan angka yang dapat diperiksa, dan menyatakan dengan jelas di mana data belum ada.

| Kolom | Nilai |
| :--- | :--- |
| **Edisi** | Baseline — laporan pertama |
| **Periode pelaporan** | Awal proyek (commit publik pertama, 2026-08-24) hingga **2026-11-09** |
| **Pemilik** | Krewire Team (maintainer [`krewire/krewire.com`](https://github.com/krewire/krewire.com)) |
| **Pembaruan berikutnya** | Pada rilis situs ini berikutnya yang di-tag |

---

## Cara membaca laporan ini

Krewire membedakan tiga tingkat, dan hanya menganggap dua tingkat terakhir sebagai dampak:

| Tingkat | Arti | Contoh |
| :--- | :--- | :--- |
| **Keluaran (output)** | Hal yang dihasilkan proyek. Perlu, tetapi bukan dampak dengan sendirinya. | Rilis, halaman dokumentasi, starter. |
| **Hasil (outcome)** | Perubahan bagi orang yang memakai apa yang dihasilkan. | Build pertama yang berhasil, proyek di produksi, kontributor eksternal. |
| **Dampak jangka panjang** | Perubahan yang bertahan pada cara perangkat lunak dibangun, selaras dengan [misi](/project). | Pembangun yang masih dapat menjalankan dan membiayai perangkat lunaknya bertahun-tahun kemudian. |

**Metrik kosmetik (vanity metrics) tidak dipakai sebagai bukti dampak.** Bintang GitHub, fork, jumlah unduhan mentah, dan page view tidak dilaporkan sebagai dampak, karena tidak menunjukkan bahwa ada yang berhasil dengan Krewire.

---

## Keluaran (diukur)

| Metrik | Definisi | Nilai | Periode | Sumber |
| :--- | :--- | :--- | :--- | :--- |
| Repositori publik | Repositori di organisasi GitHub `krewire` yang terlihat tanpa autentikasi. | 5 | Per 2026-11-09 | API organisasi GitHub (`public_repos`) |
| Versi inti yang di-tag | Git tag pada `krewire/krewire`. | 1 (`v0.2.0`) | Per 2026-11-09 | [Tag `krewire/krewire`](https://github.com/krewire/krewire/tags) |
| GitHub Release yang dipublikasikan | Rilis dengan aset unduhan di seluruh repositori publik. | 0 | Per 2026-11-09 | GitHub Releases API |
| Starter workload | Jenis workload yang dapat di-scaffold oleh `kiw new`. | 8 | Per v0.2.0 | Flag `kiw new` di `tools/kiw/internal/commands/new.go` |
| Halaman dokumentasi | Halaman Markdown per locale di `content/<locale>/docs/` situs ini (bahasa Inggris dan Indonesia dijaga 1:1). | 135 per locale | Per 2026-11-09 | `scripts/check-i18n-parity.py` di repositori ini |

---

## Hasil (sebagian besar belum diukur)

| Metrik | Definisi | Nilai | Sumber / catatan |
| :--- | :--- | :--- | :--- |
| Proyek aktif yang memakai Krewire | Proyek berbeda di luar organisasi Krewire yang dibangun dengan `kiw` dalam periode ini. | **Belum diukur** | `kiw` tidak memiliki telemetri; metode sukarela yang menghormati privasi belum didefinisikan. |
| Keberhasilan build pertama | Persentase pengguna baru yang `kiw new` + `kiw build` pertamanya berhasil. | **Belum diukur** | Belum ada studi onboarding atau pelaporan opt-in. |
| Penggunaan produksi yang terverifikasi | Deployment publik yang dibangun dengan Krewire, dapat diverifikasi melalui URL dan repositori. | 1 — situs ini | [krewire.com](https://krewire.com) dibangun dengan `kiw build` oleh workflow deploy repositori ini. Ini situs milik proyek sendiri: bukti dogfooding, bukan adopsi eksternal. |
| Kontributor eksternal | Penulis commit di luar tim maintainer, per periode. | **Belum diukur** | Belum ada daftar maintainer publik untuk memisahkan penulis eksternal dari internal. Sebagai konteks saja: branch default repositori publik memiliki 3 identitas email penulis commit yang berbeda dalam periode ini. |
| Kontributor berulang | Kontributor eksternal yang aktif dalam dua periode atau lebih. | **Belum diukur** | Membutuhkan setidaknya dua periode pelaporan. |
| Retensi maintainer & keragaman kontributor | Maintainer yang aktif lintas periode; sebaran kontributor lintas organisasi dan wilayah. | **Belum diukur** | Belum ada daftar maintainer atau survei kontributor. |
| Adopsi pendidikan / komunitas | Kursus, lokakarya, atau kelompok komunitas yang memakai Krewire. | **Belum diukur** | Belum ada yang dilaporkan ke proyek. |

---

## Dampak jangka panjang (aspirasi, belum diukur)

Aspirasi jangka panjang tercantum di [Tentang Krewire](/project). Belum ada metrik untuknya, dan aspirasi tersebut tidak akan dilaporkan sebagai tercapai tanpa bukti.

---

## Keberlanjutan

Pendapatan **bukan** ukuran keberhasilan utama ekosistem (lihat [Tentang Krewire](/project)). Informasi keuangan dipublikasikan agar pengguna dapat menilai apakah proyek ini dapat terus dirawat.

| Butir | Status (2026-11-09) | Sumber / catatan |
| :--- | :--- | :--- |
| Sumber pendanaan | **Belum ada yang dilaporkan** | Belum ada hibah, sponsorship, atau pendapatan yang diungkapkan. |
| Sponsorship | **Belum aktif** | `krewire/krewire` memuat `FUNDING.yaml` yang merujuk GitHub Sponsors, tetapi GitHub Sponsors belum diaktifkan untuk organisasi ini. |
| Pendapatan dari layanan platform | **Tidak ada** | Belum ada layanan platform yang tersedia untuk publik atau diberi harga ([Inti & Layanan Platform](/project/core-and-platform-services)). |
| Biaya operasional | **Belum dipublikasikan** | Infrastruktur yang diketahui: situs ini berjalan sebagai kontainer nginx yang di-deploy ke VPS oleh workflow deploy repositori. Biayanya belum diungkapkan. |
| Cakupan biaya operasional & keragaman pendanaan | **Belum dapat diukur** | Membutuhkan biaya dan sumber pendanaan yang dipublikasikan. |

---

## Proses pembaruan

Agar pelaporan tetap sesuai dengan kapasitas proyek saat ini:

1. **Kapan:** pada setiap rilis situs ini yang di-tag, atau setidaknya setiap enam bulan.
2. **Bagaimana:** tambahkan nilai periode baru di samping nilai sebelumnya pada tabel di atas; pertahankan nilai lama untuk perbandingan. Laporan ini berupa Markdown biasa, sehingga pembaruannya tidak memerlukan desain ulang situs.
3. **Aturan:** setiap metrik mempertahankan definisi, periode, dan sumbernya; data yang belum tersedia tetap diberi label "belum diukur"; estimasi tidak pernah disajikan sebagai hasil pengukuran.
4. **Menambah metrik:** usulkan metrik baru melalui issue di [`krewire/krewire.com`](https://github.com/krewire/krewire.com/issues) beserta definisi dan sumber data yang andal sebelum dipublikasikan.
