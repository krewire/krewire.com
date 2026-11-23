---
title: "Buku Panduan AI-Native Engineering"
description: "Buku panduan kanonikal untuk menerapkan disiplin AI-Native Engineering yang terkelola dan berbasis bukti di seluruh ekosistem Krewire."
date: "2026-11-21"
---

# Buku Panduan AI-Native Engineering

> **Inisiatif Referensi**: [Baseline AI-Native Engineering Ekosistem (internal#8)](https://github.com/krewire/internal/issues/8) · [krewire.com#31](https://github.com/krewire/krewire.com/issues/31)  
> **Status**: Panduan Rekayasa Kanonikal  
> **Penyelarasan**: SDLC Terkelola di seluruh Krew, Boost, dan Toolchain `kiw`  

---

## 1. Apa itu AI-Native Engineering?

Rekayasa perangkat lunak modern kini bertransisi melampaui **AI-assisted coding** sederhana (chatbot yang menyalin-tempel potongan kode atau saran kode sebaris tanpa verifikasi) menuju **AI-Native Engineering**.

**AI-Native Engineering** adalah disiplin rekayasa di mana manusia dan agen AI otonom berkolaborasi melalui **konteks proyek eksplisit**, **kontrak terbaca-mesin**, **perangkat eksekusi terkelola**, **gerbang verifikasi deterministik**, dan **evaluasi berkelanjutan**.

```text
┌───────────────────────────────┬───────────────────────────────────┬─────────────────────────────────┐
│ Metrik / Karakteristik        │ AI-Assisted Coding (Lama)         │ AI-Native Engineering (Krewire) │
├───────────────────────────────┼───────────────────────────────────┼─────────────────────────────────┤
│ Model Integrasi               │ Popup chat editor sementara       │ Native MCP & runtime in-process │
│ Penemuan Konteks              │ Salin-tempel prompt ad-hoc        │ SSOT terbaca-mesin (kiw)        │
│ Batasan Keamanan              │ Akses terminal tanpa batas        │ SafePath & gerbang manusia      │
│ Model Verifikasi              │ "Tampak benar di editor"          │ krew verify & rekaman bukti      │
│ Sikap Tanggung Jawab          │ Menyalahkan AI atas bug           │ Manusia memiliki hasil & mutu   │
│ Akumulasi Pengetahuan         │ Hilang saat sesi chat ditutup     │ ADR permanen & aturan konteks   │
└───────────────────────────────┴───────────────────────────────────┴─────────────────────────────────┘
```

Di Krewire, agen AI adalah partisipan rekayasa kelas satu yang tunduk pada **Tiga Pilar (3S)**:
- **Secure (S1)**: Batasan sistem berkas ketat (`SafePath`), eksekusi hak akses minimal, nol kredensial dalam kode, dan mitigasi *prompt injection*.
- **Sustainable (S2)**: Toolchain Go murni (`CGO_ENABLED=0`), nol dependensi runtime Node.js/npm, dan biner statis tunggal yang dapat berjalan lancar di VPS $5/bulan.
- **Scalable (S3)**: Pergantian penyedia yang dapat dibalikkan, layering modular satu arah ke bawah, dan evolusi arsitektur aditif tanpa penulisan ulang (*rewrite*).

---

## 2. Model Pembagian Tanggung Jawab Tiga Arah

Ekosistem Krewire mengoordinasikan rekayasa AI melalui tiga batasan yang terpisah:

```text
               ┌───────────────────────────────────────────────────┐
               │              Krewire Krew (Layer 4)               │
               │            Engine Eksekusi ReAct Otonom           │
               │    Embed · CLI · HTTP/SSE · Model Context Protocol │
               └─────────────────────────┬─────────────────────────┘
                                         │ dipandu oleh
                                         ▼
               ┌───────────────────────────────────────────────────┐
               │             Krewire Boost (Templates)            │
               │          Konstitusi & Peran Multi-Agent           │
               │     AGENTS.md · 21 Peran Spesialis · Workflows    │
               └─────────────────────────┬─────────────────────────┘
                                         │ dijalankan & diverifikasi via
                                         ▼
               ┌───────────────────────────────────────────────────┐
               │             Krewire kiw CLI (Layer 3)             │
               │          Toolchain Deterministik & SSOT           │
               │   kiw context · krew verify · krew test · Compiler  │
               └───────────────────────────────────────────────────┘
```

1. **Penyedia Kebenaran (`krew`)**:
   Menyediakan satu-satunya sumber kebenaran (*single source of truth*) yang terbaca mesin via `kiw context --json`, menjalankan kompilasi, memeriksa kesehatan paket (`krew pkg health`), dan menegakkan gerbang mutu (`krew test`, `gofmt`, `go vet`). Berjalan 100% luring tanpa butuh LLM atau koneksi internet.
2. **Konstitusi Rekayasa (`templates/boost`)**:
   Mendefinisikan aturan operasional, konvensi commit, struktur direktori, dan definisi persona agen yang tertanam langsung di repositori via `AGENTS.md` dan `.agents/`.
3. **Engine Kognitif (`krew`)**:
   Menjalankan siklus kognitif ReAct (Think $\to$ Act $\to$ Observe), memanggil alat melalui protokol standar Model Context Protocol (MCP), dan mengelola konteks percakapan.
4. **Platform Pendukung**:
   - **Dash (`dash`)**: Workspace virtual dengan intervensi manusia untuk meninjau *diff*, mengevaluasi eksekusi terminal, dan menyetujui *pull request*.
   - **Cloud (`cloud`)**: Lingkungan sandbox terisolasi untuk proses runtime dan hosting PaaS hibrida.
   - **Auth (`auth`)**: Identitas multi-tenant, verifikasi token OIDC, dan otorisasi RBAC untuk setiap aksi agen.

---

## 3. Batas Otoritas & Persetujuan Manusia

Untuk memastikan keamanan tanpa menghambat kecepatan pengembang, kapabilitas agen dibagi ke dalam tiga tingkat operasional:

- **Tingkat 1: Otonom (Tanpa Persetujuan Khusus)**:
  - Membaca file repositori, dokumentasi, dan catatan keputusan arsitektur (ADR).
  - Memeriksa konteks proyek (`kiw context --json`) dan aturan linter.
  - Menjalankan kueri baca-saja, pemeriksaan status, dan unit test (`krew test`).
- **Tingkat 2: Diawasi (Tindakan Dicatat, Ditinjau saat PR)**:
  - Membuat atau mengubah kode aplikasi, gaya CSS, template, dan unit test.
  - Melakukan *scaffolding* komponen atau rute baru (`krew make component <Nama>`).
  - Memformat kode dan merapikan impor (`gofmt`, `go vet`).
- **Tingkat 3: Wajib Persetujuan Manusia (Gated)**:
  - Perintah sistem berkas yang merusak (`rm -rf`, penghapusan file di luar target kerja).
  - Menjalankan migrasi database atau perubahan skema permanen di *staging/production*.
  - Melakukan *commit* dan *push* ke cabang publik bersama (`git push`).
  - Mengakses soket jaringan eksternal atau memicu webhook pihak ketiga.

---

## 4. Alur Kerja Standar: Dari Tugas Hingga Produksi

Berikut adalah 7 tahapan kanonikal dalam menyelesaikan fitur atau perbaikan bug menggunakan agen AI:

```text
[1. Penemuan Konteks] ──► [2. Penyelarasan] ──► [3. Rencana Bertahap] ──► [4. Bangun Berbasis Tes]
                                                                                   │
[7. Rilis & Verifikasi] ◄── [6. Audit Keamanan] ◄── [5. Gerbang Mutu] ◄───────────┘
```

### Langkah 1: Penemuan Konteks & Gerbang Awal
Sebelum menulis kode apa pun, agen memeriksa konteks proyek resmi yang terbaca-mesin:

```bash
# Menemukan jenis beban kerja, dependensi, aturan, dan cakupan
kiw context --json
krew info
```

Agen membaca `AGENTS.md`, spesifikasi aktif di GitHub Issues, dan batasan arsitektur piramida 6 layer.

### Langkah 2: Klarifikasi & Penyelarasan
Jika persyaratan masih ambigu, agen wajib mengajukan pertanyaan klarifikasi sebelum menulis kode alih-alih menebak asumsi.

### Langkah 3: Rencana Bertahap
Agen menyusun rencana implementasi terstruktur yang merinci file-file yang terdampak, analisis risiko, dan langkah validasi. Rencana berdampak besar ditinjau oleh pengembang manusia sebelum eksekusi dimulai.

### Langkah 4: Implementasi Berbasis Pengujian
Perubahan diimplementasikan dalam kelompok-kelompok kecil ($\le 300$ baris kode per commit). Jika memungkinkan, tulis pengujian yang gagal terlebih dahulu untuk mereproduksi bug atau mendefinisikan kontrak baru:

```bash
# Memverifikasi status awal
go test -v ./pkg/target/...
```

### Langkah 5: Gerbang Mutu Deterministik
Setiap perubahan harus memenuhi 4 gerbang mutu wajib secara independen:

```bash
# 1. Pemeriksaan format (wajib 0 baris)
gofmt -l .

# 2. Analisis statis (wajib keluar kode 0)
go vet ./...

# 3. Unit test dan pengujian integrasi
go test -v ./...

# 4. Kompilasi Go murni tanpa CGO
CGO_ENABLED=0 go build ./...
```

### Langkah 6: Verifikasi AI Berbasis Bukti (`krew verify`)
Agen menjalankan `krew verify` untuk menghasilkan jejak audit yang dapat dibuktikan:

```bash
krew verify --json
```

Langkah ini memvalidasi bahwa seluruh kriteria penerimaan terpenuhi, dependensi terdaftar di `go.mod`, dan tidak ada panggilan jaringan yang mencurigakan.

### Langkah 7: Commit Atomik & Push
Perubahan di-*commit* secara *bottom-up* menggunakan format *Conventional Commits* dengan tautan ke isu terkait:

```bash
git add pkg/auth/token.go pkg/auth/token_test.go
git commit -m "feat(auth): prevent token replay within clock window (ref #12)"
git push origin main
```

---

## 5. Keamanan & Mitigasi Model Ancaman

Alur kerja AI-Native memperkenalkan pertimbangan keamanan khusus yang wajib dimitigasi:

1. **Prompt Injection & Data Tidak Tepercaya**:
   - Deskripsi isu, komentar pengguna, dan halaman web eksternal adalah masukan tidak tepercaya. Agen wajib memperlakukan instruksi dalam data pengguna sebagai teks pasif, bukan sebagai perintah operasional sistem.
2. **Keamanan Traversal Berkas (`SafePath`)**:
   - Setiap manipulasi berkas harus dibatasi dalam direktori proyek. Agen dan alat wajib menggunakan `sec.SafePath` untuk menolak percobaan traversal (`../../etc/passwd`).
3. **Invarian Tanpa Kredensial Tersimpan**:
   - Kredensial tidak boleh ada di prompt, commit, maupun kode sumber. Selalu gunakan variabel lingkungan dinamis (`${env:DB_PASSWORD}`).
4. **Verifikasi Asal-Usul Dependensi**:
   - Untuk mencegah serangan *dependency confusion* atau paket bayangan (*phantom*), verifikasi seluruh modul:
   ```bash
   go mod verify
   krew pkg health
   ```

---

## 6. Template Artefak Siap Pakai

### Template Ringkasan Tugas (`TASK.md`)

```markdown
# Tugas: [Judul Fitur / Perbaikan Bug]

## Tujuan
[Deskripsi jelas dalam 1-2 kalimat mengenai definisi selesai.]

## Batasan & Lingkup Arsitektur
- Repositori: [misal: krewire/pkg/web]
- Jenis Beban Kerja: [misal: app / service / worker]
- Tingkat Layer: [Layer 1-3 Inti / Layer 4 Platform / Layer 5 Portal]

## Kriteria Penerimaan
- [ ] Persyaratan 1: [Perilaku spesifik yang dapat diverifikasi]
- [ ] Persyaratan 2: [Penanganan kasus batas tertentu]
- [ ] Unit test lolos: go test -v ./...
- [ ] Kompilasi Go murni: CGO_ENABLED=0 go build ./...

## Di Luar Lingkup (Out of Scope)
- [Batasan eksplisit untuk mencegah pelebaran lingkup kerja]
```

### Template Catatan Keputusan Arsitektur (`ADR-XXXXX.md`)

```markdown
# ADR-XXXXX: [Judul Keputusan Arsitektur]

| Kolom | Nilai |
|---|---|
| Status | Proposed / Accepted / Superseded |
| Tanggal | YYYY-MM-DD |
| Pengambil Keputusan | [Nama / Peran Pengelola] |

## Konteks
[Pernyataan masalah, batasan teknologi, dan latar belakang.]

## Keputusan
[Arsitektur, pola, atau invarian dependensi yang dipilih.]

## Konsekuensi
- Positif: [Keuntungan dan peningkatan yang diperoleh.]
- Negatif: [Kompromi, kompleksitas, atau biaya pemeliharaan yang diterima.]
- Invarian: [Aturan mutlak yang ditegakkan ke layer di bawahnya.]
```

---

## 7. Matriks Kesiapan Kapabilitas Produksi

| Kapabilitas | Komponen | Status | Kesiapan Produksi |
|---|---|---|---|
| Native MCP Server (`stdio`) | `krew` | Rilis & Terverifikasi | Ya |
| Unified AI SDK (`pkg/ai`) | `krewire/pkg/ai` | Rilis (v0.3+) | Ya |
| Konstitusi 21 Agen Boost | `templates/boost` | Rilis & Terverifikasi | Ya |
| Verifikasi AI (`krew verify`) | `krew` | Rilis (v0.3.2) | Ya |
| Sandbox Web PTY (Bubblewrap) | `dash/pkg/sandbox` | Rilis (Linux amd64/arm64) | Ya (Fail-closed) |
| Multi-Agent Dynamic Swarms | `krew/pkg/sdlc` | Rilis (ref #15) | Ya |
| Migrasi Beban Kerja Otonom | `cloud/pkg/supervisor` | Pengembangan Aktif | Single-Tenant Saja |

---

## 8. Rangkuman Tindakan untuk Tim

- [x] Konfigurasikan `.cursor/mcp.json` atau `opencode.json` dengan `krew mcp --stdio`.
- [x] Pertahankan `AGENTS.md` yang otoritatif di root repositori Anda.
- [x] Terapkan 4 gerbang mutu dalam alur *continuous integration*.
- [x] Pastikan manusia memegang kendali persetujuan atas tindakan berisiko tinggi.
- [x] Validasi kontribusi eksternal terhadap [Kebijakan Kontribusi Berbantuan AI](https://open.krewire.com/docs/ai-contributions).
