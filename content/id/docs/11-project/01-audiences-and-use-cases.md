---
title: "Audiens & Kasus Penggunaan"
description: "Untuk siapa Krewire, masalah konkret yang ditanganinya, contoh kasus penggunaan yang dipetakan ke workload dan dokumentasi, serta cara mengadopsinya secara bertahap berdampingan dengan stack Anda."
date: "2026-11-09"
---

# Audiens & Kasus Penggunaan

Halaman ini membantu Anda cepat menilai apakah Krewire relevan dengan pekerjaan Anda dan dari mana harus memulai. Halaman ini hanya memetakan kasus penggunaan ke kapabilitas yang sudah ada di repositori inti publik saat ini; kematangan tiap workload tercantum di [Kematangan, Roadmap & Dukungan](/project/maturity-and-roadmap).

---

## Audiens utama

Fokus awal Krewire adalah **pengembang independen** dan **tim engineering kecil**.

### Pengembang independen

**Masalah:** satu orang yang merawat sebuah produk sering harus mempelajari dan memperbarui beberapa toolchain sekaligus — bundler JavaScript untuk front end, framework backend terpisah, generator dokumentasi, antrean tugas, dan skrip deployment.

**Yang ditawarkan Krewire:** satu bahasa (Go), satu CLI (`kiw`), dan satu berkas konfigurasi (`krewire.yaml`) untuk delapan jenis workload, dengan keluaran binary tunggal atau statis yang berjalan di perangkat keras sederhana.

| Kasus penggunaan | Mulai dengan | Dokumentasi |
| :--- | :--- | :--- |
| Menerbitkan landing page, portofolio, atau blog | `kiw new my-site --site` | [Situs Statis (`site`)](/workloads/site) |
| Menerbitkan dokumentasi produk atau buku teknis | `kiw new my-docs --book` | [Buku Dokumentasi (`book`)](/workloads/book) |
| Merilis aplikasi web kecil atau JSON API sebagai satu binary | `kiw new my-app --app` | [Monolit Fullstack (`app`)](/workloads/app) |
| Membangun alat baris perintah | `kiw new my-tool --cli` | [CLI & TUI (`cli`)](/workloads/cli) |

### Tim engineering kecil

**Masalah:** tim kecil menanggung biaya koordinasi tinggi ketika setiap layanan memakai konvensi, format konfigurasi, dan pipeline deployment yang berbeda.

**Yang ditawarkan Krewire:** konvensi direktori, konfigurasi, dan perintah siklus hidup yang sama (`kiw dev`, `kiw build`, `kiw test`, `kiw deploy`) di seluruh workload, sehingga berpindah antar-layanan tidak berarti berganti model mental.

| Kasus penggunaan | Mulai dengan | Dokumentasi |
| :--- | :--- | :--- |
| Memindahkan pekerjaan lambat (email, laporan) keluar dari jalur request | `kiw new jobs --worker` | [Worker Latar Belakang (`worker`)](/workloads/worker) |
| Menambahkan layanan headless di samping aplikasi yang ada | `kiw new api --service` | [Microservices (`service`)](/workloads/service) |
| Menaruh satu titik masuk di depan beberapa layanan upstream | `kiw new edge --gateway` | [Gateway API (`gateway`)](/workloads/gateway) |
| Deploy ke VPS tanpa platform kontainer | `kiw deploy --target systemd` | [Deployment & Operasi](/guides/deployment-and-operations) |
| Menjaga pengembangan berbantuan agen AI tetap konsisten | `AGENTS.md` + starter Boost | [Pengembangan Agentik](/getting-started/agentic-development) |

---

## Audiens sekunder

Audiens berikut mendapat manfaat dari fondasi terbuka yang sama, tetapi situs ini belum menyediakan materi khusus untuk mereka.

| Audiens | Contoh kasus penggunaan | Titik awal |
| :--- | :--- | :--- |
| **Pelajar & pendidik** | Membaca dan mengajarkan basis kode Go lengkap berlisensi MIT yang mencakup web, CLI, worker, dan infrastruktur. | [Filosofi](/overview/philosophy), [Disiplin Arsitektur](/architecture/layered-pyramid), kode sumber di `krewire/krewire` |
| **Maintainer open-source** | Membangun situs dokumentasi proyek dengan `site` atau `book`, seperti situs ini. | [Situs Statis (`site`)](/workloads/site), [Kontribusi & Komunitas](/project/contributing) |
| **Organisasi** | Mengevaluasi keluaran binary tunggal yang dapat di-host sendiri serta batas open-core sebelum berkomitmen. | [Keterbukaan & Self-Hosting](/project/openness-and-self-hosting), [Inti & Layanan Platform](/project/core-and-platform-services) |

---

## Adopsi bertahap — tanpa migrasi seluruh stack

Krewire dirancang untuk diadopsi satu workload demi satu workload. Anda **tidak** perlu memigrasikan seluruh stack.

1. **Mulai dari satu workload yang terisolasi.** Situs dokumentasi (`book`) atau landing page (`site`) tidak memiliki ketergantungan runtime pada sistem Anda yang lain.
2. **Tempatkan berdampingan, jangan mengganti.** Workload Krewire dapat berada di direktorinya sendiri (dengan `go.mod` dan `krewire.yaml` sendiri) di dalam repositori Node.js, Laravel, atau PHP yang sudah ada. Menurut panduan interoperabilitas di repositori inti, `kiw` mendeteksi proyek host tersebut (`package.json`, `composer.json`) tetapi tidak pernah mengubah manifes host maupun direktori dependensinya.
3. **Integrasikan lewat antarmuka standar.** Workload berkomunikasi melalui HTTP atau input/output standar, dan keluarannya berupa binary Go biasa atau berkas statis yang dapat disajikan web server mana pun. Situs ini, misalnya, dibangun dengan `kiw build` dan disajikan oleh kontainer nginx standar (lihat `Dockerfile` di repositori).
4. **Tumbuh hanya saat dibutuhkan.** Tambahkan workload `worker`, `service`, atau `gateway` ketika kebutuhan menuntutnya; paket yang tidak dipakai tidak dikompilasi ke dalam keluaran.

```text
existing-project/            # Node.js, Laravel, PHP … (tidak diubah)
├── package.json
├── src/
└── docs/                    # baru: workload book Krewire
    ├── krewire.yaml
    └── content/
```

### Hal yang tidak dilakukan Krewire

- Krewire tidak mentranspilasi JavaScript atau PHP ke Go.
- Krewire tidak memerlukan layanan yang di-host Krewire, akun, atau kunci lisensi (lihat [Keterbukaan & Self-Hosting](/project/openness-and-self-hosting)).
- Krewire belum menerbitkan binary `kiw` siap pakai; skrip instalasi saat ini beralih ke `go install`, yang memerlukan Go 1.27.1+ (lihat [Instalasi](/getting-started/installation)).

---

## Apakah Krewire cocok untuk Anda?

| Krewire kemungkinan cocok jika… | Krewire mungkin belum cocok jika… |
| :--- | :--- |
| Anda nyaman dengan Go, atau ingin menstandarkan pada Go. | Tim Anda tidak dapat mengadopsi Go 1.27.1+. |
| Anda menghargai dependensi minimal dan keluaran yang dapat di-host sendiri. | Anda membutuhkan jaminan kompatibilitas API jangka panjang — Krewire masih pra-1.0 (v0.2.x). |
| Anda ingin mulai kecil (satu situs, satu alat) lalu berkembang. | Anda membutuhkan layanan hosting terkelola hari ini — belum ada yang tersedia untuk publik. |

*Pemilik konten: Krewire Team. Perbarui halaman ini ketika workload, starter, atau label kematangan berubah.*
