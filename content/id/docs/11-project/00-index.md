---
title: "Tentang Krewire"
description: "Misi Krewire, penerima manfaat yang dituju, filosofi yang mengutamakan dampak, peran keberlanjutan finansial, dan visi jangka panjang."
date: "2026-11-09"
---

# Tentang Krewire

Bab ini menjelaskan **mengapa Krewire ada**, untuk siapa Krewire dibuat, apa saja yang terbuka, bagaimana layanan platform opsional berhubungan dengan inti open-source, cara berpartisipasi, serta bagaimana proyek melaporkan kematangan dan dampaknya. Setiap pernyataan di sini menggambarkan **komitmen saat ini** yang dapat diperiksa terhadap repositori publik, atau **aspirasi jangka panjang** yang diberi label secara eksplisit.

---

## Misi

> **Misi Krewire adalah membuatnya mendekati gratis — baik dalam biaya finansial, beban operasional, maupun beban kognitif — bagi siapa pun untuk membangun, memiliki, dan mengoperasikan produk digital yang aman, berkelanjutan, dan berskala besar.**

Kalimat yang sama menjadi landasan bab [Filosofi](/overview/philosophy) dan halaman utama. Misi inilah alasan ekosistem ini ada; setiap alat dan layanan adalah sarana untuk mencapainya, bukan tujuan itu sendiri.

### Untuk siapa misi ini

| Penerima manfaat | Mengapa Krewire relevan | Prioritas |
| :--- | :--- | :--- |
| **Pengembang independen** | Satu bahasa (Go), satu CLI (`kiw`), dan satu berkas konfigurasi (`krewire.yaml`) mengurangi jumlah alat yang harus dipelajari dan dirawat oleh satu orang. | Utama |
| **Tim engineering kecil** | Konvensi yang sama untuk delapan jenis workload mengurangi peralihan konteks dan biaya orientasi anggota baru. | Utama |
| **Pelajar dan pendidik** | Kode sumber inti bersifat publik dan berlisensi MIT, sehingga dapat dibaca, diajarkan, dan dimodifikasi. | Sekunder |
| **Maintainer open-source** | Repositori publik, Conventional Commits, dan kontrak rilis yang terdokumentasi (`kiw release`, `kiw compat`). | Sekunder |
| **Organisasi** | Binary statis tunggal dan keluaran yang dapat di-host sendiri mengurangi ketergantungan pada runtime proprietary. | Sekunder |

Lihat [Audiens & Kasus Penggunaan](/project/audiences-and-use-cases) untuk kasus penggunaan konkret per audiens.

---

## Dampak lebih dulu, keberlanjutan untuk melayani misi

Tujuan utama Krewire adalah **dampak jangka panjang pada cara perangkat lunak dibangun**: menjadikan default yang aman, arsitektur yang mudah dirawat, dan operasi berbiaya rendah sebagai jalan yang mudah, bukan jalan yang mahal.

- **Dampak adalah ukuran keberhasilan.** Pendapatan bukan metrik keberhasilan utama ekosistem. Cara proyek berniat mengukur dampak — dan pengukuran mana yang belum ada — dipublikasikan di [Laporan Dampak & Keberlanjutan](/project/impact-report).
- **Akses terbuka adalah default.** Repositori inti (`krewire`, `krew`, `ui`, `krewire.com`, `krewire.github.io`) bersifat publik dan berlisensi MIT. Membangun, menjalankan, dan men-deploy dengan tooling inti tidak memerlukan akun Krewire, kunci lisensi, ataupun layanan yang di-host Krewire.
- **Keberlanjutan finansial ada untuk mendanai misi.** Merawat sebuah ekosistem membutuhkan waktu dan infrastruktur. Layanan platform opsional (Dash, Cloud, Auth, Forge) dikembangkan di repositori privat dan mungkin ditawarkan secara komersial di masa depan untuk membantu mendanai pemeliharaan dan pertumbuhan komunitas. Ini **tidak** berarti semua layanan harus gratis, dan **tidak** menjadikan inti open-source sebagai produk berbayar.

> **Status saat ini (2026-11-09):** belum ada paket berbayar, harga, atau lisensi komersial yang dipublikasikan, dan belum ada sumber pendanaan yang dilaporkan. Jika hal itu berubah, perubahannya akan didokumentasikan di [Inti & Layanan Platform](/project/core-and-platform-services) dan [Laporan Dampak & Keberlanjutan](/project/impact-report).

### Misi ekosistem vs. proposisi nilai produk

| | Misi ekosistem | Proposisi nilai produk individual |
| :--- | :--- | :--- |
| **Pertanyaan yang dijawab** | Mengapa Krewire ada? | Mengapa seseorang memakai alat atau layanan tertentu ini? |
| **Cakupan** | Seluruh repositori, dokumentasi, dan kerja komunitas. | Satu produk (misalnya CLI `kiw` atau Krewire Cloud). |
| **Ukuran keberhasilan** | Dampak ekosistem jangka panjang (lihat laporan dampak). | Kegunaan khusus produk; untuk layanan opsional, mungkin juga pendapatan yang mendanai ekosistem. |
| **Komersial?** | Tidak. | Tooling open-source: tidak. Layanan platform opsional: mungkin, di masa depan. |

---

## Keselarasan dengan tiga pilar

Misi adalah *mengapa*; tiga pilar engineering adalah *bagaimana*. Setiap kapabilitas diharapkan lolos ketiganya sebelum dirilis (lihat [Filosofi](/overview/philosophy)):

| Pilar | Bagaimana pilar melayani misi |
| :--- | :--- |
| **Aman (Secure)** | Keamanan adalah default, bukan tambahan berbayar, sehingga pembangun tanpa tim keamanan tetap memulai dari default yang masuk akal. |
| **Berkelanjutan (Sustainable)** | Perangkat lunak yang tetap dapat dibaca, dijalankan, dan dibiayai oleh satu pembangun selama bertahun-tahun — binary tunggal, modul opsional, inti berlisensi MIT. |
| **Berskala (Scalable)** | Pertumbuhan bersifat aditif, bukan penulisan ulang, sehingga pilihan awal tidak memaksa migrasi mahal di kemudian hari. |

---

## Visi jangka panjang (aspirasi)

Berikut adalah **aspirasi jangka panjang**, bukan klaim saat ini atau hasil yang sudah diukur:

- Satu pembangun masih dapat membaca, menjalankan, dan membiayai apa yang ia bangun dengan Krewire sepuluh tahun dari sekarang.
- Belajar membangun dengan Krewire menumbuhkan keterampilan engineering yang dapat dipindahkan (Go, standard library, arsitektur eksplisit), bukan ketergantungan pada Krewire itu sendiri.
- Ekosistem ditopang oleh komunitas kontributor dan maintainer di luar penulis aslinya.
- Layanan platform opsional mendanai pemeliharaan inti open-source tanpa pernah menjadi prasyarat baginya.

Kemajuan menuju aspirasi ini belum diukur. [Laporan Dampak & Keberlanjutan](/project/impact-report) mencantumkan apa yang diukur saat ini dan apa yang belum.

---

## Hal yang tidak diklaim bab ini

- Bab ini tidak melaporkan angka adopsi, jumlah pengguna, atau hasil sosial — belum ada yang diukur.
- Bab ini tidak menjelaskan badan hukum, struktur nirlaba, atau badan tata kelola formal — belum ada yang didefinisikan secara publik.
- Bab ini tidak mempublikasikan harga — belum ada harga yang dipublikasikan.

---

## Isi bab ini

1. [**Audiens & Kasus Penggunaan**](/project/audiences-and-use-cases) — Untuk siapa Krewire dan cara mengadopsinya secara bertahap.
2. [**Keterbukaan & Self-Hosting**](/project/openness-and-self-hosting) — Lisensi, repositori sumber, dan apa yang berjalan tanpa layanan yang di-host Krewire.
3. [**Inti & Layanan Platform**](/project/core-and-platform-services) — Batas antara inti open-source dan layanan opsional.
4. [**Kontribusi & Komunitas**](/project/contributing) — Cara berpartisipasi, tidak hanya lewat kode.
5. [**Laporan Dampak & Keberlanjutan**](/project/impact-report) — Apa yang diukur, apa yang belum, dan bagaimana proyek didanai.
6. [**Kematangan, Roadmap & Dukungan**](/project/maturity-and-roadmap) — Label status, roadmap, ekspektasi pemeliharaan, dan jalur pelaporan.

*Pemilik konten: Krewire Team (maintainer [`krewire/krewire.com`](https://github.com/krewire/krewire.com)). Ditinjau pada setiap rilis situs yang di-tag.*
