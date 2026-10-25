---
title: "Ringkasan Ekosistem"
description: "Selamat datang di Krewire — Pustaka Go Modular untuk Setiap Workload. Pelajari arsitektur inti, filosofi, jenis workload, dan alur kerja pengembang."
date: "2026-09-29"
---

# Ringkasan Ekosistem (Overview)

Selamat datang di **Dokumentasi Krewire**.

**Krewire** adalah ekosistem **SDLC digital menyeluruh** open-source yang disajikan sebagai **Pustaka Modular**: paket Go independen dengan satu tujuan yang Anda komposisi di balik satu CLI (**`kiw`**) dan satu konfigurasi terpadu (**`krewire.yaml`**) untuk membangun, menguji, dan men-deploy **delapan workload berbeda**:

1. **Monolit Web Fullstack (`app`)**
2. **Antarmuka Terminal & TUI (`cli`)**
3. **Situs Statis Berkecepatan Tinggi (`site`)**
4. **Buku Dokumentasi Terstruktur (`book`)**
5. **Worker Latar Belakang & Antrean Tugas (`worker`)**
6. **Microservices Terdistribusi (`service`)**
7. **Infrastruktur Cloud sebagai Kode (`infra`)**
8. **Gateway API (`gateway`)**

---

## Isi Bab Ini

Bab **Ringkasan Ekosistem** ini memperkenalkan filosofi inti, arsitektur teknis, dan alur kerja rekayasa di Krewire:

<div class="overview-grid">

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.1</span>
      <h3 class="chapter-card-title">
        <a href="/overview/why-krewire">Mengapa Krewire? <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pahami bagaimana Krewire memecahkan fragmentasi toolchain modern, menggantikan tumpukan multi-bahasa, dan memungkinkan pertumbuhan sistem progresif tanpa penulisan ulang (rewrite).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.2</span>
      <h3 class="chapter-card-title">
        <a href="/overview/philosophy">Filosofi Inti <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Telusuri filosofi dasar, Tiga Pilar (Secure, Sustainable, Scalable), dan tujuh prinsip rekayasa yang mendasari Krewire.
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.3</span>
      <h3 class="chapter-card-title">
        <a href="/overview/krewire-workloads">Workload Krewire <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Pelajari 8 tipe workload resmi secara mendalam (<code>app</code>, <code>cli</code>, <code>site</code>, <code>book</code>, <code>worker</code>, <code>service</code>, <code>gateway</code>, <code>infra</code>).
    </p>
  </div>

  <div class="chapter-card">
    <div class="chapter-card-head">
      <span class="chapter-card-num">1.4</span>
      <h3 class="chapter-card-title">
        <a href="/overview/upgrade-guide">Panduan Pembaruan <span class="arrow">→</span></a>
      </h3>
    </div>
    <p class="chapter-card-desc">
      Ketahui janji kompatibilitas SemVer tanpa breaking change, migrasi konfigurasi <code>krewire.yaml</code>, dan cara memperbarui CLI serta dependensi pustaka.
    </p>
  </div>

</div>

---

## Tiga Pilar

Setiap kapabilitas yang dirilis oleh Krewire wajib memenuhi tiga pilar utama — **secure**, **sustainable**, dan **scalable** — sebagai standar kelayakan arsitektur:

| Pilar | Ringkasan |
| :--- | :--- |
| **Secure** | Keamanan adalah standar utama: grafik dependensi mengutamakan pustaka standar Go (stdlib), kontrol selaras OWASP/CWE, dan rahasia direferensikan via environment — tidak pernah disimpan di kode. |
| **Sustainable** | Biaya mendekati nol yang dirancang secara sadar: binary statis tunggal, baterai opsional (opt-in), dan tanpa biaya lisensi permanen. |
| **Scalable** | Bertumbuh tanpa menulis ulang: alur progresif dari situs statis hingga layanan mikro dan infrastruktur cloud, dari VPS $5 hingga skala besar. |

## Prinsip Inti Rekayasa

Krewire dibangun berlandaskan empat prinsip rekayasa utama:

1. **Go sebagai Fondasi Arsitektural**  
   Standar Go (`net/http`, `html/template`, `embed`, `context`, `slog`, `flag`) mentenagai setiap lapisan. Deployment binary tunggal dengan aset statis tersemat menjamin startup instan, penggunaan RAM sangat rendah (~15MB), dan bebas dependensi runtime eksternal.
2. **Modular di Setiap Cakupan (SRP & SoC)**  
   Dirancang menggunakan prinsip Tanggung Jawab Tunggal (*Single Responsibility Principle*) dan Kohesi Tinggi (`KWF-5ZHQV`). Mulailah dari monolit modular sederhana, lalu ekstrak worker atau microservice secara bersih saat lalu lintas meningkat tanpa merombak logika bisnis.
3. **Bebas Kelelahan JavaScript (Zero JS Fatigue)**  
   Hasilkan situs statis, portal dokumentasi, dan aplikasi fullstack super cepat menggunakan berkas komponen `.kiw` tanpa memerlukan Node.js, `npm`, webpack, atau pohon dependensi `node_modules` yang rentan rusak.
4. **Pengalaman Pengembang Terpadu**  
   Kelola seluruh siklus hidup perangkat lunak — inisialisasi, pengembangan lokal, kompilasi aset, pengujian, hingga deployment — melalui **satu binary CLI tunggal (`kiw`)** dan satu berkas konfigurasi devtool (**`krewire.yaml`**).

---

## Peta Dokumentasi

Pelajari panduan lengkap untuk ekosistem Krewire:

| Bab | Judul | Fokus |
| :--- | :--- | :--- |
| **1.** | [**Ringkasan Ekosistem**](/) | Arsitektur ekosistem, prinsip inti, tiga pilar, dan gambaran umum workload. |
| **2.** | [**Memulai (Getting Started)**](/getting-started) | Prasyarat, panduan kilat 5 menit, anatomi proyek, SDLC berbasis agen AI, dan DSL `.kiw`. |
| **3.** | [**Workload Mendalam**](/workloads) | Arsitektur rinci untuk 8 jenis workload (`app`, `cli`, `site`, `book`, `worker`, `service`, `gateway`, `infra`). |
| **4.** | [**Referensi CLI**](/cli-reference) | Katalog perintah `kiw`, spesifikasi skema `krewire.yaml`, dan otomasi tugas. |
| **5.** | [**Paket & Pustaka**](/packages) | Katalog pustaka domain: `model`, `lifecycle`, `web`, `app`, `sec`, `auth`, `ui`, `tui`, `storage`, `resilience`, dan `testing`. |
| **6.** | [**Panduan & Arsitektur**](/guides) | Desain arsitektur, ketentuan lisensi open-core, dan panduan deployment produksi. |

---

## Langkah Selanjutnya

Untuk mulai membangun bersama Krewire, telusuri sub-bab berikut secara berurutan:

- Lanjutkan ke [**1.1 Mengapa Krewire? →**](/overview/why-krewire)
- Langsung ke [**2. Memulai (Getting Started) →**](/getting-started)
- Atau pelajari [**3. Workload Mendalam →**](/workloads)
