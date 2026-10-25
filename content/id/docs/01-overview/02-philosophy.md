---
title: "Filosofi Inti"
description: "Filosofi dasar, tiga pilar, dan prinsip rekayasa yang mendasari desain Krewire serta ekosistem SDLC digitalnya."
date: "2026-10-14"
---

# Filosofi Inti

Krewire dipandu oleh sebuah keyakinan fundamental: **membuatnya mendekati gratis — baik dalam biaya finansial, beban operasional, maupun beban kognitif — bagi siapa pun untuk membangun, memiliki, dan mengoperasikan produk digital yang aman, berkelanjutan, dan berskala besar.**

Perangkat lunak seharusnya mengubah pembangun menjadi produsen yang berdaulat, bukan penyewa abadi di atas platform kepemilikan tertutup (proprietary). Krewire dihadirkan sebagai **Pustaka Modular**: paket Go independen dengan satu tujuan yang Anda padukan di balik satu berkas konfigurasi (`krewire.yaml`) dan satu CLI terpadu (`kiw`).

---

## 1. Tiga Pilar

Setiap keputusan arsitektur, paket pustaka, dan kapabilitas dalam ekosistem Krewire harus memenuhi tiga pilar tanpa kompromi:

| Pilar | Prinsip | Penerapan Nyata di Krewire |
| :--- | :--- | :--- |
| **Secure** | Keamanan adalah standar bawaan, bukan fitur berbayar tambahan. | Mengutamakan pustaka standar Go, grafik dependensi menurun (downward-only), primitif HTTP aman secara default, kepatuhan OWASP/CWE, runtime terisolasi (sandbox), dan rahasia yang direferensikan via `${env:NAME}` tanpa disimpan di kode. |
| **Sustainable** | Perangkat lunak yang masih dapat dibaca, dijalankan, dan dibiayai pengembang puluhan tahun kemudian. | Binary statis tunggal dengan template dan aset tersemat, bebas lock-in runtime, baterai opsional tanpa beban saat tak terpakai, dan fondasi open-core berlisensi MIT permanen. |
| **Scalable** | Pertumbuhan bertahap dan aditif, bukan perombakan arsitektur total. | Kontrak provider yang dapat dibalikkan, tipe workload terstandarisasi, dan alur progresif yang memungkinkan aplikasi tumbuh dari VPS $5 hingga hyperscale tanpa penulisan ulang logika bisnis. |

---

## 2. Tujuh Prinsip Rekayasa Mendasar

Tiga pilar di atas adalah hasil akhir (outcomes); tujuh prinsip berikut adalah cara rekayasa teknis untuk mencapainya:

### 1. Kepemilikan Penuh adalah Kedaulatan (Ownership is Power)

Perangkat lunak yang Anda sewa pada akhirnya akan mendikte Anda. Ketika platform pihak ketiga mengendalikan alur build, runtime, atau target deployment, mereka dapat menentukan margin harga, biaya, dan aturannya.

Krewire merancang alat yang hasilnya dimiliki sepenuhnya oleh pembuatnya:
- **Deployment Binary Tunggal:** Aplikasi dikompilasi menjadi kode mesin mandiri yang terhubung secara statis.
- **Aset Tersemat:** Template, migrasi database, dan aset web berada di dalam binary via paket `embed` Go.
- **Data Berdaulat Sendiri:** Database, penyimpanan objek, dan penyedia cache terlepas dari keterikatan vendor (vendor lock-in).
- **Nol Dependensi Runtime:** Tidak memerlukan interpreter bahasa asing, runtime container berat, atau cloud engine tertutup.

### 2. Biaya adalah Keputusan Desain (Cost is a Design Decision)

Biaya mendekati nol dirancang secara sadar, bukan dari potongan diskon sementara. Setiap dependensi yang tidak perlu, lapisan indireksi rumit, dan trik pemrograman yang berlebihan akan membebankan bunga selamanya — berupa tagihan hosting, beban pemeliharaan, audit keamanan, dan waktu onboarding.

Krewire menolak tumpukan teknologi yang secara diam-diam membebani pengembang dengan biaya operasional tinggi:
- **Pustaka Standar Pertama (Stdlib First):** Maksimalkan pustaka bawaan Go yang berperforma tinggi (`net/http`, `html/template`, `log/slog`, `context`, `flag`) sebelum memasukkan dependensi eksternal pihak ketiga.
- **Kemasan Modular:** Proyek hanya mengimpor apa yang digunakan. Jika aplikasi hanya memerlukan pembangkit situs statis, modul antrean worker tidak akan pernah ditautkan ke hasil kompilasi.
- **Komputasi Efisien:** Penggunaan memori yang ramping (10–30 MB) dan waktu booting sub-milidetik memungkinkan aplikasi produksi berjalan stabil di perangkat keras ekonomis.
- **Open Core Permanen:** Tingkat fondasi tetap berlisensi MIT selamanya. Tidak ada lisensi berbayar yang menjadi prasyarat wajib untuk membangun, men-deploy, atau mengoperasikan sistem lengkap.

### 3. Satu Jalur yang Jelas Melampaui Banyak Pilihan (One Path Beats Many Choices)

Fragmentasi membebani paling berat mereka yang memiliki sumber daya terbatas — tim pemula, pelajar, dan pengembang independen. Pilihan yang tak terbatas menciptakan kelelahan dalam mengambil keputusan (decision fatigue).

Berkomitmen pada satu toolchain yang koheren mengubah eksplorasi yang terpecah menjadi penguasaan yang terakumulasi:
- **Satu Konfigurasi:** `krewire.yaml` adalah satu-satunya sumber kebenaran (SSOT) untuk seluruh definisi workload, binding environment, dan target build.
- **Satu CLI:** `kiw` menangani pembuatan proyek, pengembangan lokal, pengujian, kompilasi, dan deployment di seluruh tipe workload.
- **Satu Bahasa:** Go mentenagai seluruh tumpukan — dari layanan backend dan antrean worker hingga situs statis dan runtime WebAssembly client-side.
- **Model Mental Universal:** Delapan jenis workload berbagi semantik direktori, sintaks konfigurasi, dan fase siklus hidup yang identik.

### 4. Cukup Sederhana untuk Diajarkan, Cukup Kokoh untuk Dipercaya

Keamanan, keberlanjutan, dan keandalan adalah hasil dari disiplin arsitektural, bukan fitur yang ditambahkan belakangan. Keamanan berasal dari permukaan serangan yang minimal dan pengaturan bawaan yang waras; keberlanjutan lahir dari komponen yang teruji; skalabilitas tercipta dari kejujuran rekayasa tanpa abstraksi spekulatif.

- **Eksplisit Melampaui Magis:** Tidak ada trik reflection tersembunyi, metaprogramming yang sulit dibaca, atau state global implisit.
- **Arsitektur yang Terbaca:** Setiap paket memiliki tanggung jawab tunggal (SRP) dan batas kohesi yang terang.
- **Kemudahan Audit:** Basis kode yang dibangun bersama Krewire dapat dibaca, dipahami, dan diaudit oleh insinyur junior maupun peninjau keamanan pihak ketiga tanpa perlu pelatihan platform tertutup.

### 5. Belajar Sambil Merilis, Merilis Sambil Belajar (Learn by Shipping, Ship by Learning)

Kemampuan rekayasa yang sesungguhnya dibentuk di atas sistem produksi nyata, bukan lewat latihan sintetis semata. Pembelajar harus membangun perangkat lunak yang benar-benar dijalankan oleh bisnis nyata, dan bisnis selayaknya mengoperasikan perangkat lunak yang dapat dibedah dan dipahami oleh pembelajar.

- **Scaffolding Kelas Produksi:** Perintah `kiw new` langsung menyediakan fondasi arsitektur nyata dengan penanganan error terstruktur, health check, dan logging sejak menit pertama.
- **Dokumentasi sebagai Produk Kelas Utama:** Panduan komprehensif berversi, dokumen keputusan arsitektur (ADR), dan contoh yang dapat dieksekusi menyertai setiap paket.
- **SDLC Berbasis Komunitas:** Membangun ekosistem pengembang inklusif tempat kolaborasi terbuka menjembatani dunia pendidikan dan industri nyata.

### 6. Keterbukaan adalah Cara Bekerja, Bukan Sekadar Lisensi

Kepercayaan dibangun di ruang terbuka. Krewire berkomitmen pada transparansi penuh di seluruh basis kode, peta jalan (roadmap), dan standar rekayasanya.

- **Evolusi Transparan:** Setiap keputusan, kompromi arsitektural, dan pilihan desain dicatat secara publik dalam catatan yang terkontrol versinya.
- **Jenjang Open-Core yang Jujur:** Tangga lisensi (`free` → `pro` → `team` → `enterprise`) bersifat aditif secara ketat. Tingkat yang lebih tinggi dapat membungkus, memperkuat keamanan, atau menambahkan otomasi kepatuhan untuk enterprise, namun modul tingkat gratis tidak akan pernah dibatasi atau dipotong fiturnya demi memaksa upgrade.
- **Standar yang Sama:** Seluruh pengujian, linter, dan gerbang mutu berjalan identik baik di laptop pengembang lokal maupun pada pipeline CI/CD terpusat.

### 7. Pustaka Modular Melampaui Framework Monolitik

Framework memaksakan konvensi kaku, jebakan siklus hidup, dan hierarki kelas yang mengikat ke dalam aplikasi Anda. Saat kebutuhan bisnis berkembang, tim sering kali terjebak dalam perombakan kode besar-besaran.

Krewire mendekati arsitektur melalui pustaka modular yang dapat dikomposisikan:
- **Kendali Kontrol Tetap di Tangan Anda:** Anda merangkai paket-paket Krewire ke dalam `main.go` aplikasi Anda sendiri, bukan mewarisi wadah runtime beropini kaku.
- **Evolusi Progresif:** Mulai dari situs dokumentasi statis atau landing page (`site`), tambahkan rute server dan model database (`app`), ekstrak komputasi berat ke worker latar belakang (`worker`), dan pecah menjadi microservices otonom (`service`) hanya ketika skala organisasi membutuhkannya.
- **Nol Biaya Penulisan Ulang:** Seluruh komponen, logika domain, dan konfigurasi dapat digunakan kembali secara utuh di setiap tahapan ekspansi sistem.

---

## 3. Rangkuman Pemetaan

Ketujuh prinsip rekayasa di atas secara langsung menopang dan mewujudkan Tiga Pilar Krewire: **Secure**, **Sustainable**, dan **Scalable**.

Selanjutnya: telusuri [**1.3 Workload Krewire →**](/overview/krewire-workloads) atau pelajari [**1.4 Panduan Pembaruan →**](/overview/upgrade-guide).
